#!/usr/bin/env bash

echo "==> Initializing influxdb instance with its default data"
echo "Initializing terraform"
cd scripts/
terraform init -plugin-dir=../examples/terraform.d/plugins/linux_amd64

echo "Planning terraform"
terraform plan -out=plan

echo "Applying configuration"
terraform apply plan

exit 0