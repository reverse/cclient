package cclient

import (
	"net/http"

	cookiejar "net/http/cookiejar"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/proxy"
)

func NewClient(cookieJar *cookiejar.Jar, clientHello utls.ClientHelloID, proxyUrl ...string) (http.Client, error) {
	if cookieJar != nil {
		if len(proxyUrl) > 0 {
			dialer, err := newConnectDialer(proxyUrl[0])
			if err != nil {
				return http.Client{}, err
			}
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer),
			}, nil
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
		}, nil
	}

	if len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, dialer),
			Jar:       cookieJar,
		}, nil
	}
	return http.Client{
		Transport: newRoundTripper(clientHello, proxy.Direct),
		Jar:       cookieJar,
	}, nil

}
