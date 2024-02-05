package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAddressInfo(zipcode string) (*Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiAddress APIAddress
	if err := json.NewDecoder(resp.Body).Decode(&apiAddress); err != nil {
		return nil, err
	}

	address := convertFromAPIAddressToAddress(apiAddress, "")

	return address, nil
}
