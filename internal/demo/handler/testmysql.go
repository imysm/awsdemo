package handler

import (
	"awsdemo/internal/demo/model"
	"awsdemo/internal/pkg/errno"
	"awsdemo/pkg/connector"
	"awsdemo/pkg/secret"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
)

func Insert(c *gin.Context) {
	v := secret.GetMySQLOption(viper.GetString("secretMySQLName"), viper.GetString("region"))

	optsMySQL := &connector.MySQLOptions{Host: v.Host, Port: v.Port, Database: v.Database,
		Username: v.Username, Password: v.Password}
	//log.Println(optsMySQL)
	db, err := connector.NewMySQL(optsMySQL)
	if err != nil {
		SendResponse(c, errno.ErrDBConn, err)
		log.Fatalln("mysql connect failed.", err)
	}

	model.Migrate(db)

	// 创建一个User实例
	test := model.Test{Name: c.Query("name"), Message: c.Query("message"), CreatedAt: time.Now(), LastUpdatedAt: time.Now()}

	// 使用Create方法插入数据
	result := db.Create(&test)

	// 检查是否有错误发生
	if result.Error != nil {
		SendResponse(c, errno.ErrCrtRecord, result.Error)
		fmt.Println("Failed to create record")
	} else {
		SendResponse(c, nil, fmt.Sprintf("Record created with ID: %v\n", test.ID))
		fmt.Printf("Record created with ID: %v\n", test.ID)
	}
}
