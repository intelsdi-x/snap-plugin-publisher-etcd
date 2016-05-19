/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2015 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package etcd

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	// Name of plugin
	Name = "etcd"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.PublisherPluginType
)

var (
	errNoHost    = errors.New("setting metric types requires an etcd host")
	errBadHost   = errors.New("failed to parse given etcd_host")
	errReqFailed = errors.New("request to etcd api failed")
)

type Etcd struct{}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		Name,
		Version,
		Type,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.Unsecure(true),
		plugin.RoutingStrategy(plugin.DefaultRouting),
	)
}

// GetConfigPolicy returns plugin mandatory fields as the config policy
func (e *Etcd) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	rule, err := cpolicy.NewStringRule("etcd_host", true)
	handleErr(err)
	p := cpolicy.NewPolicyNode()
	p.Add(rule)
	c.Add([]string{"intel", "etcd"}, p)
	return c, nil
}

// Publish publishes metric data to etcd
func (e *Etcd) Publish(contentType string, content []byte, config map[string]ctypes.ConfigValue) error {
	logger := getLogger(config)
	var metrics []plugin.MetricType
	switch contentType {
	case plugin.SnapGOBContentType:
		dec := gob.NewDecoder(bytes.NewBuffer(content))
		if err := dec.Decode(&metrics); err != nil {
			logger.WithFields(log.Fields{
				"err": err,
			}).Error("decoding error")
			return err
		}
	default:
		logger.Errorf("unknown content type '%v'", contentType)
		return fmt.Errorf("Unknown content type '%s'", contentType)
	}

	hostcfg, ok := config["etcd_host"]
	if !ok {
		return errNoHost
	}
	host, ok := hostcfg.(ctypes.ConfigValueStr)
	if !ok {
		return errBadHost
	}

	c, err := NewEtcdClient(host.Value)
	if err != nil {
		return err
	}

	return e.saveMetrics(metrics, c)
}

func (e *Etcd) saveMetrics(mts []plugin.MetricType, c *etcdClient) error {

	errs := []string{}
	var err error

	for _, m := range mts {
		err = c.worker(m)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		err = fmt.Errorf(strings.Join(errs, ";"))
	}
	return err
}

func getLogger(config map[string]ctypes.ConfigValue) *log.Entry {
	logger := log.WithFields(log.Fields{
		"plugin-name":    Name,
		"plugin-version": Version,
		"plugin-type":    Type.String(),
	})

	// default
	log.SetLevel(log.WarnLevel)

	if debug, ok := config["debug"]; ok {
		switch v := debug.(type) {
		case ctypes.ConfigValueBool:
			if v.Value {
				log.SetLevel(log.DebugLevel)
				return logger
			}
		default:
			logger.WithFields(log.Fields{
				"field":         "debug",
				"type":          v,
				"expected type": "ctypes.ConfigValueBool",
			}).Error("invalid config type")
		}
	}

	if loglevel, ok := config["log-level"]; ok {
		switch v := loglevel.(type) {
		case ctypes.ConfigValueStr:
			switch strings.ToLower(v.Value) {
			case "warn":
				log.SetLevel(log.WarnLevel)
			case "error":
				log.SetLevel(log.ErrorLevel)
			case "debug":
				log.SetLevel(log.DebugLevel)
			case "info":
				log.SetLevel(log.InfoLevel)
			default:
				log.WithFields(log.Fields{
					"value":             strings.ToLower(v.Value),
					"acceptable values": "warn, error, debug, info",
				}).Warn("invalid config value")
			}
		default:
			logger.WithFields(log.Fields{
				"field":         "log-level",
				"type":          v,
				"expected type": "ctypes.ConfigValueStr",
			}).Error("invalid config type")
		}
	}

	return logger
}

func handleErr(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
