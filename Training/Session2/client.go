package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:1234")

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

	payload := []byte("Alo!")
	_, err = dial.Write(payload)

	if err != nil{
		fmt.Println(err)
		return
	}

}