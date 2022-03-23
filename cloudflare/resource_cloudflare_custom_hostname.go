package cloudflare

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomHostname() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareCustomHostnameSchema(),
		Create: resourceCloudflareCustomHostnameCreate,
		Read:   resourceCloudflareCustomHostnameRead,
		Update: resourceCloudflareCustomHostnameUpdate,
		Delete: resourceCloudflareCustomHostnameDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomHostnameImport,
		},
	}
}

func resourceCloudflareCustomHostnameRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostnameID := d.Id()

	customHostname, err := client.CustomHostname(context.Background(), zoneID, hostnameID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error reading custom hostname %q", hostnameID))
	}

	d.Set("hostname", customHostname.Hostname)
	d.Set("custom_origin_server", customHostname.CustomOriginServer)
	d.Set("custom_origin_sni", customHostname.CustomOriginSNI)
	var sslConfig []map[string]interface{}

	if !reflect.ValueOf(customHostname.SSL).IsNil() {
		ssl := map[string]interface{}{
			"type":                  customHostname.SSL.Type,
			"method":                customHostname.SSL.Method,
			"wildcard":              customHostname.SSL.Wildcard,
			"status":                customHostname.SSL.Status,
			"certificate_authority": customHostname.SSL.CertificateAuthority,
			"custom_certificate":    customHostname.SSL.CustomCertificate,
			"custom_key":            customHostname.SSL.CustomKey,
			"settings": []map[string]interface{}{{
				"http2":           customHostname.SSL.Settings.HTTP2,
				"tls13":           customHostname.SSL.Settings.TLS13,
				"min_tls_version": customHostname.SSL.Settings.MinTLSVersion,
				"ciphers":         customHostname.SSL.Settings.Ciphers,
				"early_hints":     customHostname.SSL.Settings.EarlyHints,
			}},
		}
		if !reflect.ValueOf(customHostname.SSL.ValidationErrors).IsNil() {
			errors := []map[string]interface{}{}
			for _, e := range customHostname.SSL.ValidationErrors {
				errors = append(errors, map[string]interface{}{"message": e.Message})
			}
			ssl["validation_errors"] = errors
		}
		if !reflect.ValueOf(customHostname.SSL.ValidationRecords).IsNil() {
			records := []map[string]interface{}{}
			for _, e := range customHostname.SSL.ValidationRecords {
				records = append(records,
					map[string]interface{}{
						"cname_name":   e.CnameName,
						"cname_target": e.CnameTarget,
						"txt_name":     e.TxtName,
						"txt_value":    e.TxtValue,
						"http_body":    e.HTTPBody,
						"http_url":     e.HTTPUrl,
						"emails":       e.Emails,
					})
			}
			ssl["validation_records"] = records
		}
		sslConfig = append(sslConfig, ssl)
	}

	if err := d.Set("ssl", sslConfig); err != nil {
		return fmt.Errorf("failed to set ssl")
	}

	ownershipVerificationCfg := map[string]interface{}{
		"type":  customHostname.OwnershipVerification.Type,
		"value": customHostname.OwnershipVerification.Value,
		"name":  customHostname.OwnershipVerification.Name,
	}
	if err := d.Set("ownership_verification", ownershipVerificationCfg); err != nil {
		return fmt.Errorf("failed to set ownership_verification: %v", err)
	}

	ownershipVerificationHTTPCfg := map[string]interface{}{
		"http_body": customHostname.OwnershipVerificationHTTP.HTTPBody,
		"http_url":  customHostname.OwnershipVerificationHTTP.HTTPUrl,
	}
	if err := d.Set("ownership_verification_http", ownershipVerificationHTTPCfg); err != nil {
		return fmt.Errorf("failed to set ownership_verification_http: %v", err)
	}

	return nil
}

func resourceCloudflareCustomHostnameDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostnameID := d.Id()

	err := client.DeleteCustomHostname(context.Background(), zoneID, hostnameID)
	if err != nil {
		return errors.Wrap(err, "failed to delete custom hostname certificate")
	}

	return nil
}

func resourceCloudflareCustomHostnameCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	certificate := buildCustomHostname(d)

	newCertificate, err := client.CreateCustomHostname(context.Background(), zoneID, certificate)
	if err != nil {
		return errors.Wrap(err, "failed to create custom hostname certificate")
	}

	d.SetId(newCertificate.Result.ID)

	return resourceCloudflareCustomHostnameRead(d, meta)
}

func resourceCloudflareCustomHostnameUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostnameID := d.Id()
	certificate := buildCustomHostname(d)

	_, err := client.UpdateCustomHostname(context.Background(), zoneID, hostnameID, certificate)
	if err != nil {
		return errors.Wrap(err, "failed to update custom hostname certificate")
	}

	return resourceCloudflareCustomHostnameRead(d, meta)
}

func resourceCloudflareCustomHostnameImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/customHostnameID\"", d.Id())
	}

	zoneID, hostnameID := idAttr[0], idAttr[1]

	log.Printf("[DEBUG] Importing Cloudflare Custom Hostname: id %s for zone %s", hostnameID, zoneID)

	d.Set("zone_id", zoneID)
	d.SetId(hostnameID)

	return []*schema.ResourceData{d}, nil
}

// buildCustomHostname takes the existing schema and returns a
// `cloudflare.CustomHostname`.
func buildCustomHostname(d *schema.ResourceData) cloudflare.CustomHostname {
	ch := cloudflare.CustomHostname{
		Hostname:           d.Get("hostname").(string),
		CustomOriginServer: d.Get("custom_origin_server").(string),
		CustomOriginSNI:    d.Get("custom_origin_sni").(string),
	}

	if _, ok := d.GetOk("ssl"); ok {
		ch.SSL = &cloudflare.CustomHostnameSSL{
			Method:               d.Get("ssl.0.method").(string),
			Type:                 d.Get("ssl.0.type").(string),
			Wildcard:             cloudflare.BoolPtr(d.Get("ssl.0.wildcard").(bool)),
			CustomCertificate:    d.Get("ssl.0.custom_certificate").(string),
			CustomKey:            d.Get("ssl.0.custom_key").(string),
			CertificateAuthority: d.Get("ssl.0.certificate_authority").(string),
			Settings: cloudflare.CustomHostnameSSLSettings{
				HTTP2:         d.Get("ssl.0.settings.0.http2").(string),
				TLS13:         d.Get("ssl.0.settings.0.tls13").(string),
				MinTLSVersion: d.Get("ssl.0.settings.0.min_tls_version").(string),
				Ciphers:       expandInterfaceToStringList(d.Get("ssl.0.settings.0.ciphers").(*schema.Set).List()),
				EarlyHints:    d.Get("ssl.0.settings.0.early_hints").(string),
			},
		}
	}

	return ch
}
