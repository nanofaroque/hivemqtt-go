package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)
type Sensor struct {
	Name string
	Body string
	Time int64
}
func main() {
	done := make(chan bool)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("third")
	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := c.Subscribe("$share/group/test/topic", 1, msgRcvdSubs1); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	<-done
}

func msgRcvdSubs1(client mqtt.Client, message mqtt.Message) {
	//fmt.Print(client)
	var m Sensor
	err := json.Unmarshal(message.Payload(), &m)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

