// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_devices_dex_tests

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r ZeroTrustDevicesDEXTestsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"dex_test_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Optional:    true,
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Determines whether or not the test is active.",
				Required:    true,
			},
			"interval": schema.StringAttribute{
				Description: "How often the test will run.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the DEX test. Must be unique.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Optional:    true,
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "Device settings profiles targeted by this test",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the device settings profile",
							Optional:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the profile is the account default",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile",
							Optional:    true,
						},
					},
				},
			},
			"targeted": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}
