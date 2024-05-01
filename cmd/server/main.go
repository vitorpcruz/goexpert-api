package main

import (
	"fmt"

	"github.com/vitorpcruz/goexper/9-APIS/configs"
)

func main() {
	config, _ := configs.LoadConfig("./")
	fmt.Println(config)
}
