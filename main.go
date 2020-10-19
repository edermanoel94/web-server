package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

const (
	index = "GET / HTTP/1.1\r\n"
	sleep = "GET /sleep HTTP/1.1\r\n"
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

	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)

	if err != nil {
		fmt.Printf("Error reading request from client: %s \n", err.Error())
		return
	}

	page, statusCode := getPageAndStatusCode(string(buffer))

	response, err := getContentPage(page, statusCode)

	if err != nil {
		fmt.Printf("Error on open page html: %s \n", err.Error())
		return
	}

	_, err = io.WriteString(conn, response)

	if err != nil {
		fmt.Printf("Some shit happens on write to connection: %s \n", err.Error())
		return
	}
}

func getContentPage(page string, statusCode string) (string, error) {

	content, err := ioutil.ReadFile(fmt.Sprintf("./static/%s", page))

	if err != nil {
		return "", err
	}

	response := fmt.Sprintf("HTTP/1.1 %s\r\nContent-Length: %d\r\n\r\n%s",
		statusCode, len(content), content)

	return response, nil
}

func getPageAndStatusCode(header string) (string, string) {

	if strings.HasPrefix(header, index) {
		return "index.html", "200 OK"
	} else if strings.HasPrefix(header, sleep) {
		time.Sleep(2 * time.Second)
		return "sleep.html", "200 OK"
	} else {
		return "404.html", "404 NOT FOUND"
	}
}
