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
		AccessKey:  "pQhxWBSwRT7njZdks9t8rblM",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := sm.DeleteAdmin("neelima.rai")
	if err != nil {
		fmt.Println(err)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r), resp.StatusCode)
	//logout: terminate the session
	if err := sm.Logout(); err != nil {
		fmt.Println(err)
	}
}
