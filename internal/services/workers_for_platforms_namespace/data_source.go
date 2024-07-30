// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WorkersForPlatformsNamespaceDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &WorkersForPlatformsNamespaceDataSource{}

func NewWorkersForPlatformsNamespaceDataSource() datasource.DataSource {
	return &WorkersForPlatformsNamespaceDataSource{}
}

func (d *WorkersForPlatformsNamespaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_for_platforms_namespace"
}

func (r *WorkersForPlatformsNamespaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *WorkersForPlatformsNamespaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WorkersForPlatformsNamespaceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		res := new(http.Response)
		env := WorkersForPlatformsNamespaceResultDataSourceEnvelope{*data}
		_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Get(
			ctx,
			data.DispatchNamespace.ValueString(),
			workers_for_platforms.DispatchNamespaceGetParams{
				AccountID: cloudflare.F(data.AccountID.ValueString()),
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
		items := &[]*WorkersForPlatformsNamespaceDataSourceModel{}
		env := WorkersForPlatformsNamespaceResultListDataSourceEnvelope{items}

		page, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.List(ctx, workers_for_platforms.DispatchNamespaceListParams{
			AccountID: cloudflare.F(data.Filter.AccountID.ValueString()),
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
