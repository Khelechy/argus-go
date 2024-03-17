package argus

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Argus struct {
	Username string
	Password string
	Host     string
	Port     string
	Events   chan Event
	Messages chan string
	Errors   chan error
}

type Event struct {
	Action            string
	ActionDescription string
	Name              string
	Timestamp         time.Time
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

	argus.Events = make(chan Event)
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

				isJson, event, str := isJsonString(data)

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

func isJsonString(str string) (bool, Event, string) {
	var event Event
	if json.Unmarshal([]byte(str), &event) == nil {
		return true, event, str
	}

	return false, event, str
}
