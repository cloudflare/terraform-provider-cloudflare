// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package record

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/dns"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type RecordDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &RecordDataSource{}

func NewRecordDataSource() datasource.DataSource {
	return &RecordDataSource{}
}

func (d *RecordDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record"
}

func (r *RecordDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RecordDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy == nil {
		res := new(http.Response)
		env := RecordResultDataSourceEnvelope{*data}
		_, err := r.client.DNS.Records.Get(
			ctx,
			data.DNSRecordID.ValueString(),
			dns.RecordGetParams{
				ZoneID: cloudflare.F(data.ZoneID.ValueString()),
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
		items := &[]*RecordDataSourceModel{}
		env := RecordResultListDataSourceEnvelope{items}

		page, err := r.client.DNS.Records.List(ctx, dns.RecordListParams{
			ZoneID: cloudflare.F(data.FindOneBy.ZoneID.ValueString()),
			Comment: cloudflare.F(dns.RecordListParamsComment{
				Absent:     cloudflare.F(data.FindOneBy.Comment.Absent.ValueString()),
				Contains:   cloudflare.F(data.FindOneBy.Comment.Contains.ValueString()),
				Endswith:   cloudflare.F(data.FindOneBy.Comment.Endswith.ValueString()),
				Exact:      cloudflare.F(data.FindOneBy.Comment.Exact.ValueString()),
				Present:    cloudflare.F(data.FindOneBy.Comment.Present.ValueString()),
				Startswith: cloudflare.F(data.FindOneBy.Comment.Startswith.ValueString()),
			}),
			Content:   cloudflare.F(data.FindOneBy.Content.ValueString()),
			Direction: cloudflare.F(shared.SortDirection(data.FindOneBy.Direction.ValueString())),
			Match:     cloudflare.F(dns.RecordListParamsMatch(data.FindOneBy.Match.ValueString())),
			Name:      cloudflare.F(data.FindOneBy.Name.ValueString()),
			Order:     cloudflare.F(dns.RecordListParamsOrder(data.FindOneBy.Order.ValueString())),
			Page:      cloudflare.F(data.FindOneBy.Page.ValueFloat64()),
			PerPage:   cloudflare.F(data.FindOneBy.PerPage.ValueFloat64()),
			Proxied:   cloudflare.F(data.FindOneBy.Proxied.ValueBool()),
			Search:    cloudflare.F(data.FindOneBy.Search.ValueString()),
			Tag: cloudflare.F(dns.RecordListParamsTag{
				Absent:     cloudflare.F(data.FindOneBy.Tag.Absent.ValueString()),
				Contains:   cloudflare.F(data.FindOneBy.Tag.Contains.ValueString()),
				Endswith:   cloudflare.F(data.FindOneBy.Tag.Endswith.ValueString()),
				Exact:      cloudflare.F(data.FindOneBy.Tag.Exact.ValueString()),
				Present:    cloudflare.F(data.FindOneBy.Tag.Present.ValueString()),
				Startswith: cloudflare.F(data.FindOneBy.Tag.Startswith.ValueString()),
			}),
			TagMatch: cloudflare.F(dns.RecordListParamsTagMatch(data.FindOneBy.TagMatch.ValueString())),
			Type:     cloudflare.F(dns.RecordListParamsType(data.FindOneBy.Type.ValueString())),
		})
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
