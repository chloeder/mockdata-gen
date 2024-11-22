package data

import (
	"fmt"
	"math/rand"
)

func Generator(dataType string) string {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_BIRTHDATE:
		return generateBirthdate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}

	return ""
}
func generateName() string {
	nameLen := len(name)
	return name[rand.Intn(nameLen)]
}
func generateBirthdate() string {
	year := rand.Intn(50) + 1950
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1

	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func generateAddress() string {
	streetLen := len(address[SUBTYPE_ADDRESS_STREET])
	cityLen := len(address[SUBTYPE_ADDRESS_CITY])

	streetIndex := rand.Intn(streetLen)
	cityIndex := rand.Intn(cityLen)
	randNumber := rand.Intn(100)

	return fmt.Sprintf("%s No. %d, %s", address[SUBTYPE_ADDRESS_STREET][streetIndex], randNumber, address[SUBTYPE_ADDRESS_CITY][cityIndex])
}
func generatePhone() string {
	return fmt.Sprintf("08%d", rand.Intn(1000000000))
}
