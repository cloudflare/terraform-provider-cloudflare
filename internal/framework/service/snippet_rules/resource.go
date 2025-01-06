package snippet_rules

import (
	"context"
	"fmt"
	"strings"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SnippetRulesResource{}
var _ resource.ResourceWithImportState = &SnippetRulesResource{}

// SnippetRulesResource defines the resource implementation.
type SnippetRulesResource struct {
	client *muxclient.Client
}

func NewResource() resource.Resource {
	return &SnippetRulesResource{}
}

func (r *SnippetRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet_rules"
}

func (r *SnippetRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func resourceFromAPIResponse(data *SnippetRules, rules []cfv1.SnippetRule) {
	result := make([]SnippetRule, len(rules))
	for i, rule := range rules {
		result[i].Description = types.StringValue(rule.Description)
		result[i].Expression = types.StringValue(rule.Expression)
		result[i].Enabled = types.BoolPointerValue(rule.Enabled)
		result[i].SnippetName = types.StringValue(rule.SnippetName)
	}
	data.Rules = result
}

func requestFromResource(data *SnippetRules) []cfv1.SnippetRule {
	rules := make([]cfv1.SnippetRule, len(data.Rules))
	for i, rule := range data.Rules {
		rules[i] = cfv1.SnippetRule{
			Enabled:     rule.Enabled.ValueBoolPointer(),
			Expression:  rule.Expression.ValueString(),
			SnippetName: rule.SnippetName.ValueString(),
			Description: rule.Description.ValueString(),
		}
	}
	return rules
}

func (r *SnippetRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SnippetRules

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.UpdateZoneSnippetsRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create snippet rules", err.Error())
		return
	}

	resourceFromAPIResponse(data, rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SnippetRules

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.ListZoneSnippetsRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("failed reading snippet rules", err.Error())
		return
	}
	resourceFromAPIResponse(data, rules)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SnippetRules

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.UpdateZoneSnippetsRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create snippet rules", err.Error())
		return
	}

	resourceFromAPIResponse(data, rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SnippetRules

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.UpdateZoneSnippetsRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), []cfv1.SnippetRule{})

	if err != nil {
		resp.Diagnostics.AddError("failed to delete snippet rules", err.Error())
		return
	}
}

func (r *SnippetRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing snippet rule", `invalid ID specified. Please specify the ID as "<zone_id>/<snippet_rule_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
