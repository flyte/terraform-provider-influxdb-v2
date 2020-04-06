provider "influxdbv2" {
    url = "http://influxdbv2.example.com:9999"
    username = "influxdbUsername"
    password = "influxdbPassword"
    token = "influxdbToken"
}

data "influxdbv2_getbucketsinsource" "test"{
    id = "020f755c3c082000"
}

output "links_self" {
    value = data.influxdbv2_getbucketsinsource.test.output["Links"].Self
}