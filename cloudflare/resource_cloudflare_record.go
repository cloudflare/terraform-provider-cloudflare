package cloudflare

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCloudFlareRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareRecordCreate,
		Read:   resourceCloudFlareRecordRead,
		Update: resourceCloudFlareRecordUpdate,
		Delete: resourceCloudFlareRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudFlareRecordImport,
		},

		SchemaVersion: 1,
		MigrateState:  resourceCloudFlareRecordMigrateState,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(i interface{}) string {
					return strings.ToLower(i.(string))
				},
			},

			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"data"},
			},

			"data": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"value"},
			},

			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"proxied": {
				Default:  false,
				Optional: true,
				Type:     schema.TypeBool,
			},

			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
			},

			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"proxiable": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudFlareRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newRecord := cloudflare.DNSRecord{
		Type:     d.Get("type").(string),
		Name:     d.Get("name").(string),
		Proxied:  d.Get("proxied").(bool),
		ZoneName: d.Get("domain").(string),
	}

	value, valueOk := d.GetOk("value")
	if valueOk {
		newRecord.Content = value.(string)
	}

	data, dataOk := d.GetOk("data")
	if dataOk {
		newRecord.Data = data
	}

	if valueOk == dataOk {
		return fmt.Errorf(
			"either 'value' (present: %t) or 'data' (present: %t) must be provided",
			valueOk, dataOk)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newRecord.Priority = priority.(int)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newRecord.TTL = ttl.(int)
	}

	// Validate value based on type
	if err := validateRecordName(newRecord.Type, newRecord.Content); err != nil {
		return fmt.Errorf("Error validating record name %q: %s", newRecord.Name, err)
	}

	// Validate type
	if err := validateRecordType(newRecord.Type, newRecord.Proxied); err != nil {
		return fmt.Errorf("Error validating record type %q: %s", newRecord.Type, err)
	}

	zoneID, err := client.ZoneIDByName(newRecord.ZoneName)
	if err != nil {
		return fmt.Errorf("Error finding zone %q: %s", newRecord.ZoneName, err)
	}

	d.Set("zone_id", zoneID)
	newRecord.ZoneID = zoneID

	log.Printf("[DEBUG] CloudFlare Record create configuration: %#v", newRecord)

	r, err := client.CreateDNSRecord(zoneID, newRecord)
	if err != nil {
		return fmt.Errorf("Failed to create record: %s", err)
	}

	// In the Event that the API returns an empty DNS Record, we verify that the
	// ID returned is not the default ""
	if r.Result.ID == "" {
		return fmt.Errorf("Failed to find record in Create response; Record was empty")
	}

	d.SetId(r.Result.ID)

	log.Printf("[INFO] CloudFlare Record ID: %s", d.Id())

	return resourceCloudFlareRecordRead(d, meta)
}

func resourceCloudFlareRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	record, err := client.DNSRecord(zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "Invalid dns record identifier") ||
			strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing record from state because it's not found in API")
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(record.ID)
	d.Set("hostname", record.Name)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("ttl", record.TTL)
	d.Set("priority", record.Priority)
	d.Set("proxied", record.Proxied)
	d.Set("created_on", record.CreatedOn.Format(time.RFC3339Nano))
	d.Set("data", expandStringMap(record.Data))
	d.Set("modified_on", record.ModifiedOn.Format(time.RFC3339Nano))
	if err := d.Set("metadata", expandStringMap(record.Meta)); err != nil {
		log.Printf("[WARN] Error setting metadata: %s", err)
	}
	d.Set("proxiable", record.Proxiable)

	return nil
}

func resourceCloudFlareRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	updateRecord := cloudflare.DNSRecord{
		ID:       d.Id(),
		Type:     d.Get("type").(string),
		Name:     d.Get("name").(string),
		Content:  d.Get("value").(string),
		ZoneName: d.Get("domain").(string),
		ZoneID:   zoneID,
		Proxied:  false,
	}

	if priority, ok := d.GetOk("priority"); ok {
		updateRecord.Priority = priority.(int)
	}

	if proxied, ok := d.GetOk("proxied"); ok {
		updateRecord.Proxied = proxied.(bool)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		updateRecord.TTL = ttl.(int)
	}

	log.Printf("[DEBUG] CloudFlare Record update configuration: %#v", updateRecord)
	err := client.UpdateDNSRecord(zoneID, d.Id(), updateRecord)
	if err != nil {
		return fmt.Errorf("Failed to update CloudFlare Record: %s", err)
	}

	return resourceCloudFlareRecordRead(d, meta)
}

func resourceCloudFlareRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	log.Printf("[INFO] Deleting CloudFlare Record: %s, %s", zoneID, d.Id())

	err := client.DeleteDNSRecord(zoneID, d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting CloudFlare Record: %s", err)
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

func resourceCloudFlareRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneName string
	var recordId string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		recordId = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"zoneName/recordId\" for import", d.Id())
	}

	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}

	record, err := client.DNSRecord(zoneId, recordId)
	if err != nil {
		return nil, fmt.Errorf("Unable to find record with ID %q: %q", d.Id(), err)
	}

	log.Printf("[INFO] Found record: %s", record.Name)
	name := strings.TrimSuffix(record.Name, "."+zoneName)

	d.Set("name", name)
	d.Set("domain", zoneName)
	d.Set("zone_id", zoneId)
	d.SetId(recordId)

	return []*schema.ResourceData{d}, nil
}
