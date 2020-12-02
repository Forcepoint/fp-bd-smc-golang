package main

import (
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"io/ioutil"
)

func main() {
	sm := smc.Smc{
		Hostname:   "37.59.133.20",
		Port:       "8082",
		APIVersion: "6.7",
		AccessKey:  "8DAnAJHiyl37XJ2s35gdthlM",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	ad := smc.ActiveDirectoryLDAPS{
		Address:                   "52.232.56.48",
		BaseDn:                    "DC=corkbizdev,DC=onmicrosoft,DC=com",
		BindPassword:              "Kobani83!!1",
		BindUserId:                "CN=Dlo Bagari,OU=AADDC Users,DC=corkbizdev,DC=onmicrosoft,DC=com",
		Name:                      "newLdapServer2",
		Protocol:                  "ldaps",
		Port:                      636,
		Timeout:                   10,
		Retries:                   2,
		AuthPort:                  1812,
		ClientCertBasedUserSearch: "",
		GroupObjectClass: []string{"sggroup", "organizationalUnit", "organization", "groupOfNames",
			"group", "country"},
		PageSize:        1000,
		UserObjectClass: []string{"sguser", "person", "organizationalPerson", "inetOrgPerson"},
	}
	response, err := sm.CreateActiveDirectoryLdap(&ad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.StatusCode)
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r))
}
