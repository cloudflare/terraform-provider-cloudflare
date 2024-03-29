// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r PagerulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"pagerule_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"actions": schema.ListNestedAttribute{
				Description: "The set of actions to perform if the targets of this rule match the request. Actions can redirect to another URL or override settings, but not both.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"modified_on": schema.StringAttribute{
							Description: "The timestamp of when the override was last modified.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The type of route.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("forward_url"),
							},
						},
						"value": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: "The response type for the URL redirect.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("temporary", "permanent"),
									},
								},
								"url": schema.StringAttribute{
									Description: "The URL to redirect the request to.\nNotes: ${num} refers to the position of '*' in the constraint value.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"targets": schema.ListNestedAttribute{
				Description: "The rule targets to evaluate on each request.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"constraint": schema.SingleNestedAttribute{
							Description: "String constraint.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"operator": schema.StringAttribute{
									Description: "The matches operator can use asterisks and pipes as wildcard and 'or' operators.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("matches", "contains", "equals", "not_equal", "not_contain"),
									},
								},
								"value": schema.StringAttribute{
									Description: "The URL pattern to match against the current request. The pattern may contain up to four asterisks ('*') as placeholders.",
									Required:    true,
								},
							},
						},
						"target": schema.StringAttribute{
							Description: "A target based on the URL of the request.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("url"),
							},
						},
					},
				},
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the rule, used to define which Page Rule is processed over another. A higher number indicates a higher priority. For example, if you have a catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to take precedence (rule B: `/images/special/*`), specify a higher priority for rule B so it overrides rule A.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the Page Rule.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "disabled"),
				},
			},
		},
	}
}
