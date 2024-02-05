package user

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAddressInfo(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mockResponse := `{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP",
			"ibge": "3550308",
			"gia": "1004",
			"ddd": "11",
			"siafi": "7107"
		}`

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))

	defer mockServer.Close()

	tests := []struct {
		name            string
		zipcode         string
		expectedAddress *APIAddress
	}{
		{
			name:    "Valid_ZipCode",
			zipcode: "01001000",
			expectedAddress: &APIAddress{
				Logradouro: "Praça da Sé",
				Complement: "lado ímpar",
				Bairro:     "Sé",
				Localidade: "São Paulo",
				UF:         "SP",
				CEP:        "01001-000",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := GetAddressInfo(tt.zipcode)
			if err != nil {
				t.Errorf("GetAddressInfo returned an error: %v", err)
			}

			if !reflect.DeepEqual(address, tt.expectedAddress) {
				t.Errorf("GetAddressInfo returned unexpected address: got %v, want %v", address, tt.expectedAddress)
			}
		})
	}
}
