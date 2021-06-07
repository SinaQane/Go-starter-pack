/* First task: creating a rate limiter for clients' requests */

package main

import (
	"net/http"
	"sync"
	"time"
)

var lock = sync.Mutex{}
var breakLoop = make(chan struct{})

var ipAddresses = make(map[string]int)
var requestLimit = 10
var restoreTime = 10000

func readUserIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

func rateLimiter(ip string, w http.ResponseWriter) bool {
	lock.Lock()
	if cnt, ok := ipAddresses[ip]; cnt < requestLimit {
		if ok {
		ipAddresses[ip] = cnt + 1
		} else {
			ipAddresses[ip] = 1
		}
		lock.Unlock()
		return true
	}
	lock.Unlock()
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte("You've attempted too many requests..."))
	return false

}

func restoreLimit()  {
	for {
		time.Sleep(time.Duration(restoreTime) * time.Millisecond)
		select {
		case <- breakLoop:
			break
		default:
			lock.Lock()
			for k, v := range ipAddresses {
				if v > 0 {
					ipAddresses[k] = v - 1
					if ipAddresses[k] == 0 {
						delete(ipAddresses, k)
					}
				}
			}
			lock.Unlock()
		}
	}
}

func root(w http.ResponseWriter, r *http.Request)  {
	ip := readUserIP(r)
	ok := rateLimiter(ip, w)
	if ok {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Good to Go..."))
	}
}

func main() {
	go restoreLimit()

	http.HandleFunc("/", root)
	_ = http.ListenAndServe("127.0.0.1:8080", nil)

	breakLoop <- struct{}{}
}
