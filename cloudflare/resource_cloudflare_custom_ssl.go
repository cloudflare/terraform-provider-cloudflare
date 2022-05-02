package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomSsl() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareCustomSslCreate,
		ReadContext: resourceCloudflareCustomSslRead,
		UpdateContext: resourceCloudflareCustomSslUpdate,
		DeleteContext: resourceCloudflareCustomSslDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomSslImport,
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
	}
}

func resourceCloudflareCustomSslCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)
	zcso, err := expandToZoneCustomSSLOptions(d)
	if err != nil {
		return fmt.Errorf("failed to create custom ssl cert: %s", err)
	}

	res, err := client.CreateSSL(context.Background(), zoneID, zcso)
	if err != nil {
		return fmt.Errorf("failed to create custom ssl cert: %s", err)
	}

	if res.ID == "" {
		return fmt.Errorf("failed to find custom ssl in Create response: id was empty")
	}

	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		cert, err := client.SSLDetails(context.Background(), zoneID, res.ID)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("failed to fetch custom ssl cert: %s", err))
		}

		if cert.Status != "active" {
			return resource.RetryableError(fmt.Errorf("waiting for certificate to become active"))
		}

		d.SetId(res.ID)

		resourceCloudflareCustomSslRead(d, meta)
		return nil
	})
}

func resourceCloudflareCustomSslUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()
	var uErr error
	var reErr error
	var updateErr = false
	var reprioritizeErr = false
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	if d.HasChange("custom_ssl_options") {
		zcso, err := expandToZoneCustomSSLOptions(d)
		if err != nil {
			return fmt.Errorf("failed to update custom ssl cert: %s", err)
		}

		res, uErr := client.UpdateSSL(context.Background(), zoneID, certID, zcso)
		if uErr != nil {
			log.Printf("[DEBUG] Failed to update custom ssl cert: %s", uErr)
			updateErr = true
		} else {
			log.Printf("[DEBUG] Custom SSL set to: %s", res.ID)
		}

	}

	if d.HasChange("custom_ssl_priority") {
		zcsp, err := expandToZoneCustomSSLPriority(d)
		if err != nil {
			log.Printf("Failed to update custom ssl cert: %s", err)
		}

		resList, reErr := client.ReprioritizeSSL(context.Background(), zoneID, zcsp)
		if err != nil {
			log.Printf("Failed to update / reprioritize custom ssl cert: %s", reErr)
			reprioritizeErr = true
		} else {
			log.Printf("[DEBUG] Custom SSL reprioritized to: %#v", resList)
		}
	}

	if updateErr && reprioritizeErr {
		return fmt.Errorf("failed to update and reprioritize custom ssl cert: %s, %s", uErr, reErr)
	}

	return resourceCloudflareCustomSslRead(d, meta)
}

func resourceCloudflareCustomSslRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	// update all possible schema attributes with fields from api response
	record, err := client.SSLDetails(context.Background(), zoneID, certID)
	if err != nil {
		log.Printf("[WARN] Removing record from state because it's not found in API")
		d.SetId("")
		return nil
	}
	zcso, err := expandToZoneCustomSSLOptions(d)
	if err != nil {
		log.Printf("[WARN] Problem setting zone options not read from state %s", err)
	}
	zcso.BundleMethod = record.BundleMethod
	customSslOpts := flattenCustomSSLOptions(zcso)

	d.SetId(record.ID)
	d.Set("hosts", record.Hosts)
	d.Set("issuer", record.Issuer)
	d.Set("signature", record.Signature)
	if err := d.Set("custom_ssl_options", []interface{}{customSslOpts}); err != nil {
		return fmt.Errorf("[WARN] Error reading custom ssl opts %q: %s", d.Id(), err)
	}
	d.Set("status", record.Status)
	d.Set("uploaded_on", record.UploadedOn.Format(time.RFC3339Nano))
	d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
	d.Set("modified_on", record.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("priority", record.Priority)
	return nil
}

func resourceCloudflareCustomSslDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	log.Printf("[DEBUG] Deleting SSL cert %s for zone %s", certID, zoneID)

	err := client.DeleteSSL(context.Background(), zoneID, certID)
	if err != nil {
		errors.Wrap(err, "failed to delete custom ssl cert setting")
	}
	return nil
}

func resourceCloudflareCustomSslImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/certID\"", d.Id())
	}

	zoneID, certID := idAttr[0], idAttr[1]

	log.Printf("[DEBUG] Importing Cloudflare Custom SSL Cert: id %s for zone %s", certID, zoneID)

	d.Set("zone_id", zoneID)
	d.SetId(certID)

	resourceCloudflareCustomSslRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func expandToZoneCustomSSLPriority(d *schema.ResourceData) ([]cloudflare.ZoneCustomSSLPriority, error) {
	data, dataOk := d.GetOk("custom_ssl_priority")
	log.Printf("[DEBUG] Custom SSL priority found in config: %#v", data)
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
				return mtSlice, fmt.Errorf("Failed to create custom ssl priorities: %s", err)
			}
			// map -> json -> struct
			json.Unmarshal(zcspJSON, &zcsp)
			mtSlice = append(mtSlice, zcsp)
		}
	}
	log.Printf("[DEBUG] Custom SSL priority list creating: %#v", mtSlice)
	return mtSlice, nil
}

func expandToZoneCustomSSLOptions(d *schema.ResourceData) (cloudflare.ZoneCustomSSLOptions, error) {
	data, dataOk := d.GetOk("custom_ssl_options")
	log.Printf("[DEBUG] Custom SSL options found in config: %#v", data)

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
		return zcso, fmt.Errorf("Failed to create custom ssl options: %s", err)
	}

	log.Printf("[DEBUG] Custom SSL JSON: %s", string(zcsoJSON))

	// map -> json -> struct
	json.Unmarshal(zcsoJSON, &zcso)
	log.Printf("[DEBUG] Custom SSL options creating: %#v", zcso)
	return zcso, nil
}

func flattenCustomSSLOptions(sslopt cloudflare.ZoneCustomSSLOptions) map[string]interface{} {
	data := map[string]interface{}{
		"certificate":   sslopt.Certificate,
		"private_key":   sslopt.PrivateKey,
		"bundle_method": sslopt.BundleMethod,
		"type":          sslopt.Type,
	}

	if sslopt.GeoRestrictions.Label != "" && sslopt.GeoRestrictions.Label != "custom" {
		data["geo_restrictions"] = sslopt.GeoRestrictions.Label
	}

	return data
}
