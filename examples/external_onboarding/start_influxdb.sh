#!/usr/bin/.env bash

echo "1) launching influx"
docker run -d --name tf_example_influxdb -p 9999:9999 quay.io/influxdb/influxdb:2.0.0-beta
while ! $(curl -sS 'http://localhost:9999/ready' | grep -q ready); do echo 'Waiting for influx...'; sleep 1; done

echo "2) onboarding"
onboard=$(curl -fsSL -X POST --data '{"username":"admin", "password":"password", "org":"testorg", "bucket":"testbucket", "retentionPeriodHrs":0}' "http://localhost:9999/api/v2/setup")
echo $onboard | jq -Mcr '.bucket'
token=$(echo $onboard | jq -Mcr '.auth.token')
bucketid=$(echo $onboard | jq -Mcr '.bucket.id')
orgid=$(echo $onboard | jq -Mcr '.bucket.orgID')

echo "3) exporting env var for terraform acceptance tests"
export INFLUXDB_V2_URL="http://localhost:9999"
export INFLUXDB_V2_TOKEN="$token"
export TF_VAR_influx_org_id="$orgid"
