package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

const (
	get   = "GET / HTTP/1.1"
	sleep = "GET /sleep HTTP/1.1"
)

func main() {

	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	fmt.Println("Starting server on port 8080....")

	for {

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Some shit on connection is happening...")
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

	defer conn.Close()

	buffer := bufio.NewScanner(conn)

	var data string

	for buffer.Scan() {

		data += buffer.Text()
	}

	page, statusCode := getPageAndStatusCode(data)

	content, err := ioutil.ReadFile(fmt.Sprintf("./static/%s", page))

	if err != nil {
		log.Printf("Error on open page html: %s \n", err.Error())
		return
	}

	response := fmt.Sprintf("HTTP/1.1 %s\r\nContent-Length: %d\r\n\r\n%s",
		statusCode, len(content), content)

	_, err = fmt.Fprintf(conn, response)

	if err != nil {
		fmt.Printf("Some shit happens on write to connection: %s \n", err.Error())
		return
	}
}

func getPageAndStatusCode(header string) (string, string) {

	if strings.HasPrefix(header, get) {
		return "index.html", "200 OK"
	} else if strings.HasPrefix(header, sleep) {
		time.Sleep(2 * time.Second)
		return "sleep.html", "200 OK"
	} else {
		return "404.html", "404 NOT FOUND"
	}
}
