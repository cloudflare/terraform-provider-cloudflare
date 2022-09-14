package provider

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	accountTypeStandard   = "standard"
	accountTypeEnterprise = "enterprise"
)

func resourceCloudflareAccount() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccountSchema(),
		CreateContext: resourceCloudflareAccountCreate,
		ReadContext:   resourceCloudflareAccountRead,
		UpdateContext: resourceCloudflareAccountUpdate,
		DeleteContext: resourceCloudflareAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Account resource. Account is the basic resource for
			working with Cloudflare zones, teams and users. 
		`),
	}
}

func resourceCloudflareAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountName := d.Get("name").(string)
	accountType := d.Get("type").(string)

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare Account: name %s", accountName))

	account := cloudflare.Account{
		Name: accountName,
		Type: accountType,
	}
	acc, err := client.CreateAccount(ctx, account)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating account %q: %w", accountName, err))
	}

	d.SetId(acc.ID)

	return resourceCloudflareAccountRead(ctx, d, meta)
}

func resourceCloudflareAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Id()

	foundAcc, _, err := client.Account(ctx, accountID)
	if err != nil || foundAcc.ID == "" {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Account %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Account %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("AccountDetails: %#v", foundAcc))

	d.Set("name", foundAcc.Name)
	d.Set("type", foundAcc.Type)
	d.Set("enforce_twofactor", foundAcc.Settings.EnforceTwoFactor)

	return nil
}

func resourceCloudflareAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Id()
	foundAcc, _, err := client.Account(ctx, accountID)
	if err != nil || foundAcc.ID == "" {
		return diag.FromErr(fmt.Errorf("error finding Account %q: %w", d.Id(), err))
	}

	tflog.Debug(ctx, fmt.Sprintf("[INFO] Updating Cloudflare Account: id %s", accountID))

	if accountName, ok := d.GetOk("name"); ok && d.HasChange("name") {
		foundAcc.Name = accountName.(string)
	}

	if enforce_twofactor, ok := d.GetOk("enforce_twofactor"); ok && d.HasChange("enforce_twofactor") {
		foundAcc.Settings.EnforceTwoFactor = enforce_twofactor.(bool)
	}

	_, err = client.UpdateAccount(ctx, accountID, foundAcc)
	if err != nil {
		if ferr, ok := err.(*cloudflare.RequestError); ok {
			errCodes := ferr.ErrorCodes()
			if len(errCodes) == 1 && errCodes[0] == 1001 {
				tflog.Debug(ctx, "Ignoring error 1001: Updating account type is not supported from client api")
				return resourceCloudflareAccountRead(ctx, d, meta)
			}
		}
		tflog.Error(ctx, fmt.Sprintf("%#v", err))
		return diag.FromErr(fmt.Errorf("error updating Account %q: %w", d.Id(), err))
	} else {
		return resourceCloudflareAccountRead(ctx, d, meta)
	}
}

func resourceCloudflareAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Account: id %s", accountID)

	err := client.DeleteAccount(ctx, accountID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare Account: %w", err))
	}

	return nil
}
