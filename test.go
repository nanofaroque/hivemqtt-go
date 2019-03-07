package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

func main()  {
	done := make(chan bool)

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("sample")

	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	if token := c.Publish("test/topic", 0, true, "Example Payload for Bryan"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}


	if token := c.Subscribe("test/topic", 0, msgRcvd); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	<-done
}

func msgRcvd(client mqtt.Client, message mqtt.Message) {
	fmt.Print(client)
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}