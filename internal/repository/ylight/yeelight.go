package ylight

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

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Command struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type YLight struct {
	Location string
}

func NewYLight() *YLight {
	return &YLight{}
}

func (y *YLight) Toggle(loc string) (Response, error) {
	y.Location = loc

	cmd := Command{
		ID:     1,
		Method: "toggle",
		Params: []string{},
	}

	return y.request(cmd)
}

func (y *YLight) SetBrightness(loc string, brightness, duration int) (Response, error) {
	y.Location = loc

	var effect string

	if duration > 0 {
		effect = "smooth"
	} else {
		effect = "sudden"
		duration = 0
	}

	cmd := Command{
		ID:     5,
		Method: "set_bright",
		Params: []interface{}{brightness, effect, duration},
	}

	return y.request(cmd)
}

func (y *YLight) request(cmd Command) (Response, error) {
	conn, err := net.DialTimeout("tcp", y.Location, timeout)
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
