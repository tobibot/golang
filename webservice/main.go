package main

import (
	"fmt"
	"net/http"

	"go.heying.info/learning/webservice/controllers"
)

const (
	port = "3000"
)

func main() {
	sport := fmt.Sprintf(":%v", port)

	println("listening on port", port)

	controllers.RegisterControllers()
	http.ListenAndServe(sport, nil)
}
