// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_rule

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type RulesetRuleDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*RulesetRuleDataSource)(nil)

func NewRulesetRuleDataSource() datasource.DataSource {
	return &RulesetRuleDataSource{}
}

func (d *RulesetRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset_rule"
}

func (d *RulesetRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *RulesetRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RulesetRuleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare parameters for reading the parent ruleset
	params, diags := data.toReadParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the entire ruleset to find the specific rule
	res := new(http.Response)
	rulesetEnv := ruleset.RulesetResultDataSourceEnvelope{}

	_, err := d.client.Rulesets.Get(
		ctx,
		data.RulesetID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	err = apijsoncustom.UnmarshalComputed(bytes, &rulesetEnv)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Find the specific rule within the ruleset
	ruleID := data.RuleID.ValueString()
	found := false

	// Convert the custom field list to a slice
	rules, diags := rulesetEnv.Result.Rules.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, rule := range rules {
		if rule.ID.ValueString() == ruleID {
			// Copy rule data to our model (preserve lookup fields)
			data.ID = rule.ID
			data.Action = rule.Action
			data.ActionParameters = rule.ActionParameters
			data.Description = rule.Description
			data.Enabled = rule.Enabled
			data.ExposedCredentialCheck = rule.ExposedCredentialCheck
			data.Expression = rule.Expression
			data.Logging = rule.Logging
			data.Ratelimit = rule.Ratelimit
			data.Ref = rule.Ref
			data.Categories = rule.Categories
			found = true
			break
		}
	}

	if !found {
		resp.Diagnostics.AddError(
			"rule not found",
			fmt.Sprintf("Rule with ID %s not found in ruleset %s", ruleID, data.RulesetID.ValueString()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
