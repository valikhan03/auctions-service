package models

import (
	"io/ioutil"
	"log"
	"github.com/joho/godotenv"
	"os"

	"gopkg.in/yaml.v3"
)

type serverConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var ServerConfigGlobal serverConfig

type elasticConfig struct {
	Addrs    []string `yaml:"addrs"`
	Username string   `yaml:"username"`
	Password string   `yaml:"-"`
}

var ElasticConfigGlobal elasticConfig

type redisConfig struct {
	Network    string `yaml:"network"`
	Addr       string `yaml:"addr"`
	ClientName string `yaml:"client_name"`
	Username   string `yaml:"username"`
	Password   string `yaml:"-"`
	DB         int    `yaml:"db"`
}

var RedisConfigGlobal redisConfig


type KafkaConfigs struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

func GetKafkaConfigs() *KafkaConfigs {
	data, err := ioutil.ReadFile("configs/kafka.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var configs KafkaConfigs

	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatal(err)
	}

	return &configs
}


func InitConfigs() {
	godotenv.Load()

	elastic, err := ioutil.ReadFile("configs/elastic.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(elastic, &ElasticConfigGlobal)
	if err != nil {
		log.Fatal(err)
	}
	ElasticConfigGlobal.Password = os.Getenv("ELASTIC_PASSWORD")

	redis, err := ioutil.ReadFile("configs/redis.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(redis, &RedisConfigGlobal)
	if err != nil {
		log.Fatal(err)
	}
	RedisConfigGlobal.Password = os.Getenv("REDIS_PASSWORD")

	server, err := ioutil.ReadFile("configs/server.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(server, &ServerConfigGlobal)
	if err != nil {
		log.Fatal(err)
	}
}
