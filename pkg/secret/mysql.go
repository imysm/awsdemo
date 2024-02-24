package secret

import (
	"encoding/json"
	"log"
)

type MySQLOption struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Engine     string `json:"engine"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   string `json:"dbname"`
	Identifier string `json:"dbInstanceIdentifier"`
}

func GetMySQLOption(secretName, region string) *MySQLOption {
	secretString := getSecretValue(secretName, region)
	var r MySQLOption
	err := json.Unmarshal([]byte(secretString), &r)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &r

}
