package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessServiceToken() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessServiceTokenSchema(),
		CreateContext: resourceCloudflareAccessServiceTokenCreate,
		ReadContext:   resourceCloudflareAccessServiceTokenRead,
		UpdateContext: resourceCloudflareAccessServiceTokenUpdate,
		DeleteContext: resourceCloudflareAccessServiceTokenDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessServiceTokenImport,
		},
		Description: heredoc.Doc(`
			Access Service Tokens are used for service-to-service communication
			when an application is behind Cloudflare Access.
		`),
	}
}

func resourceCloudflareAccessServiceTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	// The Cloudflare API doesn't support fetching a single service token
	// so instead we loop over all the service tokens and only continue
	// when we have a match.
	var serviceTokens []cloudflare.AccessServiceToken
	if identifier.Type == AccountType {
		serviceTokens, _, err = client.AccessServiceTokens(ctx, identifier.Value)
	} else {
		serviceTokens, _, err = client.ZoneLevelAccessServiceTokens(ctx, identifier.Value)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching access service tokens: %w", err))
	}
	for _, token := range serviceTokens {
		if token.ID == d.Id() {
			zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
			accountID := d.Get(consts.AccountIDSchemaKey).(string)
			mindays := d.Get("min_days_for_renewal").(int)
			if mindays > 0 {
				expires_at := d.Get("expires_at").(string)

				if expires_at != "" {
					expected_expiration_date, _ := time.Parse(time.RFC3339, expires_at)

					expiration_date := time.Now().Add(time.Duration(mindays) * 24 * time.Hour)

					if expiration_date.After(expected_expiration_date) {
						var refreshedToken cloudflare.AccessServiceTokenRefreshResponse
						var err error

						if accountID != "" {
							refreshedToken, err = client.RefreshAccessServiceToken(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
						} else {
							refreshedToken, err = client.RefreshAccessServiceToken(ctx, cloudflare.ZoneIdentifier(zoneID), d.Id())
						}

						if err != nil {
							return diag.FromErr(fmt.Errorf("failed to automatically refresh token %q: %w", d.Id(), err))
						}

						token.ExpiresAt = refreshedToken.ExpiresAt
					}
				}
			}
			d.Set("name", token.Name)
			d.Set("client_id", token.ClientID)
			d.Set("expires_at", token.ExpiresAt.Format(time.RFC3339))
		}
	}

	return nil
}

func resourceCloudflareAccessServiceTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	tokenName := d.Get("name").(string)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var serviceToken cloudflare.AccessServiceTokenCreateResponse
	if identifier.Type == AccountType {
		serviceToken, err = client.CreateAccessServiceToken(ctx, identifier.Value, tokenName)
	} else {
		serviceToken, err = client.CreateZoneLevelAccessServiceToken(ctx, identifier.Value, tokenName)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating access service token: %w", err))
	}

	d.SetId(serviceToken.ID)
	d.Set("name", serviceToken.Name)
	d.Set("client_id", serviceToken.ClientID)
	d.Set("client_secret", serviceToken.ClientSecret)
	d.Set("expires_at", serviceToken.ExpiresAt.Format(time.RFC3339))

	resourceCloudflareAccessServiceTokenRead(ctx, d, meta)

	return nil
}

func resourceCloudflareAccessServiceTokenUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	tokenName := d.Get("name").(string)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var serviceToken cloudflare.AccessServiceTokenUpdateResponse
	if identifier.Type == AccountType {
		serviceToken, err = client.UpdateAccessServiceToken(ctx, identifier.Value, d.Id(), tokenName)
	} else {
		serviceToken, err = client.UpdateZoneLevelAccessServiceToken(ctx, identifier.Value, d.Id(), tokenName)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating access service token: %w", err))
	}

	d.Set("name", serviceToken.Name)

	return resourceCloudflareAccessServiceTokenRead(ctx, d, meta)
}

func resourceCloudflareAccessServiceTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if identifier.Type == AccountType {
		_, err = client.DeleteAccessServiceToken(ctx, identifier.Value, d.Id())
	} else {
		_, err = client.DeleteZoneLevelAccessServiceToken(ctx, identifier.Value, d.Id())
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting access service token: %w", err))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareAccessServiceTokenImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/serviceTokenID\"", d.Id())
	}

	d.Set(consts.AccountIDSchemaKey, attributes[0])
	d.SetId(attributes[1])

	resourceCloudflareAccessServiceTokenRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
