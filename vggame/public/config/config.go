package iconfig

const (
	Config_System = iota
	Config_Item
	Config_Level
	Config_Vip

	Config_Count
)

type IConfig interface {
	OnLoadConfig() error
}
