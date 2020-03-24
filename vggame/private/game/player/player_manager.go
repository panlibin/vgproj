package player

import (
	"sync"
	ec "vgproj/common/define/err_code"
	"vgproj/common/util"
	"vgproj/vggame/public"
	iplayer "vgproj/vggame/public/game/player"
	igate "vgproj/vggame/public/gate"
)

const ServerIdOffset int64 = 10000000

type insertContext struct {
	playerId  int64
	accountId int64
	serverId  int32
	name      string
	head      int32
	cb        func(interface{}, iplayer.IPlayer, int32)
	ctx       interface{}
}

type PlayerManager struct {
	mapPlayer          map[int64]*Player
	mapOnlinePlayer    map[int64]iplayer.IPlayer
	mapNamePlayerId    map[string]int64
	mapAccountPlayerId map[int32]map[int64]int64
	maxPlayerId        int64
	msgRouter          *messageRouter
}

func NewPlayerManager(msgDesc *util.MessageDescriptor) *PlayerManager {
	pObj := new(PlayerManager)
	pObj.msgRouter = &messageRouter{}
	pObj.msgRouter.init(msgDesc)
	pObj.mapPlayer = make(map[int64]*Player)
	pObj.mapOnlinePlayer = make(map[int64]iplayer.IPlayer)
	pObj.mapNamePlayerId = make(map[string]int64)
	pObj.mapAccountPlayerId = make(map[int32]map[int64]int64)
	arrServerId := public.Server.GetServerIdArray()
	for _, serverId := range arrServerId {
		pObj.mapAccountPlayerId[serverId] = make(map[int64]int64)
	}

	return pObj
}

func (pm *PlayerManager) OnLoadData() error {
	rows, err := public.Server.GetGlobalDb().Query(0, "select player_id,account_id,server_id,name from player_data")
	if err != nil {
		return err
	}
	var tmpPlayerId int64
	var tmpAccountId int64
	var tmpName string
	var tmpMaxPlayerId int64
	var tmpServerId int32
	for rows.Next() {
		if err = rows.Scan(&tmpPlayerId, &tmpAccountId, &tmpServerId, &tmpName); err != nil {
			return err
		}
		pm.mapNamePlayerId[tmpName] = tmpPlayerId
		pm.mapPlayer[tmpPlayerId] = nil
		pm.mapAccountPlayerId[tmpServerId][tmpAccountId] = tmpPlayerId
		tmpMaxPlayerId = tmpPlayerId % ServerIdOffset
		if tmpMaxPlayerId > pm.maxPlayerId {
			pm.maxPlayerId = tmpMaxPlayerId
		}
	}
	rows.Close()

	return nil
}

func (pm *PlayerManager) OnInit() error {
	// public.Server.GetGameManager().GetEventManager().Register(igame.EventType_DailyRefresh, pm.onDailyRefresh)
	return nil
}

func (pm *PlayerManager) OnRelease() {

}

func (pm *PlayerManager) CreatePlayer(accountId int64, serverId int32, name string, head int32, ctx interface{}, cb func(ctx interface{}, pPlayer iplayer.IPlayer, errCode int32)) {
	var errCode int32 = ec.Unknown

	for {
		if pm.GetPlayerIdByAccountId(accountId, serverId) > 0 {
			errCode = ec.DuplicateCreateCharacter
			break
		}
		if pm.GetPlayerIdByName(name) > 0 {
			errCode = ec.DuplicatePlayerName
			break
		}
		playerId := pm.genPlayerId(serverId)
		pPlayer := newPlayer(playerId)
		pm.mapPlayer[playerId] = pPlayer
		pm.mapNamePlayerId[name] = playerId
		pm.mapAccountPlayerId[serverId][accountId] = playerId
		pInstertCtx := &insertContext{
			playerId:  playerId,
			accountId: accountId,
			serverId:  serverId,
			name:      name,
			head:      head,
			cb:        cb,
			ctx:       ctx,
		}
		pPlayer.insert(accountId, serverId, name, head, pInstertCtx, pm.createPlayerInsertCallback)

		errCode = ec.Success
		break
	}

	if errCode != ec.Success {
		cb(ctx, nil, errCode)
	}
}

func (pm *PlayerManager) createPlayerInsertCallback(ctx interface{}, iPlayer iplayer.IPlayer) {
	if iPlayer == nil {
		pInsertCtx := ctx.(insertContext)
		delete(pm.mapPlayer, pInsertCtx.playerId)
		delete(pm.mapAccountPlayerId[pInsertCtx.serverId], pInsertCtx.accountId)
		delete(pm.mapNamePlayerId, pInsertCtx.name)

		pInsertCtx.cb(pInsertCtx.ctx, nil, ec.CreateCharacterFail)
		return
	}

	iPlayer.(*Player).init(ctx, pm.createPlayerInitCallback)
}

func (pm *PlayerManager) createPlayerInitCallback(ctx interface{}, iPlayer iplayer.IPlayer) {
	pInsertCtx := ctx.(insertContext)
	if iPlayer == nil {
		pm.mapPlayer[pInsertCtx.playerId] = nil

		pInsertCtx.cb(pInsertCtx.ctx, nil, ec.CreateCharacterFail)
		return
	}

	pPlayer := iPlayer.(*Player)
	pPlayer.createInit()
	pInsertCtx.cb(pInsertCtx.ctx, iPlayer, ec.Success)
}

func (pm *PlayerManager) GetPlayerIdByName(name string) int64 {
	playerId, exist := pm.mapNamePlayerId[name]
	if !exist {
		return 0
	}
	return playerId
}

func (pm *PlayerManager) GetPlayerIdByAccountId(accountId int64, serverId int32) int64 {
	mapPlayerId, exist := pm.mapAccountPlayerId[serverId]
	if !exist {
		return 0
	}
	playerId, exist := mapPlayerId[accountId]
	if !exist {
		return 0
	}
	return playerId
}

func (pm *PlayerManager) GetPlayer(playerId int64, ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	pPlayer, exist := pm.mapPlayer[playerId]
	if !exist {
		cb(ctx, nil)
		return
	}
	if pPlayer != nil {
		if pPlayer.status == EPlayerStatus_Loading {
			pPlayer.waitLoadFinish(ctx, cb)
		} else {
			cb(ctx, pPlayer)
		}
	} else {
		pm.loadPlayer(playerId, ctx, cb)
	}
}

func (pm *PlayerManager) GetPlayers(arrPlayerId []int64, ctx interface{}, cb func(interface{}, map[int64]iplayer.IPlayer)) {
	wg := sync.WaitGroup{}
	mapPlayer := make(map[int64]iplayer.IPlayer)
	for _, playerId := range arrPlayerId {
		pPlayer, exist := pm.mapPlayer[playerId]
		if !exist {
			continue
		}

		if pPlayer != nil {
			if pPlayer.status == EPlayerStatus_Loading {
				wg.Add(1)
				pPlayer.waitLoadFinish(nil, func(_ interface{}, pPlayer iplayer.IPlayer) {
					if pPlayer != nil {
						mapPlayer[playerId] = pPlayer
					}
					wg.Done()
				})
			} else {
				mapPlayer[playerId] = pPlayer
			}
		} else {
			wg.Add(1)
			pm.loadPlayer(playerId, nil, func(_ interface{}, pPlayer iplayer.IPlayer) {
				if pPlayer != nil {
					mapPlayer[playerId] = pPlayer
				}
				wg.Done()
			})
		}
	}

	public.Server.AsyncTask(func([]interface{}) {
		wg.Wait()

		public.Server.SyncTask(func([]interface{}) {
			cb(ctx, mapPlayer)
		})
	})
}

func (pm *PlayerManager) GetOnlinePlayers() map[int64]iplayer.IPlayer {
	return pm.mapOnlinePlayer
}

func (pm *PlayerManager) SetPlayerOnline(playerId int64) {
	pPlayer, exist := pm.mapPlayer[playerId]
	if !exist || pPlayer == nil {
		return
	}
	pm.mapOnlinePlayer[playerId] = pPlayer
}

func (pm *PlayerManager) SetPlayerOffline(playerId int64) {
	delete(pm.mapOnlinePlayer, playerId)
}

func (pm *PlayerManager) loadPlayer(playerId int64, ctx interface{}, cb func(interface{}, iplayer.IPlayer)) {
	pPlayer := newPlayer(playerId)
	pm.mapPlayer[playerId] = pPlayer
	pPlayer.init(nil, func(_ interface{}, iPlayer iplayer.IPlayer) {
		if iPlayer == nil {
			pm.mapPlayer[playerId] = nil
			cb(ctx, nil)
		} else {
			cb(ctx, iPlayer)
		}
	})
}

func (pm *PlayerManager) genPlayerId(serverId int32) int64 {
	pm.maxPlayerId++
	return int64(serverId)*ServerIdOffset + pm.maxPlayerId
}

func (pm *PlayerManager) GetMessageRouter() igate.IMessageRouter {
	return pm.msgRouter
}
