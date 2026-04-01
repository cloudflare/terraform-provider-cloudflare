package origin_cloud_regions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const apiPath = "/zones/%s/cache/origin_public_cloud_region"
const apiPathWithIP = "/zones/%s/cache/origin_public_cloud_region/%s"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*OriginCloudRegionsResource)(nil)
var _ resource.ResourceWithImportState = (*OriginCloudRegionsResource)(nil)

type OriginCloudRegionsResource struct {
	client *cloudflare.Client
}

func NewResource() resource.Resource {
	return &OriginCloudRegionsResource{}
}

func (r *OriginCloudRegionsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_origin_cloud_regions"
}

func (r *OriginCloudRegionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OriginCloudRegionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OriginCloudRegionsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := data.ZoneID.ValueString()
	for _, m := range data.Mappings {
		body := OriginCloudRegionUpsertRequest{
			IP:     m.OriginIP.ValueString(),
			Vendor: m.Vendor.ValueString(),
			Region: m.Region.ValueString(),
		}
		if err := r.patchMapping(ctx, zoneID, body); err != nil {
			resp.Diagnostics.AddError("failed to create origin cloud region mapping", err.Error())
			return
		}
	}

	data.ID = types.StringValue(zoneID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OriginCloudRegionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OriginCloudRegionsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	mappings, err := r.readMappings(ctx, data.ZoneID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to read origin cloud region mappings", err.Error())
		return
	}

	data.Mappings = mappings
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OriginCloudRegionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OriginCloudRegionsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := plan.ZoneID.ValueString()

	// Index desired state by IP — reused for both upsert and delete passes.
	planByIP := make(map[string]OriginCloudRegionMappingModel, len(plan.Mappings))
	for _, m := range plan.Mappings {
		planByIP[m.OriginIP.ValueString()] = m
	}

	// Index current state by IP — used to detect new vs changed entries.
	stateByIP := make(map[string]OriginCloudRegionMappingModel, len(state.Mappings))
	for _, m := range state.Mappings {
		stateByIP[m.OriginIP.ValueString()] = m
	}

	// PATCH new and changed entries — PATCH upserts (creates if not exists, updates if exists).
	for ip, pm := range planByIP {
		sm, exists := stateByIP[ip]
		if !exists || sm.Vendor != pm.Vendor || sm.Region != pm.Region {
			body := OriginCloudRegionUpsertRequest{
				IP:     ip,
				Vendor: pm.Vendor.ValueString(),
				Region: pm.Region.ValueString(),
			}
			if err := r.patchMapping(ctx, zoneID, body); err != nil {
				resp.Diagnostics.AddError("failed to upsert origin cloud region mapping", err.Error())
				return
			}
		}
	}

	// DELETE removed entries — planByIP reused here.
	for _, m := range state.Mappings {
		if _, exists := planByIP[m.OriginIP.ValueString()]; !exists {
			if err := r.deleteMapping(ctx, zoneID, m.OriginIP.ValueString()); err != nil {
				resp.Diagnostics.AddError("failed to delete origin cloud region mapping", err.Error())
				return
			}
		}
	}

	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *OriginCloudRegionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OriginCloudRegionsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := data.ZoneID.ValueString()
	for _, m := range data.Mappings {
		if err := r.deleteMapping(ctx, zoneID, m.OriginIP.ValueString()); err != nil {
			resp.Diagnostics.AddError("failed to delete origin cloud region mapping", err.Error())
			return
		}
	}
}

func (r *OriginCloudRegionsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data OriginCloudRegionsModel

	path := ""
	diags := importpath.ParseImportID(req.ID, "<zone_id>", &path)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path)

	mappings, err := r.readMappings(ctx, path)
	if err != nil {
		resp.Diagnostics.AddError("failed to read origin cloud region mappings during import", err.Error())
		return
	}

	data.ID = data.ZoneID
	data.Mappings = mappings
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// readMappings fetches the current list of mappings from the API.
func (r *OriginCloudRegionsResource) readMappings(ctx context.Context, zoneID string) ([]OriginCloudRegionMappingModel, error) {
	res := new(http.Response)
	err := r.client.Get(
		ctx,
		fmt.Sprintf(apiPath, zoneID),
		nil,
		&res,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var env OriginCloudRegionsResultEnvelope
	if err := json.Unmarshal(body, &env); err != nil {
		return nil, fmt.Errorf("failed to deserialize response: %w", err)
	}

	mappings := make([]OriginCloudRegionMappingModel, 0, len(env.Result.Value))
	for _, entry := range env.Result.Value {
		mappings = append(mappings, OriginCloudRegionMappingModel{
			OriginIP: types.StringValue(entry.OriginIP),
			Vendor:   types.StringValue(entry.Vendor),
			Region:   types.StringValue(entry.Region),
		})
	}
	return mappings, nil
}

// patchMapping issues a PATCH for an existing mapping.
func (r *OriginCloudRegionsResource) patchMapping(ctx context.Context, zoneID string, body OriginCloudRegionUpsertRequest) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return r.client.Patch(
		ctx,
		fmt.Sprintf(apiPath, zoneID),
		nil,
		nil,
		option.WithRequestBody("application/json", bytes.NewReader(bodyBytes)),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
}

// deleteMapping issues DELETE for a single IP entry.
func (r *OriginCloudRegionsResource) deleteMapping(ctx context.Context, zoneID, ip string) error {
	return r.client.Delete(
		ctx,
		fmt.Sprintf(apiPathWithIP, zoneID, ip),
		nil,
		nil,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
}
