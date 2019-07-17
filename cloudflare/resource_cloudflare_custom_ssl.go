package cloudflare

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"log"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomSsl() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareCustomSslCreate,
		Read:   resourceCloudflareCustomSslRead,
		Update: resourceCloudflareCustomSslUpdate,
		Delete: resourceCloudflareCustomSslDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomSslImport,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_ssl_priority": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"custom_ssl_options": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate": {
							Type:     schema.TypeString,
							Optional: false,
						},
						"private_key": {
							Type:      schema.TypeString,
							Optional:  false,
							Sensitive: true,
						},
						"bundle_method": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"ubiquitous", "optimal", "force"}, false),
						},
					},
				},
			},
			"hosts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bundle_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uploaded_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
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
		return fmt.Errorf("Failed to create custom ssl cert: %s", err)
	}

	res, err := client.CreateSSL(zoneID, zcso)
	if err != nil {
		return fmt.Errorf("Failed to create custom ssl cert: %s", err)
	}

	if res.ID == "" {
		return fmt.Errorf("Failed to find custom ssl in Create response: id was empty")
	}

	d.SetId(res.ID)

	log.Printf("[INFO] Cloudflare Custom SSL ID: %s", d.Id())

	return resourceCloudflareCustomSslRead(d, meta)
}

func resourceCloudflareCustomSslUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	zcso, err := expandToZoneCustomSSLOptions(d)
	if err != nil {
		return fmt.Errorf("Failed to update custom ssl cert: %s", err)
	}

	res, err := client.UpdateSSL(zoneID, certID, zcso)
	if err != nil {
		fmt.Printf("Failed to update custom ssl cert: %s", err)
	}

	log.Printf("[DEBUG] Custom SSL set to: %s", res.ID)

	zcsp, err := expandToZoneCustomSSLPriority(d)
	if err != nil {
		return fmt.Errorf("Failed to update custom ssl cert: %s", err)
	}

	resList, err := client.ReprioritizeSSL(zoneID, zcsp)
	if err != nil {
		return fmt.Errorf("Failed to update / reprioritize custom ssl cert: %s", err)
	}
	log.Printf("[DEBUG] Custom SSL reprioritized to: %#v", resList)

	return resourceCloudflareCustomSslRead(d, meta)
}

func resourceCloudflareCustomSslRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	record, err := client.SSLDetails(zoneID, certID)
	if err != nil {
		log.Printf("[WARN] Removing record from state because it's not found in API")
		d.SetId("")
		return nil
	}

	d.SetId(record.ID)
	d.Set("hosts", record.Hosts)
	d.Set("issuer", record.Issuer)
	d.Set("signature", record.Signature)
	d.Set("status", record.Status)
	d.Set("bundle_method", record.BundleMethod)
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

	err := client.DeleteSSL(zoneID, certID)
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
			zcspJson, err := json.Marshal(newData)
			if err != nil {
				return mtSlice, fmt.Errorf("Failed to create custom ssl priorities: %s", err)
			}
			// map -> json -> struct
			json.Unmarshal(zcspJson, &zcsp)
			mtSlice = append(mtSlice, zcsp)
		}
	}
	log.Printf("[DEBUG] Custom SSL priority list creating: %#v", mtSlice)
	return mtSlice, nil
}

func expandToZoneCustomSSLOptions(d *schema.ResourceData) (cloudflare.ZoneCustomSSLOptions, error) {
	data, dataOk := d.GetOk("custom_ssl_options")
	log.Printf("[DEBUG] Custom SSL options found in config: %#v", data)

	newData := make(map[string]string)
	if dataOk {
		for id, value := range data.(map[string]interface{}) {
			newValue := value.(string)
			newData[id] = newValue
		}
	}

	zcso := cloudflare.ZoneCustomSSLOptions{}
	zcsoJson, err := json.Marshal(newData)
	if err != nil {
		return zcso, fmt.Errorf("Failed to create custom ssl options: %s", err)
	}
	// map -> json -> struct
	json.Unmarshal(zcsoJson, &zcso)
	log.Printf("[DEBUG] Custom SSL options creating: %#v", zcso)
	return zcso, nil
}
