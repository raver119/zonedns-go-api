package api

import "net"

type IPv4 string
type IPv6 string

type Zone struct {
	Id   int64  `json:"Id"`
	Name string `json:"name"`
	A    []IPv4 `json:"a"`
	AAAA []IPv6 `json:"aaaa"`
}

type Domain struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	ZoneID int64  `json:"zoneId"`
	Txt    string `json:"txt"`
}

/*
	This method allows to "move" domain to another zone
*/
func (d *Domain) ChangeZone(zone Zone) {
	d.ZoneID = zone.Id
}

func validateIPv4(v IPv4) bool {
	return net.ParseIP(string(v)) != nil
}

func validateIPv6(v IPv6) bool {
	return net.ParseIP(string(v)) != nil
}

/*
	This function creates new zone
*/
func NewZone(id int64, name string, a []IPv4, aaaa []IPv6) Zone {
	return Zone{
		Id:   id,
		Name: name,
		A:    a,
		AAAA: aaaa,
	}
}

/*
	This function creates new zone, with IPv4 addresses only
*/
func NewZone4(id int64, name string, a []IPv4) Zone {
	return Zone{
		Id:   id,
		Name: name,
		A:    a,
	}
}

/*
	This function creates new zone, with IPv6 addresses only
*/
func NewZone6(id int64, name string, aaaa []IPv6) Zone {
	return Zone{
		Id:   id,
		Name: name,
		AAAA: aaaa,
	}
}

/*
	This function creates new Domain instance within given Zone
*/
func NewDomain(domain string, zone Zone) Domain {
	return Domain{
		Name:   domain,
		ZoneID: zone.Id,
		Txt:    "",
	}
}
