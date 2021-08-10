#!/usr/bin/env bash

echo "launching influx..."
docker run -d --name tf_example_influxdb -p 8086:8086 quay.io/influxdb/influxdb:v2.0.7
while ! $(curl -sS 'http://localhost:8086/ready' | grep -q ready); do echo 'Waiting for influx...'; sleep 1; done
echo "influx ready"
