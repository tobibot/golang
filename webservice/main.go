package main

import (
	"net/http"

	"go.heying.info/learning/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
