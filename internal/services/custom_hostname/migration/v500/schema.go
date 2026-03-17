package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UnionV0Schema(buildSchema func(context.Context) schema.Schema, ctx context.Context) *schema.Schema {
	s := buildSchema(ctx)
	s.Version = 0
	delete(s.Attributes, "ssl")
	s.Attributes["wait_for_ssl_pending_validation"] = schema.BoolAttribute{Optional: true}
	return &s
}

// SourceCustomHostnameSchema returns the minimal v4 schema for reading legacy state.
func SourceCustomHostnameSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id":       schema.StringAttribute{Computed: true},
			"zone_id":  schema.StringAttribute{Required: true},
			"hostname": schema.StringAttribute{Required: true},
			"ssl": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"status":                schema.StringAttribute{Computed: true},
						"bundle_method":         schema.StringAttribute{Optional: true},
						"method":                schema.StringAttribute{Optional: true},
						"type":                  schema.StringAttribute{Optional: true},
						"certificate_authority": schema.StringAttribute{Optional: true, Computed: true},
						"validation_records": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cname_target": schema.StringAttribute{Optional: true, Computed: true},
									"cname_name":   schema.StringAttribute{Optional: true, Computed: true},
									"txt_name":     schema.StringAttribute{Optional: true, Computed: true},
									"txt_value":    schema.StringAttribute{Optional: true, Computed: true},
									"http_url":     schema.StringAttribute{Optional: true, Computed: true},
									"http_body":    schema.StringAttribute{Optional: true, Computed: true},
									"emails":       schema.ListAttribute{ElementType: types.StringType, Optional: true, Computed: true},
								},
							},
						},
						"validation_errors": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"message": schema.StringAttribute{Computed: true},
								},
							},
						},
						"wildcard":           schema.BoolAttribute{Optional: true},
						"custom_certificate": schema.StringAttribute{Optional: true},
						"custom_key":         schema.StringAttribute{Optional: true},
						"settings": schema.ListNestedAttribute{
							Optional: true,
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"http2":           schema.StringAttribute{Optional: true},
									"tls13":           schema.StringAttribute{Optional: true},
									"min_tls_version": schema.StringAttribute{Optional: true},
									"ciphers":         schema.SetAttribute{ElementType: types.StringType, Optional: true},
									"early_hints":     schema.StringAttribute{Optional: true},
								},
							},
						},
					},
				},
			},
			"custom_origin_server":            schema.StringAttribute{Optional: true},
			"custom_origin_sni":               schema.StringAttribute{Optional: true},
			"custom_metadata":                 schema.MapAttribute{ElementType: types.StringType, Optional: true},
			"status":                          schema.StringAttribute{Computed: true},
			"ownership_verification":          schema.MapAttribute{ElementType: types.StringType, Computed: true},
			"ownership_verification_http":     schema.MapAttribute{ElementType: types.StringType, Computed: true},
			"wait_for_ssl_pending_validation": schema.BoolAttribute{Optional: true},
		},
	}
}
