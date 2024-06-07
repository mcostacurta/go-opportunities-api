package main

import "github.com/mcostacurta/go-opportunities-api/router"

func main() {
	router.Initialize() // listen and serve on 0.0.0.0:8080
}
