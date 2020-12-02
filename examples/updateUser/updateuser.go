package main

import (
	"encoding/json"
	"fmt"
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc"
	"io/ioutil"
)

func main() {
	js := `
{Name:Ldap Test Enabled:false AllowSudo:false ConsoleSuperuser:false AllowedToLoginInShared:false EngineTarget:[] LocalAdmin:false Superuser:false CanUseApi:false Comment:<nil> AuthMethod: LdapUser: IsUserLocked:false Key:0 Permissions:map[permission:[{GrantedDomainRef:http://217.182.25.38:8082/6.7/elements/admin_domain/1 GrantedElements:[http://217.182.25.38:8082/6.7/elements/access_control_list/7] RoleRef:http://217.182.25.38:8082/6.7/elements/role/4} {GrantedDomainRef:http://217.182.25.38:8082/6.7/elements/admin_domain/1 GrantedElements:[http://217.182.25.38:8082/6.7/elements/access_control_list/7] RoleRef:http://217.182.25.38:8082/6.7/elements/role/7}]] ReadOnly:false System:false SystemKey:0}{"name":"Ldap Test","enabled":false,"allow_sudo":false,"console_superuser":false,"allowed_to_login_in_shared":false,"engine_target":null,"local_admin":false,"superuser":false,"can_use_api":false,"comment":null,"auth_method":"","ldap_user":"","is_user_locked":false,"key":0,"permissions":{"permission":[{"granted_domain_ref":"http://217.182.25.38:8082/6.7/elements/admin_domain/1","granted_elements":["http://217.182.25.38:8082/6.7/elements/access_control_list/7"],"role_ref":"http://217.182.25.38:8082/6.7/elements/role/4"},{"granted_domain_ref":"http://217.182.25.38:8082/6.7/elements/admin_domain/1","granted_elements":["http://217.182.25.38:8082/6.7/elements/access_control_list/7"],"role_ref":"http://217.182.25.38:8082/6.7/elements/role/7"}]},"read_only":false,"system":false,"system_key":0}
`
	var user smc.UserData
	_ = json.Unmarshal([]byte(js), &user)
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
	response, err := sm.UpdateUser(&user)
	bus, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bus), response.StatusCode)
}
