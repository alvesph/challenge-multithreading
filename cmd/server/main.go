package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alvesph/challenge-multithreading.git/configs"
	"github.com/alvesph/challenge-multithreading.git/internal/entity"
)

type ApiResponse struct {
	Source string
	Data   []byte
}

func main() {
	config, _ := configs.LoadConfig(".")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	responseCh := make(chan ApiResponse, 2)

	go func() {
		result := entity.ReqViaCep(ctx, config.CEP)
		if result != nil {
			responseCh <- ApiResponse{Source: "ViaCep", Data: result}
		}
	}()

	go func() {
		result := entity.ReqBrasilapi(ctx, config.CEP)
		if result != nil {
			responseCh <- ApiResponse{Source: "BrasilApi", Data: result}
		}
	}()

	select {
	case response := <-responseCh:
		cancel()
		fmt.Printf("Response from %s:\n%s\n", response.Source, string(response.Data))
	case <-ctx.Done():
		fmt.Println("Request timed out or was cancelled.")
	}
}
