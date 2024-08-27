package cloud_connector_rules

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
var _ resource.Resource = &CloudConnectorRulesResource{}
var _ resource.ResourceWithImportState = &CloudConnectorRulesResource{}

// CloudConnectorRulesResource defines the resource implementation.
type CloudConnectorRulesResource struct {
	client *muxclient.Client
}

func NewResource() resource.Resource {
	return &CloudConnectorRulesResource{}
}

func (r *CloudConnectorRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_connector_rules"
}

func (r *CloudConnectorRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func resourceFromAPIResponse(data *CloudConnectorRules, rules []cfv1.CloudConnectorRule) {
	result := make([]CloudConnectorRule, len(rules))
	for i, rule := range rules {
		result[i].Description = types.StringValue(rule.Description)
		result[i].Expression = types.StringValue(rule.Expression)
		result[i].Provider = types.StringValue(rule.Provider)
		result[i].Enabled = types.BoolPointerValue(rule.Enabled)
		result[i].Parameters.Host = types.StringValue(rule.Parameters.Host)
	}
	data.Rules = result
}

func requestFromResource(data *CloudConnectorRules) []cfv1.CloudConnectorRule {
	rules := make([]cfv1.CloudConnectorRule, len(data.Rules))
	for i, rule := range data.Rules {
		rules[i] = cfv1.CloudConnectorRule{
			Enabled:    rule.Enabled.ValueBoolPointer(),
			Expression: rule.Expression.ValueString(),
			Provider:   rule.Provider.ValueString(),
			Parameters: cfv1.CloudConnectorRuleParameters{
				Host: rule.Parameters.Host.ValueString(),
			},
			Description: rule.Description.ValueString(),
		}
	}
	return rules
}

func (r *CloudConnectorRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudConnectorRules

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.UpdateZoneCloudConnectorRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create cloud connector rules", err.Error())
		return
	}

	resourceFromAPIResponse(data, rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudConnectorRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudConnectorRules

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.ListZoneCloudConnectorRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("failed reading cloud connector rules", err.Error())
		return
	}
	resourceFromAPIResponse(data, rules)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudConnectorRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudConnectorRules

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rules, err := r.client.V1.UpdateZoneCloudConnectorRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create cloud connector rules", err.Error())
		return
	}

	resourceFromAPIResponse(data, rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudConnectorRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudConnectorRules

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.UpdateZoneCloudConnectorRules(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), []cfv1.CloudConnectorRule{})

	if err != nil {
		resp.Diagnostics.AddError("failed to delete cloud connector rules", err.Error())
		return
	}
}

func (r *CloudConnectorRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing cloud connector rule", `invalid ID specified. Please specify the ID as "<zone_id>/<cloud_connector_rule_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
