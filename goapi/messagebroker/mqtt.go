package messagebroker

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConnection struct {
	mqttClient mqtt.Client
}

func NewConnectionMQTT() (conn *MqttConnection) {
	var broker = "192.168.1.22"
	var port = 1888

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("bot")
	opts.SetPassword("P@ssw0rd")
	opts.AutoReconnect = true
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	conn = &MqttConnection{client}
	return conn
}

func (conn *MqttConnection) IsConnected() bool {
	connected := conn.mqttClient.IsConnected()
	if !connected {
		log.Println("Healthcheck MQTT fails")
	}
	return connected
}

// Subscriber
func (conn *MqttConnection) Subscribe(topic string) {
	token := conn.mqttClient.Subscribe(topic, 1, onMessageReceived())
	token.Wait()
	log.Println("Subscribed to topic: ", topic)
}

// Callback ; Take action when receive payload from robot
func onMessageReceived() func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

		// event := string(msg.Payload())
		// fmt.Println("event = ", event)
	}
}

// Publisher
func (conn *MqttConnection) Publish(topic string, payload interface{}) {
	token := conn.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()
	log.Println("Publish to topic: ", topic)
}

// Disconnect
func (conn *MqttConnection) Disconnect(quiesce ...uint) {
	// the specified number of milliseconds to
	// wait for existing work to be completed.
	var default_quiesce uint = 250
	if len(quiesce) > 0 {
		default_quiesce = quiesce[0]
	}
	conn.mqttClient.Disconnect(default_quiesce)
	log.Println("MQTT disconnected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt connected")
}
