package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ConfigPackage struct {
	appConfigObjectPointer interface{}
	appConfigFile          string
}

func InitConfig(appConfigObjectPointer interface{}, appConfigFile string) *ConfigPackage {
	configPackageObject := &ConfigPackage{}
	configPackageObject.SetAppConfigObjectPointer(appConfigObjectPointer)
	configPackageObject.SetAppConfigFile(appConfigFile)
	return configPackageObject
}

func (configPackageObject *ConfigPackage) SetAppConfigObjectPointer(appConfigPointer interface{}) error {
	if appConfigPointer == nil {
		return fmt.Errorf("appConfigPointer cannot be nil")
	}

	refType := reflect.TypeOf(appConfigPointer)
	if refType.Kind() != reflect.Ptr {
		return fmt.Errorf("appConfigPointer must be a pointer")
	}

	if refType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("appConfigPointer must be a pointer to a struct")
	}

	configPackageObject.appConfigObjectPointer = appConfigPointer
	return nil
}
func (configPackageObject *ConfigPackage) GetAppConfigPointer() interface{} {
	return configPackageObject.appConfigObjectPointer
}

func (configPackageObject *ConfigPackage) SetAppConfigFile(appConfigFile string) error {
	if strings.Compare(appConfigFile, "") == 0 {
		return errors.New("empty config file path")
	}

	configPackageObject.appConfigFile = appConfigFile
	return nil
}
func (configPackageObject *ConfigPackage) GetAppConfigFile() string {
	return configPackageObject.appConfigFile
}
