package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

/*func main()  {
	done := make(chan bool)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("sample")
	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	payload:="Hello World"
	for i:=0;i<10 ;i++  {
		if token := c.Publish("test/topic", 0, true, payload+strconv.Itoa(i)); token.Wait() && token.Error() != nil {
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
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}*/

type SensorData struct {
	temp  string `json:"sensor"`
	pmcId string `json:"pmc_id"`
	id    int32  `json:"id"`
}

func main() {
	done := make(chan bool)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("sample")
	opts.SetUsername("admin")
	opts.SetPassword("hivemq")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	payload := SensorData{
		"55",
		"starwood",
		1,
	}
	for i := 0; i < 10; i++ {
		fmt.Println(payload)
		bytes, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		if token := c.Publish("json/topic", 0, true, bytes); token.Wait() && token.Error() != nil {
			fmt.Println("Error: ",token.Error())
		}
		if token := c.Subscribe("json/topic", 0, jsonRcvd); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
	<-done
}

func jsonRcvd(client mqtt.Client, message mqtt.Message) {
	fmt.Print(client)
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())

}