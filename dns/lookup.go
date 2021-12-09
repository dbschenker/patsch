package dns

import (
	"net"
	"net/url"
	"sort"
)

func Lookup(site string) ([]string, error) {
	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	addrs, err := net.LookupHost(u.Hostname())
	if err != nil {
		return nil, err
	}

	sort.Strings(addrs)
	return addrs, nil
}
