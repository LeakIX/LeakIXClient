package LeakIXClient

import (
	"time"
)

// Single leak occurence on a single service
type SearchResult struct {
	Ip          string                    `json:"ip"`
	Port        string                    `json:"port"`
	Type        string                    `json:"type"`
	Time        time.Time                 `json:"time"`
	Date        int64                     `json:"date"`
	Data        string                    `json:"data"`
	Headers     map[string][]string       `json:"headers"`
	Plugin      string                    `json:"plugin"`
	Network     Network                   `json:"network"`
	GeoLocation GeoLocation               `json:"geoip"`
	Credentials []*HostServiceCredentials `json:"credentials"`
	Software    Software                  `json:"software"`
	Reverse     string                    `json:"reverse"`
	Hostname    string                    `json:"hostname"`
}

type HostServiceCredentials struct {
	NoAuth   bool   `json:"noauth"`
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Raw      []byte `json:"raw"`
}

type Network struct {
	OrganisationName string `json:"organization_name"`
	ASN              int    `json:"asn"`
}

type GeoLocation struct {
	ContinentName  string   `json:"continent_name"`
	RegionISOCode  string   `json:"region_iso_code"`
	CityName       string   `json:"city_name"`
	CountryISOCode string   `json:"country_iso_code"`
	CountryName    string   `json:"country_name"`
	RegionName     string   `json:"region_name"`
	GeoPoint       GeoPoint `json:"location"`
}

type GeoPoint struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Software struct {
	Name            string     `json:"name"`
	Version         string     `json:"version"`
	OperatingSystem string     `json:"os"`
	Modules         []Software `json:"modules"`
}
