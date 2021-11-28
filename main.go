package main

import "Backend/router"

func main() {
	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
