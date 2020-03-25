package mail

import (
	"encoding/json"
	"sort"
	"vgproj/vggame/public"
	imail "vgproj/vggame/public/game/mail"
	iplayer "vgproj/vggame/public/game/player"

	"github.com/panlibin/virgo/util/vgtime"
)

type MailManager struct {
	maxMailId int32
	arrMail   MailArray
}

func NewMailManager() *MailManager {
	pObj := new(MailManager)
	pObj.arrMail = make(MailArray, 0)

	return pObj
}

func (mm *MailManager) OnLoadData() error {
	rows, err := public.Server.GetGlobalDb().Query(0, sqlLoadGlobalMail)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		pMail := new(imail.GlobalMailDef)
		var tmpTitleParams []byte
		var tmpContentParams []byte
		var tmpAttachments []byte

		err = rows.Scan(&pMail.GlobalMailId, &pMail.Source, &pMail.SourceExt, &pMail.Ts, &pMail.FirstType, &pMail.SecondType, &pMail.Title, &tmpTitleParams,
			&pMail.Content, &tmpContentParams, &tmpAttachments, &pMail.VipLevLimit)
		if err != nil {
			break
		}

		if pMail.GlobalMailId > mm.maxMailId {
			mm.maxMailId = pMail.GlobalMailId
		}

		err = json.Unmarshal(tmpTitleParams, &pMail.TitleParams)
		if err != nil {
			break
		}
		err = json.Unmarshal(tmpContentParams, &pMail.ContentParams)
		if err != nil {
			break
		}
		arrAttachments := make([]*imail.MailAttachment, 0)
		err = json.Unmarshal(tmpAttachments, &arrAttachments)
		if err != nil {
			break
		}

		mm.arrMail = append(mm.arrMail, pMail)
	}

	return nil
}

func (mm *MailManager) OnInit() error {
	sort.Sort(mm.arrMail)
	return nil
}

func (mm *MailManager) OnRelease() {

}

func (mm *MailManager) GetGlobalMailId() int32 {
	return mm.maxMailId
}

func (mm *MailManager) GetGlobalMail(globalMailId int32) []*imail.GlobalMailDef {
	idx := mm.arrMail.Search(globalMailId)
	if idx >= len(mm.arrMail) {
		return nil
	}
	return mm.arrMail[idx:]
}

func (mm *MailManager) SendGlobalMail(pMailDef *imail.GlobalMailDef) {
	pMailDef.GlobalMailId = mm.genGlobalMailId()
	pMailDef.Ts = vgtime.Now()

	mm.arrMail = append(mm.arrMail, pMailDef)

	mm.insertMail(pMailDef)

	mapPlayers := public.Server.GetGameManager().GetPlayerManager().GetOnlinePlayers()
	for _, pPlayer := range mapPlayers {
		pMailModule := pPlayer.GetModule(iplayer.PlayerModule_Mail).(iplayer.IMailModule)
		pMailModule.SendGlobalMail(pMailDef)
	}
}

func (mm *MailManager) genGlobalMailId() int32 {
	mm.maxMailId++
	return mm.maxMailId
}

func (mm *MailManager) insertMail(pMailDef *imail.GlobalMailDef) {
	jsonTitleParams, _ := json.Marshal(pMailDef.TitleParams)
	jsonContentParams, _ := json.Marshal(pMailDef.ContentParams)
	jsonAttachments, _ := json.Marshal(pMailDef.Attachment)

	public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlInsertGlobalMail, pMailDef.GlobalMailId, pMailDef.Source, pMailDef.SourceExt, pMailDef.Ts, pMailDef.FirstType,
		pMailDef.Title, jsonTitleParams, pMailDef.Content, jsonContentParams, jsonAttachments, pMailDef.VipLevLimit)
}

func (mm *MailManager) deleteMail(mailId int32) {
	idx := mm.arrMail.Search(mailId)
	mm.arrMail = append(mm.arrMail[:idx], mm.arrMail[idx+1:]...)
	public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlDeleteGlobalMail, mailId)
}
