// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*DNSRecordResource)(nil)
var _ resource.ResourceWithModifyPlan = (*DNSRecordResource)(nil)
var _ resource.ResourceWithImportState = (*DNSRecordResource)(nil)

func NewResource() resource.Resource {
	return &DNSRecordResource{}
}

// DNSRecordResource defines the resource implementation.
type DNSRecordResource struct {
	client *cloudflare.Client
}

func (r *DNSRecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_record"
}

func (r *DNSRecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DNSRecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DNSRecordModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Proxied.ValueBool() && data.TTL.ValueFloat64() != 1 {
		resp.Diagnostics.AddError(
			"ttl must be set to 1 when `proxied` is true",
			"When a DNS record is marked as `proxied` the TTL must be 1 as Cloudflare will control the TTL internally.",
		)
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := DNSRecordResultEnvelope{*data}
	_, err = r.client.DNS.Records.New(
		ctx,
		dns.RecordNewParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSRecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DNSRecordModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *DNSRecordModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := DNSRecordResultEnvelope{*data}
	_, err = r.client.DNS.Records.Update(
		ctx,
		data.ID.ValueString(),
		dns.RecordUpdateParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSRecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DNSRecordModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := DNSRecordResultEnvelope{*data}
	_, err := r.client.DNS.Records.Get(
		ctx,
		data.ID.ValueString(),
		dns.RecordGetParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSRecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DNSRecordModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DNS.Records.Delete(
		ctx,
		data.ID.ValueString(),
		dns.RecordDeleteParams{
			ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSRecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *DNSRecordModel = new(DNSRecordModel)

	path_zone_id := ""
	path_dns_record_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>/<dns_record_id>",
		&path_zone_id,
		&path_dns_record_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path_zone_id)
	data.ID = types.StringValue(path_dns_record_id)

	res := new(http.Response)
	env := DNSRecordResultEnvelope{*data}
	_, err := r.client.DNS.Records.Get(
		ctx,
		path_dns_record_id,
		dns.RecordGetParams{
			ZoneID: cloudflare.F(path_zone_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSRecordResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Only proceed if we have a plan (not destroying)
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan DNSRecordModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the current state if it exists
	var state *DNSRecordModel
	if !req.State.Raw.IsNull() {
		state = &DNSRecordModel{}
		resp.Diagnostics.Append(req.State.Get(ctx, state)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Handle name normalization for FQDN vs subdomain
	// The API returns FQDN but users often configure subdomain
	if state != nil && !plan.Name.IsUnknown() && !state.Name.IsNull() {
		planName := plan.Name.ValueString()
		stateName := state.Name.ValueString()

		// Remove trailing dots for comparison
		planNameNorm := strings.TrimSuffix(planName, ".")
		stateNameNorm := strings.TrimSuffix(stateName, ".")

		// Check if plan is "@" (apex) and state is the zone name
		if planName == "@" && stateName != "@" && strings.Contains(stateName, ".") {
			plan.Name = state.Name
		} else if strings.HasPrefix(stateNameNorm, planNameNorm+".") {
			// State is FQDN, plan is subdomain - keep state to prevent drift
			plan.Name = state.Name
		} else if planNameNorm != stateNameNorm {
			// Check if they're semantically the same record
			planParts := strings.Split(planNameNorm, ".")
			stateParts := strings.Split(stateNameNorm, ".")
			if len(planParts) < len(stateParts) {
				matches := true
				for i, part := range planParts {
					if i >= len(stateParts) || part != stateParts[i] {
						matches = false
						break
					}
				}
				if matches {
					plan.Name = state.Name
				}
			}
		}
	}

	// Preserve computed fields from state during updates
	if state != nil {
		// Check if there are actual configuration changes
		hasConfigChanges := false

		// Check for changes in user-configurable fields
		// For CAA and other records that use data field, content might be computed
		// so we need to be careful about comparing it
		contentChanged := false
		if plan.Data == nil {
			// Regular record using content field
			// Special handling for CNAME records: DNS is case-insensitive
			if !plan.Type.IsNull() && plan.Type.ValueString() == "CNAME" &&
				!plan.Content.IsNull() && !state.Content.IsNull() {
				// Do case-insensitive comparison for CNAME content
				planContent := strings.ToLower(plan.Content.ValueString())
				stateContent := strings.ToLower(state.Content.ValueString())
				contentChanged = planContent != stateContent
				// If only case differs, preserve the state value to prevent drift
				if !contentChanged && plan.Content.ValueString() != state.Content.ValueString() {
					plan.Content = state.Content
				}
			} else {
				contentChanged = !plan.Content.Equal(state.Content)
			}
		} else {
			// Record using data field (like CAA), content is computed
			// Don't consider content changes for these records
			contentChanged = false
		}

		// Special handling for tags: treat empty list and null as equivalent
		// Also, when tags is unknown (marked by Terraform as "known after apply"),
		// don't consider it as a change if state is empty
		tagsChanged := false

		if plan.Tags.IsUnknown() {
			// When plan tags is unknown and state is empty, no real change
			stateTagsEmpty := state.Tags.IsNull() || (!state.Tags.IsUnknown() && len(state.Tags.Elements()) == 0)
			tagsChanged = !stateTagsEmpty
		} else {
			// Normal comparison when plan tags is known
			planTagsEmpty := plan.Tags.IsNull() || len(plan.Tags.Elements()) == 0
			stateTagsEmpty := state.Tags.IsNull() || (!state.Tags.IsUnknown() && len(state.Tags.Elements()) == 0)

			if planTagsEmpty && stateTagsEmpty {
				// Both are empty, no change
				tagsChanged = false
			} else {
				// At least one is not empty, use regular comparison
				tagsChanged = !plan.Tags.Equal(state.Tags)
			}
		}

		if !plan.Name.Equal(state.Name) || !plan.Type.Equal(state.Type) ||
			contentChanged || !plan.TTL.Equal(state.TTL) ||
			!plan.Proxied.Equal(state.Proxied) || !plan.Priority.Equal(state.Priority) ||
			!plan.Comment.Equal(state.Comment) || tagsChanged {
			hasConfigChanges = true
		}

		// Check if data field has changes (for CAA records)
		// This is complex because we'd need to compare all subfields
		// For now, assume data hasn't changed if both plan and state have it
		if (plan.Data == nil) != (state.Data == nil) {
			// One is nil and the other is not
			hasConfigChanges = true
		}

		// Always preserve created_on since it never changes
		if plan.CreatedOn.IsUnknown() {
			plan.CreatedOn = state.CreatedOn
		}

		// Only preserve modified_on if there are no config changes
		if plan.ModifiedOn.IsUnknown() && !hasConfigChanges {
			plan.ModifiedOn = state.ModifiedOn
		}

		// Preserve proxiable flag
		if plan.Proxiable.IsUnknown() {
			plan.Proxiable = state.Proxiable
		}

		// Preserve meta field
		if plan.Meta.IsUnknown() {
			plan.Meta = state.Meta
		}

		// For CAA records and others that use data field, preserve computed content
		if plan.Content.IsUnknown() && plan.Data != nil {
			plan.Content = state.Content
		}

		// Handle settings: preserve from state if not explicitly set
		if plan.Settings.IsUnknown() || plan.Settings.IsNull() {
			plan.Settings = state.Settings
		} else if !plan.Settings.IsNull() {
			// If settings is set in plan, ensure nested fields have defaults
			var settingsData DNSRecordSettingsModel
			resp.Diagnostics.Append(plan.Settings.As(ctx, &settingsData, basetypes.ObjectAsOptions{})...)
			if !resp.Diagnostics.HasError() {
				updated := false
				// Set defaults for unset boolean fields to false
				if settingsData.IPV4Only.IsUnknown() {
					settingsData.IPV4Only = types.BoolValue(false)
					updated = true
				}
				if settingsData.IPV6Only.IsUnknown() {
					settingsData.IPV6Only = types.BoolValue(false)
					updated = true
				}
				if settingsData.FlattenCNAME.IsUnknown() {
					settingsData.FlattenCNAME = types.BoolValue(false)
					updated = true
				}
				if updated {
					plan.Settings = customfield.NewObjectMust(ctx, &settingsData)
				}
			}
		}

		// Handle tags: preserve empty set from state to avoid showing as unknown
		if plan.Tags.IsUnknown() {
			if state.Tags.IsNull() || len(state.Tags.Elements()) == 0 {
				plan.Tags = state.Tags
			}
		}
	}

	// Handle comment_modified_on drift: similar to tags_modified_on
	commentIsEmpty := plan.Comment.IsNull() || (!plan.Comment.IsUnknown() && plan.Comment.ValueString() == "")

	if commentIsEmpty && plan.CommentModifiedOn.IsUnknown() {
		// Set comment_modified_on to null when comment is empty, preventing drift
		plan.CommentModifiedOn = timetypes.NewRFC3339Null()
	} else if !commentIsEmpty && plan.CommentModifiedOn.IsUnknown() && state != nil {
		// If comment hasn't changed, preserve comment_modified_on from state
		if plan.Comment.Equal(state.Comment) {
			plan.CommentModifiedOn = state.CommentModifiedOn
		}
		// Otherwise let it be unknown (will be updated by the API)
	}

	// Handle tags_modified_on drift: if tags is empty/null, ensure tags_modified_on is null
	// This works around terraform-plugin-framework issue #898 where computed fields adjacent
	// to optional+computed collections show as "known after apply"
	tagsIsEmpty := plan.Tags.IsNull() || (!plan.Tags.IsUnknown() && len(plan.Tags.Elements()) == 0)

	if tagsIsEmpty && plan.TagsModifiedOn.IsUnknown() {
		// Set tags_modified_on to null when tags are empty, preventing drift
		plan.TagsModifiedOn = timetypes.NewRFC3339Null()
	} else if !tagsIsEmpty && plan.TagsModifiedOn.IsUnknown() && state != nil {
		// If tags haven't changed, preserve tags_modified_on from state
		if plan.Tags.Equal(state.Tags) {
			plan.TagsModifiedOn = state.TagsModifiedOn
		}
		// Otherwise let it be unknown (will be updated by the API)
	}

	// Set the updated plan
	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
}
