package etcd

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
	"github.com/intelsdi-x/snap/control/plugin"
)

type etcdClient struct {
	keysAPI client.KeysAPI
}

// Config the go etcd API
func NewEtcdClient(host string) (*etcdClient, error) {
	cfg := client.Config{
		Endpoints: []string{host},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	newClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return &etcdClient{client.NewKeysAPI(newClient)}, nil
}

// works insert data into etcd only when the data is valid
func (c *etcdClient) worker(m plugin.MetricType) error {
	_, err := c.keysAPI.Set(context.Background(), fmt.Sprintf("%v/%v", m.Namespace(), m.Timestamp().UTC().UnixNano()), fmt.Sprintf("%v", m.Data()), nil)
	return err
}
