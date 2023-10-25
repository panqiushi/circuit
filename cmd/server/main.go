package main

import (
	"circuit.io/circuit/internal/router"
)

func main() {
	err := router.Router.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
