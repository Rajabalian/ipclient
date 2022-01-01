package ipclient

import (
	"context"
	"net/http"
)

// holds the config to be pass to the plugin
type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}
 
type UIDdemo struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &UIDdemo{
		next: next,
		name: name,
	}, nil
}

func (u *UIDdemo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ip := req.Header.Get("X-Forwarded-For")

	req.Header.Set("X-User-Ip", ip)

	u.next.ServeHTTP(res, req)
}
