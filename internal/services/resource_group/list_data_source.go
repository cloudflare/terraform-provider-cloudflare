// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ResourceGroupsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*ResourceGroupsDataSource)(nil)

func NewResourceGroupsDataSource() datasource.DataSource {
	return &ResourceGroupsDataSource{}
}

func (d *ResourceGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_groups"
}

func (d *ResourceGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ResourceGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ResourceGroupsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := ResourceGroupsResultListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	page, err := d.client.IAM.ResourceGroups.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, env.Result.Elements()...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

    uniqueAcc := []attr.Value{}
    seenIDs := make(map[string]bool)

    for _, item := range acc {
        obj, ok := item.(types.Object)
        if !ok {
            continue
        }

        var itemModel ResourceGroupsResultDataSourceModel

        diags := obj.As(ctx, &itemModel, basetypes.ObjectAsOptions{
            UnhandledNullAsEmpty:    true,
            UnhandledUnknownAsEmpty: true,
        })

        if diags.HasError() {
            continue
        }

        id := itemModel.ID.ValueString()
        if !seenIDs[id] {
            uniqueAcc = append(uniqueAcc, item)
            seenIDs[id] = true
        }
    }
    acc = uniqueAcc

	acc = acc[:min(len(acc), maxItems)]
	result, diags := customfield.NewObjectListFromAttributes[ResourceGroupsResultDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Result = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
