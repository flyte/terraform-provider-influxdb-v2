provider "influxdbv2" {
    url = "http://localhost:9999"
    username = "influxdbUsername"
    password = "influxdbPassword"
    token = "SaQ3En2d6aWQxLlQZWcyDk5zPA-rZUPcYj-VUsLillGRifxDxB2tjDty_KPU7UopWvuwz_GUyFDiH674rp53xw=="
}

data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
    value = data.influxdbv2_ready.test.output["status"]
}

output "ready_started"  {
    value = data.influxdbv2_ready.test.output["started"]
}

output "ready_up"  {
    value = data.influxdbv2_ready.test.output["up"]
}

resource "influxdbv2_bucket" "initial" {
    description = ""
    name = "le bucket de test terraform"
    org_id = "1de630828fbc3210"
    retention_rules {
        every_seconds = 45
        type = "expire"
    }
    rp = ""
}

output "created_at" {
    value = influxdbv2_bucket.initial.created_at
}
output "updated_at" {
    value = influxdbv2_bucket.initial.updated_at
}
output "type" {
    value = influxdbv2_bucket.initial.type
}
output "retention_rules" {
    value = influxdbv2_bucket.initial.retention_rules
}
output "rp" {
    value = influxdbv2_bucket.initial.rp
}