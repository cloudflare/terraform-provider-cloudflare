package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UnionV0Schema builds a PriorSchema for version 0 that can parse BOTH V4 and V5 state.
//
// The problem: V4 and V5 both use schema_version=0, but have incompatible state formats.
// V4 stores `placement` as an array, V5 stores it as an object. The solution: define
// `placement` as ListNestedAttribute which handles V4 arrays and V5 null. The handler
// converts list→object for V5 state before delegating to the run_worker_first upgrade.
//
// Detection: V4 state has `name` field, V5 state has `script_name` field.
// The handler checks which is populated to determine the upgrade path.
//
// buildSchema is the ResourceSchema function from the parent package, passed to avoid circular imports.
func UnionV0Schema(buildSchema func(context.Context) schema.Schema, ctx context.Context) *schema.Schema {
	// Start from V5 schema (has all V5 fields)
	s := buildSchema(ctx)
	s.Version = 0

	// Change run_worker_first from Dynamic to Bool (V5 version 0 had it as bool)
	s.Attributes["assets"].(schema.SingleNestedAttribute).
		Attributes["config"].(schema.SingleNestedAttribute).
		Attributes["run_worker_first"] = schema.BoolAttribute{Optional: true}

	// Change placement from SingleNestedAttribute to ListNestedAttribute.
	// V4 (SDKv2) stores placement as array: [] or [{"mode":"smart"}]
	// V5 stores it as object: null or {"mode":"smart",...}
	// ListNestedAttribute handles V4 arrays and V5 null.
	// V5 non-null object placement at version 0 is extremely rare (requires
	// configuring Smart Placement before the V0→V1 run_worker_first update
	// and never running apply since).
	s.Attributes["placement"] = schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"mode": schema.StringAttribute{Optional: true},
			},
		},
	}

	// Add V4-specific fields (all Optional — null for V5 state)
	s.Attributes["name"] = schema.StringAttribute{Optional: true}
	s.Attributes["module"] = schema.BoolAttribute{Optional: true}
	s.Attributes["tags"] = schema.SetAttribute{Optional: true, ElementType: types.StringType}
	s.Attributes["dispatch_namespace"] = schema.StringAttribute{Optional: true}

	// Add V4 binding arrays (null for V5 state)
	s.Attributes["plain_text_binding"] = v4BindingSchema("name", "text")
	s.Attributes["secret_text_binding"] = v4BindingSchema("name", "text")
	s.Attributes["kv_namespace_binding"] = v4BindingSchema("name", "namespace_id")
	s.Attributes["webassembly_binding"] = v4BindingSchema("name", "module")
	s.Attributes["service_binding"] = v4BindingSchema("name", "service", "environment")
	s.Attributes["r2_bucket_binding"] = v4BindingSchema("name", "bucket_name")
	s.Attributes["analytics_engine_binding"] = v4BindingSchema("name", "dataset")
	s.Attributes["queue_binding"] = v4BindingSchema("binding", "queue")
	s.Attributes["d1_database_binding"] = v4BindingSchema("name", "database_id")
	s.Attributes["hyperdrive_config_binding"] = v4BindingSchema("binding", "id")

	return &s
}

// v4BindingSchema creates a ListNestedAttribute schema for a V4 binding type.
func v4BindingSchema(fields ...string) schema.ListNestedAttribute {
	attrs := make(map[string]schema.Attribute, len(fields))
	for _, f := range fields {
		attrs[f] = schema.StringAttribute{Optional: true}
	}
	return schema.ListNestedAttribute{
		Optional:     true,
		NestedObject: schema.NestedAttributeObject{Attributes: attrs},
	}
}

// SourceWorkerScriptSchema returns the v4 cloudflare_worker_script schema.
// Schema version: 1 (v4 SDKv2 provider)
// Resource type: cloudflare_worker_script (singular)
//
// This minimal schema is used only for reading v4 state during MoveState migration.
// Validators, PlanModifiers, Defaults, and Descriptions are intentionally omitted.
//
// Note: Uses ListNestedAttribute (not ListNestedBlock) for binding blocks because SDKv2
// stores block data in the attributes section of the state JSON.
func SourceWorkerScriptSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"content": schema.StringAttribute{
				Optional: true,
			},
			"module": schema.BoolAttribute{
				Optional: true,
			},
			"tags": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"dispatch_namespace": schema.StringAttribute{
				Optional: true,
			},

			// V4 binding blocks — stored as attribute arrays in SDKv2 state
			"plain_text_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{Optional: true},
						"text": schema.StringAttribute{Optional: true},
					},
				},
			},
			"secret_text_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{Optional: true},
						"text": schema.StringAttribute{Optional: true},
					},
				},
			},
			"kv_namespace_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":         schema.StringAttribute{Optional: true},
						"namespace_id": schema.StringAttribute{Optional: true},
					},
				},
			},
			"webassembly_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":   schema.StringAttribute{Optional: true},
						"module": schema.StringAttribute{Optional: true}, // renamed to "part" in v5
					},
				},
			},
			"service_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":        schema.StringAttribute{Optional: true},
						"service":     schema.StringAttribute{Optional: true},
						"environment": schema.StringAttribute{Optional: true},
					},
				},
			},
			"r2_bucket_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":        schema.StringAttribute{Optional: true},
						"bucket_name": schema.StringAttribute{Optional: true},
					},
				},
			},
			"analytics_engine_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":    schema.StringAttribute{Optional: true},
						"dataset": schema.StringAttribute{Optional: true},
					},
				},
			},
			"queue_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"binding": schema.StringAttribute{Optional: true}, // renamed to "name" in v5
						"queue":   schema.StringAttribute{Optional: true}, // renamed to "queue_name" in v5
					},
				},
			},
			"d1_database_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":        schema.StringAttribute{Optional: true},
						"database_id": schema.StringAttribute{Optional: true}, // renamed to "id" in v5
					},
				},
			},
			"hyperdrive_config_binding": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"binding": schema.StringAttribute{Optional: true}, // renamed to "name" in v5
						"id":      schema.StringAttribute{Optional: true},
					},
				},
			},

			// V4 placement — stored as single-element array in SDKv2 state
			"placement": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mode": schema.StringAttribute{Optional: true},
					},
				},
			},
		},
	}
}
