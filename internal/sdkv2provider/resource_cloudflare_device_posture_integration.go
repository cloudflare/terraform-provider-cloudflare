package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ws1         = "workspace_one"
	crowdstrike = "crowdstrike_s2s"
	uptycs      = "uptycs"
	intune      = "intune"
	kolide      = "kolide"
	sentinelone = "sentinelone_s2s"
	tanium      = "tanium_s2s"
)

func resourceCloudflareDevicePostureIntegration() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareDevicePostureIntegrationSchema(),
		CreateContext: resourceCloudflareDevicePostureIntegrationCreate,
		ReadContext:   resourceCloudflareDevicePostureIntegrationRead,
		UpdateContext: resourceCloudflareDevicePostureIntegrationUpdate,
		DeleteContext: resourceCloudflareDevicePostureIntegrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareDevicePostureIntegrationImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Device Posture Integration resource. Device
			posture integrations configure third-party data providers for device
			posture rules.
		`),
	}
}

func resourceCloudflareDevicePostureIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	newDevicePostureIntegration := cloudflare.DevicePostureIntegration{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		Interval: d.Get("interval").(string),
	}

	err := setDevicePostureIntegrationConfig(&newDevicePostureIntegration, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture integration with provided config: %w", err))
	}
	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Device Posture Integration from struct: %+v\n", newDevicePostureIntegration))

	// The API does not return the client_secret so it must be stored in the state func on resource create.
	savedSecret := newDevicePostureIntegration.Config.ClientSecret

	newDevicePostureIntegration, err = client.CreateDevicePostureIntegration(ctx, accountID, newDevicePostureIntegration)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture Rule for account %q: %w %+v", accountID, err, newDevicePostureIntegration))
	}

	d.SetId(newDevicePostureIntegration.IntegrationID)

	return diag.FromErr(devicePostureIntegrationReadHelper(ctx, d, meta, savedSecret))
}

func resourceCloudflareDevicePostureIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Client secret is always read from the local state.
	secret, _ := d.Get("config.0.client_secret").(string)
	return diag.FromErr(devicePostureIntegrationReadHelper(ctx, d, meta, secret))
}

func devicePostureIntegrationReadHelper(ctx context.Context, d *schema.ResourceData, meta interface{}, secret string) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	devicePostureIntegration, err := client.DevicePostureIntegration(ctx, accountID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Device posture integration %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding device posture integration %q: %w", d.Id(), err)
	}

	devicePostureIntegration.Config.ClientSecret = secret
	d.Set("name", devicePostureIntegration.Name)
	d.Set("type", devicePostureIntegration.Type)
	d.Set("interval", devicePostureIntegration.Interval)
	d.Set("config", convertIntegrationConfigToSchema(devicePostureIntegration.Config))

	return nil
}

func resourceCloudflareDevicePostureIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	updatedDevicePostureIntegration := cloudflare.DevicePostureIntegration{
		IntegrationID: d.Id(),
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Interval:      d.Get("interval").(string),
	}

	err := setDevicePostureIntegrationConfig(&updatedDevicePostureIntegration, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Device Posture Rule with provided match input: %w", err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare device posture integration from struct: %+v", updatedDevicePostureIntegration))

	devicePostureIntegration, err := client.UpdateDevicePostureIntegration(ctx, accountID, updatedDevicePostureIntegration)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating device posture integration for account %q: %w", accountID, err))
	}

	if devicePostureIntegration.IntegrationID == "" {
		return diag.FromErr(fmt.Errorf("failed to find device posture integration_id in update response; resource was empty"))
	}

	return resourceCloudflareDevicePostureIntegrationRead(ctx, d, meta)
}

func resourceCloudflareDevicePostureIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare device posture integration using ID: %s", appID))

	err := client.DeleteDevicePostureIntegration(ctx, accountID, appID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Device Posture Rule for account %q: %w", accountID, err))
	}

	resourceCloudflareDevicePostureIntegrationRead(ctx, d, meta)

	return nil
}

func resourceCloudflareDevicePostureIntegrationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/devicePostureIntegrationID\"", d.Id())
	}

	accountID, devicePostureIntegrationID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare device posture integration: id %s for account %s", devicePostureIntegrationID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(devicePostureIntegrationID)

	resourceCloudflareDevicePostureIntegrationRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func setDevicePostureIntegrationConfig(integration *cloudflare.DevicePostureIntegration, d *schema.ResourceData) error {
	if _, ok := d.GetOk("config"); ok {
		config := cloudflare.DevicePostureIntegrationConfig{}
		switch integration.Type {
		case ws1:
			if config.ClientID, ok = d.Get("config.0.client_id").(string); !ok {
				return fmt.Errorf("client_id has to be of type string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.AuthUrl, ok = d.Get("config.0.auth_url").(string); !ok {
				return fmt.Errorf("auth_url has to be of type string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url has to be of type string")
			}
			integration.Config = config
		case crowdstrike:
			if config.ClientID, ok = d.Get("config.0.client_id").(string); !ok {
				return fmt.Errorf("client_id has to be of type string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.CustomerID, ok = d.Get("config.0.customer_id").(string); !ok {
				return fmt.Errorf("customer_id has to be of type string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url has to be of type string")
			}
			integration.Config = config
		case uptycs:
			if config.ClientKey, ok = d.Get("config.0.client_key").(string); !ok {
				return fmt.Errorf("client_id has to be of type string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.CustomerID, ok = d.Get("config.0.customer_id").(string); !ok {
				return fmt.Errorf("customer_id has to be of type string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url has to be of type string")
			}
			integration.Config = config
		case intune:
			if config.ClientID, ok = d.Get("config.0.client_id").(string); !ok {
				return fmt.Errorf("client_id has to be of type string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.CustomerID, ok = d.Get("config.0.customer_id").(string); !ok {
				return fmt.Errorf("customer_id has to be of type string")
			}
			integration.Config = config
		case kolide:
			if config.ClientID, ok = d.Get("config.0.client_id").(string); !ok {
				return fmt.Errorf("client_id has to be of type string")
			}
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			integration.Config = config
		case sentinelone:
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url has to be of type string")
			}
			integration.Config = config
		case tanium:
			if config.ClientSecret, ok = d.Get("config.0.client_secret").(string); !ok {
				return fmt.Errorf("client_secret has to be of type string")
			}
			if config.ApiUrl, ok = d.Get("config.0.api_url").(string); !ok {
				return fmt.Errorf("api_url has to be of type string")
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
		"client_key":    input.ClientKey,
		"customer_id":   input.CustomerID,
	}
	return []interface{}{m}
}
