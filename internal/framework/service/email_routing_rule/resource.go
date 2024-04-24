package email_routing_rule

import (
	"context"
	"fmt"
	"strings"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/expanders"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EmailRoutingRuleResource{}
var _ resource.ResourceWithImportState = &EmailRoutingRuleResource{}

func NewResource() resource.Resource {
	return &EmailRoutingRuleResource{}
}

// EmailRoutingRuleResource defines the resource implementation.
type EmailRoutingRuleResource struct {
	client *muxclient.Client
}

func (r *EmailRoutingRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email_routing_rule"
}

func (r *EmailRoutingRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EmailRoutingRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *EmailRoutingRuleModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	matcherModels, actionModels := buildMatchersAndActions(ctx, data)

	routingRule, err := r.client.V1.CreateEmailRoutingRule(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		cfv1.CreateEmailRoutingRuleParameters{
			Name:     data.Name.ValueString(),
			Priority: int(data.Priority.ValueInt64()),
			Enabled:  data.Enabled.ValueBoolPointer(),
			Matchers: matcherModels,
			Actions:  actionModels,
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create email routing rule", err.Error())
		return
	}
	data = buildRoutingRuleModel(data.ZoneID.ValueString(), routingRule)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailRoutingRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *EmailRoutingRuleModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	routingRule, err := r.client.V1.GetEmailRoutingRule(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), data.Tag.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed reading email routing rule", err.Error())
		return
	}
	data = buildRoutingRuleModel(data.ZoneID.ValueString(), routingRule)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailRoutingRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *EmailRoutingRuleModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	matcherModels, actionModels := buildMatchersAndActions(ctx, data)

	routingRule, err := r.client.V1.UpdateEmailRoutingRule(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		cfv1.UpdateEmailRoutingRuleParameters{
			RuleID:   data.Tag.ValueString(),
			Name:     data.Name.ValueString(),
			Priority: int(data.Priority.ValueInt64()),
			Enabled:  data.Enabled.ValueBoolPointer(),
			Matchers: matcherModels,
			Actions:  actionModels,
		},
	)

	if err != nil {
		resp.Diagnostics.AddError("failed updating email routing rule", err.Error())
		return
	}
	data = buildRoutingRuleModel(data.ZoneID.ValueString(), routingRule)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailRoutingRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *EmailRoutingRuleModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1.DeleteEmailRoutingRule(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), data.Tag.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("failed to email routing rule", err.Error())
		return
	}
}

func buildMatchersAndActions(ctx context.Context, data *EmailRoutingRuleModel) ([]cfv1.EmailRoutingRuleMatcher, []cfv1.EmailRoutingRuleAction) {
	var matcherModels []cfv1.EmailRoutingRuleMatcher
	var actionModels []cfv1.EmailRoutingRuleAction

	for _, matcher := range data.Matcher {
		matcherModels = append(matcherModels, cfv1.EmailRoutingRuleMatcher{
			Type:  matcher.Type.ValueString(),
			Field: matcher.Field.ValueString(),
			Value: matcher.Value.ValueString(),
		})
	}

	for _, action := range data.Action {
		actionModels = append(actionModels, cfv1.EmailRoutingRuleAction{
			Type:  action.Type.ValueString(),
			Value: expanders.StringSet(ctx, action.Value),
		})
	}

	return matcherModels, actionModels
}

func buildRoutingRuleModel(zoneID string, routingRule cfv1.EmailRoutingRule) *EmailRoutingRuleModel {
	var matcherModels []*EmailRoutingRuleMatcherModel
	var actionModels []*EmailRoutingRuleActionModel

	for _, matcher := range routingRule.Matchers {
		matcherModels = append(matcherModels, &EmailRoutingRuleMatcherModel{
			Type:  types.StringValue(matcher.Type),
			Field: types.StringValue(matcher.Field),
			Value: types.StringValue(matcher.Value),
		})
	}
	for _, action := range routingRule.Actions {
		var values []attr.Value
		for _, value := range action.Value {
			values = append(values, types.StringValue(value))
		}

		actionModels = append(actionModels, &EmailRoutingRuleActionModel{
			Type:  types.StringValue(action.Type),
			Value: flatteners.StringSet(values),
		})
	}

	return &EmailRoutingRuleModel{
		ZoneID:   types.StringValue(zoneID),
		ID:       types.StringValue(routingRule.Tag),
		Tag:      types.StringValue(routingRule.Tag),
		Name:     types.StringValue(routingRule.Name),
		Priority: types.Int64Value(int64(routingRule.Priority)),
		Enabled:  types.BoolValue(cfv1.Bool(routingRule.Enabled)),
		Matcher:  matcherModels,
		Action:   actionModels,
	}
}

func (r *EmailRoutingRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing email routing rule", `invalid ID specified. Please specify the ID as "<zone_id>/<email_routing_rule_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("tag"), idparts[1],
	)...)
}
