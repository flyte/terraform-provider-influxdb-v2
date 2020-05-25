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