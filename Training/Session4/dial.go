package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		return
	}
	defer cancel()

	client, err := net.Dial("udp", serverAddr.String())
	if err != nil {
		return
	}
	defer func() { _ = client.Close() }()

	interloper, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		return
	}

	interrupt := []byte("interrupt")
	n, err := interloper.WriteTo(interrupt, client.LocalAddr())
	if err != nil {
		return
	}

	_ = interloper.Close()

	if len(interrupt) != n{
		fmt.Printf("wrote %d bytes, expected %d\n", n, len(interrupt))
	}

	ping := []byte("ping")
	_, err = client.Write(ping)
	if err != nil {
		return
	}

	buf := make([]byte, 1024)
	n, err = client.Read(buf)

	fmt.Println(string(buf[:n]), "kakoi")

	if err != nil{
		fmt.Println(err)
	}

	if !bytes.Equal(ping, buf[:n]){
		fmt.Printf("expected %q, got %q\n", ping, buf[:n])
	}

	err = client.SetDeadline(time.Now().Add(time.Second))
	if err != nil{
		fmt.Println(err)
	}

	n, err = client.Read(buf)
	fmt.Println(buf[:n], err)

	if err == nil{
		fmt.Println("expected error")
	}

}