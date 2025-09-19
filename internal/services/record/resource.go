// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package record

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*RecordResource)(nil)
var _ resource.ResourceWithModifyPlan = (*RecordResource)(nil)
var _ resource.ResourceWithImportState = (*RecordResource)(nil)

func NewResource() resource.Resource {
	return &RecordResource{}
}

// RecordResource defines the resource implementation.
type RecordResource struct {
	client *cloudflare.Client
}

func (r *RecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record"
}

func (r *RecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provides a Cloudflare DNS Record resource. This is the legacy resource that only supports Read operations.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the DNS record.",
			},
			"zone_id": schema.StringAttribute{
				Required:    true,
				Description: "The Zone ID to use for this endpoint.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the record.",
			},
			"hostname": schema.StringAttribute{
				Computed:    true,
				Description: "The FQDN of the record.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the record.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "The value of the record.",
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Description: "The content of the record.",
			},
			"data": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Map of attributes that constitute the record value.",
			},
			"ttl": schema.Int64Attribute{
				Computed:    true,
				Description: "The TTL of the record.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "The priority of the record.",
			},
			"proxied": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the record gets Cloudflare's origin protection.",
			},
			"created_on": schema.StringAttribute{
				Computed:    true,
				Description: "The RFC3339 timestamp of when the record was created.",
			},
			"metadata": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "A key-value map of string metadata Cloudflare associates with the record.",
			},
			"modified_on": schema.StringAttribute{
				Computed:    true,
				Description: "The RFC3339 timestamp of when the record was last modified.",
			},
			"proxiable": schema.BoolAttribute{
				Computed:    true,
				Description: "Shows whether this record can be proxied.",
			},
			"allow_overwrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Allow creation of this record in Terraform to overwrite an existing record, if any.",
			},
			"comment": schema.StringAttribute{
				Computed:    true,
				Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Custom tags for the DNS record.",
			},
		},
	}
}

func (r *RecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Create operation not supported",
		"This resource only supports Read operations. Create operations are not supported for the legacy cloudflare_record resource.",
	)
}

func (r *RecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RecordModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DNS.Records.Get(
		ctx,
		data.ID.ValueString(),
		dns.RecordGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithResponseBodyInto(&resp.State),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
}

func (r *RecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update operation not supported",
		"This resource only supports Read operations. Update operations are not supported for the legacy cloudflare_record resource.",
	)
}

func (r *RecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddError(
		"Delete operation not supported",
		"This resource only supports Read operations. Delete operations are not supported for the legacy cloudflare_record resource.",
	)
}

func (r *RecordResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// No modification needed for read-only resource
}

func (r *RecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Simple import that just sets the ID
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
