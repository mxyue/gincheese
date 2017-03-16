package config

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	loadYml()
	fmt.Println("settings: ", Settings)
	for k, v := range Settings {
		fmt.Fprintf(gin.DefaultWriter, "key:  %s ,value: %s \r\n", k, v)

	}
}
