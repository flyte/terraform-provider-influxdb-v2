---
layout: "influxdb-v2"
page_title: "InfluxDB V2: influxdb-v2_authorization"
sidebar_current: "docs-influxdb-v2-resource-authorization"
description: |-
  The influxdb-v2_authorization resource manages influxdb v2 authorizations.
---

## Example Usage

```hcl
resource "influxdb-v2_authorization" "my_service" {
    org_id = <related organization id>
        description = "a token for my service"
        status = "active"
        permissions {
            action = "read"
            resource {
                id = <some bucket id>
                org_id = <related organization id>
                type = "buckets"
            }
        }
        permissions {
            action = "write"
            resource {
                id = <some bucket id>
                org_id =  <related organization id>
                type = "buckets"
            }
        }
}
```

## Argument Reference

The following arguments are supported: 

* ``org_id`` (Required) The organization id to which the authorization will be linked.
* ``permissions`` (Required) Permission array of the authorization.
    * ``action`` (Required) Action of the permission, can be "read" or "write".
    * ``resource`` (Required) Permission resource
        * ``id`` (Required) ID of the resource to which the permission is linked
        * ``orgID`` (Required) Organization ID to link to.
        * ``type`` (Required) The type of authorization, can be "authorizations" "buckets" "dashboards" "orgs" "sources" "tasks" "telegrafs" "users" "variables" "scrapers" "secrets" "labels" "views" "documents" "notificationRules" "notificationEndpoints" "checks" "dbrp"
        * ``name`` (Optional) Name of the resource 
        * ``org`` (Optional) Name of the organization with orgID.
* ``status`` (Optional) Status of the authorization, can be "active" or "inactive" - Default "active"
* ``description`` (Optional) The description of the bucket.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* ``user_id`` - The user ID which is created with the authorization.
* ``user_org_id`` - The org ID linked to the user of that authorization.
* ``token`` - The token newly created.
