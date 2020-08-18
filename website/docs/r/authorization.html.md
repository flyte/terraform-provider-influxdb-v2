---
layout: "influxdbv2"
page_title: "InfluxDB V2: influxdbv2_authorization"
sidebar_current: "docs-influxdbv2-resource-authorization"
description: |-
  The influxdbv2_authorization resource allows an InfluxDB to create, update and delete authorizations.
---

## Example Usage

### Create a new authorization
```hcl
resource "influxdbv2_authorization" "new_authorization" {
    org_id = influxdbv2_bucket.test.org_id
        description = "new token"
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
```

### Output usage
```hcl
output "Token" {
    value = influxdbv2_authorization.new_authorization.token
}
output "status" {
    value = influxdbv2_authorization.new_authorization.status
}
output "user_id" {
    value = influxdbv2_authorization.new_authorization.user_id
}
output "user_org_id" {
    value = influxdbv2_authorization.new_authorization.user_org_id
}
```

## Argument Reference

The following arguments are supported: 

* ``description`` (Optional) The description of the bucket.
* ``org_id`` (Required) The organization id the bucket is linked.
* ``status`` (Optional) Status of the authorization, can be "active" or "inactive" - Default "active"
* ``permissions`` (Required) Permissions array of the authorization.
    * ``action`` (Required) Action of the permission, can be "read" or "write".
    * ``resource`` (Required) Object Resource
        * ``id`` (Required) ID of the resource the permission is linked
        * ``name`` (Optional) Name of the resource 
        * ``org`` (Optional) Name of the organization with orgID.
        * ``orgID`` (Required) Org ID to link to.
        * ``type`` (Required) The type of authorization, can be "authorizations" "buckets" "dashboards" "orgs" "sources" "tasks" "telegrafs" "users" "variables" "scrapers" "secrets" "labels" "views" "documents" "notificationRules" "notificationEndpoints" "checks" "dbrp"

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* ``user_id`` - The user ID which is created with the authorization.
* ``user_org_id`` - The org ID linked to the user of that authorization.
* ``token`` - The token newly created.
