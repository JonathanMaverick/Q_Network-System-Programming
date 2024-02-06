package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ViewProducts(){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1212/products/view-products", nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer func() {resp.Body.Close()}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(body))
}

func AddProduct(){
	var productName string
	var productPrice float64

	reader := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("Nombre del producto: ")
		productName, _ = reader.ReadString('\n')
		productName = strings.TrimSpace(productName)

		if productName == "" {
			fmt.Println("Nombre del producto no puede estar vacío")
			return
		}else{
			break
		}
	}

	for{
		fmt.Print("Precio del producto: ")
		_, err := fmt.Scanln(&productPrice)
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		if productPrice <= 0 {
			fmt.Println("Precio del producto no puede ser menor o igual a 0")
		}else{
			break
		}
	}

	reqBody := new(bytes.Buffer)
	w := multipart.NewWriter(reqBody)
	
	nameField, err := w.CreateFormField("name")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	_, err = nameField.Write([]byte(productName))
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	priceField, err := w.CreateFormField("price")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	_, err = priceField.Write([]byte(fmt.Sprintf("%f", productPrice)))
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	file, err := os.Open("./files/productos.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileField, err := w.CreateFormFile("file", file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:1212/products/add-product", reqBody)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer func() { resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func UpdateProduct(){
	var productID int
	var productPrice float64

	for{
		fmt.Print("ID del producto: ")
		_, err := fmt.Scanln(&productID)
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		if productID <= 0 {
			fmt.Println("ID del producto no puede ser menor o igual a 0")
		}else{
			break
		}
	}

	for{
		fmt.Print("Precio del producto: ")
		_, err := fmt.Scanln(&productPrice)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		
		if productPrice <= 0 {
			fmt.Println("Precio del producto no puede ser menor o igual a 0")
		}else{
			break
		}
	}

	jsonStr := fmt.Sprintf(`{"id": %d, "price": %f}`, productID, productPrice)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, "http://localhost:1212/products/update-product/" + strconv.Itoa(int(productID)), strings.NewReader(jsonStr))
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println("Error", err)
		return
	}

	defer func() { resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(body))
}

func main(){
	for{
		fmt.Println("Hola VKlient!")
		fmt.Println("1. Ver productos")
		fmt.Println("2. Agregar productos")
		fmt.Println("3. Actualizar productos")
		fmt.Println("4. Salir")

		var choice int
		fmt.Print(">> ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error", err)
		}

		switch choice {
		case 1:
			ViewProducts()
		case 2:
			AddProduct()
		case 3:
			UpdateProduct()
		case 4:
			fmt.Println("Adios")
			return
		default:
			fmt.Println("Opción no válida")
		}

	}
}