package imail

import (
	"vgproj/proto/msg"
	"vgproj/vggame/public/define/filler"
)

type MailAttachment struct {
	Id    int32
	Num   int64
	Extra []int64
}

type GlobalMailDef struct {
	GlobalMailId  int32
	Source        int32
	SourceExt     string
	Ts            int64
	FirstType     int32
	SecondType    int32
	Title         string
	TitleParams   []*filler.Params
	Content       string
	ContentParams []*filler.Params
	Attachment    []*MailAttachment
	VipLevLimit   int32
}

type PlayerMailDef struct {
	Source        int32
	SourceExt     string
	FirstType     int32
	SecondType    int32
	Title         string
	TitleParams   []*filler.Params
	Content       string
	ContentParams []*filler.Params
	Attachment    []*MailAttachment
}

func FormatAttachmentToMessage(args []*MailAttachment) []*msg.MAIL_ATTACHMENT {
	ret := make([]*msg.MAIL_ATTACHMENT, len(args))
	for i, v := range args {
		pMsgAtt := new(msg.MAIL_ATTACHMENT)
		pMsgAtt.Id = v.Id
		pMsgAtt.Num = v.Num
		pMsgAtt.Extra = v.Extra
		ret[i] = pMsgAtt
	}
	return ret
}
