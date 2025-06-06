package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	resp, err := brightnessEdgardBulb("192.168.0.42:55443", 30, 5000)
	if err != nil {
		log.Fatal("fatal", err)
	}

	fmt.Println(resp)
	return
}

type Commandd struct {
	ID     int32       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

func brightnessEdgardBulb(location string, brightness, duration int) (string, error) {
	effect := "sudden"

	if duration > 0 {
		effect = "smooth"
	}

	cmd := Commandd{
		Method: "set_bright",
		Params: []interface{}{brightness, effect, duration},
	}

	if cmd.ID == 0 {
		r := rand.NewSource(time.Now().UnixNano())
		cmd.ID = rand.New(r).Int31()
	}

	conn, err := net.Dial("tcp", location)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	time.Sleep(300 * time.Millisecond)

	cmdJSON, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}
	if _, err = fmt.Fprintf(conn, "%s\r\n", cmdJSON); err != nil {
		return "", err
	}

	return bufio.NewReader(conn).ReadString('\n')
}
