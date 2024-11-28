package leaked_credential_check_rules

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &LeakedCredentialCheckRulesResource{}
	_ resource.ResourceWithImportState = &LeakedCredentialCheckRulesResource{}
)

func NewResource() resource.Resource {
	return &LeakedCredentialCheckRulesResource{}
}

type LeakedCredentialCheckRulesResource struct {
	client *muxclient.Client
}

func (r *LeakedCredentialCheckRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_leaked_credential_check_rules"
}

func (r *LeakedCredentialCheckRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LeakedCredentialCheckRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LeakedCredentialCheckRulesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := plan.ZoneID.ValueString()
	var createdRules []LCCRuleValueModel
	if len(plan.Rules) > 0 {
		for _, rule := range plan.Rules {
			createParam := cloudflare.LeakedCredentialCheckCreateDetectionParams{
				Username: rule.Username.ValueString(),
				Password: rule.Password.ValueString(),
			}
			detection, err := r.client.V1.LeakedCredentialCheckCreateDetection(ctx, cloudflare.ZoneIdentifier(zoneID), createParam)
			if err != nil {
				resp.Diagnostics.AddError("Error creating a user-defined detection patter for Leaked Credential Check", err.Error())
				return
			}
			createdRules = append(createdRules, LCCRuleValueModel{ID: types.StringValue(detection.ID), Username: rule.Username, Password: rule.Password})
		}

	}
	plan.Rules = createdRules

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LeakedCredentialCheckRulesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	rules, err := r.client.V1.LeakedCredentialCheckListDetections(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error listing Leaked Credential Check user-defined detection patterns", err.Error())
		return
	}
	var readRules []LCCRuleValueModel
	for _, rule := range rules {
		readRules = append(readRules, buildLCCRuleValueModel(rule))
	}
	state.Rules = readRules
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan LeakedCredentialCheckRulesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(plan.ZoneID.ValueString())
	// fetch existing rules from API
	existing_rules, err := r.client.V1.LeakedCredentialCheckListDetections(ctx, zoneID, cloudflare.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error fetching Leaked Credential Check user-defined detection patterns", err.Error())
		return
	}
	// compare and create/delete accordingly
	var createdRules []LCCRuleValueModel
	toAdd, toRemove, toKeep := diffRules(plan.Rules, existing_rules)
	// create
	for _, createParam := range toAdd {
		detection, err := r.client.V1.LeakedCredentialCheckCreateDetection(ctx, zoneID, createParam)
		if err != nil {
			resp.Diagnostics.AddError("Error creating a user-defined detection patter for Leaked Credential Check", err.Error())
			return
		}
		createdRules = append(createdRules, buildLCCRuleValueModel(detection))
	}
	plan.Rules = createdRules // update plan rules with the newly created rules
	// delete
	for _, deleteParam := range toRemove {
		_, err := r.client.V1.LeakedCredentialCheckDeleteDetection(ctx, zoneID, deleteParam)
		if err != nil {
			resp.Diagnostics.AddError("Error deleting a user-defined detection patter for Leaked Credential Check", err.Error())
			return
		}
	}
	// add the existing rules we kept to the plan, if any
	for _, kr := range toKeep {
		plan.Rules = append(plan.Rules, buildLCCRuleValueModel(kr))
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete all user-defined detection rules
func (r *LeakedCredentialCheckRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LeakedCredentialCheckRulesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(state.ZoneID.ValueString())
	// fetch existing rules from API
	rules, err := r.client.V1.LeakedCredentialCheckListDetections(ctx, zoneID, cloudflare.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error fetching Leaked Credential Check user-defined detection patterns", err.Error())
		return
	}
	for _, rule := range rules {
		deleteParam := cloudflare.LeakedCredentialCheckDeleteDetectionParams{DetectionID: rule.ID}
		_, err := r.client.V1.LeakedCredentialCheckDeleteDetection(ctx, zoneID, deleteParam)
		if err != nil {
			resp.Diagnostics.AddError("Error deleting a user-defined detection patter for Leaked Credential Check", err.Error())
			return
		}

	}

}

func (r *LeakedCredentialCheckRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// req.ID is the zoneID for which you want to import the state of the
	// Leaked Credential Check Rules
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), req.ID)...)
	return
}

func diffRules(desired []LCCRuleValueModel, current []cloudflare.LeakedCredentialCheckDetectionEntry) (toAdd []cloudflare.LeakedCredentialCheckCreateDetectionParams, toRemove []cloudflare.LeakedCredentialCheckDeleteDetectionParams, toKeep []cloudflare.LeakedCredentialCheckDetectionEntry) {
	// Create a map for the desired rules
	desiredMap := make(map[string]struct{})
	for _, d := range desired {
		key := d.Username.ValueString() + ":" + d.Password.ValueString() // Use a unique key for comparison
		desiredMap[key] = struct{}{}
	}

	// Create a map for the current rules
	currentMap := make(map[string]struct{})
	for _, c := range current {
		key := c.Username + ":" + c.Password
		currentMap[key] = struct{}{}

		// If a rule exists in current but not in desired, mark it for removal
		if _, exists := desiredMap[key]; !exists {
			toRemove = append(toRemove, cloudflare.LeakedCredentialCheckDeleteDetectionParams{DetectionID: c.ID})
		} else {
			// If a rule exists both in current and desired, mark it as to keep
			toKeep = append(
				toKeep,
				cloudflare.LeakedCredentialCheckDetectionEntry{ID: c.ID, Username: c.Username, Password: c.Password},
			)
		}
	}

	// Find rules that exist in desired but not in current
	for _, d := range desired {
		key := d.Username.ValueString() + ":" + d.Password.ValueString()
		if _, exists := currentMap[key]; !exists {
			toAdd = append(toAdd, cloudflare.LeakedCredentialCheckCreateDetectionParams{
				Username: d.Username.ValueString(),
				Password: d.Password.ValueString(),
			})
		}
	}

	return
}

func buildLCCRuleValueModel(lccEntry cloudflare.LeakedCredentialCheckDetectionEntry) LCCRuleValueModel {
	rule := LCCRuleValueModel{
		ID:       types.StringValue(lccEntry.ID),
		Username: types.StringValue(lccEntry.Username),
		Password: types.StringValue(lccEntry.Password),
	}
	return rule
}
