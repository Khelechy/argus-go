package argus

import (
	"fmt"
	"net"

	"github.com/khelechy/argus/models"
	"github.com/khelechy/argus/utils"
)

type Argus struct {
	Username string
	Password string
	Host     string
	Port     string
	Events   chan models.Event
	Messages chan string
	Errors   chan error
}


func Connect(argusConfig *Argus) (*Argus, error) {

	argus := &Argus{}

	connectionString := fmt.Sprintf("<ArgusAuth>%s:%s</ArgusAuth>", argusConfig.Username, argusConfig.Password)

	//Resolve host and port
	if len(argusConfig.Host) == 0 {
		argusConfig.Host = "localhost"
	}

	if len(argusConfig.Port) == 0 {
		argusConfig.Port = "1337"
	}

	addr := fmt.Sprintf("%s:%s", argusConfig.Host, argusConfig.Port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	sendAuthData(conn, connectionString)

	argus.Events = make(chan models.Event)
	argus.Errors = make(chan error)
	argus.Messages = make(chan string)

	go func(conn net.Conn) {

		defer conn.Close()

		// Create a buffer to read data into
		buffer := make([]byte, 1024)

		for {
			// Read data from the client
			n, err := conn.Read(buffer)
			if err != nil {
				argus.Errors <- err
			}

			data := string(buffer[:n])

			if len(data) > 0 {

				isJson, event, str := utils.IsJsonString(data)

				if isJson {
					// Push event to event channel
					argus.Events <- event
				} else {

					argus.Messages <- fmt.Sprintf("Received: %s\n", str)
				}
			}

		}
	}(conn)

	return argus, nil
}

func sendAuthData(conn net.Conn, connectionString string) {
	data := []byte(connectionString)
	_, _ = conn.Write(data)
}
