package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"telegram-callback-svc/dto"
	"telegram-callback-svc/properties"
	"telegram-callback-svc/client"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"encoding/json"
)

var appProps *properties.AppProperties
var rabbit *rabbitmq.Conn

func main() {
	router := gin.Default()
    router.POST("/", callbackHandler)
	var err error
	appProps, err = properties.LoadApplicationProperties()
	if (nil != err) {
		log.Fatal(err)
	}
	rabbit, err = client.CreateConnection(appProps.Rabbit)
	if (nil != err) {
		log.Fatal(err)
	}
    router.Run("localhost:8080")	
}

func callbackHandler(context *gin.Context) {
	if !isAuthenticated(context) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})	
		return
	}
	var dto dto.UpdateTelegram
	if err := context.BindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing dto"})	
		return
	}
	headers := map[string]interface{} {
		"cmd": dto.GetMessage().GetEntity().Type,
	}
	data, err := json.Marshal(dto)
	log.Println(string(data))
	if (nil != err) {
		context.JSON(http.StatusInternalServerError, nil)
		log.Fatal(err)
		return
	}
	err = client.Publish(rabbit, appProps.Rabbit.Exchange, headers, data)
	if (nil != err) {
		context.JSON(http.StatusInternalServerError, nil)
		log.Fatal(err)
		return
	}
	context.JSON(http.StatusOK, dto)
}

func isAuthenticated(context *gin.Context) bool {
	token := context.Request.Header["X-Telegram-Bot-Api-Secret-Token"]
	return nil != token && len(token) > 0 && token[0] == appProps.Token  
}

