package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const (
	timeout = 1 * time.Second
)

type Response struct {
	ID     int         `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error  Error       `json:"error,omitempty"`
}

// Error struct is used on the ResponseError payload
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Command struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

func main() {
	cmd := Command{
		ID:     7,
		Method: "toggle",
		Params: []string{},
	}

	resp, err := request(cmd, "192.168.0.15:55443")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func request(cmd Command, location string) (Response, error) {
	conn, err := net.DialTimeout("tcp", location, timeout)
	if err != nil {
		return Response{}, fmt.Errorf("failed to connect to light: %w", err)
	}
	defer conn.Close()

	cmdJSON, err := json.Marshal(cmd)
	if err != nil {
		return Response{}, fmt.Errorf("failed to marshal command: %w", err)
	}

	time.Sleep(500 * time.Millisecond)

	_, err = fmt.Fprintf(conn, "%s\r\n", cmdJSON)
	if err != nil {
		return Response{}, fmt.Errorf("failed to send command: %w", err)
	}

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return Response{}, fmt.Errorf("failed to read response: %w", err)
	}

	//parse response
	resp := Response{}
	err = json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}
