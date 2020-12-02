package main

import (
	"encoding/json"
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"io/ioutil"
	"net/http"
)

func main() {
	sm := smc.Smc{
		Hostname:   "217.182.25.38",
		Port:       "8082",
		APIVersion: "6.7",
		AccessKey:  "5gKwrtyDLWrvA7gvgRzuKWHS",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	smcRoles := make(map[string][]map[string]string)
	roles := make(map[string]string)
	response, err := sm.GetHttp(fmt.Sprintf("http://%s:%s/%s/elements/role",
		sm.Hostname, sm.Port, sm.APIVersion))
	if err != nil {
	}
	if response == nil {
		fmt.Println("error: no response")
	}
	if err == nil && response.StatusCode == http.StatusOK {

		buff, err := ioutil.ReadAll(response.Body)
		if err != nil {
		}
		if err := json.Unmarshal(buff, &smcRoles); err != nil {
		}
	}
	for _, role := range smcRoles["result"] {
		roles[role["name"]] = role["href"]
	}
	fmt.Println(roles)
}
