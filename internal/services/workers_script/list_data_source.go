// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WorkersScriptsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*WorkersScriptsDataSource)(nil)

func NewWorkersScriptsDataSource() datasource.DataSource {
	return &WorkersScriptsDataSource{}
}

func (d *WorkersScriptsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_scripts"
}

func (d *WorkersScriptsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *WorkersScriptsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WorkersScriptsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	items := customfield.NullObjectList[WorkersScriptsResultDataSourceModel](ctx)
	env := WorkersScriptsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}

	page, err := d.client.Workers.Scripts.List(ctx, params)
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
		acc = append(acc, items.Elements()...)
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
	result, diags := customfield.NewObjectListFromAttributes[WorkersScriptsResultDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Result = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
