package main

type WhoIsInfo struct {
	Address        string `json:"Address"`
	Cidr           string `json:"CIDR"`
	City           string `json:"City"`
	Comment        string `json:"Comment"`
	Country        string `json:"Country"`
	NetHandle      string `json:"NetHandle"`
	NetName        string `json:"NetName"`
	NetRange       string `json:"NetRange"`
	NetType        string `json:"NetType"`
	OrgAbuseEmail  string `json:"OrgAbuseEmail"`
	OrgAbuseHandle string `json:"OrgAbuseHandle"`
	OrgAbuseName   string `json:"OrgAbuseName"`
	OrgAbusePhone  string `json:"OrgAbusePhone"`
	OrgAbuseRef    string `json:"OrgAbuseRef"`
	OrgID          string `json:"OrgId"`
	OrgName        string `json:"OrgName"`
	OrgTechEmail   string `json:"OrgTechEmail"`
	OrgTechHandle  string `json:"OrgTechHandle"`
	OrgTechName    string `json:"OrgTechName"`
	OrgTechPhone   string `json:"OrgTechPhone"`
	OrgTechRef     string `json:"OrgTechRef"`
	Organization   string `json:"Organization"`
	OriginAS       string `json:"OriginAS"`
	Parent         string `json:"Parent"`
	PostalCode     string `json:"PostalCode"`
	Ref            string `json:"Ref"`
	RegDate        string `json:"RegDate"`
	StateProv      string `json:"StateProv"`
	Updated        string `json:"Updated"`
}

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
}
