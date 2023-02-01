package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareMTLSCertificate() *schema.Resource {
	return &schema.Resource{
		// You cannot edit mTLS certificates, rather, only upload new ones.
		CreateContext: resourceCloudflareMTLSCertificateCreate,
		ReadContext:   resourceCloudflareMTLSCertificateRead,
		DeleteContext: resourceCloudflareMTLSCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareMTLSCertificateImport,
		},

		Schema: resourceCloudflareMTLSCertificateSchema(),
		Description: heredoc.Doc(`
			Provides a Cloudflare mTLS certificate resource. These certificates may be used with mTLS enabled Cloudflare services.
		`),
	}
}

func resourceCloudflareMTLSCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	accountIDrc := cloudflare.AccountIdentifier(accountID)

	certificate := cloudflare.CreateMTLSCertificateParams{
		Name:         d.Get("name").(string),
		Certificates: d.Get("certificates").(string),
		PrivateKey:   d.Get("private_key").(string),
		CA:           d.Get("ca").(bool),
	}

	newCertificate, err := client.CreateMTLSCertificate(ctx, accountIDrc, certificate)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to create mtls certificate"))
	}

	d.SetId(newCertificate.ID)

	return resourceCloudflareMTLSCertificateRead(ctx, d, meta)
}

func resourceCloudflareMTLSCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	accountIDrc := cloudflare.AccountIdentifier(accountID)
	certID := d.Id()

	record, err := client.GetMTLSCertificate(ctx, accountIDrc, certID)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("mTLS certificate %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding mTLS certificate %q: %w", d.Id(), err))
	}
	d.Set("name", record.Name)
	d.Set("issuer", record.Issuer)
	d.Set("signature", record.Signature)
	d.Set("serial_number", record.SerialNumber)
	d.Set("uploaded_on", record.UploadedOn.Format(time.RFC3339Nano))
	d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))

	return nil
}

func resourceCloudflareMTLSCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	accountIDrc := cloudflare.AccountIdentifier(accountID)
	certID := d.Id()

	_, err := client.DeleteMTLSCertificate(ctx, accountIDrc, certID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting mTLS certificate in account %q: %w", accountID, err))
	}
	return nil
}

func resourceCloudflareMTLSCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/certID\"", d.Id())
	}
	accountID, certID := idAttr[0], idAttr[1]
	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(certID)

	resourceCloudflareMTLSCertificateRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
