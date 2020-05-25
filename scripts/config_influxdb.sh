#!/usr/bin/env bash

echo "==> Initializing influxdb instance with its default data"
echo "Initializing terraform"
terraform init -plugin-dir=scripts/terraform.d/plugins/linux_amd64 scripts/

terraform plan -out=plan scripts/

echo "Applying configuration"
terraform apply plan

exit 0