package dns_record

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ---------------------------------------------------------------------------
// Private-state helpers – persist zone_name across plan cycles so ModifyPlan
// can distinguish FQDN drift from genuine name changes without an extra API
// call.  The zone_name field is part of every DNS record API response.
// ---------------------------------------------------------------------------

// privateStateKeyZoneName is the key used to store the zone's domain name in
// Terraform private state.
const privateStateKeyZoneName = "zone_name"

// privateStateReadWriter combines the Get/Set capabilities needed for private
// state operations.
type privateStateReadWriter interface {
	GetKey(ctx context.Context, key string) ([]byte, diag.Diagnostics)
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

// deriveAndSaveZoneName compares the prior state name with the FQDN returned
// by the API to derive the zone suffix, then stores it in private state.
//
// Example: priorName="sub.app", apiFQDN="sub.app.zone.com" → zone="zone.com"
//
// This is called from Read (and ImportState) where we have both the prior
// state name and the fresh API response name.  Create/Update use
// UnmarshalComputed which does not overwrite the name field, so the FQDN
// is only available in Read.
func deriveAndSaveZoneName(ctx context.Context, priorName, apiFQDN string, private privateStateReadWriter) {
	priorNorm := strings.TrimSuffix(priorName, ".")
	fqdnNorm := strings.TrimSuffix(apiFQDN, ".")

	var zoneName string
	if priorNorm != "" && fqdnNorm != priorNorm && strings.HasPrefix(fqdnNorm, priorNorm+".") {
		// Prior state had subdomain form, API returned FQDN → derive zone suffix.
		zoneName = fqdnNorm[len(priorNorm)+1:]
	} else if priorNorm == fqdnNorm && strings.Contains(fqdnNorm, ".") {
		// Prior state was already the FQDN (e.g. after import or previous Read).
		// Try to recover zone name from private state; if present, keep it.
		existing := getZoneNameFromPrivateState(ctx, private)
		if existing != "" {
			return // already stored, nothing to update
		}
		// Can't derive zone name when prior == FQDN and no prior zone stored.
		return
	} else {
		return
	}

	if zoneName != "" {
		jsonVal, _ := json.Marshal(zoneName)
		if diags := private.SetKey(ctx, privateStateKeyZoneName, jsonVal); diags.HasError() {
			tflog.Warn(ctx, "failed to save zone_name to private state", map[string]interface{}{
				"error": diags.Errors()[0].Detail(),
			})
		}
	}
}

// getZoneNameFromPrivateState reads the zone name that was previously stored
// in private state.  Returns "" if unavailable.
func getZoneNameFromPrivateState(ctx context.Context, private privateStateReadWriter) string {
	if private == nil {
		return ""
	}
	raw, diags := private.GetKey(ctx, privateStateKeyZoneName)
	if diags.HasError() || len(raw) == 0 {
		return ""
	}
	var name string
	if err := json.Unmarshal(raw, &name); err != nil {
		return ""
	}
	return name
}

// ---------------------------------------------------------------------------
// ModifyPlan – all plan-modification logic lives here to avoid conflicts with
// code-generated CRUD methods in resource.go.
// ---------------------------------------------------------------------------

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
		} else if planNameNorm != stateNameNorm && strings.HasPrefix(stateNameNorm, planNameNorm+".") {
			// State name starts with plan name — this could be:
			// a) FQDN drift: user configured "sub" and API returned "sub.zone.com" (suppress diff)
			// b) Real name change: user changed "sub.app" to "sub" and state is "sub.app.zone.com" (keep change)
			//
			// To distinguish, read the zone name (previously saved from the DNS record
			// API response) and verify the suffix matches exactly.
			zoneName := strings.TrimSuffix(getZoneNameFromPrivateState(ctx, req.Private), ".")
			suffix := stateNameNorm[len(planNameNorm)+1:]
			if zoneName != "" && suffix == zoneName {
				// The plan name + zone name == state FQDN: this is just FQDN drift, suppress it
				plan.Name = state.Name
			} else if zoneName == "" {
				// No zone name in private state (e.g. first plan after import or
				// upgrade from a version that didn't store it).  Fall back to the
				// prefix heuristic to avoid breaking existing setups.
				plan.Name = state.Name
			}
			// Otherwise it's a real change — the user shortened the name; let Terraform see the diff
		}
	}

	// Preserve computed fields from state during updates
	if state != nil {
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

		// Check if any user-configurable fields have actually changed
		hasChanges := !plan.Name.Equal(state.Name) || !plan.Type.Equal(state.Type) ||
			contentChanged || !plan.TTL.Equal(state.TTL) ||
			!plan.Proxied.Equal(state.Proxied) || !plan.Priority.Equal(state.Priority) ||
			!plan.Comment.Equal(state.Comment) || tagsChanged

		// For Data field (CAA, LOC, etc records), check if it actually changed
		if (plan.Data == nil) != (state.Data == nil) {
			hasChanges = true
		} else if plan.Data != nil && state.Data != nil && !plan.Data.Equal(state.Data) {
			hasChanges = true
		}
		// Note: If both have data, the contentChanged check above already covers it
		// since content is computed from data for these record types

		// Check settings changes - treat empty object as equivalent to null
		planSettingsEmpty := plan.Settings.IsNull() || plan.Settings.IsUnknown()
		stateSettingsEmpty := state.Settings.IsNull() || state.Settings.IsUnknown()

		// Check if plan settings is an empty object {}
		if !plan.Settings.IsNull() && !plan.Settings.IsUnknown() {
			var planSettingsData DNSRecordSettingsModel
			diags := plan.Settings.As(ctx, &planSettingsData, basetypes.ObjectAsOptions{})
			if !diags.HasError() {
				// If all fields are null/unknown, treat as empty
				if (planSettingsData.IPV4Only.IsNull() || planSettingsData.IPV4Only.IsUnknown()) &&
					(planSettingsData.IPV6Only.IsNull() || planSettingsData.IPV6Only.IsUnknown()) &&
					(planSettingsData.FlattenCNAME.IsNull() || planSettingsData.FlattenCNAME.IsUnknown()) {
					planSettingsEmpty = true
				}
			}
		}

		// Only consider it a change if one is empty and the other is not,
		// or if both are non-empty and different
		if !planSettingsEmpty && !stateSettingsEmpty {
			hasChanges = hasChanges || !plan.Settings.Equal(state.Settings)
		} else if planSettingsEmpty != stateSettingsEmpty {
			// One is empty, other is not - but need to check if the non-empty one
			// only has default/null values
			if !stateSettingsEmpty {
				var stateSettingsData DNSRecordSettingsModel
				diags := state.Settings.As(ctx, &stateSettingsData, basetypes.ObjectAsOptions{})
				if !diags.HasError() {
					// Check if state settings only has false/null values (defaults)
					hasActualSettings := false
					if !stateSettingsData.IPV4Only.IsNull() && !stateSettingsData.IPV4Only.IsUnknown() && stateSettingsData.IPV4Only.ValueBool() {
						hasActualSettings = true
					}
					if !stateSettingsData.IPV6Only.IsNull() && !stateSettingsData.IPV6Only.IsUnknown() && stateSettingsData.IPV6Only.ValueBool() {
						hasActualSettings = true
					}
					if !stateSettingsData.FlattenCNAME.IsNull() && !stateSettingsData.FlattenCNAME.IsUnknown() && stateSettingsData.FlattenCNAME.ValueBool() {
						hasActualSettings = true
					}
					// Only consider it a change if state has actual non-default settings
					hasChanges = hasChanges || hasActualSettings
				}
			}
		}

		// Always preserve created_on since it never changes
		if plan.CreatedOn.IsUnknown() {
			plan.CreatedOn = state.CreatedOn
		}

		// Handle modified_on: preserve from state if no actual changes
		// This prevents showing as "known after apply" when nothing changed
		if plan.ModifiedOn.IsUnknown() {
			if !hasChanges {
				// No actual changes, preserve modified_on from state
				plan.ModifiedOn = state.ModifiedOn
			}
			// Otherwise let it be unknown (will be updated by the API)
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

		// Handle settings: preserve from state if not explicitly set or if empty object
		if plan.Settings.IsUnknown() || plan.Settings.IsNull() {
			plan.Settings = state.Settings
		} else if !plan.Settings.IsNull() && !plan.Settings.IsUnknown() {
			// Check if user provided empty settings {}
			var planSettingsData DNSRecordSettingsModel
			diags := plan.Settings.As(ctx, &planSettingsData, basetypes.ObjectAsOptions{})
			if !diags.HasError() {
				// If all fields are null/unknown (empty object), preserve state
				if (planSettingsData.IPV4Only.IsNull() || planSettingsData.IPV4Only.IsUnknown()) &&
					(planSettingsData.IPV6Only.IsNull() || planSettingsData.IPV6Only.IsUnknown()) &&
					(planSettingsData.FlattenCNAME.IsNull() || planSettingsData.FlattenCNAME.IsUnknown()) {
					// User provided empty settings {}, preserve whatever is in state
					if state.Settings.IsNull() || state.Settings.IsUnknown() {
						// State has no settings, keep it that way
						plan.Settings = state.Settings
					} else {
						// State has settings, check if they're just defaults
						var stateSettingsData DNSRecordSettingsModel
						diags := state.Settings.As(ctx, &stateSettingsData, basetypes.ObjectAsOptions{})
						if !diags.HasError() {
							// Check if state only has false values (defaults)
							allDefaults := true
							if !stateSettingsData.IPV4Only.IsNull() && !stateSettingsData.IPV4Only.IsUnknown() && stateSettingsData.IPV4Only.ValueBool() {
								allDefaults = false
							}
							if !stateSettingsData.IPV6Only.IsNull() && !stateSettingsData.IPV6Only.IsUnknown() && stateSettingsData.IPV6Only.ValueBool() {
								allDefaults = false
							}
							if !stateSettingsData.FlattenCNAME.IsNull() && !stateSettingsData.FlattenCNAME.IsUnknown() && stateSettingsData.FlattenCNAME.ValueBool() {
								allDefaults = false
							}
							if allDefaults {
								// State only has defaults, preserve it to avoid drift
								plan.Settings = state.Settings
							}
						}
					}
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

	// For CNAME records, resolve null/unknown settings sub-fields to false in the plan.
	// These fields (ipv4_only, ipv6_only, flatten_cname) only apply to CNAME records.
	// The API never returns them for other record types, so we only inject defaults here
	// for CNAME to avoid spurious diffs on A/AAAA/MX/etc. records with settings = {}.
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() && plan.Type.ValueString() == "CNAME" {
		if !plan.Settings.IsUnknown() {
			var s DNSRecordSettingsModel
			if !plan.Settings.IsNull() {
				plan.Settings.As(ctx, &s, basetypes.ObjectAsOptions{})
			}
			changed := false
			if s.IPV4Only.IsNull() || s.IPV4Only.IsUnknown() {
				s.IPV4Only = types.BoolValue(false)
				changed = true
			}
			if s.IPV6Only.IsNull() || s.IPV6Only.IsUnknown() {
				s.IPV6Only = types.BoolValue(false)
				changed = true
			}
			if s.FlattenCNAME.IsNull() || s.FlattenCNAME.IsUnknown() {
				s.FlattenCNAME = types.BoolValue(false)
				changed = true
			}
			if changed {
				if normalized, diags := customfield.NewObject[DNSRecordSettingsModel](ctx, &s); !diags.HasError() {
					plan.Settings = normalized
				}
			}
		}
	}

	// Set the updated plan
	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
}

// ---------------------------------------------------------------------------
// normalizeSettings – called from Read, Create, Update, and ImportState in
// resource.go to stabilize the settings object in state.
// ---------------------------------------------------------------------------

// normalizeSettings unconditionally materializes the settings object with false defaults
// for any sub-fields that are null. This handles two cases:
//  1. The API returns a settings object but omits some sub-fields (null sub-fields).
//  2. The API omits the settings object entirely (data.Settings.IsNull()), which it does
//     for non-proxied CNAME records where these fields are not applicable.
//
// In both cases we treat absent as false so that state is stable against a config of
// settings = { flatten_cname = false } (or ipv4_only / ipv6_only = false).
func normalizeSettings(ctx context.Context, data *DNSRecordModel, diags *diag.Diagnostics) {
	// settings sub-fields (ipv4_only, ipv6_only, flatten_cname) only apply to CNAME
	// records. For all other record types the API never returns a settings object, so
	// leave state as null to avoid injecting spurious settings into non-CNAME records.
	if data.Type.IsNull() || data.Type.IsUnknown() || data.Type.ValueString() != "CNAME" {
		return
	}

	var s DNSRecordSettingsModel
	if !data.Settings.IsNull() && !data.Settings.IsUnknown() {
		// Populate s from the existing object; ignore errors (s stays zero-value).
		data.Settings.As(ctx, &s, basetypes.ObjectAsOptions{})
	}
	// s sub-fields are null either because Settings was null (API omitted the object
	// entirely for non-proxied CNAME records) or because the API returned the object
	// without those particular sub-fields. Treat absent as false.
	if s.IPV4Only.IsNull() {
		s.IPV4Only = types.BoolValue(false)
	}
	if s.IPV6Only.IsNull() {
		s.IPV6Only = types.BoolValue(false)
	}
	if s.FlattenCNAME.IsNull() {
		s.FlattenCNAME = types.BoolValue(false)
	}
	normalized, d := customfield.NewObject[DNSRecordSettingsModel](ctx, &s)
	diags.Append(d...)
	if !diags.HasError() {
		data.Settings = normalized
	}
}

// ---------------------------------------------------------------------------
// DNSRecordDataModel.Equal – semantic equality for the data sub-object.
// ---------------------------------------------------------------------------

// Equal compares two DNSRecordDataModel instances for equality
func (d *DNSRecordDataModel) Equal(other *DNSRecordDataModel) bool {
	if d == nil && other == nil {
		return true
	}
	if d == nil || other == nil {
		return false
	}

	// Use semantic equality for Flags to handle Int64(0) == Float64(0.0)
	ctx := context.Background()
	flagsEqual := true
	if !d.Flags.IsNull() || !other.Flags.IsNull() {
		var diags diag.Diagnostics
		flagsEqual, diags = d.Flags.DynamicSemanticEquals(ctx, other.Flags)
		if diags.HasError() {
			flagsEqual = d.Flags.Equal(other.Flags)
		}
	}

	return flagsEqual &&
		d.Tag.Equal(other.Tag) &&
		d.Value.Equal(other.Value) &&
		d.Algorithm.Equal(other.Algorithm) &&
		d.Certificate.Equal(other.Certificate) &&
		d.KeyTag.Equal(other.KeyTag) &&
		d.Type.Equal(other.Type) &&
		d.Protocol.Equal(other.Protocol) &&
		d.PublicKey.Equal(other.PublicKey) &&
		d.Digest.Equal(other.Digest) &&
		d.DigestType.Equal(other.DigestType) &&
		d.Priority.Equal(other.Priority) &&
		d.Target.Equal(other.Target) &&
		d.Altitude.Equal(other.Altitude) &&
		d.LatDegrees.Equal(other.LatDegrees) &&
		d.LatDirection.Equal(other.LatDirection) &&
		d.LatMinutes.Equal(other.LatMinutes) &&
		d.LatSeconds.Equal(other.LatSeconds) &&
		d.LongDegrees.Equal(other.LongDegrees) &&
		d.LongDirection.Equal(other.LongDirection) &&
		d.LongMinutes.Equal(other.LongMinutes) &&
		d.LongSeconds.Equal(other.LongSeconds) &&
		d.PrecisionHorz.Equal(other.PrecisionHorz) &&
		d.PrecisionVert.Equal(other.PrecisionVert) &&
		d.Size.Equal(other.Size) &&
		d.Order.Equal(other.Order) &&
		d.Preference.Equal(other.Preference) &&
		d.Regex.Equal(other.Regex) &&
		d.Replacement.Equal(other.Replacement) &&
		d.Service.Equal(other.Service) &&
		d.MatchingType.Equal(other.MatchingType) &&
		d.Selector.Equal(other.Selector) &&
		d.Usage.Equal(other.Usage) &&
		d.Port.Equal(other.Port) &&
		d.Weight.Equal(other.Weight) &&
		d.Fingerprint.Equal(other.Fingerprint)
}
