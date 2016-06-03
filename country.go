package ipisp

import "github.com/johngb/langreg"

//Country encapsulates an ISO_3166-1 country and name
type Country struct {
	Code string
	Name string
}

//NewCountryFromCode returns a country from a country code
func NewCountryFromCode(code string) *Country {
	country := &Country{}
	country.Name, _ = langreg.RegionName(code)
	return country
}
