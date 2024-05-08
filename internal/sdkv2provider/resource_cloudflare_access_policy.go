package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessPolicy() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessPolicySchema(),
		CreateContext: resourceCloudflareAccessPolicyCreate,
		ReadContext:   resourceCloudflareAccessPolicyRead,
		UpdateContext: resourceCloudflareAccessPolicyUpdate,
		DeleteContext: resourceCloudflareAccessPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessPolicyImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Policy resource. Access Policies are
			used in conjunction with Access Applications to restrict access to
			a particular resource.
		`),
	}
}

func apiAccessPolicyApprovalGroupToSchema(approvalGroup cloudflare.AccessApprovalGroup) map[string]interface{} {
	data := make(map[string]interface{})
	data["approvals_needed"] = approvalGroup.ApprovalsNeeded

	if approvalGroup.EmailAddresses != nil {
		data["email_addresses"] = approvalGroup.EmailAddresses
	}

	if approvalGroup.EmailListUuid != "" {
		data["email_list_uuid"] = approvalGroup.EmailListUuid
	}
	return data
}

func schemaAccessPolicyApprovalGroupToAPI(data map[string]interface{}) cloudflare.AccessApprovalGroup {
	var approvalGroup cloudflare.AccessApprovalGroup

	approvalGroup.ApprovalsNeeded, _ = data["approvals_needed"].(int)
	approvalGroup.EmailListUuid, _ = data["email_list_uuid"].(string)

	if emailAddresses, ok := data["email_addresses"].([]interface{}); ok {
		approvalGroup.EmailAddresses = expandInterfaceToStringList(emailAddresses)
	}

	return approvalGroup
}

func apiCloudflareAccessPolicyToResource(ctx context.Context, d *schema.ResourceData, appID string, accessPolicy cloudflare.AccessPolicy) diag.Diagnostics {
	if appID != "" {
		// policy is tied to a single application (legacy) and its execution precedence
		// within the app can be set.
		d.Set("precedence", accessPolicy.Precedence)
	}

	d.Set("name", accessPolicy.Name)
	d.Set("decision", accessPolicy.Decision)

	if err := d.Set("require", TransformAccessGroupForSchema(ctx, accessPolicy.Require)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set require attribute: %w", err))
	}

	if err := d.Set("exclude", TransformAccessGroupForSchema(ctx, accessPolicy.Exclude)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set exclude attribute: %w", err))
	}

	if err := d.Set("include", TransformAccessGroupForSchema(ctx, accessPolicy.Include)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set include attribute: %w", err))
	}

	d.Set("isolation_required", accessPolicy.IsolationRequired)
	d.Set("purpose_justification_required", accessPolicy.PurposeJustificationRequired)
	d.Set("purpose_justification_prompt", accessPolicy.PurposeJustificationPrompt)
	d.Set("approval_required", accessPolicy.ApprovalRequired)

	if accessPolicy.SessionDuration != nil {
		d.Set("session_duration", accessPolicy.SessionDuration)
	}

	if len(accessPolicy.ApprovalGroups) != 0 {
		approvalGroups := make([]map[string]interface{}, 0, len(accessPolicy.ApprovalGroups))
		for _, apiApprovalGroup := range accessPolicy.ApprovalGroups {
			approvalGroups = append(approvalGroups, apiAccessPolicyApprovalGroupToSchema(apiApprovalGroup))
		}
		if err := d.Set("approval_group", approvalGroups); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set approval_group attribute: %w", err))
		}
	}
	return nil
}

func resourceCloudflareAccessPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	appID, _ := d.Get("application_id").(string)
	policy, err := client.GetAccessPolicy(ctx, identifier, cloudflare.GetAccessPolicyParams{
		ApplicationID: appID,
		PolicyID:      d.Id(),
	})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Policy %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Access Policy %q: %w", d.Id(), err))
	}
	return apiCloudflareAccessPolicyToResource(ctx, d, appID, policy)
}

func resourceCloudflareAccessPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID, _ := d.Get("application_id").(string)
	precedence, _ := d.Get("precedence").(int)
	newAccessPolicy := cloudflare.CreateAccessPolicyParams{
		ApplicationID:   appID,
		Precedence:      precedence,
		Name:            d.Get("name").(string),
		Decision:        d.Get("decision").(string),
		SessionDuration: cloudflare.StringPtr(d.Get("session_duration").(string)),
	}

	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		if value != nil {
			newAccessPolicy.Exclude = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		if value != nil {
			newAccessPolicy.Require = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		if value != nil {
			newAccessPolicy.Include = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	isolationRequired := d.Get("isolation_required").(bool)
	newAccessPolicy.IsolationRequired = &isolationRequired

	purposeJustificationRequired := d.Get("purpose_justification_required").(bool)
	newAccessPolicy.PurposeJustificationRequired = &purposeJustificationRequired

	purposeJustificationPrompt := d.Get("purpose_justification_prompt").(string)
	newAccessPolicy.PurposeJustificationPrompt = &purposeJustificationPrompt

	approvalRequired := d.Get("approval_required").(bool)
	newAccessPolicy.ApprovalRequired = &approvalRequired

	approvalGroups := d.Get("approval_group").([]interface{})
	for _, approvalGroup := range approvalGroups {
		approvalGroupAsMap := approvalGroup.(map[string]interface{})
		newAccessPolicy.ApprovalGroups = append(newAccessPolicy.ApprovalGroups, schemaAccessPolicyApprovalGroupToAPI(approvalGroupAsMap))
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Policy from struct: %+v", newAccessPolicy))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessPolicy, err := client.CreateAccessPolicy(ctx, identifier, newAccessPolicy)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Policy for ID %q: %w", d.Id(), err))
	}

	d.SetId(accessPolicy.ID)

	return apiCloudflareAccessPolicyToResource(ctx, d, appID, accessPolicy)
}

func resourceCloudflareAccessPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	appID, _ := d.Get("application_id").(string)
	precedence, _ := d.Get("precedence").(int)
	updateReq := cloudflare.UpdateAccessPolicyParams{
		ApplicationID:   appID,
		Precedence:      precedence,
		PolicyID:        d.Id(),
		Name:            d.Get("name").(string),
		Decision:        d.Get("decision").(string),
		SessionDuration: cloudflare.StringPtr(d.Get("session_duration").(string)),
	}

	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		if value != nil {
			updateReq.Exclude = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		if value != nil {
			updateReq.Require = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		if value != nil {
			updateReq.Include = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	isolationRequired := d.Get("isolation_required").(bool)
	updateReq.IsolationRequired = &isolationRequired

	purposeJustificationRequired := d.Get("purpose_justification_required").(bool)
	updateReq.PurposeJustificationRequired = &purposeJustificationRequired

	purposeJustificationPrompt := d.Get("purpose_justification_prompt").(string)
	updateReq.PurposeJustificationPrompt = &purposeJustificationPrompt

	approvalRequired := d.Get("approval_required").(bool)
	updateReq.ApprovalRequired = &approvalRequired

	approvalGroups := d.Get("approval_group").([]interface{})
	for _, approvalGroup := range approvalGroups {
		approvalGroupAsMap := approvalGroup.(map[string]interface{})
		updateReq.ApprovalGroups = append(updateReq.ApprovalGroups, schemaAccessPolicyApprovalGroupToAPI(approvalGroupAsMap))
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Policy from struct: %+v", updateReq))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	updatedPolicy, err := client.UpdateAccessPolicy(ctx, identifier, updateReq)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Policy for ID %q: %w", d.Id(), err))
	}

	return apiCloudflareAccessPolicyToResource(ctx, d, updateReq.ApplicationID, updatedPolicy)
}

func resourceCloudflareAccessPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	appID, _ := d.Get("application_id").(string)
	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Policy using ID: %s", d.Id()))
	err = client.DeleteAccessPolicy(ctx, identifier, cloudflare.DeleteAccessPolicyParams{
		ApplicationID: appID,
		PolicyID:      d.Id(),
	})
	if err != nil {
		// If the resource is already deleted, we should return without an error
		// according to the terraform spec
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Policy %s no longer exists", d.Id()))
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting Access Policy for ID %q: %w", d.Id(), err))
	}

	resourceCloudflareAccessPolicyRead(ctx, d, meta)

	return nil
}

func resourceCloudflareAccessPolicyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 4)

	if len(attributes) < 3 {
		return nil, fmt.Errorf(
			"invalid id (%q) specified, should be in format %q, %q or %q",
			d.Id(),
			"account/accountID/accessPolicyID",
			"account/accountID/accessApplicationID/accessPolicyID",
			"zone/zoneID/accessApplicationID/accessPolicyID",
		)
	}

	identifierType, identifierID := attributes[0], attributes[1]
	if len(attributes) == 4 {
		// Legacy policy tied to a single application.
		accessAppID, accessPolicyID := attributes[2], attributes[3]
		tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Policy: %s %q, appID %q, accessPolicyID %q", identifierType, identifierID, accessAppID, accessPolicyID))
		d.Set("application_id", accessAppID)
		d.SetId(accessPolicyID)
	} else {
		// Standalone reusable policy
		accessPolicyID := attributes[2]
		tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Policy: %s %q, accessPolicyID %q", identifierType, identifierID, accessPolicyID))
		d.SetId(accessPolicyID)
	}

	//lintignore:R001
	d.Set(fmt.Sprintf("%s_id", identifierType), identifierID)

	resourceCloudflareAccessPolicyRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
