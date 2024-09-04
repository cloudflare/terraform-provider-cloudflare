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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsCertificate() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsCertificateSchema(),
		CreateContext: resourceCloudflareTeamsCertificateCreate,
		UpdateContext: resourceCloudflareTeamsCertificateUpdate,
		ReadContext:   resourceCloudflareTeamsCertificateRead,
		DeleteContext: resourceCloudflareTeamsCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsCertificateImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Teams Gateway Certificate resource. A Teams Certificate can
			be specified for Gateway TLS interception and block pages.
		`),
	}
}

func resourceCloudflareTeamsCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	if d.Get("gateway_managed").(bool) {
		newTeamsCertificate := cloudflare.TeamsCertificateCreateRequest{
			ValidityPeriodDays: d.Get("validity_period_days").(int),
		}

		tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams Certificate from struct: %+v", newTeamsCertificate))

		certificate, err := client.TeamsGenerateCertificate(ctx, accountID, newTeamsCertificate)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating Teams Certificate for account %q: %w", accountID, err))
		}
		d.SetId(certificate.ID)
	}

	if d.Get("activate").(bool) {
		certificate, err := client.TeamsActivateCertificate(ctx, accountID, d.Get("id").(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error activating Teams Certificate with id %q for account %q: %w", d.Get("id"), accountID, err))
		}
		d.Set("binding_status", certificate.BindingStatus)
		d.Set("qs_pack_id", certificate.QsPackId)
	}

	return resourceCloudflareTeamsCertificateRead(ctx, d, meta)
}

func resourceCloudflareTeamsCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	if d.Get("activate").(bool) {
		certificate, err := client.TeamsActivateCertificate(ctx, accountID, d.Get("id").(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error activating Teams Certificate with id %q for account %q: %w", d.Get("id"), accountID, err))
		}
		d.Set("binding_status", certificate.BindingStatus)
		d.Set("qs_pack_id", certificate.QsPackId)
	} else {
		certificate, err := client.TeamsDeactivateCertificate(ctx, accountID, d.Get("id").(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error deactivating Teams Certificate with id %q for account %q: %w", d.Get("id"), accountID, err))
		}
		d.Set("binding_status", certificate.BindingStatus)
		d.Set("qs_pack_id", certificate.QsPackId)
	}

	return resourceCloudflareTeamsCertificateRead(ctx, d, meta)
}

func resourceCloudflareTeamsCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	Certificate, err := client.TeamsCertificate(ctx, accountID, d.Id())

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Teams Certificate %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams Certificate %q: %w", d.Id(), err))
	}

	d.Set("id", Certificate.ID)
	d.Set("in_use", Certificate.InUse)
	d.Set("binding_status", Certificate.BindingStatus)
	d.Set("qs_pack_id", Certificate.QsPackId)
	certType := Certificate.Type
	if certType == "gateway_managed" {
		d.Set("gateway_managed", true)
	} else if certType == "custom" {
		d.Set("custom", true)
	} else {
		return diag.FromErr(fmt.Errorf("error reading Teams Certificate type %q: %w", certType, err))
	}
	if Certificate.UploadedOn != nil {
		d.Set("uploaded_on", Certificate.UploadedOn.Format(time.RFC3339Nano))
	}
	if Certificate.CreatedAt != nil {
		d.Set("created_at", Certificate.CreatedAt.Format(time.RFC3339Nano))
	}
	if Certificate.ExpiresOn != nil {
		d.Set("expires_on", Certificate.ExpiresOn.Format(time.RFC3339Nano))
	}

	return nil
}

func resourceCloudflareTeamsCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	certID := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Teams Certificate using ID: %s", certID))

	if d.Get("custom").(bool) {
		return diag.FromErr(fmt.Errorf("error deleting Teams Certificate with id %q: custom certificates must be deleted via the mTLS certificate manager api", d.Get("id")))
	}

	if d.Get("activate").(bool) {
		return diag.FromErr(fmt.Errorf("error deleting Teams Certificate with id %q: certificate must be deactivated before it can be deleted", d.Get("id")))
	} else {
		_, err := client.TeamsDeactivateCertificate(ctx, accountID, d.Get("id").(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error deactivating Teams Certificate with id %q for account %q: %w", d.Get("id"), accountID, err))
		}
	}

	err := client.TeamsDeleteCertificate(ctx, accountID, certID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Teams Certificate for account %q: %w", accountID, err))
	}

	resourceCloudflareTeamsCertificateRead(ctx, d, meta)

	return nil
}

func resourceCloudflareTeamsCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsCertificateID\"", d.Id())
	}

	accountID, teamsCertificateID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Teams Certificate: id %s for account %s", teamsCertificateID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(teamsCertificateID)

	resourceCloudflareTeamsCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
