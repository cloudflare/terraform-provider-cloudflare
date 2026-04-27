// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetRuleDataSource struct {
}

var _ datasource.DataSource = (*RulesetRuleDataSource)(nil)

func NewRulesetRuleDataSource() datasource.DataSource {
	return &RulesetRuleDataSource{}
}

func (d *RulesetRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset_rules"
}

func (d *RulesetRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RulesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bytes, err := apijsoncustom.Marshal(&data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize data", err.Error())
		return
	}

	data.ID = types.StringNull()
	h := sha256.New()
	hash := h.Sum(bytes)
	data.ID = types.StringValue(fmt.Sprintf("%x", hash))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
