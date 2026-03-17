package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceSchemaV0 returns the schema for version 0 (earliest v5 releases: v5.10, v5.11).
// Schema version: 0
// Resource type: cloudflare_account_token
//
// This schema is used only for reading v0 state during migration.
// Validators, PlanModifiers (except UseStateForUnknown on id), and Descriptions
// are intentionally minimal — only the structural shape matters for state parsing.
//
// Key differences from v500:
// - policies is ListNestedAttribute (not Set)
// - policies[].id exists (computed, removed in v500)
// - policies[].permission_groups is SetNestedAttribute with meta + name fields
// - policies[].resources is MapAttribute (not StringAttribute)
// - status is Optional+Computed (same, but included for completeness)
//
// Reference: cloudflare/cloudflare provider v5.10.0 / v5.11.0 schema
func SourceSchemaV0() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
			"policies": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
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
									"meta": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Optional: true,
											},
											"value": schema.StringAttribute{
												Optional: true,
											},
										},
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"resources": schema.MapAttribute{
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"value": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
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

// SourceSchemaV1 returns the schema for version 1 (dormant v5 state before migration activation).
// Schema version: 1
// Resource type: cloudflare_account_token
//
// This schema is compatible with v500 — the upgrade is a no-op (version bump only).
//
// Key differences from v0:
// - policies is SetNestedAttribute (not List)
// - policies[].id is removed
// - policies[].permission_groups has only id (no meta, no name)
// - policies[].resources is StringAttribute (JSON string, not Map)
func SourceSchemaV1() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
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
				Computed:  true,
				Sensitive: true,
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
