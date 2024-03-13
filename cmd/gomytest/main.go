package main

import (
	"fmt"
	"os"

	"github.com/prasadjivane/gomytest/pkg/gomytest"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: gomytest <method> <url> [data]")
		os.Exit(1)
	}

	method := os.Args[1]
	url := os.Args[2]
	var data []byte
	if len(os.Args) > 3 {
		data = []byte(os.Args[3])
	}

	switch method {
	case "GET":
		err := gomytest.TestGET(url)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "POST":
		err := gomytest.TestPOST(url, data)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "PUT":
		err := gomytest.TestPUT(url, data)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "DELETE":
		err := gomytest.TestDELETE(url)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid method. Available methods: GET, POST, PUT, DELETE")
		os.Exit(1)
	}
}
