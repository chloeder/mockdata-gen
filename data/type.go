package data

const (
	TYPE_NAME      = "name"
	TYPE_BIRTHDATE = "birthdate"
	TYPE_ADDRESS   = "address"
	TYPE_PHONE     = "phone"
)

const (
	SUBTYPE_ADDRESS_STREET = "street"
	SUBTYPE_ADDRESS_CITY   = "city"
)

var Supported = map[string]bool{
	TYPE_NAME:      true,
	TYPE_BIRTHDATE: true,
	TYPE_ADDRESS:   true,
	TYPE_PHONE:     true,
}
