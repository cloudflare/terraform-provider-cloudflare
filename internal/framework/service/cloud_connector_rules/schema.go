package cloud_connector_rules

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *CloudConnectorRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Cloud Connector Rules](https://developers.cloudflare.com/rules/cloud-connector/) resource allows you to create and manage cloud connector rules for a zone.
		`),
		Version: 1,

		Attributes: map[string]schema.Attribute{
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.ZoneIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				MarkdownDescription: "List of Cloud Connector Rules",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.IsRequired(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"provider": schema.StringAttribute{
							Required: true,
							MarkdownDescription: fmt.Sprintf("Type of provider. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{
								"aws_s3",
								"cloudflare_r2",
								"azure_storage",
								"gcp_storage",
							})),
							Validators: []validator.String{
								stringvalidator.OneOf(
									"aws_s3",
									"cloudflare_r2",
									"azure_storage",
									"gcp_storage",
								),
							},
						},
						"enabled": schema.BoolAttribute{
							Optional:            true,
							MarkdownDescription: "Whether the headers rule is active.",
						},
						"expression": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Criteria for an HTTP request to trigger the cloud connector rule. Uses the Firewall Rules expression language based on Wireshark display filters.",
						},
						"description": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Brief summary of the cloud connector rule and its intended use.",
						},
					},
					Blocks: map[string]schema.Block{
						"parameters": schema.SingleNestedBlock{
							MarkdownDescription: "Cloud Connector Rule Parameters",
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "Host parameter for cloud connector rule",
								},
							},
						},
					},
				},
			},
		},
	}
}
