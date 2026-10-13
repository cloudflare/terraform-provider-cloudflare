package zero_trust_access_application

import (
	"context"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var infrastructureDurationRegex = regexp.MustCompile(`^(?:0|[-+]?(\d+(?:\.\d*)?|\.\d+)(?:s|m|h)(?:(\d+(?:\.\d*)?|\.\d+)(?:s|m|h))*)$`)

type infrastructureSessionDurationValidator struct{}

var _ validator.String = infrastructureSessionDurationValidator{}

func (infrastructureSessionDurationValidator) Description(context.Context) string {
	return "For infrastructure applications, session_duration must use s, m, or h and resolve to whole seconds."
}

func (v infrastructureSessionDurationValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v infrastructureSessionDurationValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	var appType types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("type"), &appType)...)
	if resp.Diagnostics.HasError() || appType.IsNull() || appType.IsUnknown() || appType.ValueString() != "infrastructure" {
		return
	}
	value := req.ConfigValue.ValueString()
	duration, err := time.ParseDuration(value)
	if !infrastructureDurationRegex.MatchString(value) || err != nil || duration%time.Second != 0 {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid Infrastructure Session Duration", v.Description(ctx))
	}
}
