package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/api/forms"
	"goapi/caching"
	"goapi/messagebroker"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RobotHandler struct {
	mqttConnected  *messagebroker.MqttConnection
	redisConnected *caching.RedisConnection
}

func NewRobotHandler(mqttConnected *messagebroker.MqttConnection, redisConnected *caching.RedisConnection) RobotHandler {
	return RobotHandler{
		mqttConnected:  mqttConnected,
		redisConnected: redisConnected,
	}
}

// Handshake
func (h RobotHandler) Handshake(c *gin.Context) {
	var p forms.ControlReceiver

	// Bind request body - JSON
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("robot username = ", p.Username)
	// fmt.Println("robot message = ", p.Message)

	// Caching ; Get
	mac_addr, redisErr := h.redisConnected.GetRedis(p.Username)
	if redisErr != nil {
		// Handle error
		log.Println("[key not found] ", redisErr)
		c.JSON(http.StatusNotFound, gin.H{"error": "please connect first"})
		return
	}

	// Topic - handshake
	// JSON format ; Example
	// { "username": "dolly", "message": "Ahoy!" }
	pub_topic := fmt.Sprintf("/mR_robot/%s", mac_addr)
	// fmt.Println("robot pub_topic = ", pub_topic)

	pub_payload := map[string]string{
		"username": p.Username,
		"message":  "Ahoy!",
	}
	// fmt.Println("robot pub_payload = ", pub_payload)

	// Convert map to json
	jsonString, marshalErr := json.Marshal(pub_payload)
	if marshalErr != nil {
		log.Println("JSON Error : ", marshalErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not send message to robot"})
		return
	}

	// MQTT Publish
	h.mqttConnected.Publish(pub_topic, jsonString)

	// MQTT Disconnect
	// h.mqttConnected.Disconnect()
}

// Control robot
func (h RobotHandler) Movement(c *gin.Context) {
	var p forms.ControlReceiver

	// Bind request body - JSON
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("robot username = ", p.Username)
	// fmt.Println("robot message = ", p.Message)

	// Caching ; Get
	mac_addr, redisErr := h.redisConnected.GetRedis(p.Username)
	if redisErr != nil {
		// Handle error
		log.Println("[key not found] ", redisErr)
		c.JSON(http.StatusNotFound, gin.H{"error": "please connect first"})
		return
	}

	// Topic - movement
	// JSON format ; Example
	// { "username": "dolly", "message": "forward" }
	pub_topic := fmt.Sprintf("/mR_robot/%s", mac_addr)
	// fmt.Println("robot pub_topic = ", pub_topic)

	pub_payload := map[string]string{
		"username": strings.ToLower(p.Username),
		"message":  strings.ToLower(p.Message),
	}
	// fmt.Println("robot pub_payload = ", pub_payload)

	// Convert map to json
	jsonString, marshalErr := json.Marshal(pub_payload)
	if marshalErr != nil {
		log.Println("JSON Error : ", marshalErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can not send message to robot"})
		return
	}

	// MQTT Publish
	h.mqttConnected.Publish(pub_topic, jsonString)
}
