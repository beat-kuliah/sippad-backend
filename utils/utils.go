package utils

var Currencies = map[string]string{
	"USD": "USD",
	"NGN": "NGN",
	"IDR": "IDR",
}

func IsValidCurrency(currency string) bool {
	if _, ok := Currencies[currency]; ok {
		return true
	}
	return false
}
