package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:5678")

	if err != nil{
		return;
	}

	defer dial.Close()

	// payload := "Alo!"
	// _, err = dial.Write([]byte(payload))
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }

	payload := Binary("Alo! boku no namaewa, Jotaro Kujo!")
	_, err = payload.WriteTo(dial)

	if err != nil{
		fmt.Println(err)
		return
	}

	p, err := Decode(dial)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(p)

}