package cloudflare

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareApiToken() *schema.Resource {
	p := schema.Resource{
		Schema: map[string]*schema.Schema{
			"resources": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"permission_groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	return &schema.Resource{
		Create: resourceCloudflareApiTokenCreate,
		Read:   resourceCloudflareApiTokenRead,
		Update: resourceCloudflareApiTokenUpdate,
		Delete: resourceCloudflareApiTokenDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policy": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashResource(&p),
				Elem:     &p,
			},
			"request_ip_in": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"request_ip_not_in": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"value": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issued_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudflareApiTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	name := d.Get("name").(string)

	log.Printf("[INFO] Creating Cloudflare ApiToken: name %s", name)

	t := cloudflare.APIToken{
		Name:      name,
		Policies:  resourceDataToApiTokenPolices(d),
		Condition: resourceDataToApiTokenCondition(d),
	}

	t, err := client.CreateAPIToken(t)
	if err != nil {
		return fmt.Errorf("Error creating Cloudflare ApiToken %q: %s", name, err)
	}

	d.SetId(t.ID)
	d.Set("status", t.Status)
	d.Set("issued_on", t.IssuedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", t.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("value", t.Value)

	return nil
}

func resourceDataToApiTokenCondition(d *schema.ResourceData) *cloudflare.APITokenCondition {
	ipIn := []string{}
	ipNotIn := []string{}
	if ips, ok := d.GetOk("request_ip_in"); ok {
		ipIn = expandInterfaceToStringList(ips)
	}
	if ips, ok := d.GetOk("request_ip_not_in"); ok {
		ipNotIn = expandInterfaceToStringList(ips)
	}

	return &cloudflare.APITokenCondition{
		RequestIP: &cloudflare.APITokenRequestIPCondition{
			In:    ipIn,
			NotIn: ipNotIn,
		},
	}
}

func resourceDataToApiTokenPolices(d *schema.ResourceData) []cloudflare.APITokenPolicies {
	policies := d.Get("policy").(*schema.Set).List()
	var cfPolicies []cloudflare.APITokenPolicies

	for _, p := range policies {
		policy := p.(map[string]interface{})

		resources := expandInterfaceToStringList(policy["resources"])
		cfResources := map[string]interface{}{}
		for _, r := range resources {
			cfResources[r] = "*"
		}

		permissionGroups := expandInterfaceToStringList(policy["permission_groups"])
		var cfPermissionGroups []cloudflare.APITokenPermissionGroups
		for _, pg := range permissionGroups {
			cfPermissionGroups = append(cfPermissionGroups, cloudflare.APITokenPermissionGroups{
				ID: pg,
			})
		}

		cfPolicies = append(cfPolicies, cloudflare.APITokenPolicies{
			Effect:           "allow",
			Resources:        cfResources,
			PermissionGroups: cfPermissionGroups,
		})
	}

	return cfPolicies
}

func resourceCloudflareApiTokenRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cloudflare.API)
	tokenID := d.Id()

	t, err := client.GetAPIToken(tokenID)

	log.Printf("[DEBUG] Cloudflare APIToken: %+v", t)
	log.Printf("[DEBUG] Cloudflare APIToken error: %#v", err)

	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Cloudflare ApiToken %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Cloudflare ApiToken %q: %s", d.Id(), err)
	}

	policies := []map[string]interface{}{}

	for _, p := range t.Policies {
		resources := []string{}
		for k, _ := range p.Resources {
			resources = append(resources, k)
		}

		permissionGroups := []string{}
		for _, v := range p.PermissionGroups {
			permissionGroups = append(permissionGroups, v.ID)
		}

		policies = append(policies, map[string]interface{}{
			"resources":         resources,
			"permission_groups": permissionGroups,
		})
	}

	d.Set("name", t.Name)
	d.Set("policies", policies)
	d.Set("status", t.Status)
	d.Set("issued_on", t.IssuedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", t.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("request_ip_in", t.Condition.RequestIP.In)
	d.Set("request_ip_not_in", t.Condition.RequestIP.NotIn)

	return nil
}

func resourceCloudflareApiTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	name := d.Get("name").(string)
	tokenID := d.Id()

	t := cloudflare.APIToken{
		Name:      d.Get("name").(string),
		Policies:  resourceDataToApiTokenPolices(d),
		Condition: resourceDataToApiTokenCondition(d),
	}

	log.Printf("[INFO] Updating Cloudflare ApiToken: name %s", name)

	t, err := client.UpdateAPIToken(tokenID, t)
	if err != nil {
		return fmt.Errorf("Error updating Cloudflare ApiToken %q: %s", name, err)
	}

	d.Set("modified_on", t.ModifiedOn.Format(time.RFC3339Nano))

	return nil
}

func resourceCloudflareApiTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	tokenID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare APIToken: id %s", tokenID)

	err := client.DeleteAPIToken(tokenID)
	if err != nil {
		return fmt.Errorf("Error deleting Cloudflare ApiToken: %s", err)
	}

	return nil
}
