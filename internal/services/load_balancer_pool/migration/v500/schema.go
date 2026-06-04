// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareLoadBalancerPoolSchema returns the source schema for legacy load_balancer_pool resource.
// Schema version: 0 (v4 had no explicit schema version, defaults to 0)
// Resource type: cloudflare_load_balancer_pool
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Note: SDK v2 storage quirks:
// - TypeSet fields are stored as arrays in state
// - TypeList MaxItems:1 fields are stored as arrays with single element: [{...}]
func SourceCloudflareLoadBalancerPoolSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 had no explicit schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"origins": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"address": schema.StringAttribute{
							Required: true,
						},
						"virtual_network_id": schema.StringAttribute{
							Optional: true,
						},
						"weight": schema.Float64Attribute{
							Optional: true,
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"header": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"header": schema.StringAttribute{
										Required: true,
									},
									"values": schema.SetAttribute{
										Required:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"minimum_origins": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"latitude": schema.Float64Attribute{
				Optional: true,
			},
			"longitude": schema.Float64Attribute{
				Optional: true,
			},
			"check_regions": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"monitor": schema.StringAttribute{
				Optional: true,
			},
			"notification_email": schema.StringAttribute{
				Optional: true,
			},
			"load_shedding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_percent": schema.Float64Attribute{
							Optional: true,
							Computed: true,
						},
						"default_policy": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"session_percent": schema.Float64Attribute{
							Optional: true,
							Computed: true,
						},
						"session_policy": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"origin_steering": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"policy": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
