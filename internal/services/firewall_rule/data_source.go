// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rule

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type FirewallRuleDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &FirewallRuleDataSource{}

func NewFirewallRuleDataSource() datasource.DataSource {
	return &FirewallRuleDataSource{}
}

func (d *FirewallRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_rule"
}

func (r *FirewallRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *FirewallRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FirewallRuleDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := FirewallRuleResultDataSourceEnvelope{*data}
		_, err := r.client.Firewall.Rules.Get(
			ctx,
			data.ZoneIdentifier.ValueString(),
			firewall.RuleGetParams{
				PathID: cloudflare.F(data.PathID.ValueString()),
			},
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}
		bytes, _ := io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
			return
		}
		data = &env.Result
	} else {
		items := &[]*FirewallRuleDataSourceModel{}
		env := FirewallRuleResultListDataSourceEnvelope{items}

		page, err := r.client.Firewall.Rules.List(
			ctx,
			data.FindOneBy.ZoneIdentifier.ValueString(),
			firewall.RuleListParams{
				ID:          cloudflare.F(data.FindOneBy.ID.ValueString()),
				Action:      cloudflare.F(data.FindOneBy.Action.ValueString()),
				Description: cloudflare.F(data.FindOneBy.Description.ValueString()),
				Page:        cloudflare.F(data.FindOneBy.Page.ValueFloat64()),
				Paused:      cloudflare.F(data.FindOneBy.Paused.ValueBool()),
				PerPage:     cloudflare.F(data.FindOneBy.PerPage.ValueFloat64()),
			},
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(*items); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		data = (*items)[0]
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
