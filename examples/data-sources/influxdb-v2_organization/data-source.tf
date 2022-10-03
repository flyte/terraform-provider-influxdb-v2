data "influxdb-v2_organization" "organization" {
  name = "my-org"
}

output "influxdb-v2_organization_id" {
  value = data.influxdb-v2_organization.organization.id
}