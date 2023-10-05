package main

import (
	"code.circuit.io/circuit/internal/router"
)

func main() {
	err := router.Router.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
