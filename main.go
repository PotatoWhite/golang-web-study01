package main

import (
	"fmt"
	"github.com/potatowhite/web/study01/myapp"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", myapp.NewHttpHandler())
	if err != nil {
		fmt.Printf("An error occurred %s", err)
		return
	}
}
