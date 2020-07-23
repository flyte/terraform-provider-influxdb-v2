provider "influxdbv2-onboarding" {
    url = "http://localhost:9999"
}

resource "influxdbv2-onboarding_setup" "setup" {
    username = "test"
    password = "test1234"
    bucket = "test-bucket"
    org = "test-org"
    retention_period = 4
}

provider "influxdbv2" {
    url = "http://localhost:9999/"
    token = influxdbv2-onboarding_setup.setup.token
}

data "influxdbv2_ready" "test" {}

output "influxdbv2_ready" {
    value = data.influxdbv2_ready.test.output["url"]
}

resource "influxdbv2_bucket" "test" {
    description = "Le bucket terraform"
    name = "le bucket de test terraform modifié "
    org_id = influxdbv2-onboarding_setup.setup.org_id
    retention_rules {
        every_seconds = 10
    }
    rp = ""
}

resource "influxdbv2_authorization" "fred" {
    org_id = influxdbv2_bucket.test.org_id
    description = "fred token"
    status = "inactive"
    permissions {
        action = "read"
        resource {
            id = influxdbv2_bucket.test.id
            org_id = influxdbv2_bucket.test.org_id
            type = "buckets"
        }
    }
    permissions {
        action = "write"
        resource {
            id = influxdbv2_bucket.test.id
            org_id = influxdbv2_bucket.test.org_id
            type = "buckets"
        }
    }
}

output "org_id" {
    value = influxdbv2_bucket.test.org_id
}
output "name" {
    value = influxdbv2_bucket.test.name
}
output "description" {
    value = influxdbv2_bucket.test.description
}

output "Token" {
    value = influxdbv2_authorization.fred.token
}
output "status" {
    value = influxdbv2_authorization.fred.status
}
output "user_id" {
    value = influxdbv2_authorization.fred.user_id
}
output "user_org_id" {
    value = influxdbv2_authorization.fred.user_org_id
}
output "auth_id" {
    value = influxdbv2_authorization.fred.id
}


output "created_at" {
    value = influxdbv2_bucket.test.created_at
}
output "updated_at" {
    value = influxdbv2_bucket.test.updated_at
}
output "type" {
    value = influxdbv2_bucket.test.type
}
output "retention_rules" {
    value = influxdbv2_bucket.test.retention_rules
}
output "rp" {
    value = influxdbv2_bucket.test.rp
}