package player

import (
	"encoding/json"
	"vgproj/proto/msg"
	"vgproj/vggame/public"
	"vgproj/vggame/public/define/filler"
	imail "vgproj/vggame/public/game/mail"
)

type Mail struct {
	playerId      int64
	mailId        int64
	source        int32
	sourceExt     string
	ts            int64
	firstType     int32
	secondType    int32
	title         string
	titleParams   []*filler.Params
	content       string
	contentParams []*filler.Params
	attachment    []*imail.MailAttachment
	isNew         int32
	isRead        int32
	isGot         int32
}

func newMail() *Mail {
	pObj := new(Mail)
	pObj.isNew = 1
	return pObj
}

func (m *Mail) insert() {
	jsonTitleParams, _ := json.Marshal(m.titleParams)
	jsonContentParams, _ := json.Marshal(m.contentParams)
	jsonAttachments, _ := json.Marshal(m.attachment)

	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(m.playerId), sqlInsertMail, m.playerId, m.mailId, m.source, m.sourceExt, m.ts, m.firstType,
		m.secondType, m.title, jsonTitleParams, m.content, jsonContentParams, jsonAttachments, m.isNew, m.isRead, m.isGot)
}

func (m *Mail) save() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(m.playerId), sqlUpdateMailStatus, m.isNew, m.isRead, m.isGot, m.playerId, m.mailId)
}

func (m *Mail) delete() {
	public.Server.GetDataDb().AsyncExec(nil, nil, uint32(m.playerId), sqlDeleteMail, m.playerId, m.mailId)
}

func (m *Mail) formatToMessage() *msg.MAIL_INFO {
	pMsgMail := new(msg.MAIL_INFO)
	pMsgMail.MailId = m.mailId
	pMsgMail.FirstType = m.firstType
	pMsgMail.SecondType = m.secondType
	pMsgMail.MailTitle = m.title
	pMsgMail.TitleParams = filler.FormatToMessage(m.titleParams)
	pMsgMail.MailDesc = m.content
	pMsgMail.DescParams = filler.FormatToMessage(m.contentParams)
	pMsgMail.Ts = m.ts
	pMsgMail.Items = imail.FormatAttachmentToMessage(m.attachment)
	pMsgMail.IsNew = m.isNew != 0
	pMsgMail.IsItemGot = m.isGot != 0
	pMsgMail.IsReaded = m.isRead != 0
	return pMsgMail
}
