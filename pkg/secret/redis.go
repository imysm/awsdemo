package secret

import (
	"encoding/json"
	"log"
)

type RedisOption struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func GetRedisOption(secretName, region string) *RedisOption {
	secretString := getSecretValue(secretName, region)
	var r RedisOption
	err := json.Unmarshal([]byte(secretString), &r)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &r

}
