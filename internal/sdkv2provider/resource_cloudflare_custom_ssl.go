package sdkv2provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomSsl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareCustomSslCreate,
		ReadContext:   resourceCloudflareCustomSslRead,
		UpdateContext: resourceCloudflareCustomSslUpdate,
		DeleteContext: resourceCloudflareCustomSslDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareCustomSslImport,
		},

		SchemaVersion: 1,

		Schema: resourceCloudflareCustomSslSchema(),

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceCloudflareCustomSSLV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceCloudflareCustomSSLStateUpgradeV1,
				Version: 0,
			},
		},

		Description: "Provides a Cloudflare custom SSL resource.",
	}
}

func resourceCloudflareCustomSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	tflog.Debug(ctx, fmt.Sprintf("zone ID: %s", zoneID))
	zcso, err := expandToZoneCustomSSLOptions(ctx, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create custom ssl cert: %w", err))
	}

	res, err := client.CreateSSL(ctx, zoneID, zcso)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create custom ssl cert: %w", err))
	}

	if res.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find custom ssl in Create response: id was empty"))
	}

	retry := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		cert, err := client.SSLDetails(ctx, zoneID, res.ID)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("failed to fetch custom ssl cert: %w", err))
		}

		if cert.Status != "active" {
			return resource.RetryableError(fmt.Errorf("waiting for certificate to become active"))
		}

		d.SetId(res.ID)

		resourceCloudflareCustomSslRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareCustomSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certID := d.Id()
	var uErr error
	var reErr error
	var updateErr = false
	var reprioritizeErr = false
	tflog.Debug(ctx, fmt.Sprintf("zone ID: %s", zoneID))

	if d.HasChange("custom_ssl_options") {
		zcso, err := expandToZoneCustomSSLOptions(ctx, d)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update custom ssl cert: %w", err))
		}

		res, uErr := client.UpdateSSL(ctx, zoneID, certID, zcso)
		if uErr != nil {
			tflog.Debug(ctx, fmt.Sprintf("Failed to update custom ssl cert: %s", uErr))
			updateErr = true
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Custom SSL set to: %s", res.ID))
			if res.ID != certID {
				d.SetId(res.ID)
			}
		}
	}

	if d.HasChange("custom_ssl_priority") {
		zcsp, err := expandToZoneCustomSSLPriority(ctx, d)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("Failed to update custom ssl cert: %s", err))
		}

		resList, reErr := client.ReprioritizeSSL(ctx, zoneID, zcsp)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("Failed to update / reprioritize custom ssl cert: %s", reErr))
			reprioritizeErr = true
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Custom SSL reprioritized to: %#v", resList))
		}
	}

	if updateErr && reprioritizeErr {
		return diag.Errorf("failed to update and reprioritize custom ssl cert: %s, %s", uErr, reErr)
	}

	return resourceCloudflareCustomSslRead(ctx, d, meta)
}

func resourceCloudflareCustomSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certID := d.Id()

	// update all possible schema attributes with fields from api response
	record, err := client.SSLDetails(ctx, zoneID, certID)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Removing record from state because it's not found in API"))
		d.SetId("")
		return nil
	}
	zcso, err := expandToZoneCustomSSLOptions(ctx, d)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Problem setting zone options not read from state %s", err))
	}
	zcso.BundleMethod = record.BundleMethod
	customSslOpts := flattenCustomSSLOptions(zcso)

	d.SetId(record.ID)
	d.Set("hosts", record.Hosts)
	d.Set("issuer", record.Issuer)
	d.Set("signature", record.Signature)
	if err := d.Set("custom_ssl_options", []interface{}{customSslOpts}); err != nil {
		return diag.FromErr(fmt.Errorf("[WARN] Error reading custom ssl opts %q: %w", d.Id(), err))
	}
	d.Set("status", record.Status)
	d.Set("uploaded_on", record.UploadedOn.Format(time.RFC3339Nano))
	d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
	d.Set("modified_on", record.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("priority", record.Priority)
	return nil
}

func resourceCloudflareCustomSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certID := d.Id()

	tflog.Debug(ctx, fmt.Sprintf("Deleting SSL cert %s for zone %s", certID, zoneID))

	err := client.DeleteSSL(ctx, zoneID, certID)
	if err != nil {
		errors.Wrap(err, "failed to delete custom ssl cert setting")
	}
	return nil
}

func resourceCloudflareCustomSslImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/certID\"", d.Id())
	}

	zoneID, certID := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Custom SSL Cert: id %s for zone %s", certID, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(certID)

	resourceCloudflareCustomSslRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func expandToZoneCustomSSLPriority(ctx context.Context, d *schema.ResourceData) ([]cloudflare.ZoneCustomSSLPriority, error) {
	data, dataOk := d.GetOk("custom_ssl_priority")
	tflog.Debug(ctx, fmt.Sprintf("Custom SSL priority found in config: %#v", data))
	var mtSlice []cloudflare.ZoneCustomSSLPriority
	if dataOk {
		for _, innerData := range data.([]interface{}) {
			newData := make(map[string]interface{})
			for id, value := range innerData.(map[string]interface{}) {
				switch idName := id; idName {
				case "id":
					newValue := value.(string)
					newData["ID"] = newValue
				case "priority":
					newValue := value.(int)
					newData[id] = newValue
				default:
					newValue := value
					newData[id] = newValue
				}
			}
			zcsp := cloudflare.ZoneCustomSSLPriority{}
			zcspJSON, err := json.Marshal(newData)
			if err != nil {
				return mtSlice, fmt.Errorf("Failed to create custom ssl priorities: %w", err)
			}
			// map -> json -> struct
			json.Unmarshal(zcspJSON, &zcsp)
			mtSlice = append(mtSlice, zcsp)
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Custom SSL priority list creating: %#v", mtSlice))
	return mtSlice, nil
}

func expandToZoneCustomSSLOptions(ctx context.Context, d *schema.ResourceData) (cloudflare.ZoneCustomSSLOptions, error) {
	data, dataOk := d.GetOk("custom_ssl_options")
	tflog.Debug(ctx, fmt.Sprintf("Custom SSL options found in config: %#v", data))

	newData := make(map[string]interface{})
	if dataOk {
		for _, cert := range data.([]interface{}) {
			for id, value := range cert.(map[string]interface{}) {
				var newValue interface{}
				if id == "geo_restrictions" {
					newValue = cloudflare.ZoneCustomSSLGeoRestrictions{
						Label: value.(string),
					}
				} else {
					newValue = value.(string)
				}
				newData[id] = newValue
			}
		}
	}

	zcso := cloudflare.ZoneCustomSSLOptions{}
	zcsoJSON, err := json.Marshal(newData)
	if err != nil {
		return zcso, fmt.Errorf("Failed to create custom ssl options: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Custom SSL JSON: %s", string(zcsoJSON)))

	// map -> json -> struct
	json.Unmarshal(zcsoJSON, &zcso)
	tflog.Debug(ctx, fmt.Sprintf("Custom SSL options creating: %#v", zcso))
	return zcso, nil
}

func flattenCustomSSLOptions(sslopt cloudflare.ZoneCustomSSLOptions) map[string]interface{} {
	data := map[string]interface{}{
		"certificate":   sslopt.Certificate,
		"private_key":   sslopt.PrivateKey,
		"bundle_method": sslopt.BundleMethod,
		"type":          sslopt.Type,
	}

	if sslopt.GeoRestrictions != nil && sslopt.GeoRestrictions.Label != "" && sslopt.GeoRestrictions.Label != "custom" {
		data["geo_restrictions"] = sslopt.GeoRestrictions.Label
	}

	return data
}
