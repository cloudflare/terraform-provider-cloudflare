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
				ValidateFunc: validation.StringInSlice([]string{"centrify", "facebook", "google-apps", "oidc", "github", "google", "saml", "linkedin", "azureAD", "okta", "onetimepin", "onelogin", "authn", "yandex"}, false),
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
							ForceNew: true,
							// client_secret is a write only operation from the Cloudflare API
							// and once it's set, it is no longer accessible. To avoid storing
							// it and messing up the state, hardcode in the concealed version.
							StateFunc: func(val interface{}) string {
								return "**********************************"
							},
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
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	accessIdentityProvider, err := client.AccessIdentityProviderDetails(accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Access Identity Provider %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("unable to find Access Identity Provider %q: %s", d.Id(), err)
	}

	d.SetId(accessIdentityProvider.ID)
	d.Set("name", accessIdentityProvider.Name)
	d.Set("type", accessIdentityProvider.Type)

	fmt.Printf("concertStructToSchema: %+v", convertStructToSchema(d, accessIdentityProvider.Config))

	config := convertStructToSchema(d, accessIdentityProvider.Config)
	if configErr := d.Set("config", config); configErr != nil {
		return fmt.Errorf("error setting Access Identity Provider configuration: %s", configErr)
	}

	return nil
}

func resourceCloudflareAccessIdentityProviderCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	IDPConfig := cloudflare.AccessIdentityProviderConfiguration{}

	if _, ok := d.GetOk("config"); ok {
		IDPConfig.ClientID = d.Get("config.0.client_id").(string)
		IDPConfig.ClientSecret = d.Get("config.0.client_secret").(string)
	}

	identityProvider := cloudflare.AccessIdentityProvider{
		Name:   d.Get("name").(string),
		Type:   d.Get("type").(string),
		Config: IDPConfig,
	}

	log.Printf("[DEBUG] Creating Cloudflare Access Identity Provider from struct: %+v", identityProvider)

	accessPolicy, err := client.CreateAccessIdentityProvider(accountID, identityProvider)
	if err != nil {
		return fmt.Errorf("error creating Access Identity Provider for ID %q: %s", d.Id(), err)
	}

	d.SetId(accessPolicy.ID)

	return resourceCloudflareAccessIdentityProviderRead(d, meta)
}

func resourceCloudflareAccessIdentityProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	IDPConfig := cloudflare.AccessIdentityProviderConfiguration{}

	if _, ok := d.GetOk("config"); ok {
		tfAction := d.Get("config").([]interface{})[0].(map[string]interface{})

		fmt.Printf("tfAction: %+v", tfAction)

		clientID := tfAction["client_id"].(string)
		clientSecret := tfAction["client_secret"].(string)

		IDPConfig.ClientID = clientID
		IDPConfig.ClientSecret = clientSecret
	}

	log.Printf("[DEBUG] updatedConfig: %+v", IDPConfig)
	updatedAccessIdentityProvider := cloudflare.AccessIdentityProvider{
		Name:   d.Get("name").(string),
		Type:   d.Get("type").(string),
		Config: IDPConfig,
	}

	log.Printf("[DEBUG] Updating Cloudflare Access Identity Provider from struct: %+v", updatedAccessIdentityProvider)

	accessIdentityProvider, err := client.UpdateAccessIdentityProvider(accountID, d.Id(), updatedAccessIdentityProvider)
	if err != nil {
		return fmt.Errorf("error updating Access Identity Provider for ID %q: %s", d.Id(), err)
	}

	if accessIdentityProvider.ID == "" {
		return fmt.Errorf("failed to find Access Identity Provider ID in update response; resource was empty")
	}

	return resourceCloudflareAccessIdentityProviderRead(d, meta)
}

func resourceCloudflareAccessIdentityProviderDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare Access Identity Provider using ID: %s", d.Id())

	_, err := client.DeleteAccessIdentityProvider(accountID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting Access Policy for ID %q: %s", d.Id(), err)
	}

	d.SetId("")

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

// func convertSchemaToStruct(d *schema.ResourceData) (cloudflare.AccessIdentityProviderConfiguration, error) {

// 	fmt.Printf("IDPConfig: %+v", IDPConfig)

// 	return IDPConfig, nil
// }

func convertStructToSchema(d *schema.ResourceData, options cloudflare.AccessIdentityProviderConfiguration) []interface{} {
	// todo: find a better way of confirming we have options
	if options.ClientID == "" {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"client_id":     options.ClientID,
		"client_secret": options.ClientSecret,
	}

	return []interface{}{m}
}
