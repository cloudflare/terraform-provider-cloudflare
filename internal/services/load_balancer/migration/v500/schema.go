package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareLoadBalancerSchema returns the legacy cloudflare_load_balancer schema (schema_version=1).
// This is used by UpgradeState to parse state from the legacy SDKv2 provider (v4).
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_load_balancer.go
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - ElementType for collections
// - Nested attributes structure
//
// Intentionally omitted (not needed for reading state):
// - Validators
// - PlanModifiers
// - Descriptions/MarkdownDescription
// - Default values
func SourceCloudflareLoadBalancerSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			// v4 field names (will be renamed in v5)
			"default_pool_ids": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"fallback_pool_id": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			// Int64 in v4, Float64 in v5
			"ttl": schema.Int64Attribute{
				Optional: true,
			},
			"session_affinity": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			// Int64 in v4, Float64 in v5
			"session_affinity_ttl": schema.Int64Attribute{
				Optional: true,
			},
			"proxied": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"steering_policy": schema.StringAttribute{
				Optional: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			// TypeList MaxItems:1 in v4 → single object in v5
			"session_affinity_attributes": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"samesite": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"secure": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						// Int64 in v4, Float64 in v5
						"drain_duration": schema.Int64Attribute{
							Optional: true,
							Computed: true,
						},
						"zero_downtime_failover": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"headers": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"require_all_headers": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			// TypeList MaxItems:1 in v4 → single object in v5
			"adaptive_routing": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"failover_across_pools": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			// TypeList MaxItems:1 in v4 → single object in v5
			"location_strategy": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"prefer_ecs": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"mode": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			// TypeList MaxItems:1 in v4 → single object in v5
			"random_steering": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"pool_weights": schema.MapAttribute{
							ElementType: types.Float64Type,
							Optional:    true,
						},
						"default_weight": schema.Float64Attribute{
							Optional: true,
						},
					},
				},
			},

			// TypeSet of blocks in v4 → map in v5
			"region_pools": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"region": schema.StringAttribute{
							Required: true,
						},
						"pool_ids": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},

			// TypeSet of blocks in v4 → map in v5
			"pop_pools": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"pop": schema.StringAttribute{
							Required: true,
						},
						"pool_ids": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},

			// TypeSet of blocks in v4 → map in v5
			"country_pools": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"country": schema.StringAttribute{
							Required: true,
						},
						"pool_ids": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},

			// Rules list
			"rules": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"priority": schema.Int64Attribute{
							Optional: true,
						},
						"disabled": schema.BoolAttribute{
							Optional: true,
						},
						"condition": schema.StringAttribute{
							Optional: true,
						},
						"terminates": schema.BoolAttribute{
							Optional: true,
						},
					},
					Blocks: map[string]schema.Block{
						// Rules.overrides - TypeList MaxItems:1 in v4
						"overrides": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"session_affinity": schema.StringAttribute{
										Optional: true,
									},
									// Int64 in v4, Float64 in v5
									"session_affinity_ttl": schema.Int64Attribute{
										Optional: true,
									},
									// Int64 in v4, Float64 in v5
									"ttl": schema.Int64Attribute{
										Optional: true,
									},
									"steering_policy": schema.StringAttribute{
										Optional: true,
									},
									"fallback_pool": schema.StringAttribute{
										Optional: true,
									},
									"default_pools": schema.ListAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
								Blocks: map[string]schema.Block{
									// Nested blocks within overrides - same as top-level
									"session_affinity_attributes": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"samesite": schema.StringAttribute{
													Optional: true,
												},
												"secure": schema.StringAttribute{
													Optional: true,
												},
												"zero_downtime_failover": schema.StringAttribute{
													Optional: true,
												},
												"headers": schema.ListAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"require_all_headers": schema.BoolAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"adaptive_routing": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"failover_across_pools": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									"location_strategy": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"prefer_ecs": schema.StringAttribute{
													Optional: true,
												},
												"mode": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									"random_steering": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"pool_weights": schema.MapAttribute{
													ElementType: types.Float64Type,
													Optional:    true,
												},
												"default_weight": schema.Float64Attribute{
													Optional: true,
												},
											},
										},
									},
									"region_pools": schema.SetNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"region": schema.StringAttribute{
													Required: true,
												},
												"pool_ids": schema.ListAttribute{
													ElementType: types.StringType,
													Required:    true,
												},
											},
										},
									},
									"pop_pools": schema.SetNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"pop": schema.StringAttribute{
													Required: true,
												},
												"pool_ids": schema.ListAttribute{
													ElementType: types.StringType,
													Required:    true,
												},
											},
										},
									},
									"country_pools": schema.SetNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"country": schema.StringAttribute{
													Required: true,
												},
												"pool_ids": schema.ListAttribute{
													ElementType: types.StringType,
													Required:    true,
												},
											},
										},
									},
								},
							},
						},

						// Rules.fixed_response - TypeList MaxItems:1 in v4
						"fixed_response": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"message_body": schema.StringAttribute{
										Optional: true,
									},
									"status_code": schema.Int64Attribute{
										Optional: true,
									},
									"content_type": schema.StringAttribute{
										Optional: true,
									},
									"location": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
