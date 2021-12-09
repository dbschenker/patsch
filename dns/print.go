package dns

import (
	"reflect"

	"github.com/fatih/color"
)

type Record struct {
	Addr     []string
	Seen     int
	Switched bool
}

type Records map[string]*Record

var (
	db = make(Records)
)

func Print(addrs []string, site string) string {
	rec := record(addrs, site)
	var col color.Attribute
	switch rec.Switched {
	case true:
		col = color.FgGreen
	default:
		col = color.FgBlue
	}
	if rec.Seen > 5 {
		return color.New(col, color.Faint).Sprint(addrs)
	}
	return color.New(col).Sprint(addrs)
}

func record(addrs []string, site string) *Record {
	if _, ok := db[site]; !ok {
		db[site] = &Record{Addr: addrs}
	}
	if !reflect.DeepEqual(db[site].Addr, addrs) {
		db[site].Switched = true
	}
	db[site].Seen += 1
	return db[site]
}
