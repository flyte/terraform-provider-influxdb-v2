provider "influxdbv2-onboarding" {
  url = "http://localhost:9999/"
}

resource "influxdbv2-onboarding_setup" "setup" {
  username = "test"
  password = "test1234"
  bucket = "test-bucket"
  org = "test-org"
  retention_period = 4
}

output "token" {
  value = influxdbv2-onboarding_setup.setup.token
}

output "org_id" {
  value = influxdbv2-onboarding_setup.setup.org_id
}

provider "influxdbv2" {
  url = "http://localhost:9999/"
  token = influxdbv2-onboarding_setup.setup.token
}
resource "influxdbv2_bucket" "initial" {
  description = "Le bucket initial"
  name = "le bucket de deploiement initial"
  org_id = influxdbv2-onboarding_setup.setup.org_id
  retention_rules {
    every_seconds = 30
  }
  rp = ""
}