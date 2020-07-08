package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("Starting TCP chat bot client")
	//open connection
	conn, err := net.Dial("tcp", "localhost:5000")

	if err != nil {
		fmt.Println("Error connecting to server", err.Error())
		return //terminate the program
	}

	//read input from ther user and send to server

	inputreader := bufio.NewReader(os.Stdin)

	fmt.Println("Wecome to chat bot ,please enter your name")

	ClientName, _ := inputreader.ReadString('\n')

	TrimmedClient := strings.Trim(ClientName, "\r\n")

	for {
		fmt.Println("please enter your message to the server ? type Q to quit")

		intput, _ := inputreader.ReadString('\n')

		trimmedinput := strings.Trim(intput, "\n")

		if trimmedinput == "Q" {
			//close the clinet
			return
		}

		_, err = conn.Write([]byte(TrimmedClient + " says " + trimmedinput + "\n"))

	}

}
