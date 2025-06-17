package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"

	"cmd/main.go/internal/model"
)

type Response struct {
	ID     int32    `json:"id"`
	Method string   `json:"method"`
	Params Params   `json:"params"`
	Result []string `json:"result"`
	Error  Error    `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    int32  `json:"code"`
}

type Params struct {
	Power      string `json:"power"`
	Hue        int    `json:"hue"`
	Saturation int    `json:"sat"`
	RGB        int    `json:"rgb"`
	Brightness int    `json:"bright"`
}

type Command struct {
	ID     int32       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type YeeLight struct {
}

func NewYeeLight() *YeeLight {
	return &YeeLight{}
}

func (y *YeeLight) Toggle(loc string) (Response, error) {
	cmd := Command{
		Method: "toggle",
		Params: []interface{}{},
	}

	return request(loc, cmd)
}

func (y *YeeLight) SetBrightness(loc string, brightness, duration int) (Response, error) {
	cmd := Command{
		Method: "set_bright",
		Params: []interface{}{brightness, setEffect(duration), duration},
	}

	return request(loc, cmd)
}

func (y *YeeLight) SetRGB(loc string, red, green, blue, duration int) (Response, error) {
	rgb := (red << 16) + (green << 8) + blue

	cmd := Command{
		Method: "set_rgb",
		Params: []interface{}{rgb, "smooth", duration},
	}

	return request(loc, cmd)
}

func (y *YeeLight) PowerOff(loc string, duration int) (Response, error) {
	cmd := Command{
		Method: "set_power",
		Params: []interface{}{model.Off, setEffect(duration), duration},
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
		return Response{}, fmt.Errorf("could not connect to %s", loc)
	}
	defer conn.Close()

	time.Sleep(100 * time.Millisecond)

	cmdJSON, err := json.Marshal(cmd)
	if err != nil {
		return Response{}, err
	}
	if _, err = fmt.Fprintf(conn, "%s\r\n", cmdJSON); err != nil {
		return Response{}, fmt.Errorf("could not send command to %s", loc)
	}

	respStr, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return Response{}, fmt.Errorf("could not read from %s", loc)
	}

	resp := Response{}
	err = json.Unmarshal([]byte(respStr), &resp)

	return resp, nil
}

func setEffect(duration int) string {
	effect := "sudden"
	if duration > 0 {
		effect = "smooth"
	}

	return effect
}
