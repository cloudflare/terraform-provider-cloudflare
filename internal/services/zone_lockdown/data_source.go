// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

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

type ZoneLockdownDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &ZoneLockdownDataSource{}

func NewZoneLockdownDataSource() datasource.DataSource {
	return &ZoneLockdownDataSource{}
}

func (d *ZoneLockdownDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone_lockdown"
}

func (r *ZoneLockdownDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ZoneLockdownDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ZoneLockdownDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := ZoneLockdownResultDataSourceEnvelope{*data}
		_, err := r.client.Firewall.Lockdowns.Get(
			ctx,
			data.ZoneIdentifier.ValueString(),
			data.ID.ValueString(),
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
		dataFilterCreatedOn, errs := data.Filter.CreatedOn.ValueRFC3339Time()
		resp.Diagnostics.Append(errs...)
		dataFilterModifiedOn, errs := data.Filter.ModifiedOn.ValueRFC3339Time()
		resp.Diagnostics.Append(errs...)
		if resp.Diagnostics.HasError() {
			return
		}

		items := &[]*ZoneLockdownDataSourceModel{}
		env := ZoneLockdownResultListDataSourceEnvelope{items}

		page, err := r.client.Firewall.Lockdowns.List(
			ctx,
			data.Filter.ZoneIdentifier.ValueString(),
			firewall.LockdownListParams{
				CreatedOn:         cloudflare.F(dataFilterCreatedOn),
				Description:       cloudflare.F(data.Filter.Description.ValueString()),
				DescriptionSearch: cloudflare.F(data.Filter.DescriptionSearch.ValueString()),
				IP:                cloudflare.F(data.Filter.IP.ValueString()),
				IPRangeSearch:     cloudflare.F(data.Filter.IPRangeSearch.ValueString()),
				IPSearch:          cloudflare.F(data.Filter.IPSearch.ValueString()),
				ModifiedOn:        cloudflare.F(dataFilterModifiedOn),
				Page:              cloudflare.F(data.Filter.Page.ValueFloat64()),
				PerPage:           cloudflare.F(data.Filter.PerPage.ValueFloat64()),
				Priority:          cloudflare.F(data.Filter.Priority.ValueFloat64()),
				URISearch:         cloudflare.F(data.Filter.URISearch.ValueString()),
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
