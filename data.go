package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func buildResponse(ip string) []byte {
	whoIsInfo := getWhoIsInfo(ip)
	fmt.Println(whoIsInfo)

	geoIpInfo := getGeoIpInfo(ip)

	jsonResponse, err := json.Marshal(`GeoIpInfo` + geoIpInfo)

	if err != nil {
		log.Fatal(err)
	}

	return jsonResponse
}

func getWhoIsInfo(ip string) string {
	if ipValid(ip) {
		ianaServer := "whois.iana.org"
		ianaResponse := runWhoisCommand("-h", ianaServer, ip)

		whoisServer := "whois.arin.net"

		reader := bytes.NewReader(ianaResponse.Bytes())
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, "refer:") {
				whoisServer = strings.TrimPrefix(line, "refer:")
				whoisServer = strings.TrimSpace(whoisServer)
			}
		}

		whois := runWhoisCommand("-h", whoisServer, ip)

		return whois.String()
	}
	return ""
}

func getGeoIpInfo(ip string) string {
	if ipValid(ip) {
		response, err := http.Get("https://json.geoiplookup.io/" + ip)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		return string(responseData)
	}
	return ""
}
