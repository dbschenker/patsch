package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dbschenker/patsch/dns"

	"github.com/dbschenker/patsch/http"
	"github.com/dbschenker/patsch/kube"
	"github.com/dbschenker/patsch/util"
)

var (
	quiet      bool
	insecure   bool
	fail       bool
	auto       bool
	once       bool
	interval   int
	urls       []string
	kubeconfig *string
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Gets http status of URL(s) and prints an error, if the status is not okay\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTION]... URL [URL]...\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&quiet, "q", false, "quiet mode, does not print successful requests")
	flag.BoolVar(&insecure, "k", false, "ignore ssl verification errors")
	flag.BoolVar(&fail, "f", false, "fail mode, exit with an error, if any request fails")
	flag.BoolVar(&auto, "a", false, "(EXPERIMENTAL!) auto mode, finds and checks all ingress rules in current kubernetes cluster")
	flag.BoolVar(&once, "o", false, "single mode, only check once")
	flag.IntVar(&interval, "n", 2, "interval <secs>")
	kubeconfig = flag.String("kubeconfig", filepath.Join(util.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	if auto {
		urls = kube.FindIngresses(*kubeconfig)
	} else {
		urls = flag.Args()
	}

	if len(urls) == 0 {
		if auto {
			fmt.Println("Could not find ingress routes in current cluster")
		} else {
			flag.Usage()
		}
		os.Exit(1)
	}
	start := time.Now()
	ch := make(chan http.Result)

	client := http.CreateHTTPClient(insecure)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for ; true; <-ticker.C {
		for _, site := range urls {
			go http.Fetch(client, site, ch) // start a goroutine
		}
		for range urls {
			ret := <-ch
			if ret.Error != nil {
				failure(ret.Error)
			} else {
				success(ret.Message, ret.Site)
			}
		}
		if once {
			ticker.Stop()
			break
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func failure(err error) {
	_, _ = fmt.Println("❌ ", err)
	if fail {
		os.Exit(1)
	}
}

func success(msg string, site string) {
	if quiet {
		return
	}
	addrs, err := dns.Lookup(site)
	if err != nil {
		failure(err)
		return
	}

	fmt.Printf("✅ %s %s\n", msg, dns.Print(addrs, site))
}
