package vpn

import (
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
	"net/http"
)

type GatewayProfile struct {
	TypeOf       string          `json:"-" default:"gateway_profile"`
	Name         string          `json:"name"`
	Capabilities map[string]bool `json:"capabilities"`
}

type GatewaySettings struct {
	TypeOf                      string `json:"-" default:"gateway_settings"`
	Name                        string `json:"name"`
	NegotiationExpiration       int64  `json:"negotiation_expiration"`
	NegotiationRetryTimer       int64  `json:"negotiation_retry_timer"`
	NegotiationRetryMaxNumber   int8   `json:"negotiation_retry_max_number"`
	NegotiationRetryTimerMax    int64  `json:"negotiation_retry_timer_max"`
	CertificateCacheCrlValidity int32  `json:"certificate_cache_crl_validity"`
	MobikeAfterSaUpdate         bool   `json:"mobike_after_sa_update"`
	MobikeBeforeSaUpdate        bool   `json:"mobike_before_sa_update"`
	MobikeNoRrc                 bool   `json:"mobike_no_rrc"`
}

func (g *GatewaySettings) GetTypeOf() string {
	return g.TypeOf
}

func (g *GatewaySettings) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (g *GatewaySettings) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, g)
}

func (g *GatewaySettings) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, g)
}

func (g *GatewaySettings) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, g)
}

func (g *GatewaySettings) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, g)
}

func (g *GatewaySettings) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, g)
}

func (g *GatewaySettings) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, g)
}

func (g *GatewaySettings) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, g)
}

type ExternalEndpoint struct {
	TypeOf           string `json:"-" default:"external_endpoint"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Dynamic          bool   `json:"dynamic"`
	Enabled          bool   `json:"enabled"`
	NatT             bool   `json:"nat_t"`
	ForceNatT        bool   `json:"force_nat_t"`
	IpsecVpn         bool   `json:"ipsec_vpn"`
	IkePhase1IdType  int8   `json:"ike_phase1_id_type"`
	IkePhase1IdValue string `json:"ike_phase1_id_value"`
	ConnectionType   string `json:"connection_type_ref"`
}

func (e *ExternalEndpoint) GetTypeOf() string {
	return e.TypeOf
}

func (e *ExternalEndpoint) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (e *ExternalEndpoint) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, e)
}

func (e *ExternalEndpoint) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, e)
}

func (e *ExternalEndpoint) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, e)
}

func (e *ExternalEndpoint) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, e)
}

func (e *ExternalEndpoint) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, e)
}

func (e *ExternalEndpoint) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, e)
}

func (e *ExternalEndpoint) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, e)
}

type ExternalGateway struct {
	TypeOf           string              `json:"-" default:"external_gateway"`
	Name             string              `json:"name"`
	ExternalEndpoint []string            `json:"end_point"`
	VPNSite          []map[string]string `json:"site"`
	TrustAllCas      bool                `json:"trust_all_cas"`
	GatewayProfile   string              `json:"gateway_profile"`
}

func (e *ExternalGateway) GetTypeOf() string {
	return e.TypeOf
}

func (e *ExternalGateway) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (e *ExternalGateway) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, e)
}

func (e *ExternalGateway) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, e)
}

func (e *ExternalGateway) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, e)
}

func (e *ExternalGateway) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, e)
}

func (e *ExternalGateway) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, e)
}

func (e *ExternalGateway) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, e)
}

func (e *ExternalGateway) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, e)
}

/*
A VPN Profile is used to specify cryptography and other Policy VPN specific
features. Every PolicyVPN requires a VPNProfile. The system will provide
common profiles labeled as System Element that can be used without
modification.
*/
//noinspection GoNameStartsWithPackageName
type VPNProfile struct {
	TypeOf       string `json:"-" default:"vpn_profile"`
	Name         string `json:"name"`
	Comment      string `json:"comment"`
	Data         string `json:"data"`
	Capabilities map[string]bool
}

func (v *VPNProfile) GetTypeOf() string {
	return v.TypeOf
}

func (v *VPNProfile) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (v *VPNProfile) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, v)
}

func (v *VPNProfile) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, v)
}

func (v *VPNProfile) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, v)
}

func (v *VPNProfile) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, v)
}

func (v *VPNProfile) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, v)
}

func (v *VPNProfile) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, v)
}

func (v *VPNProfile) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, v)
}

/*
Introduced in SMC >= 6.5 to provide a way to group VPN element types

ConnectionTypes are used in various VPN configurations such as an
ExternalGateway endpoint element to define how the endpoint should
be treated, i.e. active, aggregate or standby.

:ivar int connectivity_group: connectivity group for this connection type
:ivar str mode: mode, valid options: 'active', 'aggregate', 'standby'
*/
type ConnectionType struct {
	TypeOf            string `json:"-" default:"connection_type"`
	Name              string `json:"name"`
	Mode              string `json:"mode"`
	Comment           string `json:"comment"`
	ConnectivityGroup int8   `json:"connectivity_group"`
}

func (c *ConnectionType) GetTypeOf() string {
	return c.TypeOf
}

func (c *ConnectionType) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (c *ConnectionType) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, c)
}

func (c *ConnectionType) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, c)
}

func (c *ConnectionType) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, c)
}

func (c *ConnectionType) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, c)
}

func (c *ConnectionType) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, c)
}

func (c *ConnectionType) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, c)
}

func (c *ConnectionType) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, c)
}
