package config

import (
	"github.com/Muhammad-Aziz/gCore/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func (configPackageObject *ConfigPackage) Load() bool {
	configPackageObject.loadConfigFile()
	configPackageObject.updateConfigVariableOnFileChange()
	return true
}

func (configPackageObject *ConfigPackage) loadConfigFile() {
	viper.SetConfigFile(configPackageObject.GetAppConfigFile())
	utils.Must(viper.ReadInConfig())
	utils.Must(viper.Unmarshal(configPackageObject.GetAppConfigPointer()))
}

func (configPackageObject *ConfigPackage) updateConfigVariableOnFileChange() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//LogWarn().Msg("Config file changed:" + e.Name)
		utils.Must(viper.Unmarshal(configPackageObject.GetAppConfigPointer()))
	})
}
