[![License](https://img.shields.io/github/license/dbschenker/patsch?color=blue)](https://github.com/dbschenker/patsch/blob/main/LICENSE)
[![Releases](https://img.shields.io/github/v/tag/dbschenker/patsch?color=blue)](https://github.com/dbschenker/patsch/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbschenker/patsch)](https://goreportcard.com/report/github.com/dbschenker/patsch)

# patsch

**P**ermanently **A**ssert **T**arget **S**ucceeds **C**heck **H**ealth

# use cases
* used by kubernetes cluster admins to quickly identify faulty ingresses
* used by kubernetes cluster admins to monitor zero-downtimeness and DNS changes during cluster blue/green deployments

# installation
download from Releases and extract into your `$PATH`

# screenshot 
cluster switch: 
from old loadbalancer (127.0.0.151 127.0.0.220 127.0.0.251) to new one (127.0.0.169 127.0.0.210 127.0.0.234)
```console
$ patsch \
    https://example.com/api/public/myapp-user/health \
    https://api.example.com/api/public/myapp-backend/health \
    https://api.example.com/api/public/db/health \
    https://example.com \
    https://api.example.com/api/public/emyapp/health \
    https://api.example.com/api/public/myapp-backend/health \
    https://api.example.com/api/public/myapp-auth/health

✅  0.88s  200 OK                https://example.com/                                      [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.88s  200 OK                https://api.example.com/api/public/myapp-backend/health  [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.97s  200 OK                https://api.example.com/api/public/redis/health           [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.97s  200 OK                https://api.example.com/api/public/db/health              [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.97s  200 OK                https://example.com                                       [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.97s  200 OK                https://api.example.com/api/public/myapp-auth/health     [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.97s  200 OK                https://api.example.com/api/public/emyapp/health         [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  1.04s  200 OK                https://api.example.com/api/public/myapp-user/health     [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.44s  200 OK                https://example.com                                       [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.44s  200 OK                https://example.com/                                      [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.58s  200 OK                https://api.example.com/api/public/emyapp/health         [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.61s  200 OK                https://api.example.com/api/public/myapp-user/health     [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.63s  200 OK                https://api.example.com/api/public/myapp-auth/health     [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.63s  200 OK                https://api.example.com/api/public/myapp-backend/health  [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.82s  200 OK                https://api.example.com/api/public/redis/health           [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.82s  200 OK                https://api.example.com/api/public/db/health              [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.69s  200 OK                https://example.com                                       [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.69s  200 OK                https://example.com/                                      [127.0.0.151 127.0.0.220 127.0.0.251] 
✅  0.94s  200 OK                https://api.example.com/api/public/myapp-user/health     [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.94s  200 OK                https://api.example.com/api/public/emyapp/health         [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  0.94s  200 OK                https://api.example.com/api/public/db/health              [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  1.14s  200 OK                https://api.example.com/api/public/myapp-auth/health     [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  1.24s  200 OK                https://api.example.com/api/public/redis/health           [127.0.0.169 127.0.0.210 127.0.0.234] 
✅  1.24s  200 OK                https://api.example.com/api/public/myapp-backend/health  [127.0.0.169 127.0.0.210 127.0.0.234]
```

# usage
```bash
patsch https://google.ie "https://auth.example.com" "https://httpbin.org/status/418" "https://httpbin.org/status/511"
```

# synopsis
```console
Gets http status of URL(s) and prints an error, if the status is not okay
Usage: patsch [OPTION]... URL [URL]...
  -a    (EXPERIMENTAL!) auto mode, finds and checks all ingress rules in current kubernetes cluster
  -f    fail mode, exit with an error, if any request fails
  -kubeconfig string
        (optional) absolute path to the kubeconfig file (default "/home/pschu/.kube/config")
  -n int
        interval <secs> (default 2)
  -o    single mode, only check once
  -q    quiet mode, does not print successful requests
```

