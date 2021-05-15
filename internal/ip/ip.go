package ip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/http_client_mock.go -package=mocks github.com/omegion/go-ddclient/internal/ip HTTPClient
// HTTPClient is an interface for http client.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// IP is struct for IP address.
type IP struct {
	Client   HTTPClient
	Provider Provider
	Address  net.IP
}

// Check checks IP address with given provider.
func (i *IP) Check() error {
	resp, err := i.Client.Get(i.Provider.GetURL().String())
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Debugln(fmt.Sprintf("IP provider %s called successfully.", i.Provider.GetName()))

	defer resp.Body.Close()

	i.Address, err = i.Provider.ExtractIP(data)
	if err != nil {
		return err
	}

	log.Debugln(fmt.Sprintf("IP provider %s returned IP %s.", i.Provider.GetName(), i.Address.String()))

	return nil
}
