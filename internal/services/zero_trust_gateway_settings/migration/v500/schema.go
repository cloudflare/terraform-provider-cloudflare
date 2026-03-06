package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceV4ZeroTrustGatewaySettingsSchema returns the schema for the v4
// cloudflare_teams_account resource (SDKv2 format, schema_version=0).
//
// This schema is used as PriorSchema in the state upgrader to parse v4 state.
// All TypeList MaxItems:1 blocks from v4 are represented as ListNestedAttribute.
func SourceV4ZeroTrustGatewaySettingsSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},

			// Flat boolean fields
			"activity_log_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"tls_decrypt_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"protocol_detection_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"url_browser_isolation_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"non_identity_browser_isolation_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},

			// MaxItems:1 blocks as ListNestedAttribute (SDKv2 stores TypeList as arrays)
			"block_page": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"background_color": schema.StringAttribute{
							Optional: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						"footer_text": schema.StringAttribute{
							Optional: true,
						},
						"header_text": schema.StringAttribute{
							Optional: true,
						},
						"logo_path": schema.StringAttribute{
							Optional: true,
						},
						"mailto_address": schema.StringAttribute{
							Optional: true,
						},
						"mailto_subject": schema.StringAttribute{
							Optional: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},

			"body_scanning": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"inspection_mode": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},

			"fips": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tls": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},

			"antivirus": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled_download_phase": schema.BoolAttribute{
							Required: true,
						},
						"enabled_upload_phase": schema.BoolAttribute{
							Required: true,
						},
						"fail_closed": schema.BoolAttribute{
							Required: true,
						},
						// notification_settings is also a TypeList MaxItems:1 in v4
						"notification_settings": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Optional: true,
									},
									// v4 field name is "message"; v5 renames it to "msg"
									"message": schema.StringAttribute{
										Optional: true,
									},
									"support_url": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"extended_email_matching": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},

			"custom_certificate": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Required: true,
						},
						"id": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						// updated_at is plain string in v4 (no CustomType)
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},

			"certificate": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},

			// Blocks dropped in v5 - included for state parsing only
			"logging": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"redact_pii": schema.BoolAttribute{
							Required: true,
						},
						"settings_by_rule_type": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"dns": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"log_all": schema.BoolAttribute{
													Required: true,
												},
												"log_blocks": schema.BoolAttribute{
													Required: true,
												},
											},
										},
									},
									"http": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"log_all": schema.BoolAttribute{
													Required: true,
												},
												"log_blocks": schema.BoolAttribute{
													Required: true,
												},
											},
										},
									},
									"l4": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"log_all": schema.BoolAttribute{
													Required: true,
												},
												"log_blocks": schema.BoolAttribute{
													Required: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"proxy": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tcp": schema.BoolAttribute{
							Required: true,
						},
						"udp": schema.BoolAttribute{
							Required: true,
						},
						"root_ca": schema.BoolAttribute{
							Required: true,
						},
						"virtual_ip": schema.BoolAttribute{
							Required: true,
						},
						"disable_for_time": schema.Int64Attribute{
							Required: true,
						},
					},
				},
			},

			"ssh_session_log": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"public_key": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},

			"payload_log": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"public_key": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}
