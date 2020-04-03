package game

const (
	Module_Player int32 = iota
	Module_CustomLanguage
	Module_Mail
	Module_Rank

	Module_Count
)

type IModule interface {
	OnLoadData() error
	OnInit() error
	OnRelease()
}
