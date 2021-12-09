package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

const (
	MaxIdleConnections int = 50
	RequestTimeout     int = 20
)

type Result struct {
	Message string
	Error   error
	Site    string
}

func CreateHTTPClient(insecure bool) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
			//nolint:gosec
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func Fetch(client *http.Client, site string, ch chan<- Result) {
	start := time.Now()
	resp, err := client.Get(site)

	if err != nil {
		ch <- Result{
			Message: "",
			Error:   err,
			Site:    site,
		}
		return
	}

	secs := time.Since(start).Milliseconds()
	out := fmt.Sprintf("%d\t  %-10.10s  %-70.70s", secs, resp.Status, site)

	if resp.StatusCode < 200 || resp.StatusCode > 499 {
		ch <- Result{
			Message: "",
			Error:   fmt.Errorf(out),
			Site:    site,
		}
		_ = resp.Body.Close() // don't leak resources
		return
	}

	// _, err = io.Copy(os.Stdout, resp.Body)
	_ = resp.Body.Close() // don't leak resources

	ch <- Result{
		Message: out,
		Error:   nil,
		Site:    site,
	}
}
