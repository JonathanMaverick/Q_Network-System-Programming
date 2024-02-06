package main

import (
	"fmt"
	"net"
	"net/http"
)

func validateMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func viewProductsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "televisión, computadora, celular")
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Producto agregado")
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/products/update-product/"):]
	_, _ = fmt.Fprintf(w, "El producto con ID: %s ha sido actualizado", id)
}

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/products/view-products", validateMethod(http.MethodGet, viewProductsHandler))
	serveMux.HandleFunc("/products/add-product", validateMethod(http.MethodPost, addProductHandler))
	serveMux.HandleFunc("/products/update-product/", validateMethod(http.MethodPut, updateProductHandler))

	httpServer := &http.Server{
		Addr:    "localhost:1212",
		Handler: serveMux,
	}

	listener, err := net.Listen("tcp", httpServer.Addr)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	err = httpServer.Serve(listener)

	if err != http.ErrServerClosed {
		fmt.Println(err)
		return
	}
}
