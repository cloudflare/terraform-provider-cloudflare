package container_application

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Terraform state model

type ContainerApplicationModel struct {
	ID                       types.String                          `tfsdk:"id"`
	AccountID                types.String                          `tfsdk:"account_id"`
	ApplicationID            types.String                          `tfsdk:"application_id"`
	ScriptName               types.String                          `tfsdk:"script_name"`
	ClassName                types.String                          `tfsdk:"class_name"`
	Name                     types.String                          `tfsdk:"name"`
	Image                    types.String                          `tfsdk:"image"`
	MaxInstances             types.Int64                           `tfsdk:"max_instances"`
	InstanceType             types.String                          `tfsdk:"instance_type"`
	Vcpu                     types.Float64                         `tfsdk:"vcpu"`
	MemoryMib                types.Int64                           `tfsdk:"memory_mib"`
	DiskSizeMb               types.Int64                           `tfsdk:"disk_size_mb"`
	SchedulingPolicy         types.String                          `tfsdk:"scheduling_policy"`
	RolloutStepPercentage    types.Int64                           `tfsdk:"rollout_step_percentage"`
	RolloutKind              types.String                          `tfsdk:"rollout_kind"`
	RolloutActiveGracePeriod types.Int64                           `tfsdk:"rollout_active_grace_period"`
	Constraints              *ContainerApplicationConstraintsModel `tfsdk:"constraints"`
	Affinities               *ContainerApplicationAffinitiesModel  `tfsdk:"affinities"`
	Observability            *ContainerApplicationObservabilityModel `tfsdk:"observability"`
	WranglerSSH              *ContainerApplicationWranglerSSHModel `tfsdk:"wrangler_ssh"`
	AuthorizedKeys           *[]*ContainerApplicationAuthorizedKeyModel  `tfsdk:"authorized_keys"`
	TrustedUserCAKeys        *[]*ContainerApplicationTrustedCAKeyModel   `tfsdk:"trusted_user_ca_keys"`
}

type ContainerApplicationConstraintsModel struct {
	Tiers   types.List `tfsdk:"tiers"`
	Regions types.List `tfsdk:"regions"`
	Cities  types.List `tfsdk:"cities"`
}

type ContainerApplicationAffinitiesModel struct {
	Colocation         types.String `tfsdk:"colocation"`
	HardwareGeneration types.String `tfsdk:"hardware_generation"`
}

type ContainerApplicationObservabilityModel struct {
	LogsEnabled types.Bool `tfsdk:"logs_enabled"`
}

type ContainerApplicationWranglerSSHModel struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Port    types.Int64 `tfsdk:"port"`
}

type ContainerApplicationAuthorizedKeyModel struct {
	Name      types.String `tfsdk:"name"`
	PublicKey types.String `tfsdk:"public_key"`
}

type ContainerApplicationTrustedCAKeyModel struct {
	Name      types.String `tfsdk:"name"`
	PublicKey types.String `tfsdk:"public_key"`
}

// API request/response types

type apiApplicationEnvelope struct {
	Result  apiApplication `json:"result"`
	Success bool           `json:"success"`
}

type apiApplicationListEnvelope struct {
	Result  []apiApplication `json:"result"`
	Success bool             `json:"success"`
}

type apiDONamespaceListEnvelope struct {
	Result  []apiDurableObjectNamespace `json:"result"`
	Success bool                        `json:"success"`
}

type apiApplication struct {
	ID                       string             `json:"id"`
	Name                     string             `json:"name"`
	Configuration            apiConfiguration   `json:"configuration"`
	Instances                int64              `json:"instances"`
	MaxInstances             int64              `json:"max_instances"`
	SchedulingPolicy         string             `json:"scheduling_policy"`
	Constraints              *apiConstraints    `json:"constraints,omitempty"`
	Affinities               *apiAffinities     `json:"affinities,omitempty"`
	DurableObjects           *apiDurableObjects `json:"durable_objects,omitempty"`
	RolloutActiveGracePeriod int64              `json:"rollout_active_grace_period"`
}

type apiConfiguration struct {
	Image         string            `json:"image"`
	InstanceType  string            `json:"instance_type,omitempty"`
	Vcpu          *float64          `json:"vcpu,omitempty"`
	MemoryMib     *int64            `json:"memory_mib,omitempty"`
	Disk          *apiDisk          `json:"disk,omitempty"`
	Observability *apiObservability `json:"observability,omitempty"`
	WranglerSSH   *apiWranglerSSH  `json:"wrangler_ssh,omitempty"`
	AuthorizedKeys    []apiAuthorizedKey `json:"authorized_keys,omitempty"`
	TrustedUserCAKeys []apiTrustedCAKey  `json:"trusted_user_ca_keys,omitempty"`
}

type apiDisk struct {
	SizeMb int64 `json:"size_mb"`
}

type apiObservability struct {
	Logs *apiObservabilityLogs `json:"logs,omitempty"`
}

type apiObservabilityLogs struct {
	Enabled bool `json:"enabled"`
}

type apiWranglerSSH struct {
	Enabled bool  `json:"enabled"`
	Port    int64 `json:"port,omitempty"`
}

type apiAuthorizedKey struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

type apiTrustedCAKey struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key"`
}

type apiConstraints struct {
	Tiers   []int64  `json:"tiers,omitempty"`
	Regions []string `json:"regions,omitempty"`
	Cities  []string `json:"cities,omitempty"`
}

type apiAffinities struct {
	Colocation         string `json:"colocation,omitempty"`
	HardwareGeneration string `json:"hardware_generation,omitempty"`
}

type apiDurableObjects struct {
	NamespaceID string `json:"namespace_id"`
}

type apiCreateApplicationRequest struct {
	Name                     string            `json:"name"`
	SchedulingPolicy         string            `json:"scheduling_policy"`
	Configuration            apiConfiguration  `json:"configuration"`
	Instances                int64             `json:"instances"`
	MaxInstances             int64             `json:"max_instances"`
	Constraints              *apiConstraints   `json:"constraints,omitempty"`
	Affinities               *apiAffinities    `json:"affinities,omitempty"`
	DurableObjects           apiDurableObjects `json:"durable_objects"`
	RolloutActiveGracePeriod int64             `json:"rollout_active_grace_period"`
}

type apiModifyApplicationRequest struct {
	Configuration            *apiConfiguration `json:"configuration,omitempty"`
	MaxInstances             *int64            `json:"max_instances,omitempty"`
	SchedulingPolicy         *string           `json:"scheduling_policy,omitempty"`
	Constraints              *apiConstraints   `json:"constraints,omitempty"`
	Affinities               *apiAffinities    `json:"affinities,omitempty"`
	RolloutActiveGracePeriod *int64            `json:"rollout_active_grace_period,omitempty"`
}

type apiCreateRolloutRequest struct {
	Description         string           `json:"description"`
	Strategy            string           `json:"strategy"`
	TargetConfiguration apiConfiguration `json:"target_configuration"`
	StepPercentage      int64            `json:"step_percentage,omitempty"`
	Kind                string           `json:"kind"`
}

type apiDurableObjectNamespace struct {
	ID     string `json:"id"`
	Class  string `json:"class"`
	Name   string `json:"name"`
	Script string `json:"script"`
}

// Conversion functions

func (m *ContainerApplicationModel) effectiveName() string {
	name := m.Name.ValueString()
	if name == "" || m.Name.IsNull() || m.Name.IsUnknown() {
		return fmt.Sprintf("%s-%s", m.ScriptName.ValueString(), m.ClassName.ValueString())
	}
	return name
}

func (m *ContainerApplicationModel) buildConfiguration() apiConfiguration {
	cfg := apiConfiguration{
		Image: m.Image.ValueString(),
	}

	// Named instance type vs custom resources (mutually exclusive)
	if !m.InstanceType.IsNull() && !m.InstanceType.IsUnknown() && m.InstanceType.ValueString() != "" {
		cfg.InstanceType = m.InstanceType.ValueString()
	}
	if !m.Vcpu.IsNull() && !m.Vcpu.IsUnknown() {
		v := m.Vcpu.ValueFloat64()
		cfg.Vcpu = &v
	}
	if !m.MemoryMib.IsNull() && !m.MemoryMib.IsUnknown() {
		v := m.MemoryMib.ValueInt64()
		cfg.MemoryMib = &v
	}
	if !m.DiskSizeMb.IsNull() && !m.DiskSizeMb.IsUnknown() && m.DiskSizeMb.ValueInt64() > 0 {
		cfg.Disk = &apiDisk{SizeMb: m.DiskSizeMb.ValueInt64()}
	}

	// Observability — default to logs enabled if not specified
	if m.Observability != nil {
		cfg.Observability = &apiObservability{
			Logs: &apiObservabilityLogs{
				Enabled: m.Observability.LogsEnabled.ValueBool(),
			},
		}
	} else {
		cfg.Observability = &apiObservability{
			Logs: &apiObservabilityLogs{Enabled: true},
		}
	}

	// SSH config
	if m.WranglerSSH != nil {
		ssh := &apiWranglerSSH{
			Enabled: m.WranglerSSH.Enabled.ValueBool(),
		}
		if !m.WranglerSSH.Port.IsNull() && !m.WranglerSSH.Port.IsUnknown() {
			ssh.Port = m.WranglerSSH.Port.ValueInt64()
		}
		cfg.WranglerSSH = ssh
	}

	if m.AuthorizedKeys != nil {
		for _, k := range *m.AuthorizedKeys {
			cfg.AuthorizedKeys = append(cfg.AuthorizedKeys, apiAuthorizedKey{
				Name:      k.Name.ValueString(),
				PublicKey: k.PublicKey.ValueString(),
			})
		}
	}

	if m.TrustedUserCAKeys != nil {
		for _, k := range *m.TrustedUserCAKeys {
			cfg.TrustedUserCAKeys = append(cfg.TrustedUserCAKeys, apiTrustedCAKey{
				Name:      k.Name.ValueString(),
				PublicKey: k.PublicKey.ValueString(),
			})
		}
	}

	return cfg
}

func (m *ContainerApplicationModel) affinitesToAPI() *apiAffinities {
	if m.Affinities == nil {
		return nil
	}
	a := &apiAffinities{}
	if !m.Affinities.Colocation.IsNull() && !m.Affinities.Colocation.IsUnknown() {
		a.Colocation = m.Affinities.Colocation.ValueString()
	}
	if !m.Affinities.HardwareGeneration.IsNull() && !m.Affinities.HardwareGeneration.IsUnknown() {
		a.HardwareGeneration = m.Affinities.HardwareGeneration.ValueString()
	}
	return a
}

func (m *ContainerApplicationModel) toCreateRequest(ctx context.Context, namespaceID string) (apiCreateApplicationRequest, error) {
	req := apiCreateApplicationRequest{
		Name:             m.effectiveName(),
		SchedulingPolicy: m.SchedulingPolicy.ValueString(),
		Configuration:    m.buildConfiguration(),
		Instances:        0,
		MaxInstances:     m.MaxInstances.ValueInt64(),
		DurableObjects: apiDurableObjects{
			NamespaceID: namespaceID,
		},
		RolloutActiveGracePeriod: m.RolloutActiveGracePeriod.ValueInt64(),
		Affinities:               m.affinitesToAPI(),
	}

	constraints, err := m.constraintsToAPI(ctx)
	if err != nil {
		return req, err
	}
	req.Constraints = constraints

	return req, nil
}

func (m *ContainerApplicationModel) toModifyRequest(ctx context.Context) (apiModifyApplicationRequest, error) {
	maxInstances := m.MaxInstances.ValueInt64()
	schedulingPolicy := m.SchedulingPolicy.ValueString()
	rolloutActiveGracePeriod := m.RolloutActiveGracePeriod.ValueInt64()
	cfg := m.buildConfiguration()

	req := apiModifyApplicationRequest{
		Configuration:            &cfg,
		MaxInstances:             &maxInstances,
		SchedulingPolicy:         &schedulingPolicy,
		RolloutActiveGracePeriod: &rolloutActiveGracePeriod,
		Affinities:               m.affinitesToAPI(),
	}

	constraints, err := m.constraintsToAPI(ctx)
	if err != nil {
		return req, err
	}
	req.Constraints = constraints

	return req, nil
}

func (m *ContainerApplicationModel) toRolloutRequest() apiCreateRolloutRequest {
	return apiCreateRolloutRequest{
		Description:         "Progressive update",
		Strategy:            "rolling",
		TargetConfiguration: m.buildConfiguration(),
		StepPercentage:      m.RolloutStepPercentage.ValueInt64(),
		Kind:                m.RolloutKind.ValueString(),
	}
}

func (m *ContainerApplicationModel) constraintsToAPI(ctx context.Context) (*apiConstraints, error) {
	if m.Constraints == nil {
		return nil, nil
	}

	constraints := &apiConstraints{}

	if !m.Constraints.Tiers.IsNull() && !m.Constraints.Tiers.IsUnknown() {
		var tiers []int64
		diags := m.Constraints.Tiers.ElementsAs(ctx, &tiers, false)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert tiers: %s", diags.Errors())
		}
		constraints.Tiers = tiers
	}

	if !m.Constraints.Regions.IsNull() && !m.Constraints.Regions.IsUnknown() {
		var regions []string
		diags := m.Constraints.Regions.ElementsAs(ctx, &regions, false)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert regions: %s", diags.Errors())
		}
		constraints.Regions = regions
	}

	if !m.Constraints.Cities.IsNull() && !m.Constraints.Cities.IsUnknown() {
		var cities []string
		diags := m.Constraints.Cities.ElementsAs(ctx, &cities, false)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert cities: %s", diags.Errors())
		}
		constraints.Cities = cities
	}

	return constraints, nil
}

func (m *ContainerApplicationModel) fromAPIApplication(app apiApplication) {
	m.ApplicationID = types.StringValue(app.ID)
	m.ID = types.StringValue(app.ID)
	m.Name = types.StringValue(app.Name)
	m.Image = types.StringValue(app.Configuration.Image)
	m.MaxInstances = types.Int64Value(app.MaxInstances)
	m.SchedulingPolicy = types.StringValue(app.SchedulingPolicy)
	m.RolloutActiveGracePeriod = types.Int64Value(app.RolloutActiveGracePeriod)

	if app.Configuration.InstanceType != "" {
		m.InstanceType = types.StringValue(app.Configuration.InstanceType)
	}
	if app.Configuration.Vcpu != nil {
		m.Vcpu = types.Float64Value(*app.Configuration.Vcpu)
	}
	if app.Configuration.MemoryMib != nil {
		m.MemoryMib = types.Int64Value(*app.Configuration.MemoryMib)
	}
	if app.Configuration.Disk != nil {
		m.DiskSizeMb = types.Int64Value(app.Configuration.Disk.SizeMb)
	}

	// Affinities
	if app.Affinities != nil {
		m.Affinities = &ContainerApplicationAffinitiesModel{}
		if app.Affinities.Colocation != "" {
			m.Affinities.Colocation = types.StringValue(app.Affinities.Colocation)
		}
		if app.Affinities.HardwareGeneration != "" {
			m.Affinities.HardwareGeneration = types.StringValue(app.Affinities.HardwareGeneration)
		}
	}

	// Observability
	if app.Configuration.Observability != nil && app.Configuration.Observability.Logs != nil {
		m.Observability = &ContainerApplicationObservabilityModel{
			LogsEnabled: types.BoolValue(app.Configuration.Observability.Logs.Enabled),
		}
	}

	// SSH config
	if app.Configuration.WranglerSSH != nil {
		m.WranglerSSH = &ContainerApplicationWranglerSSHModel{
			Enabled: types.BoolValue(app.Configuration.WranglerSSH.Enabled),
		}
		if app.Configuration.WranglerSSH.Port != 0 {
			m.WranglerSSH.Port = types.Int64Value(app.Configuration.WranglerSSH.Port)
		}
	}

	// Authorized keys
	if len(app.Configuration.AuthorizedKeys) > 0 {
		keys := make([]*ContainerApplicationAuthorizedKeyModel, len(app.Configuration.AuthorizedKeys))
		for i, k := range app.Configuration.AuthorizedKeys {
			keys[i] = &ContainerApplicationAuthorizedKeyModel{
				Name:      types.StringValue(k.Name),
				PublicKey: types.StringValue(k.PublicKey),
			}
		}
		m.AuthorizedKeys = &keys
	}

	// Trusted CA keys
	if len(app.Configuration.TrustedUserCAKeys) > 0 {
		keys := make([]*ContainerApplicationTrustedCAKeyModel, len(app.Configuration.TrustedUserCAKeys))
		for i, k := range app.Configuration.TrustedUserCAKeys {
			keys[i] = &ContainerApplicationTrustedCAKeyModel{
				Name:      types.StringValue(k.Name),
				PublicKey: types.StringValue(k.PublicKey),
			}
		}
		m.TrustedUserCAKeys = &keys
	}
}
