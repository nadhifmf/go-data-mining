package main

import(
	//"io/ioutil"
	"log"
	//"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)


func getLokalTerdaftar(context *gin.Context){
	url := "https://pse.kominfo.go.id/static/json-static/LOKAL_TERDAFTAR/0.json?page[page]=1.json?page[page]="
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var j interface{}
    err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		log.Fatalln(err)
	}
	context.IndentedJSON(http.StatusOK, j)
}

func main()  {
	router := gin.Default()
	router.GET("/pselokal", getLokalTerdaftar)
	router.Run("localhost:9011")
}