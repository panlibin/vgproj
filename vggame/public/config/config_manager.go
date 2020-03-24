package iconfig

type IConfigManager interface {
	GetConfig(confId int32) IConfig
}
