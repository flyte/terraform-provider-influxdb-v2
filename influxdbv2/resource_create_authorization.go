package influxdb-v2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/domain"
)

func ResourceAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthorizationCreate,
		Delete: resourceAuthorizationDelete,
		Read:   resourceAuthorizationRead,
		Update: resourceAuthorizationUpdate,
		Schema: map[string]*schema.Schema{
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "active",
			},
			"permissions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"org": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"org_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAuthorizationCreate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	permissions, err := getPermissions(d.Get("permissions"))
	if err != nil {
		return fmt.Errorf("error getting permissions: %v", err)
	}
	orgId := d.Get("org_id").(string)
	description := d.Get("description").(string)
	status := domain.AuthorizationUpdateRequestStatus(d.Get("status").(string))
	authorizations := domain.Authorization{
		AuthorizationUpdateRequest: domain.AuthorizationUpdateRequest{
			Description: &(description),
			Status:      &status,
		},
		OrgID:       &orgId,
		Permissions: &permissions,
	}

	result, err := influx.AuthorizationsAPI().CreateAuthorization(context.Background(), &authorizations)
	if err != nil {
		return fmt.Errorf("error creating authorization: %e", err)
	}
	d.SetId(*result.Id)
	return resourceAuthorizationRead(d, meta)
}

func resourceAuthorizationDelete(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	id := d.Id()
	authorization := domain.Authorization{
		Id: &id,
	}
	err := influx.AuthorizationsAPI().DeleteAuthorization(context.Background(), &authorization)
	if err != nil {
		return fmt.Errorf("error deleting authorization: %v", err)
	}
	return nil
}

func resourceAuthorizationRead(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	result, err := influx.AuthorizationsAPI().FindAuthorizationsByOrgID(context.Background(), d.Get("org_id").(string))
	if err != nil {
		return fmt.Errorf("error getting authorization: %v", err)
	}
	authorizations, err := getAuthorizationsById(result, d.Id())
	if err != nil {
		return fmt.Errorf("error reading authorization: %v", err)
	}
	d.Set("status", authorizations.Status)
	d.Set("user_id", authorizations.UserID)
	d.Set("user_org_id", authorizations.OrgID)
	d.Set("token", authorizations.Token)
	return nil
}

func resourceAuthorizationUpdate(d *schema.ResourceData, meta interface{}) error {
	influx := meta.(influxdb2.Client)
	id := d.Id()
	authorization := domain.Authorization{
		Id: &id,
	}
	statusUpdate := domain.AuthorizationUpdateRequestStatus(d.Get("status").(string))
	_, err := influx.AuthorizationsAPI().UpdateAuthorizationStatus(context.Background(), &authorization, statusUpdate)
	if err != nil {
		return fmt.Errorf("error updating authorization: %v", err)
	}
	return resourceAuthorizationRead(d, meta)
}

func getPermissions(input interface{}) ([]domain.Permission, error) {
	result := []domain.Permission{}
	permissionsSet := input.(*schema.Set).List()
	for _, permission := range permissionsSet {
		perm, ok := permission.(map[string]interface{})
		if ok {
			resourceSet := perm["resource"].(*schema.Set).List()
			for _, resource := range resourceSet {
				res := resource.(map[string]interface{})
				var id, org_id, name, org = "", "", "", ""
				if res["id"] != nil {
					id = res["id"].(string)
				}
				if res["org_id"] != nil {
					org_id = res["org_id"].(string)
				}
				if res["name"] != nil {
					name = res["name"].(string)
				}
				if res["org"] != nil {
					org = res["org"].(string)
				}
				Resource := domain.Resource{Type: domain.ResourceType(res["type"].(string)), Id: &id, OrgID: &org_id, Name: &name, Org: &org}
				each := domain.Permission{Action: domain.PermissionAction(perm["action"].(string)), Resource: Resource}
				result = append(result, each)
			}
		}
	}
	return result, nil
}

func getAuthorizationsById(input *[]domain.Authorization, id string) (domain.Authorization, error) {
	result := domain.Authorization{}
	for _, authorization := range *input {
		if *authorization.Id == id {
			result = authorization
		}
	}
	return result, nil
}
