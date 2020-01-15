package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomHostnames() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareCustomHostnamesCreate,
		Read:   resourceCloudflareCustomHostnamesRead,
		Update: resourceCloudflareCustomHostnamesUpdate,
		Delete: resourceCloudflareCustomHostnamesDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareCustomHostnamesImport,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_origin_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"http", "email", "cname"}, false),
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"dv"}, false),
						},
						"custom_hostname_settings": {
							Type:     schema.TypeList,
							MaxItems: 4,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http2": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									},
									"tls13": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									},
									"mintlsversion": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"1.0", "1.1", "1.2", "1.3"}, false),
									},
									"ciphers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareCustomHostnamesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)
	zcso := expandToZoneCustomHostnamesOptions(d)
	res, err := client.CreateCustomHostname(zoneID, zcso)
	if err != nil {
		return fmt.Errorf("Failed to create custom hostname cert: %s", err)
	}

	if res.Result.ID == "" {
		return fmt.Errorf("Failed to find custom ssl in Create response: id was empty")
	}

	d.SetId(res.Result.ID)

	log.Printf("[INFO] Cloudflare Custom SSL ID: %s", d.Id())

	return resourceCloudflareCustomHostnamesRead(d, meta)
}

func resourceCloudflareCustomHostnamesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()
	zcso := expandToZoneCustomHostnamesOptionsSSL(d)

	log.Printf("[INFO] Updating Custom Hostname from struct: %+v", zcso)

	_, err := client.UpdateCustomHostnameSSL(zoneID, certID, zcso)
	if err != nil {
		return errors.Wrap(err, "Error updating Custom Hostname")
	}

	return resourceCloudflareCustomHostnamesRead(d, meta)
}

func resourceCloudflareCustomHostnamesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	// update all possible schema attributes with fields from api response
	record, err := client.CustomHostname(zoneID, certID)
	if err != nil {
		log.Printf("[WARN] Removing record from state because it's not found in API")
		d.SetId("")
		return nil
	}
	d.Set("hostname", record.Hostname)
	d.Set("custom_origin_server", record.CustomOriginServer)
	return nil
}

func resourceCloudflareCustomHostnamesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	log.Printf("[DEBUG] Deleting Custom Hostname cert %s for zone %s", certID, zoneID)

	err := client.DeleteCustomHostname(zoneID, certID)
	if err != nil {
		errors.Wrap(err, "failed to delete custom hostanme cert setting")
	}
	return nil
}

func resourceCloudflareCustomHostnamesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/certID\"", d.Id())
	}

	zoneID, certID := idAttr[0], idAttr[1]

	log.Printf("[DEBUG] Importing Cloudflare Custom SSL Cert: id %s for zone %s", certID, zoneID)

	d.Set("zone_id", zoneID)
	d.SetId(certID)

	resourceCloudflareCustomHostnamesRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func expandToZoneCustomHostnamesOptions(d *schema.ResourceData) (options cloudflare.CustomHostname) {

	options = cloudflare.CustomHostname{
		Hostname:           d.Get("hostname").(string),
		CustomOriginServer: d.Get("custom_origin_server").(string),
	}

	options.SSL = expandToZoneCustomHostnamesOptionsSSL(d)

	return options
}

func expandToZoneCustomHostnamesOptionsSSL(d *schema.ResourceData) (options cloudflare.CustomHostnameSSL) {
	v, ok := d.GetOk("ssl")
	if !ok {
		return
	}
	cfg := v.([]interface{})[0].(map[string]interface{})

	sslOpt := cloudflare.CustomHostnameSSL{
		Method: cfg["method"].(string),
		Type:   cfg["type"].(string),
		Status: cfg["status"].(string),
	}

	if hostSettIface, ok := cfg["custom_hostname_settings"]; ok && len(hostSettIface.([]interface{})) > 0 {
		hostSett := hostSettIface.([]interface{})[0].(map[string]interface{})

		hostOptions := cloudflare.CustomHostnameSSLSettings{
			HTTP2:         hostSett["http2"].(string),
			TLS13:         hostSett["tls13"].(string),
			MinTLSVersion: hostSett["mintlsversion"].(string),
			Ciphers:       expandInterfaceToStringList(hostSett["ciphers"]),
		}
		sslOpt.Settings = hostOptions
	}
	return sslOpt
}
