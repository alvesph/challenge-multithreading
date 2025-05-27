package entity

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alvesph/challenge-multithreading.git/internal/utils"
)

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func ReqViaCep(cep string) []byte {
	var data Viacep
	err := utils.FetchAndUnmarshal("https://viacep.com.br/ws/"+cep+"/json/", &data)
	if err != nil {
		return nil
	}
	jsonResult, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in marshal this cep: %v\n", err)
		return nil
	}
	return jsonResult
}
