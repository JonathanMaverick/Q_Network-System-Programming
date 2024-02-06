package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:6666")

	if err != nil{
		return;
	}

	defer dial.Close()

	payload := Binary("Alo! boku no namaewa, Jotaro Kujo!")
	_, err = payload.WriteTo(dial)

	if err != nil{
		fmt.Println(err)
		return
	}

	dial.SetReadDeadline(time.Now().Add(5 * time.Second))

	p, err := Decode(dial)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(p)

}