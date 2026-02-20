package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceAccessApplicationSchema returns the v4 cloudflare_access_application schema.
// This is used by MoveState and UpgradeFromV4 to parse the source state from v4 provider.
func SourceAccessApplicationSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"session_duration": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"auto_redirect_to_identity": schema.BoolAttribute{
				Optional: true,
			},
			"enable_binding_cookie": schema.BoolAttribute{
				Optional: true,
			},
			"http_only_cookie_attribute": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"same_site_cookie_attribute": schema.StringAttribute{
				Optional: true,
			},
			"logo_url": schema.StringAttribute{
				Optional: true,
			},
			"skip_interstitial": schema.BoolAttribute{
				Optional: true,
			},
			"app_launcher_visible": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"service_auth_401_redirect": schema.BoolAttribute{
				Optional: true,
			},
			"custom_deny_message": schema.StringAttribute{
				Optional: true,
			},
			"custom_deny_url": schema.StringAttribute{
				Optional: true,
			},
			"custom_non_identity_deny_url": schema.StringAttribute{
				Optional: true,
			},
			"allowed_idps": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"self_hosted_domains": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"custom_pages": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"options_preflight_bypass": schema.BoolAttribute{
				Optional: true,
			},
			"path_cookie_attribute": schema.BoolAttribute{
				Optional: true,
			},
			"aud": schema.StringAttribute{
				Computed: true,
			},
			// v4 stores policies as a simple string array
			"policies": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			// Additional fields from v4
			"app_launcher_logo_url": schema.StringAttribute{
				Optional: true,
			},
			"header_bg_color": schema.StringAttribute{
				Optional: true,
			},
			"bg_color": schema.StringAttribute{
				Optional: true,
			},
			"skip_app_launcher_login_page": schema.BoolAttribute{
				Optional: true,
			},
			"allow_authenticate_via_warp": schema.BoolAttribute{
				Optional: true,
			},
			// Deprecated/removed in v5
			"domain_type": schema.StringAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			// v4 cors_headers is a list block with MaxItems: 1
			"cors_headers": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"allow_all_headers": schema.BoolAttribute{
							Optional: true,
						},
						"allow_all_methods": schema.BoolAttribute{
							Optional: true,
						},
						"allow_all_origins": schema.BoolAttribute{
							Optional: true,
						},
						"allow_credentials": schema.BoolAttribute{
							Optional: true,
						},
						"allowed_headers": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"allowed_methods": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"allowed_origins": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"max_age": schema.Int64Attribute{
							Optional: true,
						},
					},
				},
			},
			// v4 saas_app is a list block with MaxItems: 1
			"saas_app": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"auth_type": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"consumer_service_url": schema.StringAttribute{
							Optional: true,
						},
						"sp_entity_id": schema.StringAttribute{
							Optional: true,
						},
						"idp_entity_id": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"public_key": schema.StringAttribute{
							Computed: true,
						},
						"name_id_format": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"name_id_transform_jsonata": schema.StringAttribute{
							Optional: true,
						},
						"saml_attribute_transform_jsonata": schema.StringAttribute{
							Optional: true,
						},
						"default_relay_state": schema.StringAttribute{
							Optional: true,
						},
						"sso_endpoint": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"app_launcher_url": schema.StringAttribute{
							Optional: true,
						},
						"client_id": schema.StringAttribute{
							Computed: true,
						},
						"client_secret": schema.StringAttribute{
							Computed:  true,
							Sensitive: true,
						},
						"access_token_lifetime": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"allow_pkce_without_client_secret": schema.BoolAttribute{
							Optional: true,
						},
						"group_filter_regex": schema.StringAttribute{
							Optional: true,
						},
						"grant_types": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
						"redirect_uris": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"scopes": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
					},
					Blocks: map[string]schema.Block{
						"custom_attribute": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Optional: true,
									},
									"friendly_name": schema.StringAttribute{
										Optional: true,
									},
									"name_format": schema.StringAttribute{
										Optional: true,
									},
									"required": schema.BoolAttribute{
										Optional: true,
									},
								},
								Blocks: map[string]schema.Block{
									"source": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional: true,
												},
												"name_by_idp": schema.MapAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"custom_claim": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Optional: true,
									},
									"required": schema.BoolAttribute{
										Optional: true,
									},
									"scope": schema.StringAttribute{
										Optional: true,
									},
								},
								Blocks: map[string]schema.Block{
									"source": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional: true,
												},
												"name_by_idp": schema.MapAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"hybrid_and_implicit_options": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"return_access_token_from_authorization_endpoint": schema.BoolAttribute{
										Optional: true,
									},
									"return_id_token_from_authorization_endpoint": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
						"refresh_token_options": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"lifetime": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			// v4 scim_config is a list block with MaxItems: 1
			"scim_config": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"idp_uid": schema.StringAttribute{
							Required: true,
						},
						"remote_uri": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						"deactivate_on_delete": schema.BoolAttribute{
							Optional: true,
						},
					},
					Blocks: map[string]schema.Block{
						"authentication": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"scheme": schema.StringAttribute{
										Required: true,
									},
									"user": schema.StringAttribute{
										Optional: true,
									},
									"password": schema.StringAttribute{
										Optional:  true,
										Sensitive: true,
									},
									"token": schema.StringAttribute{
										Optional:  true,
										Sensitive: true,
									},
									"authorization_url": schema.StringAttribute{
										Optional: true,
									},
									"client_id": schema.StringAttribute{
										Optional: true,
									},
									"client_secret": schema.StringAttribute{
										Optional:  true,
										Sensitive: true,
									},
									"token_url": schema.StringAttribute{
										Optional: true,
									},
									"scopes": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
							},
						},
						"mappings": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"schema": schema.StringAttribute{
										Required: true,
									},
									"enabled": schema.BoolAttribute{
										Optional: true,
									},
									"filter": schema.StringAttribute{
										Optional: true,
									},
									"transform_jsonata": schema.StringAttribute{
										Optional: true,
									},
									"strictness": schema.StringAttribute{
										Optional: true,
									},
								},
								Blocks: map[string]schema.Block{
									"operations": schema.ListNestedBlock{
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"create": schema.BoolAttribute{
													Optional: true,
												},
												"update": schema.BoolAttribute{
													Optional: true,
												},
												"delete": schema.BoolAttribute{
													Optional: true,
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
			// v4 landing_page_design is a list block with MaxItems: 1
			"landing_page_design": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"button_color": schema.StringAttribute{
							Optional: true,
						},
						"button_text_color": schema.StringAttribute{
							Optional: true,
						},
						"image_url": schema.StringAttribute{
							Optional: true,
						},
						"message": schema.StringAttribute{
							Optional: true,
						},
						"title": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			// v4 footer_links is a set block
			"footer_links": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"url": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			// v4 destinations is a list block
			"destinations": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Optional: true,
						},
						"uri": schema.StringAttribute{
							Optional: true,
						},
						"hostname": schema.StringAttribute{
							Optional: true,
						},
						"cidr": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"port_range": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"vnet_id": schema.StringAttribute{
							Optional: true,
						},
						"l4_protocol": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			// v4 target_criteria is a list block
			"target_criteria": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Required: true,
						},
						"protocol": schema.StringAttribute{
							Required: true,
						},
					},
					Blocks: map[string]schema.Block{
						"target_attributes": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
									},
									"values": schema.ListAttribute{
										ElementType: types.StringType,
										Required:    true,
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
