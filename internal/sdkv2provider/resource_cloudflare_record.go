package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareRecordCreate,
		ReadContext:   resourceCloudflareRecordRead,
		UpdateContext: resourceCloudflareRecordUpdate,
		DeleteContext: resourceCloudflareRecordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareRecordImport,
		},
		Description:   heredoc.Doc(`Provides a Cloudflare record resource.`),
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

func resourceCloudflareRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newRecord := cloudflare.CreateDNSRecordParams{
		Type:   d.Get("type").(string),
		Name:   d.Get("name").(string),
		ZoneID: d.Get(consts.ZoneIDSchemaKey).(string),
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
	tflog.Debug(ctx, fmt.Sprintf("Data found in config: %#v", data))

	newDataMap := make(map[string]interface{})

	if dataOk {
		dataMap := data.([]interface{})[0]
		for id, value := range dataMap.(map[string]interface{}) {
			newData, err := transformToCloudflareDNSData(newRecord.Type, id, value)
			if err != nil {
				return diag.FromErr(err)
			} else if newData == nil {
				continue
			}
			newDataMap[id] = newData
		}

		newRecord.Data = newDataMap
	}

	if valueOk == dataOk {
		return diag.FromErr(fmt.Errorf(
			"either 'value' (present: %t) or 'data' (present: %t) must be provided",
			valueOk, dataOk))
	}

	if priority, ok := d.GetOkExists("priority"); ok {
		p := uint16(priority.(int))
		newRecord.Priority = &p
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		if ttl.(int) != 1 && proxiedOk && *newRecord.Proxied {
			return diag.FromErr(fmt.Errorf("error validating record %s: ttl must be set to 1 when `proxied` is true", newRecord.Name))
		}

		newRecord.TTL = ttl.(int)
	}

	if newRecord.Name == "" {
		return diag.FromErr(fmt.Errorf("record on zone %s must not have an empty name (use @ for the zone apex)", newRecord.ZoneID))
	}

	// Validate value based on type
	if err := validateRecordContent(newRecord.Type, newRecord.Content); err != nil {
		return diag.FromErr(fmt.Errorf("error validating record content of %q: %w", newRecord.Name, err))
	}

	var proxiedVal *bool
	if proxiedOk {
		proxiedVal = newRecord.Proxied
	} else {
		proxiedVal = cloudflare.BoolPtr(false)
	}

	if comment, ok := d.GetOk("comment"); ok {
		newRecord.Comment = comment.(string)
	}

	if tags, ok := d.GetOk("tags"); ok {
		for _, tag := range tags.(*schema.Set).List() {
			newRecord.Tags = append(newRecord.Tags, tag.(string))
		}
	}

	// Validate type
	if err := validateRecordType(newRecord.Type, *proxiedVal); err != nil {
		return diag.FromErr(fmt.Errorf("error validating record type %q: %w", newRecord.Type, err))
	}

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Record create configuration: %#v", newRecord))

	retry := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := client.CreateDNSRecord(ctx, cloudflare.ZoneIdentifier(newRecord.ZoneID), newRecord)
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				if d.Get("allow_overwrite").(bool) {
					var r cloudflare.ListDNSRecordsParams
					tflog.Debug(ctx, fmt.Sprintf("Cloudflare Record already exists however we are overwriting it"))
					zone, _ := client.ZoneDetails(ctx, d.Get(consts.ZoneIDSchemaKey).(string))
					if d.Get("name").(string) == "@" || d.Get("name").(string) == zone.Name {
						r = cloudflare.ListDNSRecordsParams{
							Name: zone.Name,
							Type: d.Get("type").(string),
						}
					} else {
						r = cloudflare.ListDNSRecordsParams{
							Name: d.Get("name").(string) + "." + zone.Name,
							Type: d.Get("type").(string),
						}
					}
					rs, _, _ := client.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(d.Get(consts.ZoneIDSchemaKey).(string)), r)

					if len(rs) != 1 {
						return resource.RetryableError(fmt.Errorf("attempted to override existing record however didn't find an exact match"))
					}

					// Here we need to set the ID as the state will not have one and in order
					// for Terraform to operate on it, we need an anchor.
					d.SetId(rs[0].ID)

					if updateErr := resourceCloudflareRecordUpdate(ctx, d, meta); updateErr != nil {
						return resource.NonRetryableError(errors.New("failed to update record"))
					}

					return nil
				}

				return resource.RetryableError(fmt.Errorf("expected DNS record to not already be present but already exists"))
			}

			return resource.NonRetryableError(fmt.Errorf("failed to create DNS record: %w", err))
		}

		// In the event that the API returns an empty DNS Record, we verify that the
		// ID returned is not the default ""
		if r.Result.ID == "" {
			return resource.NonRetryableError(fmt.Errorf("failed to find record in Create response; Record was empty"))
		}

		d.SetId(r.Result.ID)

		resourceCloudflareRecordRead(ctx, d, meta)

		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	record, err := client.GetDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing record from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Data found in config: %#v", record.Data))

	readDataMap := make(map[string]interface{})

	if record.Data != nil {
		dataMap := record.Data.(map[string]interface{})
		if dataMap != nil {
			for id, value := range dataMap {
				newData, err := transformToCloudflareDNSData(record.Type, id, value)
				if err != nil {
					return diag.FromErr(err)
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
		tflog.Warn(ctx, fmt.Sprintf("Error setting metadata: %s", err))
	}
	d.Set("proxiable", record.Proxiable)
	d.Set("comment", record.Comment)
	d.Set("tags", record.Tags)

	if record.Priority != nil {
		priority := record.Priority
		p := *priority
		d.Set("priority", int(p))
	}

	return nil
}

func resourceCloudflareRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	updateRecord := cloudflare.UpdateDNSRecordParams{
		ID:      d.Id(),
		Type:    d.Get("type").(string),
		Name:    d.Get("name").(string),
		Content: d.Get("value").(string),
	}

	data, dataOk := d.GetOk("data")
	tflog.Debug(ctx, fmt.Sprintf("Data found in config: %#v", data))

	newDataMap := make(map[string]interface{})

	if dataOk {
		dataMap := data.([]interface{})[0]
		for id, value := range dataMap.(map[string]interface{}) {
			newData, err := transformToCloudflareDNSData(updateRecord.Type, id, value)
			if err != nil {
				return diag.FromErr(err)
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
			return diag.FromErr(fmt.Errorf("error validating record %s: ttl must be set to 1 when `proxied` is true", updateRecord.Name))
		}

		updateRecord.TTL = ttl.(int)
	}

	if comment, ok := d.GetOk("comment"); ok {
		updateRecord.Comment = comment.(string)
	}

	tags := []string{}
	for _, tag := range d.Get("tags").(*schema.Set).List() {
		tags = append(tags, tag.(string))
	}
	updateRecord.Tags = tags

	tflog.Debug(ctx, fmt.Sprintf("Cloudflare Record update configuration: %#v", updateRecord))

	retry := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		updateRecord.ID = d.Id()
		err := client.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), updateRecord)
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				return resource.RetryableError(fmt.Errorf("expected DNS record to not already be present but already exists"))
			}

			return resource.NonRetryableError(fmt.Errorf("failed to create DNS record: %w", err))
		}

		resourceCloudflareRecordRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Record: %s, %s", zoneID, d.Id()))

	err := client.DeleteDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Record: %w", err))
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

func resourceCloudflareRecordImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can look up
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var recordID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		recordID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"zoneID/recordID\" for import", d.Id())
	}

	record, err := client.GetDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), recordID)
	if err != nil {
		return nil, fmt.Errorf("Unable to find record with ID %q: %w", d.Id(), err)
	}

	tflog.Info(ctx, fmt.Sprintf("Found record: %s", record.Name))
	name := strings.TrimSuffix(record.Name, "."+record.ZoneName)

	d.Set("name", name)
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(recordID)

	resourceCloudflareRecordRead(ctx, d, meta)

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
			strings.ToUpper(recordType) == "DNSKEY":
			newValue, err = value.(string), nil
		case strings.ToUpper(recordType) == "NAPTR":
			newValue, err = value.(string), nil
		case strings.ToUpper(recordType) == "CAA":
			// this is required because "flags" is shared however, it comes from
			// the API as a float64 but the Terraform internal type is string 😢.
			switch value.(type) {
			case float64:
				newValue, err = fmt.Sprintf("%.0f", value.(float64)), nil
			case string:
				newValue, err = value.(string), nil
			}
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

func suppressTrailingDots(k, old, new string, d *schema.ResourceData) bool {
	newTrimmed := strings.TrimSuffix(new, ".")

	// Ensure to distinguish values consists of dots only.
	if newTrimmed == "" {
		return old == new
	}

	return strings.TrimSuffix(old, ".") == newTrimmed
}
