package main

import (
	"gincheese/apis/routes"
)

func main() {
	routes.RouteEngine().Run(":8080")
}
