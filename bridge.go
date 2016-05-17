package main

import (
	"github.com/tarm/serial"
	"log"
	"net"
)

func main() {
	config := &serial.Config{Name: "COM6", Baud: 115200}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal("Error opening COM Port")
	}
	defer port.Close()
	
	tcpCon, err := net.Dial("tcp", "127.0.0.1:5511")
	if err != nil {
		log.Fatal("Error opening TCP Port")
	}
	defer tcpCon.Close()

	tcpToCom := make([]byte, 128)
	comToTcp := make([]byte, 128)

	go func() {
		for {
			nRx, err := tcpCon.Read(tcpToCom)
			if err != nil {
				log.Fatal("Error Receiving from TCP")
			}
			nTx, err := port.Write(tcpToCom[:nRx])
			if err != nil {
				log.Fatal("Error Sending to COM")
			}
			log.Println("TCP -> UART: ", nRx, nTx)
		}
	}()

	for {
		nRx, err := port.Read(comToTcp)
		if err != nil {
			log.Fatal("Error Receiving from COM")
		}
		nTx, err := tcpCon.Write(comToTcp[:nRx])
		if err != nil {
			log.Fatal("Error Sending to TCP")
		}
		log.Println("UART -> TCP: ", nRx, nTx)
	}
}
