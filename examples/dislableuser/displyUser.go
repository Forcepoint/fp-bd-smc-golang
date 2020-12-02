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
		AccessKey:  "pQhxWBSwRT7njZdks9t8rblM",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	body, err := sm.GetAllAdmins()
	result, _ := utils.ResponseToMap(body)
	for _, users := range result {
		for _, user := range users {
			if user["name"] == "student" {
				fmt.Println(user["href"])
			}
		}
	}

}
