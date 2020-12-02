package main

import (
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"io/ioutil"
)

func main() {
	sm := smc.Smc{
		Hostname:   "217.182.25.38",
		Port:       "8082",
		APIVersion: "6.7",
		AccessKey:  "kxjN8pdwipSCVcWH4B38aJnn",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	ldapAuthService, err := sm.FindExternalLdap()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ldapAuthService["href"])

	activeDirectory, err := sm.FindExternalActiveDirectory("ldap_azure")
	fmt.Println(activeDirectory)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activeDirectory["href"])

	externalLdapUser := smc.ExternalLDAPUser{
		AuthMethod: ldapAuthService["href"],
		IsDefault:  true,
		LdapServer: []string{activeDirectory["href"]},
		Name:       "Test_active_directory",
		ReadOnly:   false,
		System:     false,
	}
	response, err := sm.CreateLdapExternalUser(&externalLdapUser)
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
