package rulesets

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type sbfmDeprecationWarningValidator struct{}

func (v sbfmDeprecationWarningValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Cloudflare is going to change the way Super Bot Fight Mode managed rules are configured through Terraform and our API. No action is required at this time. " +
		" Please follow updates to our documentation regarding this here: https://developers.cloudflare.com/bots/get-started/biz-and-ent/#ruleset-engine")
}

func (v sbfmDeprecationWarningValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Cloudflare is going to change the way Super Bot Fight Mode managed rules are configured through Terraform and our API. **No action is required at this time**. " +
		" Please follow updates to our documentation regarding this [here](https://developers.cloudflare.com/bots/get-started/biz-and-ent/#ruleset-engine)")
}

func (v sbfmDeprecationWarningValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	if req.ConfigValue.ValueString() == string(cloudflare.RulesetPhaseSuperBotFightMode) {
		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			fmt.Sprintf(`%q phase will soon be deprecated in the "cloudflare_ruleset" resource`, string(cloudflare.RulesetPhaseSuperBotFightMode)),
			v.Description(ctx),
		)

		return
	}
}
