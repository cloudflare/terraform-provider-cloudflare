package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceV4RulesetSchema returns the schema for v4 cloudflare_ruleset resources (Plugin Framework format).
// This schema represents the state structure from the v4 provider (schema_version=1).
//
// In Plugin Framework, ListNestedBlock with SizeAtMost(1) is stored as an array in state JSON.
// We represent those blocks as ListNestedAttribute here so the framework can deserialize
// the state correctly. No validators are needed for migration schemas.
func SourceV4RulesetSchema() schema.Schema {
	return schema.Schema{
		Version: 1,
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
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"kind": schema.StringAttribute{
				Required: true,
			},
			"phase": schema.StringAttribute{
				Required: true,
			},
			// rules is a ListNestedBlock in v4 (no MaxItems restriction)
			"rules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"ref": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
							Computed: true,
						},
						"description": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"expression": schema.StringAttribute{
							Required: true,
						},
						"action": schema.StringAttribute{
							Optional: true,
						},
						// action_parameters is a ListNestedBlock MaxItems:1 in v4 — stored as array
						"action_parameters": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"additional_cacheable_ports": schema.SetAttribute{
										ElementType: types.Int64Type,
										Optional:    true,
									},
									"automatic_https_rewrites": schema.BoolAttribute{
										Optional: true,
									},
									"bic": schema.BoolAttribute{
										Optional: true,
									},
									"cache": schema.BoolAttribute{
										Optional: true,
									},
									"content": schema.StringAttribute{
										Optional: true,
									},
									"content_type": schema.StringAttribute{
										Optional: true,
									},
									"cookie_fields": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"disable_apps": schema.BoolAttribute{
										Optional: true,
									},
									"disable_railgun": schema.BoolAttribute{
										Optional: true,
									},
									"disable_zaraz": schema.BoolAttribute{
										Optional: true,
									},
									"disable_rum": schema.BoolAttribute{
										Optional: true,
									},
									"fonts": schema.BoolAttribute{
										Optional: true,
									},
									"email_obfuscation": schema.BoolAttribute{
										Optional: true,
									},
									"host_header": schema.StringAttribute{
										Optional: true,
									},
									"hotlink_protection": schema.BoolAttribute{
										Optional: true,
									},
									"id": schema.StringAttribute{
										Optional: true,
									},
									"increment": schema.Int64Attribute{
										Optional: true,
									},
									"mirage": schema.BoolAttribute{
										Optional: true,
									},
									"opportunistic_encryption": schema.BoolAttribute{
										Optional: true,
									},
									"origin_cache_control": schema.BoolAttribute{
										Optional: true,
									},
									"origin_error_page_passthru": schema.BoolAttribute{
										Optional: true,
									},
									"phases": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"polish": schema.StringAttribute{
										Optional: true,
									},
									"products": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"read_timeout": schema.Int64Attribute{
										Optional: true,
									},
									"request_fields": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"respect_strong_etags": schema.BoolAttribute{
										Optional: true,
									},
									"response_fields": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"rocket_loader": schema.BoolAttribute{
										Optional: true,
									},
									"redirects_for_ai_training": schema.BoolAttribute{
										Optional: true,
									},
									"rules": schema.MapAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"ruleset": schema.StringAttribute{
										Optional: true,
									},
									"rulesets": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"security_level": schema.StringAttribute{
										Optional: true,
									},
									"server_side_excludes": schema.BoolAttribute{
										Optional: true,
									},
									"ssl": schema.StringAttribute{
										Optional: true,
									},
									"status_code": schema.Int64Attribute{
										Optional: true,
									},
									"sxg": schema.BoolAttribute{
										Optional: true,
									},
									// algorithms is a ListNestedBlock (no MaxItems restriction) in v4
									"algorithms": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
									// uri is a ListNestedBlock MaxItems:1 in v4
									"uri": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"origin": schema.BoolAttribute{
													Optional: true,
												},
												// path is a ListNestedBlock MaxItems:1 inside uri
												"path": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional: true,
															},
															"expression": schema.StringAttribute{
																Optional: true,
															},
														},
													},
												},
												// query is a ListNestedBlock MaxItems:1 inside uri
												"query": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional: true,
															},
															"expression": schema.StringAttribute{
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									// headers is a ListNestedBlock (no MaxItems restriction) in v4
									"headers": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional: true,
												},
												"value": schema.StringAttribute{
													Optional: true,
												},
												"expression": schema.StringAttribute{
													Optional: true,
												},
												"operation": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									// matched_data is a ListNestedBlock MaxItems:1 in v4
									"matched_data": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"public_key": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									// response is a ListNestedBlock MaxItems:1 in v4
									"response": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"status_code": schema.Int64Attribute{
													Optional: true,
												},
												"content_type": schema.StringAttribute{
													Optional: true,
												},
												"content": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									// autominify is a ListNestedBlock MaxItems:1 in v4
									"autominify": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"html": schema.BoolAttribute{
													Optional: true,
												},
												"css": schema.BoolAttribute{
													Optional: true,
												},
												"js": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									// edge_ttl is a ListNestedBlock MaxItems:1 in v4
									"edge_ttl": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Required: true,
												},
												"default": schema.Int64Attribute{
													Optional: true,
												},
												// status_code_ttl is a ListNestedBlock (no MaxItems) inside edge_ttl
												"status_code_ttl": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"status_code": schema.Int64Attribute{
																Optional: true,
															},
															"value": schema.Int64Attribute{
																Optional: true,
															},
															// status_code_range is a ListNestedBlock MaxItems:1 inside status_code_ttl
															"status_code_range": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.Int64Attribute{
																			Optional: true,
																		},
																		"to": schema.Int64Attribute{
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
									// browser_ttl is a ListNestedBlock MaxItems:1 in v4
									"browser_ttl": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Required: true,
												},
												"default": schema.Int64Attribute{
													Optional: true,
												},
											},
										},
									},
									// serve_stale is a ListNestedBlock MaxItems:1 in v4
									"serve_stale": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"disable_stale_while_updating": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
									// cache_key is a ListNestedBlock MaxItems:1 in v4
									"cache_key": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cache_by_device_type": schema.BoolAttribute{
													Optional: true,
												},
												"ignore_query_strings_order": schema.BoolAttribute{
													Optional: true,
												},
												"cache_deception_armor": schema.BoolAttribute{
													Optional: true,
												},
												// custom_key is a ListNestedBlock MaxItems:1 inside cache_key
												"custom_key": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															// query_string is a ListNestedBlock MaxItems:1 inside custom_key
															"query_string": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"exclude": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																	},
																},
															},
															// header is a ListNestedBlock MaxItems:1 inside custom_key
															"header": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"check_presence": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"exclude_origin": schema.BoolAttribute{
																			Optional: true,
																			Computed: true,
																		},
																		// contains maps header names to sets of values
																		"contains": schema.MapAttribute{
																			ElementType: types.SetType{
																				ElemType: types.StringType,
																			},
																			Optional: true,
																		},
																	},
																},
															},
															// cookie is a ListNestedBlock MaxItems:1 inside custom_key
															"cookie": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"include": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"check_presence": schema.SetAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																	},
																},
															},
															// user is a ListNestedBlock MaxItems:1 inside custom_key
															"user": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"device_type": schema.BoolAttribute{
																			Optional: true,
																		},
																		"geo": schema.BoolAttribute{
																			Optional: true,
																		},
																		"lang": schema.BoolAttribute{
																			Optional: true,
																		},
																	},
																},
															},
															// host is a ListNestedBlock MaxItems:1 inside custom_key
															"host": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"resolved": schema.BoolAttribute{
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
									// cache_reserve is a ListNestedBlock MaxItems:1 in v4
									"cache_reserve": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"eligible": schema.BoolAttribute{
													Required: true,
												},
												"minimum_file_size": schema.Int64Attribute{
													Optional: true,
												},
											},
										},
									},
									// from_list is a ListNestedBlock MaxItems:1 in v4
									"from_list": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Optional: true,
												},
												"key": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									// from_value is a ListNestedBlock MaxItems:1 in v4
									"from_value": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"status_code": schema.Int64Attribute{
													Optional: true,
												},
												"preserve_query_string": schema.BoolAttribute{
													Optional: true,
												},
												// target_url is a ListNestedBlock MaxItems:1 inside from_value
												"target_url": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"value": schema.StringAttribute{
																Optional: true,
															},
															"expression": schema.StringAttribute{
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									// overrides is a ListNestedBlock MaxItems:1 in v4
									"overrides": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Optional: true,
												},
												"action": schema.StringAttribute{
													Optional: true,
												},
												"sensitivity_level": schema.StringAttribute{
													Optional: true,
												},
												// categories is a ListNestedBlock (no MaxItems) inside overrides
												"categories": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"category": schema.StringAttribute{
																Optional: true,
															},
															"action": schema.StringAttribute{
																Optional: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
															},
														},
													},
												},
												// rules is a ListNestedBlock (no MaxItems) inside overrides
												"rules": schema.ListNestedAttribute{
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Optional: true,
															},
															"action": schema.StringAttribute{
																Optional: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
															},
															"score_threshold": schema.Int64Attribute{
																Optional: true,
															},
															"sensitivity_level": schema.StringAttribute{
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									// origin is a ListNestedBlock MaxItems:1 in v4
									"origin": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Optional: true,
												},
												"port": schema.Int64Attribute{
													Optional: true,
												},
											},
										},
									},
									// sni is a ListNestedBlock MaxItems:1 in v4
									"sni": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						// ratelimit is a ListNestedBlock MaxItems:1 in v4
						"ratelimit": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"characteristics": schema.SetAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"period": schema.Int64Attribute{
										Optional: true,
									},
									"requests_per_period": schema.Int64Attribute{
										Optional: true,
									},
									"score_per_period": schema.Int64Attribute{
										Optional: true,
									},
									"score_response_header_name": schema.StringAttribute{
										Optional: true,
									},
									"mitigation_timeout": schema.Int64Attribute{
										Optional: true,
									},
									"counting_expression": schema.StringAttribute{
										Optional: true,
									},
									"requests_to_origin": schema.BoolAttribute{
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						// exposed_credential_check is a ListNestedBlock MaxItems:1 in v4
						"exposed_credential_check": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"username_expression": schema.StringAttribute{
										Optional: true,
									},
									"password_expression": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
						// logging is a ListNestedBlock MaxItems:1 in v4
						"logging": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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
