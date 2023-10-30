package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type contextKey int

const orgAccessImportCtxKey contextKey = iota

func resourceCloudflareAccessOrganization() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessOrganizationSchema(),
		CreateContext: resourceCloudflareAccessOrganizationCreate,
		ReadContext:   resourceCloudflareAccessOrganizationRead,
		UpdateContext: resourceCloudflareAccessOrganizationUpdate,
		DeleteContext: resourceCloudflareAccessOrganizationNoop,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessOrganizationImport,
		},
		Description: heredoc.Doc(`
			A Zero Trust organization defines the user login experience.
		`),
	}
}

func resourceCloudflareAccessOrganizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf("access organizations cannot be created and must be imported"))
}

func resourceCloudflareAccessOrganizationNoop(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceCloudflareAccessOrganizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	organization, _, err := client.GetAccessOrganization(ctx, identifier, cloudflare.GetAccessOrganizationParams{})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching access organization: %w", err))
	}

	d.Set("name", organization.Name)
	d.Set("auth_domain", organization.AuthDomain)
	d.Set("is_ui_read_only", organization.IsUIReadOnly)
	d.Set("ui_read_only_toggle_reason", organization.UIReadOnlyToggleReason)
	d.Set("user_seat_expiration_inactive_time", organization.UserSeatExpirationInactiveTime)
	d.Set("auto_redirect_to_identity", organization.AutoRedirectToIdentity)
	d.Set("session_duration", organization.SessionDuration)

	loginDesign := convertLoginDesignStructToSchema(ctx, d, &organization.LoginDesign)
	if loginDesignErr := d.Set("login_design", loginDesign); loginDesignErr != nil {
		return diag.FromErr(fmt.Errorf("error setting Access Organization Login Design configuration: %w", loginDesignErr))
	}

	if &organization.CustomPages != nil {
		customPages := convertCustomPageStructToSchema(ctx, d, &organization.CustomPages)
		if customPagesErr := d.Set("custom_pages", customPages); customPagesErr != nil {
			return diag.FromErr(fmt.Errorf("error setting Access Organization Custom Pages configuration: %w", customPagesErr))
		}
	}

	return nil
}

func resourceCloudflareAccessOrganizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	updatedAccessOrganization := cloudflare.UpdateAccessOrganizationParams{
		Name:                           d.Get("name").(string),
		AuthDomain:                     d.Get("auth_domain").(string),
		IsUIReadOnly:                   cloudflare.BoolPtr(d.Get("is_ui_read_only").(bool)),
		UserSeatExpirationInactiveTime: d.Get("user_seat_expiration_inactive_time").(string),
		AutoRedirectToIdentity:         cloudflare.BoolPtr(d.Get("auto_redirect_to_identity").(bool)),
		SessionDuration:                cloudflare.StringPtr(d.Get("session_duration").(string)),
	}
	loginDesign := convertLoginDesignSchemaToStruct(d)
	updatedAccessOrganization.LoginDesign = *loginDesign

	if _, ok := d.GetOk("custom_pages"); ok {
		customPagesStruct := convertCustomPageSchemaToStruct(d)
		updatedAccessOrganization.CustomPages = *customPagesStruct
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Organization from struct: %+v", updatedAccessOrganization))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.UpdateAccessOrganization(ctx, identifier, updatedAccessOrganization)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Organization for %s %q: %w", identifier.Level, identifier.Identifier, err))
	}

	return resourceCloudflareAccessOrganizationRead(ctx, d, meta)
}

func resourceCloudflareAccessOrganizationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	ctx = context.WithValue(ctx, orgAccessImportCtxKey, true)

	accountID := d.Id()

	tflog.Info(ctx, fmt.Sprintf("Importing Cloudflare Access Organization for account %s", accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)

	readErr := resourceCloudflareAccessOrganizationRead(ctx, d, meta)
	if readErr != nil {
		return nil, errors.New("failed to read Access Organization state")
	}

	return []*schema.ResourceData{d}, nil
}
