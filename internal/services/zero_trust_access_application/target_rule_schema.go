package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func targetRuleResourceSchema(description, name string) schema.SingleNestedAttribute {
	validators := []validator.Object{
		targetRuleSelectorsValidator{},
	}
	if name == "include" {
		validators = append(validators, objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("target_attributes")))
	} else {
		validators = append(validators,
			objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("target_attributes")),
			objectvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("include")),
		)
	}
	return schema.SingleNestedAttribute{
		Description: description,
		Optional:    true,
		Validators:  validators,
		Attributes: map[string]schema.Attribute{
			"target_attributes": schema.SingleNestedAttribute{
				Description: "Hostname selector map. Only the hostname key is supported.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"hostname": schema.ListAttribute{
						Required:    true,
						ElementType: types.StringType,
						Validators:  []validator.List{listvalidator.SizeBetween(1, 5000), listvalidator.ValueStringsAre(stringvalidator.LengthAtLeast(1))},
					},
				},
			},
			"tags": schema.MapAttribute{
				Description: "Map of target tag keys to values. Values within a key are OR'd.",
				Optional:    true,
				ElementType: types.ListType{ElemType: types.StringType},
				Validators: []validator.Map{
					mapvalidator.SizeBetween(1, 50),
					mapvalidator.KeysAre(stringvalidator.LengthAtLeast(1)),
					mapvalidator.ValueListsAre(listvalidator.SizeBetween(1, 50), listvalidator.ValueStringsAre(stringvalidator.LengthAtLeast(1))),
				},
			},
		},
	}
}

func targetRuleDataSourceSchema(ctx context.Context, description string) dsschema.SingleNestedAttribute {
	return dsschema.SingleNestedAttribute{
		Description: description,
		Computed:    true,
		CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationTargetRuleDataSourceModel](ctx),
		Attributes: map[string]dsschema.Attribute{
			"target_attributes": dsschema.SingleNestedAttribute{
				Description: "Hostname selector map. Only the hostname key is supported.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessApplicationTargetRuleAttributesDataSourceModel](ctx),
				Attributes: map[string]dsschema.Attribute{
					"hostname": dsschema.ListAttribute{Computed: true, ElementType: types.StringType, CustomType: customfield.NewListType[types.String](ctx)},
				},
			},
			"tags": dsschema.MapAttribute{
				Description: "Map of target tag keys to values. Values within a key are OR'd.",
				Computed:    true,
				CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
				ElementType: types.ListType{ElemType: types.StringType},
			},
		},
	}
}
