package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
)

type SensorData struct {
	Name string
	Body string
	Time int64
}
func main()  {
	done := make(chan bool)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("sample")
	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for i:=0;i<100 ;i++  {
		m := SensorData{"Alice", "Hello", time.Now().UnixNano()}
		b, err := json.Marshal(m)
		if err!=nil{
			panic(err)
		}
		fmt.Println(b)
		if token := c.Publish("test/topic", 0, true, b); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
		if token := c.Subscribe("test/topic", 0, msgRcvd); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
	<-done
}

func msgRcvd(client mqtt.Client, message mqtt.Message) {
	fmt.Print(client)
	var m SensorData
	err := json.Unmarshal(message.Payload(), &m)
	if err!=nil{fmt.Print(err)}
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}