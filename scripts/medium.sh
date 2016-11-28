#!/bin/bash

set -e
set -u
set -o pipefail

SNAP_ETCD_PUBLIC_IP="127.0.0.1"
SNAP_ETCD_PORT="5001"

docker_id=$(docker run -d -p 8001:8001 -p ${SNAP_ETCD_PORT}:${SNAP_ETCD_PORT} --name=etcd quay.io/coreos/etcd:v0.4.6 -peer-addr ${SNAP_ETCD_PUBLIC_IP}:8001 -addr ${SNAP_ETCD_PUBLIC_IP}:${SNAP_ETCD_PORT})

sleep 5
export SNAP_ETCD_HOST="http://$SNAP_ETCD_PUBLIC_IP:$SNAP_ETCD_PORT"

UNIT_TEST="go_test"
test_unit

_debug "stopping docker image: ${docker_id}"
docker stop "${docker_id}" > /dev/null
_debug "removing docker image: ${docker_id}"
docker rm "${docker_id}" > /dev/null
