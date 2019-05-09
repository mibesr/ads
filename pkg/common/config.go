package common

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	HttpPort        int    `json:"http_port"`
	ApiReadTimeout  int    `json:"api_read_timeout"`
	ApiWriteTimeout int    `json:"api_write_timeout"`
	MongoDBUri      string `json:"mongo_db_uri"`
	RedisUri        string `json:"redis_uri"`
	MongoDBTimeout  int    `json:"mongo_db_timeout"`
	DBName          string `json:"db_name"`
}

var (
	GConfig *Config
)

func InitConfig(filename string) (config Config, err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	return conf, err
}
