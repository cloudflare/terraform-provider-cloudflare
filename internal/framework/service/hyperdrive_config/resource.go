package hyperdrive_config

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HyperdriveConfigResource{}
var _ resource.ResourceWithImportState = &HyperdriveConfigResource{}

func NewResource() resource.Resource {
	return &HyperdriveConfigResource{}
}

// HyperdriveConfigResource defines the resource implementation for hyperdrive configs.
type HyperdriveConfigResource struct {
	client *cloudflare.API
}

func (r *HyperdriveConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hyperdrive_config"
}

func (r *HyperdriveConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *HyperdriveConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *HyperdriveConfigModel
	var caching *HyperdriveConfigCachingModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(data.Caching.As(ctx, &caching, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	if resp.Diagnostics.HasError() {
		return
	}

	config := buildHyperdriveConfigFromModel(data, caching)

	createHyperdriveConfig, err := r.client.CreateHyperdriveConfig(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()),
		cloudflare.CreateHyperdriveConfigParams{
			Name: config.Name,
			Origin: cloudflare.HyperdriveConfigOrigin{
				Database: config.Origin.Database,
				Password: config.Origin.Password,
				Host:     config.Origin.Host,
				Port:     config.Origin.Port,
				Scheme:   config.Origin.Scheme,
				User:     config.Origin.User,
			},
			Caching: cloudflare.HyperdriveConfigCaching{
				Disabled: config.Caching.Disabled,
			},
		})
	if err != nil {
		resp.Diagnostics.AddError("Error creating hyperdrive config", err.Error())
		return
	}

	var diags diag.Diagnostics
	data, diags = buildHyperdriveConfigModelFromHyperdriveConfig(ctx, data, createHyperdriveConfig)
	resp.Diagnostics.Append(diags...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HyperdriveConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *HyperdriveConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	config, err := r.client.GetHyperdriveConfig(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Error reading hyperdrive config", err.Error())
		return
	}

	var diags diag.Diagnostics
	data, diags = buildHyperdriveConfigModelFromHyperdriveConfig(ctx, data, config)
	resp.Diagnostics.Append(diags...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HyperdriveConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *HyperdriveConfigModel
	var caching *HyperdriveConfigCachingModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(data.Caching.As(ctx, &caching, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)

	if resp.Diagnostics.HasError() {
		return
	}

	config := buildHyperdriveConfigFromModel(data, caching)

	updatedConfig, err := r.client.UpdateHyperdriveConfig(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), cloudflare.UpdateHyperdriveConfigParams{
		Name: config.Name,
		Origin: cloudflare.HyperdriveConfigOrigin{
			Database: config.Origin.Database,
			Password: config.Origin.Password,
			Host:     config.Origin.Host,
			Port:     config.Origin.Port,
			Scheme:   config.Origin.Scheme,
			User:     config.Origin.User,
		},
		Caching: cloudflare.HyperdriveConfigCaching{
			Disabled: config.Caching.Disabled,
		},
	})

	if err != nil {
		resp.Diagnostics.AddError("Error updating hyperdrive config", err.Error())
		return
	}

	var diags diag.Diagnostics
	data, diags = buildHyperdriveConfigModelFromHyperdriveConfig(ctx, data, updatedConfig)
	resp.Diagnostics.Append(diags...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HyperdriveConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *HyperdriveConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteHyperdriveConfig(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting hyperdrive config", err.Error())
	}
}

func (r *HyperdriveConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError("Error importing hyperdrive config", "Invalid ID specified. Please specify the ID as \"accounts_id/hyperdrive_config_id\"")
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func buildHyperdriveConfigFromModel(config *HyperdriveConfigModel, caching *HyperdriveConfigCachingModel) cloudflare.HyperdriveConfig {
	built := cloudflare.HyperdriveConfig{
		Name: config.Name.ValueString(),
		Origin: cloudflare.HyperdriveConfigOrigin{
			Database: config.Origin.Database.ValueString(),
			Password: config.Origin.Password.ValueString(),
			Host:     config.Origin.Host.ValueString(),
			Port:     int(config.Origin.Port.ValueInt64()),
		},
	}

	if !config.Origin.Scheme.IsNull() {
		built.Origin.Scheme = config.Origin.Scheme.ValueString()
	}

	if !config.Origin.User.IsNull() {
		built.Origin.User = config.Origin.User.ValueString()
	}

	built.Caching = cloudflare.HyperdriveConfigCaching{}

	if caching != nil && !caching.Disabled.IsNull() {
		built.Caching.Disabled = cloudflare.BoolPtr(caching.Disabled.ValueBool())
	}

	return built
}

func buildHyperdriveConfigModelFromHyperdriveConfig(ctx context.Context, data *HyperdriveConfigModel, config cloudflare.HyperdriveConfig) (*HyperdriveConfigModel, diag.Diagnostics) {
	var scheme = flatteners.String("postgres")
	if data.Origin != nil {
		scheme = data.Origin.Scheme
	}

	var password = flatteners.String("")
	if data.Origin != nil {
		password = data.Origin.Password
	}

	var caching, diags = types.ObjectValueFrom(
		ctx,
		HyperdriveConfigCachingModel{}.AttributeTypes(),
		HyperdriveConfigCachingModel{
			Disabled: types.BoolValue(*config.Caching.Disabled),
		},
	)

	built := HyperdriveConfigModel{
		AccountID: data.AccountID,
		ID:        flatteners.String(config.ID),
		Name:      flatteners.String(config.Name),
		Origin: &HyperdriveConfigOriginModel{
			Database: flatteners.String(config.Origin.Database),
			Host:     flatteners.String(config.Origin.Host),
			Port:     flatteners.Int64(int64(config.Origin.Port)),
			User:     flatteners.String(config.Origin.User),
			Scheme:   scheme,
			Password: password,
		},
		Caching: caching,
	}

	return &built, diags
}
