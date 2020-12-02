package vpn

import (
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
	"net/http"
)

type PolicyVPN struct {
	TypeOf                string     `json:"-" default:"vpn"`
	Name                  string     `json:"name"`
	Nat                   bool       `json:"nat"`
	MobileVpnTopologyMode string     `json:"mobile_vpn_topology_mode"`
	VPNProfile            VPNProfile `json:"vpn_profile"`
}

func (p *PolicyVPN) GetTypeOf() string {
	return p.TypeOf
}

func (p *PolicyVPN) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (p *PolicyVPN) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, p)
}

func (p *PolicyVPN) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, p)
}

func (p *PolicyVPN) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, p)
}

func (p *PolicyVPN) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, p)
}

func (p *PolicyVPN) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, p)
}

func (p *PolicyVPN) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, p)
}

func (p *PolicyVPN) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, p)
}

type GatewayNode struct {
	Name          string   `json:"name"`
	Gateway       string   `json:"gateway"`
	EnabledSites  []string `json:"enabled_vpn_site"`
	DisabledSites []string `json:"disabled_vpn_site"`
}

type GatewayTreeNode struct {
	Name string `json:"name"`
	Data map[string]string
}
