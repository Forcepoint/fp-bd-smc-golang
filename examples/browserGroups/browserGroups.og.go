package main

import (
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
)

func main() {
	sm := smc.Smc{
		Hostname:   "217.182.25.38",
		Port:       "8082",
		APIVersion: "6.7",
		AccessKey:  " ",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	var smcUsers []string
	body, err := sm.GetAllAdmins()
	if err != nil {
		fmt.Println(err)
	}
	result, _ := utils.ResponseToMap(body)
	for _, user := range result["result"] {
		smcUsers = append(smcUsers, user["name"])
	}
	usersWithRoles := make(map[string][]string)
	for _, u := range smcUsers {
		usersWithRoles[u] = []string{}
	}
	ldapDom, err := sm.ExternalLdapDomain("azure_users")
	if err != nil {
	}
	// browser LDAP
	url := ldapDom["href"] + "/browse"
	azureAd, err := sm.GetHttp(url)
	rep, _ := utils.ResponseToMap(azureAd.Body)
	for _, r := range rep["result"] {
		// find all groups and users of azure AD
		if r["name"] == "AADDC Users" {
			groups, _ := sm.FindAllGroups(r["href"])
			for _, group := range groups {
				users, _ := sm.FindAllUsers(group["href"])
				for _, user := range users {
					u, _ := sm.ExternalAldapUser(user["href"])
					if _, ok := usersWithRoles[u.DisplayName]; ok {
						usersWithRoles[u.DisplayName] = append(usersWithRoles[u.DisplayName], group["name"])
					}
				}
			}
		}
	}
	fmt.Println(usersWithRoles)
}
