package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareRecordCreate,
		ReadContext: resourceCloudflareRecordRead,
		UpdateContext: resourceCloudflareRecordUpdate,
		DeleteContext: resourceCloudflareRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareRecordImport,
		},

		SchemaVersion: 2,
		Schema:        resourceCloudflareRecordSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceCloudflareRecordV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceCloudflareRecordStateUpgradeV2,
				Version: 1,
			},
		},
	}
}

func resourceCloudflareRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newRecord := cloudflare.DNSRecord{
		Type:   d.Get("type").(string),
		Name:   d.Get("name").(string),
		ZoneID: d.Get("zone_id").(string),
	}

	proxied, proxiedOk := d.GetOkExists("proxied")
	if proxiedOk {
		newRecord.Proxied = cloudflare.BoolPtr(proxied.(bool))
	}

	value, valueOk := d.GetOk("value")
	if valueOk {
		newRecord.Content = value.(string)
	}

	data, dataOk := d.GetOk("data")
	log.Printf("[DEBUG] Data found in config: %#v", data)

	newDataMap := make(map[string]interface{})

	if dataOk {
		dataMap := data.([]interface{})[0]
		for id, value := range dataMap.(map[string]interface{}) {
			newData, err := transformToCloudflareDNSData(newRecord.Type, id, value)
			if err != nil {
				return err
			} else if newData == nil {
				continue
			}
			newDataMap[id] = newData
		}

		newRecord.Data = newDataMap
	}

	if valueOk == dataOk {
		return fmt.Errorf(
			"either 'value' (present: %t) or 'data' (present: %t) must be provided",
			valueOk, dataOk)
	}

	if priority, ok := d.GetOkExists("priority"); ok {
		p := uint16(priority.(int))
		newRecord.Priority = &p
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		if ttl.(int) != 1 && proxiedOk && *newRecord.Proxied {
			return fmt.Errorf("error validating record %s: ttl must be set to 1 when `proxied` is true", newRecord.Name)
		}

		newRecord.TTL = ttl.(int)
	}

	// Validate value based on type
	if err := validateRecordName(newRecord.Type, newRecord.Content); err != nil {
		return fmt.Errorf("error validating record name %q: %s", newRecord.Name, err)
	}

	var proxiedVal *bool
	if proxiedOk {
		proxiedVal = newRecord.Proxied
	} else {
		proxiedVal = cloudflare.BoolPtr(false)
	}

	// Validate type
	if err := validateRecordType(newRecord.Type, *proxiedVal); err != nil {
		return fmt.Errorf("error validating record type %q: %s", newRecord.Type, err)
	}

	log.Printf("[DEBUG] Cloudflare Record create configuration: %#v", newRecord)

	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := client.CreateDNSRecord(context.Background(), newRecord.ZoneID, newRecord)
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				if d.Get("allow_overwrite").(bool) {
					var r cloudflare.DNSRecord
					log.Printf("[DEBUG] Cloudflare Record already exists however we are overwriting it")
					zone, _ := client.ZoneDetails(context.Background(), d.Get("zone_id").(string))
					if d.Get("name").(string) == "@" || d.Get("name").(string) == zone.Name {
						r = cloudflare.DNSRecord{
							Name: zone.Name,
							Type: d.Get("type").(string),
						}
					} else {
						r = cloudflare.DNSRecord{
							Name: d.Get("name").(string) + "." + zone.Name,
							Type: d.Get("type").(string),
						}
					}
					rs, _ := client.DNSRecords(context.Background(), d.Get("zone_id").(string), r)

					if len(rs) != 1 {
						return resource.RetryableError(fmt.Errorf("attempted to override existing record however didn't find an exact match"))
					}

					// Here we need to set the ID as the state will not have one and in order
					// for Terraform to operate on it, we need an anchor.
					d.SetId(rs[0].ID)

					if updateErr := resourceCloudflareRecordUpdate(d, meta); updateErr != nil {
						return resource.NonRetryableError(updateErr)
					}

					return nil
				}

				return resource.RetryableError(fmt.Errorf("expected DNS record to not already be present but already exists"))
			}

			return resource.NonRetryableError(fmt.Errorf("failed to create DNS record: %s", err))
		}

		// In the event that the API returns an empty DNS Record, we verify that the
		// ID returned is not the default ""
		if r.Result.ID == "" {
			return resource.NonRetryableError(fmt.Errorf("Failed to find record in Create response; Record was empty"))
		}

		d.SetId(r.Result.ID)

		resourceCloudflareRecordRead(d, meta)

		return nil
	})
}

func resourceCloudflareRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	record, err := client.DNSRecord(context.Background(), zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Invalid dns record identifier") ||
			strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing record from state because it's not found in API")
			d.SetId("")
			return nil
		}
		return err
	}

	data, dataOk := d.GetOk("data")
	log.Printf("[DEBUG] Data found in config: %#v", data)

	readDataMap := make(map[string]interface{})

	if dataOk {
		dataMap := data.([]interface{})[0]
		if dataMap != nil {
			for id, value := range dataMap.(map[string]interface{}) {
				newData, err := transformToCloudflareDNSData(record.Type, id, value)
				if err != nil {
					return err
				} else if newData == nil {
					continue
				}
				readDataMap[id] = newData
			}

			record.Data = []interface{}{readDataMap}
		}
	}

	d.SetId(record.ID)
	d.Set("hostname", record.Name)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("ttl", record.TTL)
	d.Set("proxied", record.Proxied)
	d.Set("created_on", record.CreatedOn.Format(time.RFC3339Nano))
	d.Set("data", record.Data)
	d.Set("modified_on", record.ModifiedOn.Format(time.RFC3339Nano))
	if err := d.Set("metadata", expandStringMap(record.Meta)); err != nil {
		log.Printf("[WARN] Error setting metadata: %s", err)
	}
	d.Set("proxiable", record.Proxiable)

	if record.Priority != nil {
		priority := record.Priority
		p := *priority
		d.Set("priority", int(p))
	}

	return nil
}

func resourceCloudflareRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	updateRecord := cloudflare.DNSRecord{
		ID:      d.Id(),
		Type:    d.Get("type").(string),
		Name:    d.Get("name").(string),
		Content: d.Get("value").(string),
		ZoneID:  zoneID,
	}

	data, dataOk := d.GetOk("data")
	log.Printf("[DEBUG] Data found in config: %#v", data)

	newDataMap := make(map[string]interface{})

	if dataOk {
		dataMap := data.([]interface{})[0]
		for id, value := range dataMap.(map[string]interface{}) {
			newData, err := transformToCloudflareDNSData(updateRecord.Type, id, value)
			if err != nil {
				return err
			} else if newData == nil {
				continue
			}
			newDataMap[id] = newData
		}

		updateRecord.Data = newDataMap
	}

	if priority, ok := d.GetOkExists("priority"); ok {
		p := uint16(priority.(int))
		updateRecord.Priority = &p
	}

	proxied, proxiedOk := d.GetOkExists("proxied")
	if proxiedOk {
		updateRecord.Proxied = cloudflare.BoolPtr(proxied.(bool))
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		if ttl.(int) != 1 && proxiedOk && *updateRecord.Proxied {
			return fmt.Errorf("error validating record %s: ttl must be set to 1 when `proxied` is true", updateRecord.Name)
		}

		updateRecord.TTL = ttl.(int)
	}

	log.Printf("[DEBUG] Cloudflare Record update configuration: %#v", updateRecord)

	return resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		err := client.UpdateDNSRecord(context.Background(), zoneID, d.Id(), updateRecord)
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				return resource.RetryableError(fmt.Errorf("expected DNS record to not already be present but already exists"))
			}

			return resource.NonRetryableError(fmt.Errorf("failed to create DNS record: %s", err))
		}

		resourceCloudflareRecordRead(d, meta)
		return nil
	})
}

func resourceCloudflareRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting Cloudflare Record: %s, %s", zoneID, d.Id())

	err := client.DeleteDNSRecord(context.Background(), zoneID, d.Id())
	if err != nil {
		return fmt.Errorf("error deleting Cloudflare Record: %s", err)
	}

	return nil
}

func expandStringMap(inVal interface{}) map[string]string {
	// although interface could hold anything
	// we assume that it is either nil or a map of interface values
	outVal := make(map[string]string)
	if inVal == nil {
		return outVal
	}
	for k, v := range inVal.(map[string]interface{}) {
		strValue := fmt.Sprintf("%v", v)
		outVal[k] = strValue
	}
	return outVal
}

func resourceCloudflareRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var recordID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		recordID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"zoneID/recordID\" for import", d.Id())
	}

	record, err := client.DNSRecord(context.Background(), zoneID, recordID)
	if err != nil {
		return nil, fmt.Errorf("Unable to find record with ID %q: %q", d.Id(), err)
	}

	log.Printf("[INFO] Found record: %s", record.Name)
	name := strings.TrimSuffix(record.Name, "."+record.ZoneName)

	d.Set("name", name)
	d.Set("zone_id", zoneID)
	d.SetId(recordID)

	resourceCloudflareRecordRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

var dnsTypeIntFields = []string{
	"algorithm",
	"key_tag",
	"type",
	"usage",
	"selector",
	"matching_type",
	"weight",
	"priority",
	"port",
	"long_degrees",
	"lat_degrees",
	"long_minutes",
	"lat_minutes",
	"protocol",
	"digest_type",
	"order",
	"preference",
}

var dnsTypeFloatFields = []string{
	"size",
	"altitude",
	"precision_horz",
	"precision_vert",
	"long_seconds",
	"lat_seconds",
}

func transformToCloudflareDNSData(recordType string, id string, value interface{}) (newValue interface{}, err error) {
	switch {
	case id == "flags":
		switch {
		case strings.ToUpper(recordType) == "SRV",
			strings.ToUpper(recordType) == "CAA",
			strings.ToUpper(recordType) == "DNSKEY":
			newValue, err = value.(string), nil
		case strings.ToUpper(recordType) == "NAPTR":
			newValue, err = value.(string), nil
		}
	case contains(dnsTypeIntFields, id):
		newValue, err = value, nil
	case contains(dnsTypeFloatFields, id):
		newValue, err = value, nil
	default:
		newValue, err = value.(string), nil
	}

	return
}

func suppressPriority(k, old, new string, d *schema.ResourceData) bool {
	recordType := d.Get("type").(string)
	if recordType != "MX" && recordType != "URI" {
		return true
	}
	return false
}
