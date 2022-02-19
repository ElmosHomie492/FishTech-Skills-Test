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

func buildResponse(ip string) GeoIpInfo {
	geoIpInfo := getGeoIpInfo(ip)

	return geoIpInfo
}

func getWhoIsInfo(ip string) WhoIsInfo {
	// In both this method and the getGeoIpInfo() method,
	// I utilized the json.Unmarshal method to unmarshal
	// the resulting byte arrays into the structs that I
	// defined in structs.go

	// NOTE: This data is no longer used for anything.
	// I had some issues getting whois to run via an
	// AWS Lambda function
	var whoIsInfo WhoIsInfo
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
		json.Unmarshal(parseWhoIsInfo(whois.String()), &whoIsInfo)
	}
	return whoIsInfo
}

func parseWhoIsInfo(data string) []byte {
	dataList := strings.Split(data, "\n")
	dataMap := make(map[string]string)

	for i := 0; i < len(dataList); i++ {
		if dataList[i] != "" {
			key := dataList[i][:strings.IndexByte(dataList[i], ':')]
			value := strings.TrimLeft(dataList[i][strings.IndexByte(dataList[i], ':')+1:], " ")

			dataMap[key] = value
		}
	}

	returnData, err := json.Marshal(dataMap)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return returnData
}

func getGeoIpInfo(ip string) GeoIpInfo {
	// Showing utilization of a third party API
	var geoIpInfo GeoIpInfo
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

		json.Unmarshal(responseData, &geoIpInfo)
	}
	return geoIpInfo
}
