package imail

type IMailManager interface {
	GetGlobalMailId() int32
	GetGlobalMail(globalMailId int32) []*GlobalMailDef
	SendGlobalMail(pMailDef *GlobalMailDef)
}
