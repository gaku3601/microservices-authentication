package config

import (
	"os"
	"testing"
)

func TestSetConfig(t *testing.T) {
	//環境:productionの場合、configが読み込めるかのテスト
	os.Setenv("AuthEnv", "production")
	err := SetConfig(".")

	if err == nil {
	} else {
		t.Log(err)
		t.Errorf("config読み取りエラー:production環境のconfigが読み込めません")
	}

	//環境:developの場合、configが読み込めるかのテスト
	os.Setenv("AuthEnv", "develop")
	err = SetConfig(".")

	if err == nil {
	} else {
		t.Errorf("config読み取りエラー:develop環境のconfigが読み込めません")
	}

	//環境:developでもproductionでもない環境変数が指定されている場合、Errorが返却されるか
	os.Setenv("AuthEnv", "abababa")
	err = SetConfig(".")
	if err.Error() == "developもしくはproductionを環境変数AuthEnvに指定して下さい。" {
	} else {
		t.Errorf("エラーハンドリングエラー")
	}
}

func TestReadConfig(t *testing.T) {
	//空filenameを指定するとerrorが返却されるか。
	err := readConfig("", "a")
	if err != nil {
	} else {
		t.Errorf("エラーが返却されていません。")
	}

	//空pathを指定するとerrorが返却されるか。
	err = readConfig("a", "")
	if err != nil {
	} else {
		t.Errorf("エラーが返却されていません。")
	}

	//存在しないfilepathとpathを指定するとerrorが返却されるか。
	err = readConfig("a", "a")
	if err != nil {
	} else {
		t.Errorf("エラーが返却されていません。")
	}
}
