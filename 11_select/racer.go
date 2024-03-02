package select_package

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("timeout occurred")

// func Racer(a, b string) string {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

//	func measureResponseTime(url string) time.Duration {
//		start := time.Now()
//		http.Get(url)
//		duration := time.Since(start)
//		return duration
//	}

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a string, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
