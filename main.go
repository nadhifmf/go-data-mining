package main

import (
	gg "go_mining/mining_app"

	"github.com/gin-gonic/gin"
)

//"io/ioutil"
//"log"
//"errors"
//"encoding/json"
//"net/http"

func main() {
	router := gin.Default()
	router.GET("/pselokal", gg.GetLokalTerdaftar)
	router.Run("localhost:9011")
}
