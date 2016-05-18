package etcd

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
	"github.com/intelsdi-x/snap/control/plugin"
)

// works insert data into etcd only when the data is valid
func worker(host string, m plugin.MetricType) error {
	cfg := client.Config{
		Endpoints: []string{host},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return err
	}
	kapi := client.NewKeysAPI(c)

	_, err = kapi.Set(context.Background(), fmt.Sprintf("%v", m.Namespace()), fmt.Sprintf("%v", m.Data()), nil)
	return err
}
