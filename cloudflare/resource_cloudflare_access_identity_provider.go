package cloudflare

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudflareAccessIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareAccessIdentityProviderCreate,
		Read:   resourceCloudflareAccessIdentityProviderRead,
		Update: resourceCloudflareAccessIdentityProviderUpdate,
		Delete: resourceCloudflareAccessIdentityProviderDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareAccessIdentityProviderImport,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"centrify", "citrix_adc_saml", "facebook", "google-apps", "oidc", "github", "google", "saml", "linkedin", "azureAD", "okta", "onetimepin", "onelogin", "authn", "yandex"}, false),
			},
			"config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apps_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attributes": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"auth_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"centrify_account": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"centrify_app_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"certs_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_secret": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"directory_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"email_attribute_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_public_cert": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"issuer_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"okta_account": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"onelogin_account": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sign_request": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"sso_target_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"support_groups": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"token_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareAccessIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cloudflare.API)
	// zoneID := d.Get("zone_id").(string)
	// appID := d.Get("application_id").(string)

	// accessPolicy, err := client.AccessPolicy(zoneID, appID, d.Id())
	// if err != nil {
	// 	if strings.Contains(err.Error(), "HTTP status 404") {
	// 		log.Printf("[INFO] Access Policy %s no longer exists", d.Id())
	// 		d.SetId("")
	// 		return nil
	// 	}
	// 	return fmt.Errorf("Error finding Access Policy %q: %s", d.Id(), err)
	// }

	// d.Set("name", accessPolicy.Name)
	// d.Set("decision", accessPolicy.Decision)
	// d.Set("precedence", accessPolicy.Precedence)
	// d.Set("require", accessPolicy.Require)
	// d.Set("exclude", accessPolicy.Exclude)
	// d.Set("include", accessPolicy.Include)

	return nil
}

func resourceCloudflareAccessIdentityProviderCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cloudflare.API)
	// appID := d.Get("application_id").(string)
	// zoneID := d.Get("zone_id").(string)
	// newAccessPolicy := cloudflare.AccessPolicy{
	// 	Name:       d.Get("name").(string),
	// 	Precedence: d.Get("precedence").(int),
	// 	Decision:   d.Get("decision").(string),
	// }

	// newAccessPolicy = appendConditionalAccessPolicyFields(newAccessPolicy, d)

	// log.Printf("[DEBUG] Creating Cloudflare Access Policy from struct: %+v", newAccessPolicy)

	// accessPolicy, err := client.CreateAccessPolicy(zoneID, appID, newAccessPolicy)
	// if err != nil {
	// 	return fmt.Errorf("error creating Access Policy for ID %q: %s", accessPolicy.ID, err)
	// }

	// d.SetId(accessPolicy.ID)

	return nil
}

func resourceCloudflareAccessIdentityProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cloudflare.API)
	// zoneID := d.Get("zone_id").(string)
	// appID := d.Get("application_id").(string)
	// updatedAccessPolicy := cloudflare.AccessPolicy{
	// 	Name:       d.Get("name").(string),
	// 	Precedence: d.Get("precedence").(int),
	// 	Decision:   d.Get("decision").(string),
	// 	ID:         d.Id(),
	// }

	// updatedAccessPolicy = appendConditionalAccessPolicyFields(updatedAccessPolicy, d)

	// log.Printf("[DEBUG] Updating Cloudflare Access Policy from struct: %+v", updatedAccessPolicy)

	// accessPolicy, err := client.UpdateAccessPolicy(zoneID, appID, updatedAccessPolicy)
	// if err != nil {
	// 	return fmt.Errorf("error updating Access Policy for ID %q: %s", d.Id(), err)
	// }

	// if accessPolicy.ID == "" {
	// 	return fmt.Errorf("failed to find Access Policy ID in update response; resource was empty")
	// }

	return resourceCloudflareAccessIdentityProviderRead(d, meta)
}

func resourceCloudflareAccessIdentityProviderDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cloudflare.API)
	// zoneID := d.Get("zone_id").(string)
	// appID := d.Get("application_id").(string)

	// log.Printf("[DEBUG] Deleting Cloudflare Access Policy using ID: %s", d.Id())

	// err := client.DeleteAccessPolicy(zoneID, appID, d.Id())
	// if err != nil {
	// 	return fmt.Errorf("error deleting Access Policy for ID %q: %s", d.Id(), err)
	// }

	// resourceCloudflareAccessIdentityProviderRead(d, meta)

	return nil
}

func resourceCloudflareAccessIdentityProviderImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// attributes := strings.SplitN(d.Id(), "/", 3)

	// if len(attributes) != 3 {
	// 	return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/accessApplicationID/accessPolicyID\"", d.Id())
	// }

	// zoneID, accessAppID, accessPolicyID := attributes[0], attributes[1], attributes[2]

	// log.Printf("[DEBUG] Importing Cloudflare Access Policy: zoneID %q, appID %q, accessPolicyID %q", zoneID, accessAppID, accessPolicyID)

	// d.Set("zone_id", zoneID)
	// d.Set("application_id", accessAppID)
	// d.SetId(accessPolicyID)

	// resourceCloudflareAccessIdentityProviderRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
