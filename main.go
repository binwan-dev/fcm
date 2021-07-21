package main

import (
	"flag"

	"github.com/Atlantis-Org/fcm/controllers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(config))

	r.GET("/client/config", controllers.GetClientConfig)
	r.GET("/app", controllers.GetAppPaged)
	r.GET("/app/:appId", controllers.GetAppForId)
	r.POST("/app/create", controllers.CreateApp)
	r.POST("/app/namespace/create", controllers.CreateAppNamespace)

	r.GET("/app/namespace", controllers.GetAppNamespacePaged)
	r.GET("/app/namespace/config", controllers.GetAppConfigs)
	r.POST("/app/namespace/config", controllers.CreateAppConfig)
	r.PUT("/app/namespace/config", controllers.ModifyAppConfig)
	r.POST("/group/create", controllers.CreateGroup)
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
