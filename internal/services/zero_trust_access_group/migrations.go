// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessGroupResource)(nil)

// zeroTrustAccessGroupResourceSchemaV0 defines the v0 schema
// Only includes simple attributes that are consistent between v4 and v5
// Complex fields (include/exclude/require) are handled via raw JSON
var zeroTrustAccessGroupResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Optional: true,
		},
	},
	// Note: include, exclude, and require are intentionally omitted
	// They have different structures in v4 (blocks with arrays) vs v5 (list of objects)
	// and will be handled via raw JSON manipulation
}

func (r *ZeroTrustAccessGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade from v0 (v4 provider) to v1 (v5 provider)
		0: {
			PriorSchema:   &zeroTrustAccessGroupResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessGroupStateV0toV1,
		},
	}
}

func upgradeZeroTrustAccessGroupStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return
	}

	// Get simple attributes from typed state
	var priorStateData struct {
		ID        types.String `tfsdk:"id"`
		AccountID types.String `tfsdk:"account_id"`
		ZoneID    types.String `tfsdk:"zone_id"`
		Name      types.String `tfsdk:"name"`
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting zero trust access group state migration from v0 to v1")

	// Create new state structure with simple attributes
	var newState ZeroTrustAccessGroupModel
	newState.ID = priorStateData.ID
	newState.AccountID = priorStateData.AccountID
	newState.ZoneID = priorStateData.ZoneID
	newState.Name = priorStateData.Name

	// Use raw JSON to handle include/exclude/require which have different formats in v4 vs v5
	rawJSON := string(req.RawState.JSON)

	// Check if this is v4 format (blocks with arrays) or v5 format (list of objects)
	includeValue := gjson.Get(rawJSON, "include")
	if includeValue.Exists() && includeValue.IsArray() && len(includeValue.Array()) > 0 {
		firstItem := includeValue.Array()[0]

		// Check if it's v4 format: include[0].email is an array
		// Or v5 format: include[0].email is an object
		if firstItem.Get("email").IsArray() {
			// V4 format - migrate from blocks with arrays
			newState.Include = migrateV4IncludeRules(rawJSON)
		} else if firstItem.Get("email").IsObject() {
			// V5 format - already in correct structure, just parse it
			newState.Include = parseV5IncludeRules(rawJSON)
		}
	}

	// Handle exclude rules
	excludeValue := gjson.Get(rawJSON, "exclude")
	if excludeValue.Exists() && excludeValue.IsArray() && len(excludeValue.Array()) > 0 {
		firstItem := excludeValue.Array()[0]

		if firstItem.Get("email").IsArray() || firstItem.Get("email_domain").IsArray() || firstItem.Get("ip").IsArray() {
			// V4 format
			newState.Exclude = migrateV4ExcludeRules(rawJSON)
		} else {
			// V5 format
			newState.Exclude = parseV5ExcludeRules(rawJSON)
		}
	}

	// Handle require rules
	requireValue := gjson.Get(rawJSON, "require")
	if requireValue.Exists() && requireValue.IsArray() && len(requireValue.Array()) > 0 {
		firstItem := requireValue.Array()[0]

		if firstItem.Get("email").IsArray() || firstItem.Get("email_domain").IsArray() || firstItem.Get("ip").IsArray() {
			// V4 format
			newState.Require = migrateV4RequireRules(rawJSON)
		} else {
			// V5 format
			newState.Require = parseV5RequireRules(rawJSON)
		}
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}

	tflog.Info(ctx, "Successfully migrated zero trust access group state from v0 to v1")
}

// migrateV4IncludeRules handles v4 format where include is a block with arrays
func migrateV4IncludeRules(rawJSON string) *[]*ZeroTrustAccessGroupIncludeModel {
	includeValue := gjson.Get(rawJSON, "include")
	if !includeValue.Exists() || !includeValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupIncludeModel

	for _, includeBlock := range includeValue.Array() {
		// Process email arrays
		emails := includeBlock.Get("email").Array()
		for _, email := range emails {
			rule := &ZeroTrustAccessGroupIncludeModel{
				Email: &ZeroTrustAccessGroupIncludeEmailModel{
					Email: types.StringValue(email.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process email_domain arrays
		domains := includeBlock.Get("email_domain").Array()
		for _, domain := range domains {
			rule := &ZeroTrustAccessGroupIncludeModel{
				EmailDomain: &ZeroTrustAccessGroupIncludeEmailDomainModel{
					Domain: types.StringValue(domain.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process IP arrays
		ips := includeBlock.Get("ip").Array()
		for _, ip := range ips {
			rule := &ZeroTrustAccessGroupIncludeModel{
				IP: &ZeroTrustAccessGroupIncludeIPModel{
					IP: types.StringValue(ip.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process boolean fields
		if includeBlock.Get("everyone").Bool() {
			rule := &ZeroTrustAccessGroupIncludeModel{
				Everyone: &ZeroTrustAccessGroupIncludeEveryoneModel{},
			}
			rules = append(rules, rule)
		}

		if includeBlock.Get("certificate").Bool() {
			rule := &ZeroTrustAccessGroupIncludeModel{
				Certificate: &ZeroTrustAccessGroupIncludeCertificateModel{},
			}
			rules = append(rules, rule)
		}

		if includeBlock.Get("any_valid_service_token").Bool() {
			rule := &ZeroTrustAccessGroupIncludeModel{
				AnyValidServiceToken: &ZeroTrustAccessGroupIncludeAnyValidServiceTokenModel{},
			}
			rules = append(rules, rule)
		}
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}

// parseV5IncludeRules handles v5 format where include is already a list of objects
func parseV5IncludeRules(rawJSON string) *[]*ZeroTrustAccessGroupIncludeModel {
	includeValue := gjson.Get(rawJSON, "include")
	if !includeValue.Exists() || !includeValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupIncludeModel

	for _, item := range includeValue.Array() {
		rule := &ZeroTrustAccessGroupIncludeModel{}

		// Check which field is present in this object
		if email := item.Get("email"); email.Exists() && email.IsObject() {
			rule.Email = &ZeroTrustAccessGroupIncludeEmailModel{
				Email: types.StringValue(email.Get("email").String()),
			}
		} else if emailDomain := item.Get("email_domain"); emailDomain.Exists() && emailDomain.IsObject() {
			rule.EmailDomain = &ZeroTrustAccessGroupIncludeEmailDomainModel{
				Domain: types.StringValue(emailDomain.Get("domain").String()),
			}
		} else if ip := item.Get("ip"); ip.Exists() && ip.IsObject() {
			rule.IP = &ZeroTrustAccessGroupIncludeIPModel{
				IP: types.StringValue(ip.Get("ip").String()),
			}
		} else if item.Get("everyone").Exists() {
			rule.Everyone = &ZeroTrustAccessGroupIncludeEveryoneModel{}
		} else if item.Get("certificate").Exists() {
			rule.Certificate = &ZeroTrustAccessGroupIncludeCertificateModel{}
		} else if item.Get("any_valid_service_token").Exists() {
			rule.AnyValidServiceToken = &ZeroTrustAccessGroupIncludeAnyValidServiceTokenModel{}
		}

		rules = append(rules, rule)
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}

// migrateV4ExcludeRules handles v4 format for exclude rules
func migrateV4ExcludeRules(rawJSON string) *[]*ZeroTrustAccessGroupExcludeModel {
	excludeValue := gjson.Get(rawJSON, "exclude")
	if !excludeValue.Exists() || !excludeValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupExcludeModel

	for _, excludeBlock := range excludeValue.Array() {
		// Process email arrays
		emails := excludeBlock.Get("email").Array()
		for _, email := range emails {
			rule := &ZeroTrustAccessGroupExcludeModel{
				Email: &ZeroTrustAccessGroupExcludeEmailModel{
					Email: types.StringValue(email.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process email_domain arrays
		domains := excludeBlock.Get("email_domain").Array()
		for _, domain := range domains {
			rule := &ZeroTrustAccessGroupExcludeModel{
				EmailDomain: &ZeroTrustAccessGroupExcludeEmailDomainModel{
					Domain: types.StringValue(domain.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process IP arrays
		ips := excludeBlock.Get("ip").Array()
		for _, ip := range ips {
			rule := &ZeroTrustAccessGroupExcludeModel{
				IP: &ZeroTrustAccessGroupExcludeIPModel{
					IP: types.StringValue(ip.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process boolean fields
		if excludeBlock.Get("everyone").Bool() {
			rule := &ZeroTrustAccessGroupExcludeModel{
				Everyone: &ZeroTrustAccessGroupExcludeEveryoneModel{},
			}
			rules = append(rules, rule)
		}

		if excludeBlock.Get("certificate").Bool() {
			rule := &ZeroTrustAccessGroupExcludeModel{
				Certificate: &ZeroTrustAccessGroupExcludeCertificateModel{},
			}
			rules = append(rules, rule)
		}

		if excludeBlock.Get("any_valid_service_token").Bool() {
			rule := &ZeroTrustAccessGroupExcludeModel{
				AnyValidServiceToken: &ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel{},
			}
			rules = append(rules, rule)
		}
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}

// parseV5ExcludeRules handles v5 format for exclude rules
func parseV5ExcludeRules(rawJSON string) *[]*ZeroTrustAccessGroupExcludeModel {
	excludeValue := gjson.Get(rawJSON, "exclude")
	if !excludeValue.Exists() || !excludeValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupExcludeModel

	for _, item := range excludeValue.Array() {
		rule := &ZeroTrustAccessGroupExcludeModel{}

		// Check which field is present in this object
		if email := item.Get("email"); email.Exists() && email.IsObject() {
			rule.Email = &ZeroTrustAccessGroupExcludeEmailModel{
				Email: types.StringValue(email.Get("email").String()),
			}
		} else if emailDomain := item.Get("email_domain"); emailDomain.Exists() && emailDomain.IsObject() {
			rule.EmailDomain = &ZeroTrustAccessGroupExcludeEmailDomainModel{
				Domain: types.StringValue(emailDomain.Get("domain").String()),
			}
		} else if ip := item.Get("ip"); ip.Exists() && ip.IsObject() {
			rule.IP = &ZeroTrustAccessGroupExcludeIPModel{
				IP: types.StringValue(ip.Get("ip").String()),
			}
		} else if item.Get("everyone").Exists() {
			rule.Everyone = &ZeroTrustAccessGroupExcludeEveryoneModel{}
		} else if item.Get("certificate").Exists() {
			rule.Certificate = &ZeroTrustAccessGroupExcludeCertificateModel{}
		} else if item.Get("any_valid_service_token").Exists() {
			rule.AnyValidServiceToken = &ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel{}
		}

		rules = append(rules, rule)
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}

// migrateV4RequireRules handles v4 format for require rules
func migrateV4RequireRules(rawJSON string) *[]*ZeroTrustAccessGroupRequireModel {
	requireValue := gjson.Get(rawJSON, "require")
	if !requireValue.Exists() || !requireValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupRequireModel

	for _, requireBlock := range requireValue.Array() {
		// Process email arrays
		emails := requireBlock.Get("email").Array()
		for _, email := range emails {
			rule := &ZeroTrustAccessGroupRequireModel{
				Email: &ZeroTrustAccessGroupRequireEmailModel{
					Email: types.StringValue(email.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process email_domain arrays
		domains := requireBlock.Get("email_domain").Array()
		for _, domain := range domains {
			rule := &ZeroTrustAccessGroupRequireModel{
				EmailDomain: &ZeroTrustAccessGroupRequireEmailDomainModel{
					Domain: types.StringValue(domain.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process IP arrays
		ips := requireBlock.Get("ip").Array()
		for _, ip := range ips {
			rule := &ZeroTrustAccessGroupRequireModel{
				IP: &ZeroTrustAccessGroupRequireIPModel{
					IP: types.StringValue(ip.String()),
				},
			}
			rules = append(rules, rule)
		}

		// Process boolean fields
		if requireBlock.Get("everyone").Bool() {
			rule := &ZeroTrustAccessGroupRequireModel{
				Everyone: &ZeroTrustAccessGroupRequireEveryoneModel{},
			}
			rules = append(rules, rule)
		}

		if requireBlock.Get("certificate").Bool() {
			rule := &ZeroTrustAccessGroupRequireModel{
				Certificate: &ZeroTrustAccessGroupRequireCertificateModel{},
			}
			rules = append(rules, rule)
		}

		if requireBlock.Get("any_valid_service_token").Bool() {
			rule := &ZeroTrustAccessGroupRequireModel{
				AnyValidServiceToken: &ZeroTrustAccessGroupRequireAnyValidServiceTokenModel{},
			}
			rules = append(rules, rule)
		}
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}

// parseV5RequireRules handles v5 format for require rules
func parseV5RequireRules(rawJSON string) *[]*ZeroTrustAccessGroupRequireModel {
	requireValue := gjson.Get(rawJSON, "require")
	if !requireValue.Exists() || !requireValue.IsArray() {
		return nil
	}

	var rules []*ZeroTrustAccessGroupRequireModel

	for _, item := range requireValue.Array() {
		rule := &ZeroTrustAccessGroupRequireModel{}

		// Check which field is present in this object
		if email := item.Get("email"); email.Exists() && email.IsObject() {
			rule.Email = &ZeroTrustAccessGroupRequireEmailModel{
				Email: types.StringValue(email.Get("email").String()),
			}
		} else if emailDomain := item.Get("email_domain"); emailDomain.Exists() && emailDomain.IsObject() {
			rule.EmailDomain = &ZeroTrustAccessGroupRequireEmailDomainModel{
				Domain: types.StringValue(emailDomain.Get("domain").String()),
			}
		} else if ip := item.Get("ip"); ip.Exists() && ip.IsObject() {
			rule.IP = &ZeroTrustAccessGroupRequireIPModel{
				IP: types.StringValue(ip.Get("ip").String()),
			}
		} else if item.Get("everyone").Exists() {
			rule.Everyone = &ZeroTrustAccessGroupRequireEveryoneModel{}
		} else if item.Get("certificate").Exists() {
			rule.Certificate = &ZeroTrustAccessGroupRequireCertificateModel{}
		} else if item.Get("any_valid_service_token").Exists() {
			rule.AnyValidServiceToken = &ZeroTrustAccessGroupRequireAnyValidServiceTokenModel{}
		}

		rules = append(rules, rule)
	}

	if len(rules) > 0 {
		return &rules
	}
	return nil
}
