package LeakIXClient

import (
	"gitlab.nobody.run/tbi/core"
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
	Network     core.Network                   `json:"network"`
	GeoLocation core.GeoLocation               `json:"geoip"`
	Credentials []*core.HostServiceCredentials `json:"credentials"`
	Software    core.Software                  `json:"software"`
	Reverse     string                    `json:"reverse"`
	Hostname    string                    `json:"hostname"`
	Dataset     core.DatasetLeak		  `json:"dataset"`
	Scheme      string					  `json:"scheme"`
	Certificate core.HostCertificate	  `json:"certificate"`
}