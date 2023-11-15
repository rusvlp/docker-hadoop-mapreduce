package main

import (
	"TCPServer/internal/entity"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	//	fmt.Print("Hello!")

	err, skinGen := entity.NewGenerator()

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	args := os.Args

	var timeout int
	if len(args) < 4 {
		timeout = 5
	} else {
		timeout, err = strconv.Atoi(args[3])
	}

	if err != nil {
		panic(err)
	}

	err, conn := openConn()

	defer conn.Close()

	fmt.Println("Starting generating JSONs with interval " + strconv.Itoa(timeout) + " seconds")
	for {

		err, skin := skinGen.GenerateSkinPrice()

		if err != nil {
			panic(err)
		}

		json, err := json.MarshalIndent(skin, "", "\t")

		_, err = conn.Write(json)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(int64(timeout)) * time.Second)
	}
}

func openConn() (error, net.Conn) {
	args := os.Args

	var host, port string

	if len(args) < 3 {
		port = "9999"
		host = "localhost"
	} else {
		port = args[2]
		host = args[1]
	}

	fmt.Println("TCP Client configuration, host: " + host + ", port: " + port)

	conn, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		return err, nil
	}

	return nil, conn
}
