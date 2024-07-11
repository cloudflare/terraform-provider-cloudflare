// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type AccessRulesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &AccessRulesDataSource{}

func NewAccessRulesDataSource() datasource.DataSource {
	return &AccessRulesDataSource{}
}

func (d *AccessRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_rules"
}

func (r *AccessRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	r.client = client
}

func (r *AccessRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AccessRulesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := firewall.AccessRuleListParams{
		Direction: cloudflare.F(firewall.AccessRuleListParamsDirection(data.Direction.ValueString())),
		EgsPagination: cloudflare.F(firewall.AccessRuleListParamsEgsPagination{
			Json: cloudflare.F(firewall.AccessRuleListParamsEgsPaginationJson{
				Page:    cloudflare.F(data.EgsPagination.Json.Page.ValueFloat64()),
				PerPage: cloudflare.F(data.EgsPagination.Json.PerPage.ValueFloat64()),
			}),
		}),
		Filters: cloudflare.F(firewall.AccessRuleListParamsFilters{
			ConfigurationTarget: cloudflare.F(firewall.AccessRuleListParamsFiltersConfigurationTarget(data.Filters.ConfigurationTarget.ValueString())),
			ConfigurationValue:  cloudflare.F(data.Filters.ConfigurationValue.ValueString()),
			Match:               cloudflare.F(firewall.AccessRuleListParamsFiltersMatch(data.Filters.Match.ValueString())),
			Mode:                cloudflare.F(firewall.AccessRuleListParamsFiltersMode(data.Filters.Mode.ValueString())),
			Notes:               cloudflare.F(data.Filters.Notes.ValueString()),
		}),
		Order:   cloudflare.F(firewall.AccessRuleListParamsOrder(data.Order.ValueString())),
		Page:    cloudflare.F(data.Page.ValueFloat64()),
		PerPage: cloudflare.F(data.PerPage.ValueFloat64()),
	}
	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	items := &[]*AccessRulesItemsDataSourceModel{}
	env := AccessRulesResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*AccessRulesItemsDataSourceModel{}

	page, err := r.client.Firewall.AccessRules.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, *items...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:maxItems]
	data.Items = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
