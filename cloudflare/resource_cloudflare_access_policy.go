package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudflareAccessPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessPolicyCreate,
		Read:   resourceCloudflareAccessPolicyRead,
		Update: resourceCloudflareAccessPolicyUpdate,
		Delete: resourceCloudflareAccessPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessPolicyImport,
		},

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"precedence": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"decision": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "non_identity", "bypass"}, false),
			},
			"require": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     policyOptionElement,
			},
			"exclude": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     policyOptionElement,
			},
			"include": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     policyOptionElement,
			},
		},
	}
}

// policyOptionElement is used by `require`, `exclude` and `include`
// attributes to build out the expected access conditions.
var policyOptionElement = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"email": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"email_domain": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"everyone": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	},
}

func resourceCloudflareAccessPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	appID := d.Get("application_id").(string)

	accessPolicy, err := client.AccessPolicy(zoneID, appID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Access Policy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Access Policy %q: %s", d.Id(), err)
	}

	d.Set("name", accessPolicy.Name)
	d.Set("decision", accessPolicy.Decision)
	d.Set("precedence", accessPolicy.Precedence)
	d.Set("require", accessPolicy.Require)
	d.Set("exclude", accessPolicy.Exclude)
	d.Set("include", accessPolicy.Include)

	return nil
}

func resourceCloudflareAccessPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	appID := d.Get("application_id").(string)
	zoneID := d.Get("zone_id").(string)
	newAccessPolicy := cloudflare.AccessPolicy{
		Name:       d.Get("name").(string),
		Precedence: d.Get("precedence").(int),
		Decision:   d.Get("decision").(string),
	}

	newAccessPolicy = appendConditionalAccessPolicyFields(newAccessPolicy, d)

	log.Printf("[DEBUG] Creating Cloudflare Access Policy from struct: %+v", newAccessPolicy)

	accessPolicy, err := client.CreateAccessPolicy(zoneID, appID, newAccessPolicy)
	if err != nil {
		return fmt.Errorf("error creating Access Policy for ID %q: %s", accessPolicy.ID, err)
	}

	d.SetId(accessPolicy.ID)

	return nil
}

func resourceCloudflareAccessPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	appID := d.Get("application_id").(string)
	updatedAccessPolicy := cloudflare.AccessPolicy{
		Name:       d.Get("name").(string),
		Precedence: d.Get("precedence").(int),
		Decision:   d.Get("decision").(string),
		ID:         d.Id(),
	}

	updatedAccessPolicy = appendConditionalAccessPolicyFields(updatedAccessPolicy, d)

	log.Printf("[DEBUG] Updating Cloudflare Access Policy from struct: %+v", updatedAccessPolicy)

	accessPolicy, err := client.UpdateAccessPolicy(zoneID, appID, updatedAccessPolicy)
	if err != nil {
		return fmt.Errorf("error updating Access Policy for ID %q: %s", d.Id(), err)
	}

	if accessPolicy.ID == "" {
		return fmt.Errorf("failed to find Access Policy ID in update response; resource was empty")
	}

	return resourceCloudflareAccessPolicyRead(d, meta)
}

func resourceCloudflareAccessPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	appID := d.Get("application_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Access Policy using ID: %s", d.Id())

	err := client.DeleteAccessPolicy(zoneID, appID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting Access Policy for ID %q: %s", d.Id(), err)
	}

	resourceCloudflareAccessPolicyRead(d, meta)

	return nil
}

func resourceCloudflareAccessPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/accessApplicationID/accessPolicyID\"", d.Id())
	}

	zoneID, accessAppID, accessPolicyID := attributes[0], attributes[1], attributes[2]

	log.Printf("[DEBUG] Importing Cloudflare Access Policy: zoneID %q, appID %q, accessPolicyID %q", zoneID, accessAppID, accessPolicyID)

	d.Set("zone_id", zoneID)
	d.Set("application_id", accessAppID)
	d.SetId(accessPolicyID)

	resourceCloudflareAccessPolicyRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

// appendConditionalAccessPolicyFields determines which of the
// conditional policy enforcement fields it should append to the
// AccessPolicy by iterating over the provided values and generating the
// correct structs.
func appendConditionalAccessPolicyFields(policy cloudflare.AccessPolicy, d *schema.ResourceData) cloudflare.AccessPolicy {
	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		policy.Exclude = buildAccessPolicyCondition(value.(map[string]interface{}))
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		policy.Require = buildAccessPolicyCondition(value.(map[string]interface{}))
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		policy.Include = buildAccessPolicyCondition(value.(map[string]interface{}))
	}

	return policy
}

// buildAccessPolicyCondition iterates the provided `map` of values and
// generates the required (repetitive) structs.
//
// Returns the intended combination structure of AccessPolicyEmail,
// AccessPolicyEmailDomain, AccessPolicyIP and AccessPolicyEveryone
// structs.
func buildAccessPolicyCondition(options map[string]interface{}) []interface{} {
	var policy []interface{}
	for accessPolicyType, values := range options {
		// Since AccessPolicyEveryone is a single boolean, we don't need to
		// iterate over the values like the others so we treat it a little differently.
		if accessPolicyType == "everyone" {
			if values == true {
				log.Printf("[DEBUG] values for everyone %s", values)
				policy = append(policy, cloudflare.AccessPolicyEveryone{})
			}
		} else {
			for _, value := range values.([]interface{}) {
				switch accessPolicyType {
				case "email":
					policy = append(policy, cloudflare.AccessPolicyEmail{Email: struct {
						Email string `json:"email"`
					}{Email: value.(string)}})
				case "email_domain":
					policy = append(policy, cloudflare.AccessPolicyEmailDomain{EmailDomain: struct {
						Domain string `json:"domain"`
					}{Domain: value.(string)}})
				case "ip":
					policy = append(policy, cloudflare.AccessPolicyIP{IP: struct {
						IP string `json:"ip"`
					}{IP: value.(string)}})
				}
			}
		}
	}

	return policy
}
