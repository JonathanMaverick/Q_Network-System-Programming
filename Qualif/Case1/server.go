package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go func(c net.Conn){
			fmt.Println("Accepted")

			payload, err := Decode(c)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(payload)
			
			p := Binary("Muda muda muda DIO! " + string(payload.Bytes()) + "Success");
			_, err = p.WriteTo(c)
			if err != nil{
				fmt.Println(err)
				return
			}
		}(conn)
	}

}