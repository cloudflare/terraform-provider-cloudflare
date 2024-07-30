// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package record

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/dns"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type RecordsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &RecordsDataSource{}

func NewRecordsDataSource() datasource.DataSource {
	return &RecordsDataSource{}
}

func (d *RecordsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_records"
}

func (r *RecordsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *RecordsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *RecordsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*RecordsResultDataSourceModel{}
	env := RecordsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*RecordsResultDataSourceModel{}

	page, err := r.client.DNS.Records.List(ctx, dns.RecordListParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		Comment: cloudflare.F(dns.RecordListParamsComment{
			Absent:     cloudflare.F(data.Comment.Absent.ValueString()),
			Contains:   cloudflare.F(data.Comment.Contains.ValueString()),
			Endswith:   cloudflare.F(data.Comment.Endswith.ValueString()),
			Exact:      cloudflare.F(data.Comment.Exact.ValueString()),
			Present:    cloudflare.F(data.Comment.Present.ValueString()),
			Startswith: cloudflare.F(data.Comment.Startswith.ValueString()),
		}),
		Content:   cloudflare.F(data.Content.ValueString()),
		Direction: cloudflare.F(shared.SortDirection(data.Direction.ValueString())),
		Match:     cloudflare.F(dns.RecordListParamsMatch(data.Match.ValueString())),
		Name:      cloudflare.F(data.Name.ValueString()),
		Order:     cloudflare.F(dns.RecordListParamsOrder(data.Order.ValueString())),
		Page:      cloudflare.F(data.Page.ValueFloat64()),
		PerPage:   cloudflare.F(data.PerPage.ValueFloat64()),
		Proxied:   cloudflare.F(data.Proxied.ValueBool()),
		Search:    cloudflare.F(data.Search.ValueString()),
		Tag: cloudflare.F(dns.RecordListParamsTag{
			Absent:     cloudflare.F(data.Tag.Absent.ValueString()),
			Contains:   cloudflare.F(data.Tag.Contains.ValueString()),
			Endswith:   cloudflare.F(data.Tag.Endswith.ValueString()),
			Exact:      cloudflare.F(data.Tag.Exact.ValueString()),
			Present:    cloudflare.F(data.Tag.Present.ValueString()),
			Startswith: cloudflare.F(data.Tag.Startswith.ValueString()),
		}),
		TagMatch: cloudflare.F(dns.RecordListParamsTagMatch(data.TagMatch.ValueString())),
		Type:     cloudflare.F(dns.RecordListParamsType(data.Type.ValueString())),
	})
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
	data.Result = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
