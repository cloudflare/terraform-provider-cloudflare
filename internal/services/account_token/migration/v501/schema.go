package v501

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceSchemaV500 returns the schema for version 500 (Set-based policies).
// Schema version: 500
// Resource type: cloudflare_account_token
//
// This schema is used only for reading v500 state during migration to v501.
// The key difference from v501 is that policies and permission_groups use
// SetNestedAttribute instead of ListNestedAttribute.
func SourceSchemaV500() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"issued_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"policies": schema.SetNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"effect": schema.StringAttribute{
							Required: true,
						},
						"permission_groups": schema.SetNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
						"resources": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"value": schema.StringAttribute{
				Computed:      true,
				Sensitive:     true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"not_before": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"expires_on": schema.StringAttribute{
				Optional:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"condition": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"request_ip": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"in": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
							"not_in": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"last_used_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}
