package worker

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// importStateWithPropagationPolicyDefault returns a synthetic WorkerModel with
// observability.traces.propagation_policy populated with the documented
// default ("authenticated"). It is used during ImportState where no prior
// state is available, so the preservePropagationPolicy fallback can still
// keep the field populated when the API omits it.
func importStateWithPropagationPolicyDefault(ctx context.Context) *WorkerModel {
	traces := WorkerObservabilityTracesModel{PropagationPolicy: types.StringValue("authenticated")}
	tracesObj, _ := customfield.NewObject(ctx, &traces)
	obs := WorkerObservabilityModel{Traces: tracesObj}
	obsObj, _ := customfield.NewObject(ctx, &obs)
	return &WorkerModel{Observability: obsObj}
}

// preservePropagationPolicy restores observability.traces.propagation_policy from
// the prior state when the API response is missing it.
//
// The Cloudflare Workers API does not consistently echo back the
// `propagation_policy` field even when it is sent in the request body. This
// causes spurious refresh plans where the field is set via Default in the
// schema, then later overwritten to null on Read because the API omitted it.
//
// This helper applies a targeted state-preservation fix: after the API
// response has been unmarshaled, if `propagation_policy` came back null and
// the prior state held a value, the prior state value is restored. When the
// API does return a value (e.g. user set it explicitly via the dashboard),
// that value is preserved as-is.
func preservePropagationPolicy(ctx context.Context, data, prior *WorkerModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil || prior == nil {
		return diags
	}
	if data.Observability.IsNull() || data.Observability.IsUnknown() {
		return diags
	}
	if prior.Observability.IsNull() || prior.Observability.IsUnknown() {
		return diags
	}

	dataObs, d := data.Observability.Value(ctx)
	diags.Append(d...)
	if diags.HasError() || dataObs == nil {
		return diags
	}
	priorObs, d := prior.Observability.Value(ctx)
	diags.Append(d...)
	if diags.HasError() || priorObs == nil {
		return diags
	}

	if dataObs.Traces.IsNull() || dataObs.Traces.IsUnknown() {
		return diags
	}
	if priorObs.Traces.IsNull() || priorObs.Traces.IsUnknown() {
		return diags
	}

	dataTraces, d := dataObs.Traces.Value(ctx)
	diags.Append(d...)
	if diags.HasError() || dataTraces == nil {
		return diags
	}
	priorTraces, d := priorObs.Traces.Value(ctx)
	diags.Append(d...)
	if diags.HasError() || priorTraces == nil {
		return diags
	}

	// Only restore when API didn't return propagation_policy AND state had a value.
	if !dataTraces.PropagationPolicy.IsNull() && !dataTraces.PropagationPolicy.IsUnknown() {
		return diags
	}
	if priorTraces.PropagationPolicy.IsNull() || priorTraces.PropagationPolicy.IsUnknown() {
		return diags
	}

	dataTraces.PropagationPolicy = priorTraces.PropagationPolicy

	newTraces, d := customfield.NewObject(ctx, dataTraces)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	dataObs.Traces = newTraces

	newObs, d := customfield.NewObject(ctx, dataObs)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Observability = newObs

	return diags
}
