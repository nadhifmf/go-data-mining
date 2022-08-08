package main

import(
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main()  {
	resp, err := http.Get("https://pse.kominfo.go.id/static/json-static/LOKAL_TERDAFTAR/0.json?page[page]=1.json?page[page]=")
	if err != nil {
	   log.Fatalln(err)
	}
 //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 //Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	router := gin.Default()
}