package player

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	imail "vgproj/vggame/public/game/mail"
	iplayer "vgproj/vggame/public/game/player"

	logger "github.com/panlibin/vglog"
	"github.com/panlibin/virgo/util/vgtime"
)

type MailModule struct {
	*playerModule

	maxMailId    int64
	globalMailId int32
	mapMail      map[int64]*Mail

	inDb bool
}

func newMailModule(pPlayer *Player) *MailModule {
	pObj := new(MailModule)
	pObj.playerModule = newPlayerModule(pPlayer)
	pObj.mapMail = make(map[int64]*Mail, 64)

	return pObj
}

func (mm *MailModule) getLoadSql() string {
	return fmt.Sprintf("select global_mail_id from player_mail_ctrl where player_id=%d;"+
		"select mail_id,source,source_ext,ts,first_type,second_type,title,title_params,content,content_params,"+
		"attachments,is_new,is_read,is_got from player_mail_list where player_id=%d;", mm.player.GetId(), mm.player.GetId())
}

func (mm *MailModule) onLoadData(rows *sql.Rows) error {
	var err error = rows.Err()
	if err != nil {
		logger.Errorf("load player_mail_ctrl error: %v", err)
		return err
	}
	if rows.Next() {
		err = rows.Scan(&mm.globalMailId)
		if err != nil {
			logger.Errorf("load player_mail_ctrl error: %v", err)
			return err
		} else {
			mm.inDb = true
		}
	}

	rows.NextResultSet()
	err = rows.Err()
	if err != nil {
		logger.Errorf("load player_mail_list error: %v", err)
		return err
	}
	for rows.Next() {
		pMail := new(Mail)
		pMail.playerId = mm.player.GetId()
		var tmpTitleParams []byte
		var tmpContentParams []byte
		var tmpAttachments []byte

		err = rows.Scan(&pMail.mailId, &pMail.source, &pMail.sourceExt, &pMail.ts, &pMail.firstType, &pMail.secondType, &pMail.title, &tmpTitleParams,
			&pMail.content, &tmpContentParams, &tmpAttachments, &pMail.isNew, &pMail.isRead, &pMail.isGot)
		if err != nil {
			break
		}

		if pMail.mailId > mm.maxMailId {
			mm.maxMailId = pMail.mailId
		}

		err = json.Unmarshal(tmpTitleParams, &pMail.titleParams)
		if err != nil {
			break
		}
		err = json.Unmarshal(tmpContentParams, &pMail.contentParams)
		if err != nil {
			break
		}
		err = json.Unmarshal(tmpAttachments, &pMail.attachment)
		if err != nil {
			break
		}
		mm.mapMail[pMail.mailId] = pMail
	}

	if err != nil {
		logger.Errorf("load player_mail_list error: %v", err)
	}
	return err
}

func (mm *MailModule) onCreate() {
	mm.globalMailId = public.Server.GetGameManager().GetMailManager().GetGlobalMailId()
	mm.saveGlobalMailId()
}

func (mm *MailModule) onLogin() {
	arrGlobalMail := public.Server.GetGameManager().GetMailManager().GetGlobalMail(mm.globalMailId)
	if arrGlobalMail != nil {
		for _, pMailDef := range arrGlobalMail {
			mm.addGlobalMail(pMailDef)
		}
		mm.saveGlobalMailId()
	}
}

func (mm *MailModule) SendMail(pMailDef *imail.PlayerMailDef) {
	pMail := mm.addPlayerMail(pMailDef)
	mm.notifyMail(pMail)
}

func (mm *MailModule) SendGlobalMail(pMailDef *imail.GlobalMailDef) {
	pMail := mm.addGlobalMail(pMailDef)
	mm.saveGlobalMailId()
	if pMail != nil {
		mm.notifyMail(pMail)
	}
}

func (mm *MailModule) addPlayerMail(pMailDef *imail.PlayerMailDef) *Mail {
	pMail := newMail()
	pMail.playerId = mm.player.GetId()
	pMail.mailId = mm.genMailId()
	pMail.source = pMailDef.Source
	pMail.sourceExt = pMailDef.SourceExt
	pMail.ts = vgtime.Now()
	pMail.firstType = pMailDef.FirstType
	pMail.secondType = pMailDef.SecondType
	pMail.title = pMailDef.Title
	pMail.titleParams = pMailDef.TitleParams
	pMail.content = pMailDef.Content
	pMail.contentParams = pMailDef.ContentParams
	pMail.attachment = pMailDef.Attachment

	mm.mapMail[pMail.mailId] = pMail

	pMail.insert()

	return pMail
}

func (mm *MailModule) canRecvGlobalMail(pMailDef *imail.GlobalMailDef) bool {
	pDataModule := mm.player.GetModule(iplayer.PlayerModule_Data).(*dataModule)
	if pDataModule.GetCreateTs() > pMailDef.Ts {
		return false
	}
	pVipModule := mm.player.GetModule(iplayer.PlayerModule_Vip).(*vipModule)
	if pVipModule.GetVipLev() < pMailDef.VipLevLimit {
		return false
	}
	return true
}

func (mm *MailModule) addGlobalMail(pMailDef *imail.GlobalMailDef) *Mail {
	if pMailDef.GlobalMailId > mm.globalMailId {
		mm.globalMailId = pMailDef.GlobalMailId
	}
	if !mm.canRecvGlobalMail(pMailDef) {
		return nil
	}

	pMail := newMail()
	pMail.playerId = mm.player.GetId()
	pMail.mailId = mm.genMailId()
	pMail.source = pMailDef.Source
	pMail.sourceExt = pMailDef.SourceExt
	pMail.ts = pMailDef.Ts
	pMail.firstType = pMailDef.FirstType
	pMail.secondType = pMailDef.SecondType
	pMail.title = pMailDef.Title
	pMail.titleParams = pMailDef.TitleParams
	pMail.content = pMailDef.Content
	pMail.contentParams = pMailDef.ContentParams
	pMail.attachment = pMailDef.Attachment

	mm.mapMail[pMail.mailId] = pMail

	pMail.insert()

	return pMail
}

func (mm *MailModule) notifyMail(pMail *Mail) {

}

func (mm *MailModule) setGlobalMailId(id int32) {
	mm.globalMailId = id
	mm.saveGlobalMailId()
}

func (mm *MailModule) genMailId() int64 {
	mm.maxMailId++
	return mm.maxMailId
}

func (mm *MailModule) insert() {
	mm.inDb = true
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(mm.player.GetId()), sqlInsertMailCtrl, mm.player.GetId(), mm.globalMailId)
}

func (mm *MailModule) saveGlobalMailId() {
	if mm.inDb {
		public.Server.GetDataDb().AsyncExec(nil, nil, uint32(mm.player.GetId()), sqlUpdateMailCtrl, mm.globalMailId, mm.player.GetId())
	} else {
		mm.insert()
	}
}

func (mm *MailModule) handleGetMailList(args []interface{}) {
	//pMsgReq := args[0].(*msg.C2S_MAIL)
	pMsgRsp := new(msg.S2C_MAIL)
	pSettingsModule := mm.player.GetModule(iplayer.PlayerModule_Settings).(*SettingsModule)
	pCustomLangMgr := public.Server.GetGameManager().GetCustomLanguageManager()
	lang := pSettingsModule.GetLanguage()
	for _, pMail := range mm.mapMail {
		pMsgMail := pMail.formatToMessage()
		if pMsgMail.SecondType == 0 {
			pMsgMail.MailTitle = pCustomLangMgr.GetLanguageValue(pMsgMail.MailTitle, lang)
			pMsgMail.MailDesc = pCustomLangMgr.GetLanguageValue(pMsgMail.MailDesc, lang)
		}
		pMsgRsp.Mails = append(pMsgRsp.Mails, pMsgMail)

		if pMail.isNew != 0 {
			pMail.isNew = 0
			pMail.save()
		}
	}
	mm.player.SendMessage(pMsgRsp)
}

func (mm *MailModule) handleDelMail(args []interface{}) {
	pMsgReq := args[0].(*msg.C2S_MAIL_DEL)
	for _, mailId := range pMsgReq.Ids {
		pMail, exist := mm.mapMail[mailId]
		if exist {
			delete(mm.mapMail, mailId)
			pMail.delete()
		}
	}
	mm.player.SendMessage(&msg.S2C_MAIL_DEL{})
}

func (mm *MailModule) handleGetAttachment(args []interface{}) {

}

func (mm *MailModule) handleReadMail(args []interface{}) {
	pMsgReq := args[0].(*msg.C2S_MAIL_READ_REQ)
	for _, mailId := range pMsgReq.Ids {
		pMail, exist := mm.mapMail[mailId]
		if exist && pMail != nil && pMail.isRead == 0 {
			pMail.isRead = 1
			pMail.save()
		}
	}
	mm.player.SendMessage(&msg.S2C_MAIL_READ_RET{Ids: pMsgReq.Ids})
}
