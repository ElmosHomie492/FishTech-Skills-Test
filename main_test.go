package main

import (
	"testing"

	"github.com/franela/goblin"
)

func Test_ipValid(t *testing.T) {
	expectedResult := true
	actualResult := ipValid("142.251.46.142")
	failedResult := ipValid("123.123.12")

	g := goblin.Goblin(t)
	g.Describe("Validate IP Address", func() {
		g.It("Should validate that a given IP address is in the correct format", func() {
			g.Assert(actualResult).Equal(expectedResult)
		})
		g.It("Should not return TRUE if an invalid IP address is passed", func() {
			g.Assert(failedResult).IsFalse()
		})
	})
}

func Test_getWhoIsInfo(t *testing.T) {
	expectedResult := "Google LLC"
	actualResult := getWhoIsInfo("142.251.46.142")

	g := goblin.Goblin(t)
	g.Describe("Validate WhoIs info that is returned", func() {
		g.It("Should return a value for OrgName if successful", func() {
			g.Assert(actualResult.OrgName).Equal(expectedResult)
		})
	})
}

func Test_getGeoIpInfo(t *testing.T) {
	expectedResult := "Google LLC"
	actualResult := getGeoIpInfo("142.251.46.142")

	g := goblin.Goblin(t)
	g.Describe("Validate GeoIP info that is returned", func() {
		g.It("Should return a value for NetName if successful", func() {
			g.Assert(actualResult.ISP).Equal(expectedResult)
		})
	})
}

func Test_buildResponse(t *testing.T) {
	expectedResult := getGeoIpInfo("142.251.46.142")
	actualResult := buildResponse("142.251.46.142")

	g := goblin.Goblin(t)
	g.Describe("Validate the response that is built", func() {
		g.It("Should match the value returned from getGeoIpInfo if using the same IP address", func() {
			g.Assert(actualResult).Equal(expectedResult)
		})
	})
}
