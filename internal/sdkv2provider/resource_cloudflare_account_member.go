package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

func resourceCloudflareAccountMember() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccountMemberSchema(),
		CreateContext: resourceCloudflareAccountMemberCreate,
		ReadContext:   resourceCloudflareAccountMemberRead,
		UpdateContext: resourceCloudflareAccountMemberUpdate,
		DeleteContext: resourceCloudflareAccountMemberDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccountMemberImport,
		},
		Description: heredoc.Doc(`
			Provides a resource which manages Cloudflare account members.
		`),
	}
}

func resourceCloudflareAccountMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	member, err := client.AccountMember(ctx, accountID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing account member from state because it's not present in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	var memberIDs []string
	for _, role := range member.Roles {
		memberIDs = append(memberIDs, role.ID)
	}

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("email_address", member.User.Email)
	d.Set("role_ids", memberIDs)
	d.Set("status", member.Status)
	d.SetId(d.Id())

	return nil
}

func resourceCloudflareAccountMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare account member ID: %s", d.Id()))

	err := client.DeleteAccountMember(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Cloudflare account member: %w", err))
	}

	return nil
}

func resourceCloudflareAccountMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	memberEmailAddress := d.Get("email_address").(string)
	requestedMemberRoles := d.Get("role_ids").(*schema.Set).List()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	var accountMemberRoleIDs []string
	for _, roleID := range requestedMemberRoles {
		accountMemberRoleIDs = append(accountMemberRoleIDs, roleID.(string))
	}

	status := d.Get("status").(string)
	r, err := client.CreateAccountMemberWithStatus(ctx, accountID, memberEmailAddress, accountMemberRoleIDs, status)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Cloudflare account member: %w", err))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find ID in create response; resource was empty"))
	}

	d.SetId(r.ID)

	return resourceCloudflareAccountMemberRead(ctx, d, meta)
}

func resourceCloudflareAccountMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountRoles := []cloudflare.AccountRole{}
	memberRoles := d.Get("role_ids").(*schema.Set).List()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	for _, r := range memberRoles {
		accountRole, _ := client.GetAccountRole(ctx, cloudflare.AccountIdentifier(accountID), r.(string))
		accountRoles = append(accountRoles, accountRole)
	}

	updatedAccountMember := cloudflare.AccountMember{Roles: accountRoles}
	_, err := client.UpdateAccountMember(ctx, accountID, d.Id(), updatedAccountMember)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update Cloudflare account member: %w", err))
	}

	return resourceCloudflareAccountMemberRead(ctx, d, meta)
}

func resourceCloudflareAccountMemberImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup the account member
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var accountMemberID string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		accountMemberID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id %q specified, should be in format \"accountID/accountMemberID\" for import", d.Id())
	}

	member, err := client.AccountMember(ctx, accountID, accountMemberID)
	if err != nil {
		return nil, fmt.Errorf("unable to find account member with ID %q: %w", accountMemberID, err)
	}

	tflog.Info(ctx, fmt.Sprintf("Found account member: %s", member.User.Email))

	var memberIDs []string
	for _, role := range member.Roles {
		memberIDs = append(memberIDs, role.ID)
	}

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("email_address", member.User.Email)
	d.Set("role_ids", memberIDs)
	d.SetId(accountMemberID)

	return []*schema.ResourceData{d}, nil
}
