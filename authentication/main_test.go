package main

import (
	"testing"

	viper "github.com/spf13/viper.git"
)

func TestReadConfig(t *testing.T) {
	readConfig("./config")
	//config読み取り
	user := viper.GetString("database.user")
	if user != "" {
	} else {
		t.Errorf("config読み取りエラー:configが読み込めません")
	}
}
