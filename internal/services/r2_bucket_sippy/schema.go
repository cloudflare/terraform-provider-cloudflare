// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketSippyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Account ID",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bucket_name": schema.StringAttribute{
				Description:   "Name of the bucket",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"destination": schema.SingleNestedAttribute{
				Description: "R2 bucket to copy objects to",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketSippyDestinationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Description: "ID of a Cloudflare API token.\nThis is the value labelled \"Access Key ID\" when creating an API\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n",
						Optional:    true,
					},
					"provider": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("r2"),
						},
					},
					"secret_access_key": schema.StringAttribute{
						Description: "Value of a Cloudflare API token.\nThis is the value labelled \"Secret Access Key\" when creating an API\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"source": schema.SingleNestedAttribute{
				Description: "AWS S3 bucket to copy objects from",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketSippySourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Description: "Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)",
						Optional:    true,
					},
					"bucket": schema.StringAttribute{
						Description: "Name of the AWS S3 bucket",
						Optional:    true,
					},
					"provider": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("aws", "gcs"),
						},
					},
					"region": schema.StringAttribute{
						Description: "Name of the AWS availability zone",
						Optional:    true,
					},
					"secret_access_key": schema.StringAttribute{
						Description: "Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)",
						Optional:    true,
					},
					"client_email": schema.StringAttribute{
						Description: "Client email of an IAM credential (ideally scoped to a single GCS bucket)",
						Optional:    true,
					},
					"private_key": schema.StringAttribute{
						Description: "Private Key of an IAM credential (ideally scoped to a single GCS bucket)",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "State of Sippy for this bucket",
				Computed:    true,
			},
		},
	}
}

func (r *R2BucketSippyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketSippyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
