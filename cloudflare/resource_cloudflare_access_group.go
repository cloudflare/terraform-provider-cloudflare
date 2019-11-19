package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareAccessGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessGroupCreate,
		Read:   resourceCloudflareAccessGroupRead,
		Update: resourceCloudflareAccessGroupUpdate,
		Delete: resourceCloudflareAccessGroupDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"require": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     groupOptionElement,
			},
			"exclude": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     groupOptionElement,
			},
			"include": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     groupOptionElement,
			},
		},
	}
}

// groupOptionElement is used by `require`, `exclude` and `include`
// attributes to build out the expected access conditions.
var groupOptionElement = &schema.Resource{
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
		"group": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	},
}

func resourceCloudflareAccessGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	accessGroup, err := client.AccessGroup(accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Access Group %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Access Group %q: %s", d.Id(), err)
	}

	d.Set("name", accessGroup.Name)
	d.Set("require", accessGroup.Require)
	d.Set("exclude", accessGroup.Exclude)
	d.Set("include", accessGroup.Include)

	return nil
}

func resourceCloudflareAccessGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	newAccessGroup := cloudflare.AccessGroup{
		Name: d.Get("name").(string),
	}

	newAccessGroup = appendConditionalAccessGroupFields(newAccessGroup, d)

	log.Printf("[DEBUG] Creating Cloudflare Access Group from struct: %+v", newAccessGroup)

	accessGroup, err := client.CreateAccessGroup(accountID, newAccessGroup)
	if err != nil {
		return fmt.Errorf("error creating Access Group for ID %q: %s", accessGroup.ID, err)
	}

	d.SetId(accessGroup.ID)
	resourceCloudflareAccessGroupRead(d, meta)

	return nil
}

func resourceCloudflareAccessGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	updatedAccessGroup := cloudflare.AccessGroup{
		Name: d.Get("name").(string),
		ID:   d.Id(),
	}

	updatedAccessGroup = appendConditionalAccessGroupFields(updatedAccessGroup, d)

	log.Printf("[DEBUG] Updating Cloudflare Access Group from struct: %+v", updatedAccessGroup)

	accessGroup, err := client.UpdateAccessGroup(accountID, updatedAccessGroup)
	if err != nil {
		return fmt.Errorf("error updating Access Group for ID %q: %s", d.Id(), err)
	}

	if accessGroup.ID == "" {
		return fmt.Errorf("failed to find Access Group ID in update response; resource was empty")
	}

	return resourceCloudflareAccessGroupRead(d, meta)
}

func resourceCloudflareAccessGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Access Group using ID: %s", d.Id())

	err := client.DeleteAccessGroup(accountID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting Access Group for ID %q: %s", d.Id(), err)
	}

	resourceCloudflareAccessGroupRead(d, meta)

	return nil
}

func resourceCloudflareAccessGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessGroupID\"", d.Id())
	}

	accountID, accessGroupID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare Access Group: accountID %q, accessGroupID %q", accountID, accessGroupID)

	d.Set("account_id", accountID)
	d.SetId(accessGroupID)

	resourceCloudflareAccessGroupRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

// appendConditionalAccessGroupFields determines which of the
// conditional group enforcement fields it should append to the
// AccessGroup by iterating over the provided values and generating the
// correct structs.
func appendConditionalAccessGroupFields(group cloudflare.AccessGroup, d *schema.ResourceData) cloudflare.AccessGroup {
	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		group.Exclude = buildAccessGroupCondition(value.(map[string]interface{}))
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		group.Require = buildAccessGroupCondition(value.(map[string]interface{}))
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		group.Include = buildAccessGroupCondition(value.(map[string]interface{}))
	}

	return group
}

// buildAccessGroupCondition iterates the provided `map` of values and
// generates the required (repetitive) structs.
//
// Returns the intended combination structure of AccessGroupEmail,
// AccessGroupEmailDomain, AccessGroupIP and AccessGroupEveryone
// structs.
func buildAccessGroupCondition(options map[string]interface{}) []interface{} {
	var group []interface{}
	for accessGroupType, values := range options {
		// Since AccessGroupEveryone is a single boolean, we don't need to
		// iterate over the values like the others so we treat it a little differently.
		if accessGroupType == "everyone" {
			if values == true {
				log.Printf("[DEBUG] values for everyone %s", values)
				group = append(group, cloudflare.AccessGroupEveryone{})
			}
		} else {
			for _, value := range values.([]interface{}) {
				switch accessGroupType {
				case "email":
					group = append(group, cloudflare.AccessGroupEmail{Email: struct {
						Email string `json:"email"`
					}{Email: value.(string)}})
				case "email_domain":
					group = append(group, cloudflare.AccessGroupEmailDomain{EmailDomain: struct {
						Domain string `json:"domain"`
					}{Domain: value.(string)}})
				case "ip":
					group = append(group, cloudflare.AccessGroupIP{IP: struct {
						IP string `json:"ip"`
					}{IP: value.(string)}})
				case "group":
					group = append(group, cloudflare.AccessGroupAccessGroup{Group: struct {
						ID string `json:"id"`
					}{ID: value.(string)}})
				}
			}
		}
	}

	return group
}
