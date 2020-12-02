package main

import (
	"encoding/json"
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
	"io/ioutil"
)

func main() {
	sm := smc.Smc{
		Hostname:   "217.182.25.38",
		Port:       "8082",
		APIVersion: "6.7",
		AccessKey:  "9kwXAXcpjKraETJMrb2yKaS7",
	}
	// login
	err := sm.Login()
	if err != nil {
		fmt.Println(err.Error())
	}
	//get all for admins
	body, err := sm.GetAllAdmins()
	if err != nil {
		fmt.Println(err)
	}
	result, _ := utils.ResponseToMap(body)
	fmt.Println("********************* Get all admins *************************************")
	fmt.Printf("%#v\n\n\n", result)
	fmt.Printf("%#v", result["result"])
	fmt.Println("************************ Create Admin *******************************")

	// create user creation object
	var userCreation smc.UserCreation
	//create user
	js := `{"name":"dlo.bagari@corkbizdev.onmicrosoft.com","enabled":true,"allow_sudo":false,"console_superuser":false,"allowed_to_login_in_shared":true,"engine_target":[],"local_admin":false,"superuser":false,"can_use_api":false,"comment":null,"auth_method":"http://192.168.122.10:8082/6.7/elements/authentication_service/2","ldap_user":"http://192.168.122.10:8082/6.7/elements/external_ldap_user/Q049ZGxvLmJhZ2FyaUBjb3JrYml6ZGV2Lm9ubWljcm9zb2Z0LmNvbSxPVT1BQUREQyBVc2VycyxEQz1jb3JrYml6ZGV2LERDPW9ubWljcm9zb2Z0LERDPWNvbSxkb21haW49Y29ya2JpemRldi5vbm1pY3Jvc29mdC5jb20=","permissions":{"permission":[{"granted_domain_ref":"http://192.168.122.10:8082/6.7/elements/admin_domain/1","granted_elements":["http://192.168.122.10:8082/6.7/elements/access_control_list/7"],"role_ref":"http://192.168.122.10:8082/6.7/elements/role/7"}]}}`
	err2 := json.Unmarshal([]byte(js), &userCreation)
	fmt.Println(err2)
	fmt.Printf("%+v\n", userCreation)
	respBody, status, _ := sm.CreateAdmin(&userCreation)
	r, err := ioutil.ReadAll(respBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r), status)
	//logout: terminate the session
	if err := sm.Logout(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n*************************** after logout **************************************")
	fmt.Println(sm)
}
