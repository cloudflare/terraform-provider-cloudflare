// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type DNSRecordsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = (*DNSRecordsDataSource)(nil)

func NewDNSRecordsDataSource() datasource.DataSource {
	return &DNSRecordsDataSource{}
}

func (d *DNSRecordsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_records"
}

func (d *DNSRecordsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DNSRecordsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DNSRecordsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := DNSRecordsResultListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	page, err := d.client.DNS.Records.List(ctx, params)
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

	acc = acc[:min(len(acc), maxItems)]

	// Records returned in the same list may have differing concrete underlying
	// types for `data.flags` (e.g. CAA records carry a number, A records omit
	// the field and decode as a typeless dynamic null). Terraform's list
	// serialiser requires every element of a list to share an identical type,
	// so the differing per-record types crash the data source with
	// "inconsistent list element types". Normalise the dynamic flags so every
	// element exposes the same underlying type before constructing the list.
	// See https://github.com/cloudflare/terraform-provider-cloudflare/issues/7004
	normalizeFlagsTypes(ctx, acc)

	result, diags := customfield.NewObjectListFromAttributes[DNSRecordsResultDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Result = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// normalizeFlagsTypes ensures every element's `data.flags` exposes the same
// concrete underlying type. When any sibling carries a known concrete value, an
// otherwise typeless dynamic null is rewrapped as a typed null of that type so
// the resulting list has a uniform element type.
func normalizeFlagsTypes(ctx context.Context, acc []attr.Value) {
	if len(acc) < 2 {
		return
	}

	models := make([]*DNSRecordsResultDataSourceModel, len(acc))

	var concreteType attr.Type
	for i, el := range acc {
		obj, ok := el.(customfield.NestedObject[DNSRecordsResultDataSourceModel])
		if !ok {
			return
		}
		m, diags := obj.Value(ctx)
		if diags.HasError() || m == nil {
			return
		}
		models[i] = m
		if concreteType != nil {
			continue
		}
		data, diags := m.Data.Value(ctx)
		if diags.HasError() || data == nil {
			continue
		}
		dyn, _ := data.Flags.ToDynamicValue(ctx)
		if dyn.IsNull() || dyn.IsUnknown() {
			continue
		}
		if underlying := dyn.UnderlyingValue(); underlying != nil && !underlying.IsNull() && !underlying.IsUnknown() {
			concreteType = underlying.Type(ctx)
		}
	}

	if concreteType == nil {
		return
	}

	typedNull, err := makeTypedDynamicNull(concreteType)
	if err {
		return
	}

	changed := false
	for _, m := range models {
		data, diags := m.Data.Value(ctx)
		if diags.HasError() || data == nil {
			continue
		}
		dyn, _ := data.Flags.ToDynamicValue(ctx)
		if !dyn.IsNull() {
			if underlying := dyn.UnderlyingValue(); underlying != nil && !underlying.IsNull() {
				continue
			}
		}
		data.Flags = typedNull
		m.Data = customfield.NewObjectMust(ctx, data)
		changed = true
	}

	if !changed {
		return
	}

	for i, m := range models {
		obj, diags := customfield.NewObject(ctx, m)
		if diags.HasError() {
			return
		}
		acc[i] = obj
	}
}

// makeTypedDynamicNull wraps a typed null of the given concrete type into a
// NormalizedDynamicValue so that ToTerraformValue serialises a typed null
// rather than the unconstrained DynamicPseudoType.
func makeTypedDynamicNull(t attr.Type) (customfield.NormalizedDynamicValue, bool) {
	var inner attr.Value
	switch t.(type) {
	case basetypes.Float64Type:
		inner = types.Float64Null()
	case basetypes.StringType:
		inner = types.StringNull()
	default:
		return customfield.NormalizedDynamicValue{}, true
	}
	return customfield.RawNormalizedDynamicValueFrom(inner), false
}
