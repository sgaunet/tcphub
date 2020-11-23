package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// YamlConfig is exported.
type YamlConfig struct {
	Server string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clientTCP(url string, c chan string, wg *sync.WaitGroup) {
	// connect to this socket
	defer wg.Done()

	for {

		conn, err := net.Dial("tcp", url)
		if err != nil {
			fmt.Println("Error when connecting to server...")
			fmt.Println("Wait 5s")
			time.Sleep(time.Second * 5)
			continue
		}
		for {
			reply := make([]byte, 1024)
			_, err = conn.Read(reply)
			if err != nil {
				fmt.Println("Error !")
				break
			}
			if len(reply) != 0 {
				//fmt.Print("Message from server: " + string(reply))
				c <- string(reply)
			}
		}
		fmt.Println("End of connection !")
	}
}

func handleConnection(socketTab map[int]net.Conn, ln net.Listener, c chan string) {
	for {
		msg := <-c
		for i, conn := range socketTab {
			_, err := conn.Write([]byte(msg))
			if err != nil {
				fmt.Printf("Client disconnection %d\n", i)
				delete(socketTab, i)
				continue
			}
		}
	}
	//ln.Close()
}

func serve(port int, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var connections = make(map[int]net.Conn)

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))

	go handleConnection(connections, ln, c)

	for {
		conn, _ := ln.Accept()
		i := 0
		for {
			if _, ok := connections[i]; !ok {
				connections[i] = conn
				fmt.Printf("Add conn to connections[%d]\n", i)
				break
			}
			i++
		}
	}
}

func main() {
	var server string
	var port int

	flag.StringVar(&server, "s", "", "TCP server <host:port>")
	flag.IntVar(&port, "p", 64000, "Port to stream")
	flag.Parse()

	c := make(chan string, 1000)
	var wg sync.WaitGroup

	go clientTCP(server, c, &wg)
	wg.Add(1)
	go serve(port, c, &wg)
	wg.Add(1)
	wg.Wait()
	os.Exit(0)
}
