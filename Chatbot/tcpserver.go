package main

import (
	"fmt"
	"net"
)

func main() {
	
	fmt.Println("starting TCP server on  localhost:5000")

	//listen on tcp port 5000
	listener,err := net.Listen("tcp","localhost:5000")

    if err != nil{
    	fmt.Println("Error Listening",err.Error())
	}

	for{

		conn, err := listener.Accept()

		if err != nil{

			return //terminate program
		}

		go serveconnection(conn)

	}

}

func serveconnection( connectin net.Conn)  {

	for {

		for{

			buff := make([] byte, 512)
			_, err := connectin.Read(buff)

			if err != nil{

				fmt.Println("Error reading buffer ",err.Error())

				return //terminate program
			}

			fmt.Printf("Receive data: %v",string(buff))

		}


	}

}
