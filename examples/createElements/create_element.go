package main

import (
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/interfaces"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/elements"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/vpn"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := initSMCClientAndLogin()

	createElement(client, &elements.Network{
		Name:        "Test Network 6",
		Ipv4Network: "1.2.5.0/24",
		Ipv6Network: nil,
		Comment:     "Test Comment 6",
	})

	createSubElement(client, &vpn.ExternalGateway{}, &vpn.ExternalEndpoint{
		Name:             "External Endpoint Test 99",
		Address:          "10.10.10.12",
		Dynamic:          false,
		Enabled:          true,
		NatT:             true,
		ForceNatT:        false,
		IpsecVpn:         true,
		IkePhase1IdType:  3,
		IkePhase1IdValue: "",
		ConnectionType:   "http://217.182.179.255:3080/6.6/elements/connection_type/1",
	}, 4)

	getSingleElement(client, &vpn.ExternalGateway{}, 4)

	getSubElement(client, &vpn.ExternalGateway{}, &vpn.ExternalEndpoint{}, 4)

	getAllElementsOfType(client, &vpn.ExternalGateway{})
}

func initSMCClientAndLogin() *smc.Smc {
	client := smc.Smc{
		APIVersion:  "6.6",
		Hostname:    "217.182.179.255",
		Port:        "3080",
		AccessKey:   "v4rSQQcz6qnh8Pxb6BR37aSn",
		EntryPoints: nil,
	}

	err := client.Login()

	if err != nil {
		log.Fatalln("error logging in to smc with credentials", err)
	}

	return &client
}

func createElement(client *smc.Smc, element interfaces.SmcElement) {
	resp, err := client.CreateElement(element)

	if err != nil {
		log.Fatalln(fmt.Sprintf("error creating element of type %s", element.GetTypeOf()), err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Println(fmt.Sprintf("error creating element of type %s Status code: %d", element.GetTypeOf(), resp.StatusCode), err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respData))
}

func createSubElement(client *smc.Smc, parentElement, subElement interfaces.SmcElement, parentId int) {
	resp, err := client.CreateSubElement(parentElement, subElement, parentId)

	if err != nil {
		log.Fatalln(fmt.Sprintf("error creating element of type %s", subElement.GetTypeOf()), err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Println(fmt.Sprintf("error creating element of type %s Status code: %d", subElement.GetTypeOf(), resp.StatusCode), err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respData))
}

func getSingleElement(client *smc.Smc, elementType interfaces.SmcElement, id int) {
	resp, err := client.GetElement(elementType, id)

	if err != nil {
		log.Fatalln(fmt.Sprintf("error retrieving element of type %s", elementType.GetTypeOf()), err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("error retrieving element of type %s Status code: %d", elementType.GetTypeOf(), resp.StatusCode), err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respData))
}

func getSubElement(client *smc.Smc, parentElement, subElement interfaces.SmcElement, id int) {
	resp, err := client.GetSubElements(parentElement, subElement, id)

	if err != nil {
		log.Fatalln(fmt.Sprintf("error retrieving element of type %s", subElement.GetTypeOf()), err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("error retrieving element of type %s Status code: %d", subElement.GetTypeOf(), resp.StatusCode), err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respData))
}

func getAllElementsOfType(client *smc.Smc, elementType interfaces.SmcElement) {
	resp, err := client.GetElements(elementType)

	if err != nil {
		log.Fatalln(fmt.Sprintf("error retrieving elements of type %s", elementType.GetTypeOf()), err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Sprintf("error retrieving elements of type %s Status code: %d", elementType.GetTypeOf(), resp.StatusCode), err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(respData))
}
