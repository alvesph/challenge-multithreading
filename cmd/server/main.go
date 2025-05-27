package main

import (
	"fmt"

	"github.com/alvesph/challenge-multithreading.git/configs"
	"github.com/alvesph/challenge-multithreading.git/internal/entity"
)

func main() {
	config, _ := configs.LoadConfig(".")
	result := entity.ReqViaCep(config.CEP)
	resultBrasilApi := entity.ReqBrasilapi(config.CEP)
	fmt.Println(string(resultBrasilApi))
	fmt.Println(string(result))
}
