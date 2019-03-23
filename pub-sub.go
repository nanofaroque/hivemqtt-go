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
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("first")
	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

/*	if token := c.Subscribe("$share/group/test/topic", 1, msgRcvd); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}*/
	for i:=0;i<1000 ;i++  {
		m := SensorData{"Alice", "Hello", time.Now().UnixNano()}
		b, err := json.Marshal(m)
		if err!=nil{
			panic(err)
		}
		fmt.Println(b)
		if token := c.Publish("test/topic", 2, true, b); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
		time.Sleep(100 * time.Microsecond)
	}
	<-done
}
/*
func msgRcvd(client mqtt.Client, message mqtt.Message) {
	//fmt.Print(client)
	var m SensorData
	err := json.Unmarshal(message.Payload(), &m)
	if err!=nil{fmt.Print(err)}
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}*/