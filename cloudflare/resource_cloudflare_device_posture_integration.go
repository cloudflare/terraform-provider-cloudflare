package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ws1 = "workspace_one"

func resourceCloudflareDevicePostureIntegration() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareDevicePostureIntegrationSchema(),
		Create: resourceCloudflareDevicePostureIntegrationCreate,
		Read:   resourceCloudflareDevicePostureIntegrationRead,
		Update: resourceCloudflareDevicePostureIntegrationUpdate,
		Delete: resourceCloudflareDevicePostureIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareDevicePostureIntegrationImport,
		},
	}
}

func resourceCloudflareDevicePostureIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	newDevicePostureIntegration := cloudflare.DevicePostureIntegration{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		Interval: d.Get("interval").(string),
	}

	err := setDevicePostureIntegrationConfig(&newDevicePostureIntegration, d)
	if err != nil {
		return fmt.Errorf("error creating Device Posture integration with provided config: %s", err)
	}
	log.Printf("[DEBUG] Creating Cloudflare Device Posture Integration from struct: %+v\n", newDevicePostureIntegration)

	// The API does not return the client_secret so it must be stored in the state func on resource create.
	savedSecret := newDevicePostureIntegration.Config.ClientSecret

	newDevicePostureIntegration, err = client.CreateDevicePostureIntegration(context.Background(), accountID, newDevicePostureIntegration)
	if err != nil {
		return fmt.Errorf("error creating Device Posture Rule for account %q: %s %+v", accountID, err, newDevicePostureIntegration)
	}

	d.SetId(newDevicePostureIntegration.IntegrationID)

	return devicePostureIntegrationReadHelper(d, meta, savedSecret)
}

func resourceCloudflareDevicePostureIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	// Client secret is always read from the local state.
	secret, _ := d.Get("config.0.client_secret").(string)
	return devicePostureIntegrationReadHelper(d, meta, secret)
}

func devicePostureIntegrationReadHelper(d *schema.ResourceData, meta interface{}, secret string) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	devicePostureIntegration, err := client.DevicePostureIntegration(context.Background(), accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Device posture integration %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding device posture integration %q: %s", d.Id(), err)
	}

	devicePostureIntegration.Config.ClientSecret = secret
	d.Set("name", devicePostureIntegration.Name)
	d.Set("type", devicePostureIntegration.Type)
	d.Set("interval", devicePostureIntegration.Interval)
	d.Set("config", convertIntegrationConfigToSchema(devicePostureIntegration.Config))

	return nil
}

func resourceCloudflareDevicePostureIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	updatedDevicePostureIntegration := cloudflare.DevicePostureIntegration{
		IntegrationID: d.Id(),
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Interval:      d.Get("interval").(string),
	}

	err := setDevicePostureIntegrationConfig(&updatedDevicePostureIntegration, d)
	if err != nil {
		return fmt.Errorf("error creating Device Posture Rule with provided match input: %s", err)
	}

	log.Printf("[DEBUG] Updating Cloudflare device posture integration from struct: %+v", updatedDevicePostureIntegration)

	devicePostureIntegration, err := client.UpdateDevicePostureIntegration(context.Background(), accountID, updatedDevicePostureIntegration)
	if err != nil {
		return fmt.Errorf("error updating device posture integration for account %q: %s", accountID, err)
	}

	if devicePostureIntegration.IntegrationID == "" {
		return fmt.Errorf("failed to find device posture integration_id in update response; resource was empty")
	}

	return resourceCloudflareDevicePostureIntegrationRead(d, meta)
}

func resourceCloudflareDevicePostureIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	appID := d.Id()
	accountID := d.Get("account_id").(string)

	log.Printf("[DEBUG] Deleting Cloudflare device posture integration using ID: %s", appID)

	err := client.DeleteDevicePostureIntegration(context.Background(), accountID, appID)
	if err != nil {
		return fmt.Errorf("error deleting Device Posture Rule for account %q: %s", accountID, err)
	}

	resourceCloudflareDevicePostureIntegrationRead(d, meta)

	return nil
}

func resourceCloudflareDevicePostureIntegrationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/devicePostureIntegrationID\"", d.Id())
	}

	accountID, devicePostureIntegrationID := attributes[0], attributes[1]

	log.Printf("[DEBUG] Importing Cloudflare device posture integration: id %s for account %s", devicePostureIntegrationID, accountID)

	d.Set("account_id", accountID)
	d.SetId(devicePostureIntegrationID)

	resourceCloudflareDevicePostureIntegrationRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func setDevicePostureIntegrationConfig(integration *cloudflare.DevicePostureIntegration, d *schema.ResourceData) error {
	if _, ok := d.GetOk("config"); ok {
		config := cloudflare.DevicePostureIntegrationConfig{}
		switch integration.Type {
		case ws1:
			if config.ClientID, ok = d.Get("config.0.client_id").(string); !ok {
				return fmt.Errorf("client_id is a string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret is a string")
			}
			if config.AuthUrl, ok = d.Get("config.0.auth_url").(string); !ok {
				return fmt.Errorf("auth_url is a string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url is a string")
			}
			integration.Config = config
		default:
			return fmt.Errorf("unsupported integration type:%s", integration.Type)
		}
	}
	return nil
}

func convertIntegrationConfigToSchema(input cloudflare.DevicePostureIntegrationConfig) []interface{} {
	m := map[string]interface{}{
		"client_id":     input.ClientID,
		"client_secret": input.ClientSecret,
		"auth_url":      input.AuthUrl,
		"api_url":       input.ApiUrl,
	}
	return []interface{}{m}
}
