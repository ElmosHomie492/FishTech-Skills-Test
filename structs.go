package main

// NOTE: Not used, but could be a stepping stone towards providing all IP information in a single api call
type GeoIpInfo struct {
	IP             string  `json:"ip"`
	ISP            string  `json:"isp"`
	Org            string  `json:"org"`
	HostName       string  `json:"hostname"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	PostalCode     string  `json:"postal_code"`
	City           string  `json:"city"`
	CountryCode    string  `json:"country_code"`
	CountryName    string  `json:"country_name"`
	ContinentCode  string  `json:"continent_code"`
	ContinentName  string  `json:"continent_name"`
	Region         string  `json:"region"`
	District       string  `json:"district"`
	TimezoneName   string  `json:"timezone_name"`
	ConnectionType string  `json:"connection_type"`
	ASNNumber      int     `json:"asn_number"`
	ASNOrg         string  `json:"asn_org"`
	ASN            string  `json:"asn"`
	CurrencyCode   string  `json:"currency_code"`
	CurrencyName   string  `json:"currency_name"`
	Success        bool    `json:"success"`
	Premium        bool    `json:"premium"`
	WhoIsData      string
}
