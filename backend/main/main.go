package main

func main() {
	DBC.Connect()
	defer DBC.Disconnect()
	SRV.Run()
}
