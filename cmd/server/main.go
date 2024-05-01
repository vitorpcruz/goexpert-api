package main

import (
	"fmt"

	"github.com/vitorpcruz/goexpert/9-APIS/configs"
)

func main() {
	config, _ := configs.LoadConfig("./")
	fmt.Println(config)
}
