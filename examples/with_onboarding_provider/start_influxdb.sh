#!/usr/bin/env bash

echo "launching influx..."
docker run -d --name tf_example_influxdb -p 9999:9999 quay.io/influxdb/influxdb:2.0.0-beta
while ! $(curl -sS 'http://localhost:9999/ready' | grep -q ready); do echo 'Waiting for influx...'; sleep 1; done
echo "influx ready"
