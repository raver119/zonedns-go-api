package api

import "net"

type IPv4 string
type IPv6 string

type Zone struct {
	id   int64
	Name string
	A    []IPv4
	AAAA []IPv6
}

type Domain struct {
	id     int64
	Name   string
	ZoneID int64
	Txt    string
}

/*
	This method allows to "move" domain to another zone
*/
func (d *Domain) ChangeZone(zone Zone) {
	d.ZoneID = zone.id
}

func (z *Zone) Id() int64 {
	return z.id
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
		id:   id,
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
		id:   id,
		Name: name,
		A:    a,
	}
}

/*
	This function creates new zone, with IPv6 addresses only
*/
func NewZone6(id int64, name string, aaaa []IPv6) Zone {
	return Zone{
		id:   id,
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
		ZoneID: zone.id,
		Txt:    "",
	}
}
