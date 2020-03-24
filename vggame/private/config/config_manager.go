package config

import iconfig "vgproj/vggame/public/config"

type ConfigManager struct {
	arrConfig [iconfig.Config_Count]iconfig.IConfig
}

func NewConfigManager() *ConfigManager {
	pObj := new(ConfigManager)

	pObj.arrConfig[iconfig.Config_System] = newSystemConfig()
	pObj.arrConfig[iconfig.Config_Item] = newItemConfig()
	pObj.arrConfig[iconfig.Config_Level] = newLevelConfig()
	pObj.arrConfig[iconfig.Config_Vip] = newVipConfig()

	return pObj
}

func (cm *ConfigManager) LoadConfig() error {
	for _, pConfig := range cm.arrConfig {
		err := pConfig.OnLoadConfig()
		if err != nil {
			return err
		}
	}
	return nil
}

func (cm *ConfigManager) GetConfig(confId int32) iconfig.IConfig {
	return cm.arrConfig[confId]
}
