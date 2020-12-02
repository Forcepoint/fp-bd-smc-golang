package elements

import (
	"github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/utils"
	"net/http"
)

type Host struct {
	TypeOf      string      `json:"-" default:"host"`
	Name        string      `json:"name, omitempty"`
	Address     string      `json:"address, omitempty"`
	Ipv6Address interface{} `json:"ipv6_address, omitempty"` //string
	Secondary   []string    `json:"secondary, omitempty"`
	Comment     string      `json:"comment, omitempty"`
}

func (h *Host) AddSecondary(address string, overwriteExisting bool) {
	if !overwriteExisting {
		h.Secondary = append(h.Secondary, address)
	} else {
		h.Secondary = []string{address}
	}
}

func (h *Host) GetTypeOf() string {
	return h.TypeOf
}

func (h *Host) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (h *Host) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, h)
}

func (h *Host) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, h)
}

func (h *Host) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, h)
}

func (h *Host) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, h)
}

func (h *Host) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, h)
}

func (h *Host) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, h)
}

func (h *Host) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, h)
}

type AddressRange struct {
	TypeOf  string `json:"-" default:"address_range"`
	Name    string `json:"name, omitempty"`
	IpRange string `json:"ip_range, omitempty"`
	Comment string `json:"comment, omitempty"`
}

func (a *AddressRange) GetTypeOf() string {
	return a.TypeOf
}

func (a *AddressRange) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (a *AddressRange) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, a)
}

func (a *AddressRange) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, a)
}

func (a *AddressRange) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, a)
}

func (a *AddressRange) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, a)
}

func (a *AddressRange) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, a)
}

func (a *AddressRange) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, a)
}

func (a *AddressRange) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, a)
}

type Router struct {
	TypeOf      string   `json:"-" default:"router"`
	Name        string   `json:"name, omitempty"`
	Address     string   `json:"address, omitempty"`
	Ipv6Address string   `json:"ipv6_address, omitempty"`
	Secondary   []string `json:"secondary, omitempty"`
	Comment     string   `json:"comment, omitempty"`
}

func (r *Router) GetTypeOf() string {
	return r.TypeOf
}

func (r *Router) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (r *Router) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, r)
}

func (r *Router) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, r)
}

func (r *Router) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, r)
}

func (r *Router) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, r)
}

func (r *Router) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, r)
}

func (r *Router) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, r)
}

func (r *Router) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, r)
}

type Network struct {
	TypeOf      string      `json:"-" default:"network"`
	Name        string      `json:"name, omitempty"`
	Ipv4Network string      `json:"ipv4_network, omitempty"`
	Ipv6Network interface{} `json:"ipv6_network, omitempty"` //string
	Comment     string      `json:"comment, omitempty"`
}

func (n *Network) GetTypeOf() string {
	return n.TypeOf
}

func (n *Network) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (n *Network) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, n)
}

func (n *Network) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, n)
}

func (n *Network) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, n)
}

func (n *Network) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, n)
}

func (n *Network) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, n)
}

func (n *Network) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, n)
}

func (n *Network) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, n)
}

type DomainName struct {
	TypeOf  string `json:"-" default:"domain_name"`
	Name    string `json:"name, omitempty"`
	Comment string `json:"comment, omitempty"`
}

func (d *DomainName) GetTypeOf() string {
	return d.TypeOf
}

func (d *DomainName) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (d *DomainName) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, d)
}

func (d *DomainName) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, d)
}

func (d *DomainName) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, d)
}

func (d *DomainName) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, d)
}

func (d *DomainName) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, d)
}

func (d *DomainName) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, d)
}

func (d *DomainName) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, d)
}

/*
URL List Application represents a list of URL's (typically by domain)
that allow for easy grouping for performing whitelist and blacklisting
note: URLListApplication requires SMC API version >= 6.1
*/
type URLListApplication struct {
	TypeOf   string   `json:"-" default:"url_list_application"`
	Name     string   `json:"name"`
	UrlEntry []string `json:"url_entry"`
	Comment  string   `json:"comment"`
}

func (u *URLListApplication) GetTypeOf() string {
	return u.TypeOf
}

func (u *URLListApplication) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (u *URLListApplication) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, u)
}

func (u *URLListApplication) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, u)
}

func (u *URLListApplication) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, u)
}

func (u *URLListApplication) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, u)
}

func (u *URLListApplication) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, u)
}

func (u *URLListApplication) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, u)
}

func (u *URLListApplication) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, u)
}

type Zone struct {
	TypeOf  string `json:"-" default:"interface_zone"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (z *Zone) GetTypeOf() string {
	return z.TypeOf
}

func (z *Zone) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (z *Zone) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, z)
}

func (z *Zone) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, z)
}

func (z *Zone) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, z)
}

func (z *Zone) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, z)
}

func (z *Zone) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, z)
}

func (z *Zone) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, z)
}

func (z *Zone) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, z)
}

type Country struct {
	TypeOf string `json:"-" default:"country"`
}

func (c *Country) GetTypeOf() string {
	return c.TypeOf
}

func (c *Country) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (c *Country) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, c)
}

func (c *Country) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, c)
}

func (c *Country) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, c)
}

func (c *Country) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, c)
}

func (c *Country) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, c)
}

func (c *Country) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, c)
}

func (c *Country) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, c)
}

type IPCountryGroup struct {
	TypeOf string `json:"-" default:"ip_country_group"`
}

func (i *IPCountryGroup) GetTypeOf() string {
	return i.TypeOf
}

func (i *IPCountryGroup) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (i *IPCountryGroup) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, i)
}

func (i *IPCountryGroup) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, i)
}

func (i *IPCountryGroup) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, i)
}

func (i *IPCountryGroup) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, i)
}

func (i *IPCountryGroup) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, i)
}

func (i *IPCountryGroup) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, i)
}

func (i *IPCountryGroup) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, i)
}

type Alias struct {
	TypeOf string `json:"-" default:"alias"`
}

func (a *Alias) GetTypeOf() string {
	return a.TypeOf
}

func (a *Alias) Get(url, cookie string) (*http.Response, error) {
	return utils.Get(url, cookie)
}

func (a *Alias) GetSubElements(url, cookie string) (*http.Response, error) {
	return utils.GetSubElements(url, cookie, a)
}

func (a *Alias) Create(url, cookie string) (*http.Response, error) {
	return utils.Create(url, cookie, a)
}

func (a *Alias) CreateSubElement(url, cookie string) (*http.Response, error) {
	return utils.CreateSubElement(url, cookie, a)
}

func (a *Alias) Update(url, cookie string) (*http.Response, error) {
	return utils.Update(url, cookie, a)
}

func (a *Alias) UpdateSubElement(url, cookie string) (*http.Response, error) {
	return utils.UpdateSubElement(url, cookie, a)
}

func (a *Alias) Delete(url, cookie string) (*http.Response, error) {
	return utils.Delete(url, cookie, a)
}

func (a *Alias) DeleteSubElement(url, cookie string) (*http.Response, error) {
	return utils.DeleteSubElement(url, cookie, a)
}
