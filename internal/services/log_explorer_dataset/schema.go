package log_explorer_dataset

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*LogExplorerDatasetResource)(nil)

type uniqueFieldNamesValidator struct{}

func (uniqueFieldNamesValidator) Description(context.Context) string {
	return "Field names must be unique without regard to case."
}

func (v uniqueFieldNamesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (uniqueFieldNamesValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	list := customfield.NestedObjectList[LogExplorerDatasetFieldsModel]{ListValue: req.ConfigValue}
	fields, diags := list.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	seen := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		if field.Name.IsNull() || field.Name.IsUnknown() {
			continue
		}
		name := strings.ToLower(field.Name.ValueString())
		if _, exists := seen[name]; exists {
			resp.Diagnostics.AddAttributeError(req.Path, "Duplicate field name", "Field names must be unique without regard to case.")
			return
		}
		seen[name] = struct{}{}
	}
}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Logs Read",
				"Logs Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique dataset ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"dataset_id": schema.StringAttribute{
				Description:   "Unique dataset ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"dataset": schema.StringAttribute{
				Description:   "Dataset type name to create (e.g. `http_requests`).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"deletion_protection": schema.BoolAttribute{
				Description:   "Set to `false` to allow deletion of this dataset.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{
				Description:   "Whether to enable or disable log ingest for this dataset.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"filter": schema.StringAttribute{
				Description:   "Optional Logpush filter predicate to restrict which events are ingested.\nOmitting this attribute means no filter. An empty string also means no filter.\nRemoving a configured filter sends an empty string to clear it in the API.\nThe provider stores missing or null API filters as empty strings in state.\nSee [Logpush filters](https://developers.cloudflare.com/logs/reference/filters/)\nfor syntax and examples.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"filter_attack_traffic": schema.BoolAttribute{
				Description:   "Whether to filter attack traffic from the Logpush job. The API defaults to true on creation for http_requests, firewall_events, and network_analytics_logs. If omitted on update, the existing setting is kept. The API omits this setting for datasets that do not support it.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"fields": schema.ListNestedAttribute{
				Description:   "Controls which fields the API ingests. Defaults to all available\nfields when absent.\nEach dataset requires its own timestamp field to remain enabled (the\nfield name varies per dataset type); omitting it from an explicit\nselection is rejected by the API.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewNestedObjectListType[LogExplorerDatasetFieldsModel](ctx),
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				Validators:    []validator.List{uniqueFieldNamesValidator{}},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Description: "Whether the API includes this field in log ingest.",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "Field name. Names are matched without regard to case.",
							Required:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Description:   "RFC3339 timestamp recording when the API created this dataset.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"object_id": schema.StringAttribute{
				Description:   "Public ID of the account or zone that owns this dataset.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"object_type": schema.StringAttribute{
				Description: "Whether this dataset belongs to an account or a zone.\nAvailable values: \"account\", \"zone\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("account", "zone"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated_at": schema.StringAttribute{
				Description: "RFC3339 timestamp recording when the API last updated this dataset.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *LogExplorerDatasetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *LogExplorerDatasetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
