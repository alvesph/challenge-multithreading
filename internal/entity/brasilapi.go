package entity

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alvesph/challenge-multithreading.git/internal/utils"
)

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func ReqBrasilapi(ctx context.Context, cep string) []byte {
	var data BrasilApi
	err := utils.FetchAndUnmarshal(ctx, "https://brasilapi.com.br/api/cep/v1/"+cep, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in search this cep: %v\n", err)
		return nil
	}

	jsonReult, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in marshal this cep: %v\n", err)
		return nil
	}
	return jsonReult
}
