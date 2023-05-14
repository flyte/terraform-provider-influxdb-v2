#!/usr/bin/env bash

echo "1) launching influx"
docker run -d --name tf_acc_tests_influxdb -p 8086:8086 influxdb:2.7.1-alpine
while ! $(curl -sS 'http://localhost:8086/ready' | grep -q ready); do echo 'Waiting for influx...'; sleep 1; done

echo "2) onboarding"
onboard=$(curl -fsSL -X POST --data '{"username":"admin", "password":"password", "org":"testorg", "bucket":"testbucket", "retentionPeriodHrs":0}' "http://localhost:8086/api/v2/setup")
echo $onboard | jq -Mcr '.bucket'
token=$(echo $onboard | jq -Mcr '.auth.token')
bucketid=$(echo $onboard | jq -Mcr '.bucket.id')
orgid=$(echo $onboard | jq -Mcr '.bucket.orgID')

echo "3) exporting env var for terraform acceptance tests"
export INFLUXDB_V2_URL="http://localhost:8086"
export INFLUXDB_V2_TOKEN="$token"
export INFLUXDB_V2_BUCKET_ID="$bucketid"
export INFLUXDB_V2_ORG_ID="$orgid"
export TF_ACC=1

