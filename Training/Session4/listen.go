package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		fmt.Println(err)
	}
	defer cancel();

	client, err := net.ListenPacket("udp","127.0.0.1:")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {_ = client.Close()}()

	interloper, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		fmt.Println(err)
	}

	interrupt := []byte("interrupt")
	n, err := interloper.WriteTo(interrupt, client.LocalAddr())
	if err != nil {
		fmt.Println(err)
	}
	if len(interrupt) != n{
		fmt.Printf("wrote %d bytes, expected %d\n", n, len(interrupt))
	}

	ping := []byte("ping")
	_, err = client.WriteTo(ping, serverAddr)
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("First Msg: ", string(buf[:n]))

	if !bytes.Equal(interrupt, buf[:n]){
		fmt.Printf("expected %q, actual reply %q\n", interrupt, buf[:n])
	}

	if addr.String() != interloper.LocalAddr().String(){
		fmt.Printf("expected %q, got %q\n", addr.String(), interloper.LocalAddr().String())
	}

	n, addr, err = client.ReadFrom(buf)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Second Msg: ", string(buf[:n]))

	if !bytes.Equal(ping, buf[:n]){
		fmt.Printf("expected %q, actual reply %q\n", ping, buf[:n])
	}

	if addr.String() != serverAddr.String(){
		fmt.Printf("expected %q, got %q\n", addr.String(), serverAddr.String())
	}

}