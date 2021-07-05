package main

import (
	"flag"

	"github.com/Atlantis-Org/fcm/controllers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	db_uri := flag.String("db", "fnlinker:123456@(dev_mysql01.fnlinker.com:6306)/fcm?charset=utf8mb4&parseTime=True", "db connect string")
	db_max_conn := flag.Int("db-maxconn", 5, "db max connection num")
	db_maxopen_conn := flag.Int("db-maxopenconn", 3, "db max open connection num")
	utils.InitDb(*db_uri, *db_max_conn, *db_maxopen_conn)
}

func main() {
	models.DbMigrate()

	r := gin.Default()
	r.GET("/client/config", controllers.GetClientConfig)
	r.POST("/app/create", controllers.CreateApp)
	r.POST("/app/namespace/create", controllers.CreateAppNamespace)
	r.POST("/app/config/create", controllers.CreateAppConfig)
	r.POST("/group/create", controllers.CreateGroup)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
