package ylight

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Response struct {
	Method string `json:"method"`
	Params Params `json:"params"`
}

type Params struct {
	Power string `json:"power"`
}

type Command struct {
	ID     int32       `json:"id"`
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
	cmd := Command{
		Method: "toggle",
		Params: []interface{}{},
	}

	return request(loc, cmd)
}

func (y *YLight) SetBrightness(loc string, brightness, duration int) (Response, error) {
	effect := "sudden"

	if duration > 0 {
		effect = "smooth"
	}

	cmd := Command{
		Method: "set_bright",
		Params: []interface{}{brightness, effect, duration},
	}

	return request(loc, cmd)
}

func request(loc string, cmd Command) (Response, error) {
	if cmd.ID == 0 {
		r := rand.NewSource(time.Now().UnixNano())
		cmd.ID = rand.New(r).Int31()
	}

	conn, err := net.Dial("tcp", loc)
	if err != nil {
		return Response{}, err
	}
	defer conn.Close()

	time.Sleep(300 * time.Millisecond)

	cmdJSON, err := json.Marshal(cmd)
	if err != nil {
		return Response{}, err
	}
	if _, err = fmt.Fprintf(conn, "%s\r\n", cmdJSON); err != nil {
		return Response{}, err
	}

	respStr, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return Response{}, err
	}

	resp := Response{}
	err = json.Unmarshal([]byte(respStr), &resp)

	return resp, nil
}
