package apishieldoperation

import (
	"context"
	"net/http"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *APIShieldOperationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Api shield operation",
		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Description: consts.IDSchemaDescription,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.ZoneIDSchemaKey: schema.StringAttribute{
				Description: consts.ZoneIDSchemaDescription,
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"method": schema.StringAttribute{
				Description: "The HTTP method used to access the endpoint",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						http.MethodGet,
						http.MethodHead,
						http.MethodPost,
						http.MethodPut,
						http.MethodPatch,
						http.MethodDelete,
						http.MethodConnect,
						http.MethodOptions,
						http.MethodTrace,
					),
				},
			},
			"host": schema.StringAttribute{
				Description: "RFC3986-compliant host",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"endpoint": schema.StringAttribute{
				Description: "The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with `{varN}`, starting with `{var1}`. This will then be [Cloudflare-normalized](https://developers.cloudflare.com/rules/normalization/how-it-works/)",
				Required:    true,
				CustomType:  EndpointType{},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
