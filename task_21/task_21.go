package main

import "fmt"

type Client struct{}

type Computer interface {
	InsertIntoLightningPort()
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer")
	com.InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)
	fmt.Println("")
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine: windowsMachine}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
