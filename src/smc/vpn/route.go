package vpn

import (
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
	"net/http"
)

type RouteVPN struct {
	TypeOf string `json:"-" default:"rbvpn_tunnel"`
}

func (r *RouteVPN) GetTypeOf() string {
	return r.TypeOf
}

func (r *RouteVPN) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (r *RouteVPN) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, r)
}

func (r *RouteVPN) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, r)
}

func (r *RouteVPN) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, r)
}

func (r *RouteVPN) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, r)
}

func (r *RouteVPN) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, r)
}

func (r *RouteVPN) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, r)
}

func (r *RouteVPN) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, r)
}

type TunnelMonitoringGroup struct {
	TypeOf string `json:"-" default:"rbvpn_tunnel_group"`
}

func (t *TunnelMonitoringGroup) GetTypeOf() string {
	return t.TypeOf
}

func (t *TunnelMonitoringGroup) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (t *TunnelMonitoringGroup) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, t)
}

func (t *TunnelMonitoringGroup) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, t)
}

func (t *TunnelMonitoringGroup) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, t)
}

func (t *TunnelMonitoringGroup) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, t)
}

func (t *TunnelMonitoringGroup) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, t)
}

func (t *TunnelMonitoringGroup) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, t)
}

func (t *TunnelMonitoringGroup) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, t)
}

type RouteVPNTunnelMonitoringGroup struct {
	TypeOf string `json:"-" default:"rbvpn_tunnel_monitoring_group"`
}

func (rt *RouteVPNTunnelMonitoringGroup) GetTypeOf() string {
	return rt.TypeOf
}

func (rt *RouteVPNTunnelMonitoringGroup) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (rt *RouteVPNTunnelMonitoringGroup) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, rt)
}

func (rt *RouteVPNTunnelMonitoringGroup) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, rt)
}

type TunnelEndpoint struct {
}
