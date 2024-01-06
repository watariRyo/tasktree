package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Env    string `json:"env"`
		Host   string `json:"host"`
		Db     Db     `yaml:"db"`
		Server Server `yaml:"server"`
		Redis  Redis  `yaml:"redis"`
		Jwt    Jwt    `yaml:"jwt"`
	}

	Db struct {
		Host      string `yaml:"host"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Port      string `yaml:"port"`
		Schema    string `yaml:"schema"`
		DebugMode bool   `yaml:"debugMode"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Key      string `yaml:"key"`
	}

	Server struct {
		Port           string `yaml:"port"`
		CorsOrigins    string `yaml:"corsOrigins"`
		FirebaseSecret string `yaml:"firebaseSecret"`
	}

	Jwt struct {
		Key string `yaml:"key"`
	}
)

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	projectRoot := projectRoot()
	viper.AddConfigPath(fmt.Sprintf("/%s/config/", projectRoot))

	// 環境変数が指定されていればそちらを優先
	viper.AutomaticEnv()
	// データ構造切り替え
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー： %s \n", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal error: %s \n", err)
	}

	return &cfg, nil
}

func projectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				return ""
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}
