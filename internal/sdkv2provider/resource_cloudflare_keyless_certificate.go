package sdkv2provider

import (
    "context"
    "errors"
    "fmt"
    "strings"

    cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareKeylessCertificate() *schema.Resource {
    return &schema.Resource{
        Schema: resourceCloudflareKeylessCertificateSchema
        CreateContext: resourceCloudflareKeylessCertificateCreate,
        ReadContext:   resourceCloudflareKeylessCertificateRead,
        UpdateContext: resourceCloudflareKeylessCertificateUpdate,
        DeleteContext: resourceCloudflareKeylessCertificateDelete,
       		Importer: &schema.ResourceImporter{
       			StateContext: resourceCloudflareKeylessCertificateImport,
       		},
       		Description: heredoc.Doc(`
       			Provides a resource, that manages Keyless certificates.
       		`),
    }
}

func resourceCloudflareKeylessCertificateCreate(d *schema.ResourceData, m interface{}) error {
   	client := meta.(*cloudflare.API)
   	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

   	request :=  cloudflare.KeylessSSLCreateRequest{
                  		Name: d.Get("name").(string),
                  		Host: d.Get("host").(string),
                  		Port: d.Get("port").(int),
                  		Certificate: d.Get("certificate").(string),
                  		Port: d.Get("bundle_method").(string),
                  	}

   	event, err := client.CreateKeylessSSL(ctx, zoneID, request)
   	if err != nil {
   		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create Keyless SSL")))
   	}

   	d.SetId(event.Result.ID)

	return resourceCloudflareKeylessCertificateRead(ctx, d, meta)
}

func resourceCloudflareKeylessCertificateRead(d *schema.ResourceData, m interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	request, err := client.KeylessSSL(ctx, zoneID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Keyless SSL %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Keyless SSL %q: %w", d.Id(), err))
	}

	d.Set("name", request.Result.Name)
	d.Set("host", request.Result.Host)
	d.Set("port", request.Result.Port)
	d.Set("status", request.Result.Status)
	d.Set("enabled", request.Result.Enabled)
	d.Set("port", request.Result.Port)
	return nil
}

func resourceCloudflareKeylessCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

   	request :=  cloudflare.KeylessSSLUpdateRequest{
                  		Name: d.Get("name").(string),
                  		Host: d.Get("host").(string),
                  		Port: d.Get("port").(int),
                  		Enabled: d.Get("enabled").(bool),
                  	}

	_, err := client.UpdateKeylessSSL(ctx, zoneID, d.Id(), request)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to update Keyless SSL")))
	}

	return resourceCloudflareKeylessCertificateRead(ctx, d, meta)
}

func resourceCloudflareKeylessCertificateDelete(d *schema.ResourceData, m interface{}) error {
 	client := meta.(*cloudflare.API)
 	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

 	_, err := client.DeleteKeylessSSL(ctx, zoneID, d.Id())
 	if err != nil {
 		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to delete Keyless SSL")))
 	}

	return resourceCloudflareKeylessCertificateRead(ctx, d, meta)
}

func resourceCloudflareKeylessCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)
	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/KeylessSSLID\"", d.Id())
	}

	zoneID, keylessSslId := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Keyless SSL: id %s for zone %s", keylessSslId, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(keylessSslId)

	resourceCloudflareKeylessCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}