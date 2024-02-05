package user

type User struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	BirthDate      string  `json:"birth_date"`
	Phone          string  `json:"phone"`
	DocumentNumber string  `json:"document_number"`
	Address        Address `json:"address"`
}

type Address struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	ZipCode    string `json:"zip_code"`
}

type APIAddress struct {
	Logradouro string `json:"logradouro"`
	Complement string `json:"complemento"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Bairro     string `json:"bairro"`
	CEP        string `json:"cep"`
}

func convertFromAPIAddressToAddress(apiAddress *APIAddress, number string) *Address {
	address := &Address{
		Street:     apiAddress.Logradouro,
		Complement: apiAddress.Complement,
		Number:     number,
		City:       apiAddress.Localidade,
		State:      apiAddress.UF,
		Country:    "Brasil",
		ZipCode:    apiAddress.CEP,
	}

	if number == "" {
		address.Number = "S/N"
	}

	return address
}

func convertFromAddressToApiAddress(address Address) *APIAddress {
	apiAddress := &APIAddress{
		Logradouro: address.Street,
		Complement: address.Complement,
		Localidade: address.City,
		UF:         address.State,
		CEP:        address.ZipCode,
	}

	return apiAddress
}
