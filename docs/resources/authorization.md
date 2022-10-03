---
layout: "influxdb-v2"
page_title: "InfluxDB V2: influxdb-v2_authorization"
sidebar_current: "docs-influxdb-v2-resource-authorization"
description: |-
The influxdb-v2_authorization resource manages influxdb v2 authorizations.
---

# influxdb-v2_authorization (Resource)



## Example Usage

```terraform
locals {
  org_id = "example_org_id"
  bucket_id = "example_bucket_id"
}

resource "influxdb-v2_authorization" "example_authorization" {
  org_id      = local.org_id
  description = "Example description"
  status      = "active"
  permissions {
    action = "read"
    resource {
      id     = local.bucket_id
      org_id = local.org_id
      type   = "buckets"
    }
  }
  permissions {
    action = "write"
    resource {
      id     = local.bucket_id
      org_id = local.org_id
      type   = "buckets"
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
        * ``type`` (Required) The type of authorization, can be `authorizations` `buckets` `dashboards` `orgs` `sources` `tasks` `telegrafs` `users` `variables` `scrapers` `secrets` `labels` `views` `documents` `notificationRules` `notificationEndpoints` `checks` `dbrp`
        * ``name`` (Optional) Name of the resource
        * ``org`` (Optional) Name of the organization with orgID.
* ``status`` (Optional) Status of the authorization, can be "active" or "inactive" - Default "active"
* ``description`` (Optional) The description of the bucket.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* ``user_id`` - The user ID which is created with the authorization.
* ``user_org_id`` - The org ID linked to the user of that authorization.
* ``token`` - The token newly created.
## Import

Import is supported using the following syntax:

```shell
terraform import influxdb-v2_authorization.example_authorization <AUTH_ID>
```
