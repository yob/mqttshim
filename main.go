package mqttshim

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/256dpi/gomqtt/packet"
)

type MqttShim struct { }

func NewMqttShim() *MqttShim {
	return &MqttShim{}
}

func (shim *MqttShim) StartServer(bindAddress string, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bindAddress, port))
	if err != nil {
		panic(err)
	}
	log.Print(fmt.Sprintf("mqttshim started on %s:%d", bindAddress, port))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error on accept: ", err)
			continue
		}
		go shim.handleConnection(conn)
	}
}

func (shim *MqttShim) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := packet.NewDecoder(conn)
	encoder := packet.NewEncoder(conn, 1 * time.Second)

	for {
		msg, err := decoder.Read()
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		switch msg.Type() {
		case packet.CONNECT:
			fmt.Printf("msg: %v\n", msg)
			connack := packet.NewConnack()
			connack.SessionPresent = false
			connack.ReturnCode = 0
			encoder.Write(connack, false)
		case packet.DISCONNECT:
			fmt.Printf("msg: %v\n", msg)
			return
		default:
			fmt.Printf("msg: %v\n", msg)
		}
	}
	fmt.Printf("closing connection\n")
}
