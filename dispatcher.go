package dispatcher

import "net/http"

// dispatch to other http.Handler implementations
type Dispatcher interface {
	Dispatch(r *http.Request) http.Handler
}

// wrapping Dispatcher into http.Handler
type DispatcherFunc func(*http.Request) http.Handler

func (d DispatcherFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d(r).ServeHTTP(w, r)
}

var _ http.Handler = DispatcherFunc(nil)
