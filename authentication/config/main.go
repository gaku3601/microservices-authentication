package config

import (
	"errors"
	"fmt"
	"os"

	viper "github.com/spf13/viper.git"
)

//configファイルの設定
func SetConfig(configPath string) error {
	if os.Getenv("AuthEnv") == "production" {
		fmt.Println("環境:production")
		return readConfig("config.production", configPath)
	} else if os.Getenv("AuthEnv") == "develop" {
		fmt.Println("環境:develop")
		return readConfig("config.develop", configPath)
	} else {
		return errors.New("developもしくはproductionを環境変数AuthEnvに指定して下さい。")
	}
	return nil
}

//configファイルの読み込み
func readConfig(filename string, path string) error {
	if filename == "" || path == "" {
		return errors.New("filenameもしくはpathが空です。")
	}
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}
