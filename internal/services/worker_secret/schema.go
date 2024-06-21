// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r WorkerSecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of this secret, this is what will be to access it inside the Worker.",
				Optional:    true,
			},
			"text": schema.StringAttribute{
				Description: "The value of the secret.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of secret to put.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("secret_text"),
				},
			},
		},
	}
}
