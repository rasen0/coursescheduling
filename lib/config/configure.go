package config

import (
	"coursesheduling/lib/util"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"

	"coursesheduling/common"
	"coursesheduling/lib/log"
	"github.com/spf13/viper"
)

type DBInfo struct {
	DBName string `json:"dbName" default:"courseschedule.db"`
	DBUser string `json:"dbUser" default:"root"`
	DBPassword string `json:"dbPassword" default:"root"`
	IpAddress string `json:"IpAddress" default:"localhost:3306"`
}

type GUIConfigure struct {
	Address string `json:"address" default:"127.0.0.1:8000"`
}

type Configure struct {
	DBInfo
	GUIConfigure
}

func InitConfigure() (config Configure,err error){
	absPath,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	if _,err = os.Stat(filepath.Join(absPath,common.Course,common.ConfigFile));os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(absPath,common.Course,common.ConfigPath),os.ModePerm)
	}
	cfgFile := filepath.Join(absPath,common.Course,common.ConfigFile)
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	config  = Configure{}
	if _,err= os.Stat(cfgFile); os.IsNotExist(err) {
		log.Debug("configure file is not exist,init file")
		util.SetDefaults(&config)
		createFile, err := os.Create(cfgFile)
		defer createFile.Close()
		if err != nil{
			err = fmt.Errorf("create file.%w",err)
			return config,err
		}
		bs, err := yaml.Marshal(&config)
		if err != nil{
			err = fmt.Errorf("marsh new config. %w",err)
			return config, err
		}
		_, err = createFile.Write(bs)
		if err != nil{
			return config,errors.New("write file")
		}
	} else {
		err := viper.ReadInConfig()
		if err != nil{
			err = fmt.Errorf("ReadInConfig fail.%w",err)
			return config, err
		}
		config.DBName = viper.GetString("dbInfo.dbName")
		config.DBUser = viper.GetString("dbInfo.dbUser")
		config.DBPassword = viper.GetString("dbInfo.dbPassword")
		config.Address = viper.GetString("GUIConfigure.address")
	}
	viper.WatchConfig()
	//viper.OnConfigChange()
	return config, nil
}