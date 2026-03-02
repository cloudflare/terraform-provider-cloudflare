package container_application

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContainerApplicationDefaultName(t *testing.T) {
	tests := []struct {
		name       string
		scriptName string
		className  string
		appName    string
		expected   string
	}{
		{
			name:       "default name from script and class",
			scriptName: "my-worker",
			className:  "MyDO",
			appName:    "",
			expected:   "my-worker-MyDO",
		},
		{
			name:       "explicit name overrides default",
			scriptName: "my-worker",
			className:  "MyDO",
			appName:    "custom-app",
			expected:   "custom-app",
		},
		{
			name:       "null name generates default",
			scriptName: "api-worker",
			className:  "Container",
			appName:    "", // will be set as null
			expected:   "api-worker-Container",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := &ContainerApplicationModel{
				ScriptName: types.StringValue(tt.scriptName),
				ClassName:  types.StringValue(tt.className),
			}
			if tt.appName != "" {
				model.Name = types.StringValue(tt.appName)
			} else {
				model.Name = types.StringNull()
			}

			assert.Equal(t, tt.expected, model.effectiveName())
		})
	}
}

func TestContainerApplicationModelToCreateRequest(t *testing.T) {
	ctx := context.Background()

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("test-account"),
		ScriptName:               types.StringValue("my-worker"),
		ClassName:                types.StringValue("MyDO"),
		Name:                     types.StringValue("my-app"),
		Image:                    types.StringValue("registry.example.com/image:v1"),
		MaxInstances:             types.Int64Value(10),
		InstanceType:             types.StringValue("basic"),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(30),
	}

	req, err := model.toCreateRequest(ctx, "ns-12345")
	require.NoError(t, err)

	assert.Equal(t, "my-app", req.Name)
	assert.Equal(t, "default", req.SchedulingPolicy)
	assert.Equal(t, "registry.example.com/image:v1", req.Configuration.Image)
	assert.Equal(t, "basic", req.Configuration.InstanceType)
	assert.Equal(t, int64(0), req.Instances)
	assert.Equal(t, int64(10), req.MaxInstances)
	assert.Equal(t, "ns-12345", req.DurableObjects.NamespaceID)
	assert.Equal(t, int64(30), req.RolloutActiveGracePeriod)
	assert.Nil(t, req.Constraints)

	// Default observability should be logs enabled
	require.NotNil(t, req.Configuration.Observability)
	require.NotNil(t, req.Configuration.Observability.Logs)
	assert.True(t, req.Configuration.Observability.Logs.Enabled)

	// Verify JSON serialization
	jsonBytes, err := json.Marshal(req)
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(jsonBytes, &parsed)
	require.NoError(t, err)

	assert.Equal(t, "my-app", parsed["name"])
	assert.Equal(t, float64(0), parsed["instances"])
	assert.Equal(t, float64(10), parsed["max_instances"])

	config := parsed["configuration"].(map[string]interface{})
	assert.Equal(t, "registry.example.com/image:v1", config["image"])
	assert.Equal(t, "basic", config["instance_type"])

	do := parsed["durable_objects"].(map[string]interface{})
	assert.Equal(t, "ns-12345", do["namespace_id"])
}

func TestContainerApplicationModelToCreateRequestWithConstraints(t *testing.T) {
	ctx := context.Background()

	tiers, _ := types.ListValueFrom(ctx, types.Int64Type, []int64{1, 2})
	regions, _ := types.ListValueFrom(ctx, types.StringType, []string{"us-east", "eu-west"})
	cities, _ := types.ListValueFrom(ctx, types.StringType, []string{})

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("test-account"),
		ScriptName:               types.StringValue("my-worker"),
		ClassName:                types.StringValue("MyDO"),
		Name:                     types.StringValue("my-app"),
		Image:                    types.StringValue("image:v1"),
		MaxInstances:             types.Int64Value(20),
		InstanceType:             types.StringValue("lite"),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(0),
		Constraints: &ContainerApplicationConstraintsModel{
			Tiers:   tiers,
			Regions: regions,
			Cities:  cities,
		},
	}

	req, err := model.toCreateRequest(ctx, "ns-abc")
	require.NoError(t, err)

	require.NotNil(t, req.Constraints)
	assert.Equal(t, []int64{1, 2}, req.Constraints.Tiers)
	assert.Equal(t, []string{"us-east", "eu-west"}, req.Constraints.Regions)
	assert.Equal(t, []string{}, req.Constraints.Cities)
}

func TestContainerApplicationModelToModifyRequest(t *testing.T) {
	ctx := context.Background()

	model := &ContainerApplicationModel{
		Image:                    types.StringValue("registry.example.com/image:v2"),
		InstanceType:             types.StringValue("standard-1"),
		MaxInstances:             types.Int64Value(50),
		SchedulingPolicy:         types.StringValue("regional"),
		RolloutActiveGracePeriod: types.Int64Value(60),
	}

	req, err := model.toModifyRequest(ctx)
	require.NoError(t, err)

	require.NotNil(t, req.Configuration)
	assert.Equal(t, "registry.example.com/image:v2", req.Configuration.Image)
	assert.Equal(t, "standard-1", req.Configuration.InstanceType)
	require.NotNil(t, req.MaxInstances)
	assert.Equal(t, int64(50), *req.MaxInstances)
	require.NotNil(t, req.SchedulingPolicy)
	assert.Equal(t, "regional", *req.SchedulingPolicy)
	require.NotNil(t, req.RolloutActiveGracePeriod)
	assert.Equal(t, int64(60), *req.RolloutActiveGracePeriod)
}

func TestContainerApplicationRolloutRequest(t *testing.T) {
	tests := []struct {
		name               string
		image              string
		instanceType       string
		stepPercentage     int64
		rolloutKind        string
		expectedStrategy   string
		expectedPercentage int64
	}{
		{
			name:               "full auto with 100%",
			image:              "image:v1",
			instanceType:       "lite",
			stepPercentage:     100,
			rolloutKind:        "full_auto",
			expectedStrategy:   "rolling",
			expectedPercentage: 100,
		},
		{
			name:               "full manual with 10%",
			image:              "image:v2",
			instanceType:       "basic",
			stepPercentage:     10,
			rolloutKind:        "full_manual",
			expectedStrategy:   "rolling",
			expectedPercentage: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := &ContainerApplicationModel{
				Image:                 types.StringValue(tt.image),
				InstanceType:         types.StringValue(tt.instanceType),
				RolloutStepPercentage: types.Int64Value(tt.stepPercentage),
				RolloutKind:          types.StringValue(tt.rolloutKind),
			}

			req := model.toRolloutRequest()

			assert.Equal(t, "Progressive update", req.Description)
			assert.Equal(t, tt.expectedStrategy, req.Strategy)
			assert.Equal(t, tt.image, req.TargetConfiguration.Image)
			assert.Equal(t, tt.instanceType, req.TargetConfiguration.InstanceType)
			assert.Equal(t, tt.expectedPercentage, req.StepPercentage)
			assert.Equal(t, tt.rolloutKind, req.Kind)
		})
	}
}

func TestContainerApplicationFromAPIApplication(t *testing.T) {
	app := apiApplication{
		ID:               "app-123",
		Name:             "my-app",
		MaxInstances:     20,
		SchedulingPolicy: "default",
		Configuration: apiConfiguration{
			Image:        "registry.example.com/image:v1",
			InstanceType: "lite",
		},
		RolloutActiveGracePeriod: 30,
	}

	model := &ContainerApplicationModel{}
	model.fromAPIApplication(app)

	assert.Equal(t, "app-123", model.ID.ValueString())
	assert.Equal(t, "app-123", model.ApplicationID.ValueString())
	assert.Equal(t, "my-app", model.Name.ValueString())
	assert.Equal(t, "registry.example.com/image:v1", model.Image.ValueString())
	assert.Equal(t, int64(20), model.MaxInstances.ValueInt64())
	assert.Equal(t, "default", model.SchedulingPolicy.ValueString())
	assert.Equal(t, "lite", model.InstanceType.ValueString())
	assert.Equal(t, int64(30), model.RolloutActiveGracePeriod.ValueInt64())
}

func TestContainerApplicationSchemaValidity(t *testing.T) {
	ctx := context.Background()
	s := ResourceSchema(ctx)

	// Verify all expected attributes exist
	expectedAttrs := []string{
		"id", "account_id", "application_id", "script_name", "class_name",
		"name", "image", "max_instances", "instance_type", "scheduling_policy",
		"rollout_step_percentage", "rollout_kind", "rollout_active_grace_period",
		"constraints", "affinities", "observability", "vcpu", "memory_mib",
		"disk_size_mb", "wrangler_ssh", "authorized_keys", "trusted_user_ca_keys",
	}

	for _, attr := range expectedAttrs {
		_, ok := s.Attributes[attr]
		assert.True(t, ok, "schema should have attribute %q", attr)
	}

	// Verify required attributes
	requiredAttrs := []string{"account_id", "script_name", "class_name", "image"}
	for _, attr := range requiredAttrs {
		a := s.Attributes[attr]
		assert.True(t, a.IsRequired(), "attribute %q should be required", attr)
	}

	// Verify computed attributes
	computedAttrs := []string{"id", "application_id"}
	for _, attr := range computedAttrs {
		a := s.Attributes[attr]
		assert.True(t, a.IsComputed(), "attribute %q should be computed", attr)
		assert.False(t, a.IsRequired(), "attribute %q should not be required", attr)
	}

	// Verify optional attributes
	optionalAttrs := []string{
		"vcpu", "memory_mib", "disk_size_mb",
		"affinities", "observability", "wrangler_ssh",
		"authorized_keys", "trusted_user_ca_keys",
	}
	for _, attr := range optionalAttrs {
		a := s.Attributes[attr]
		assert.True(t, a.IsOptional(), "attribute %q should be optional", attr)
	}
}

func TestCreateRequestJSONShape(t *testing.T) {
	ctx := context.Background()

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("acc-1"),
		ScriptName:               types.StringValue("worker"),
		ClassName:                types.StringValue("DO"),
		Name:                     types.StringNull(),
		Image:                    types.StringValue("ccr.ccs.cloudflare.com/acc-1/myimage:latest"),
		MaxInstances:             types.Int64Value(20),
		InstanceType:             types.StringValue("lite"),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(0),
	}

	req, err := model.toCreateRequest(ctx, "ns-do-123")
	require.NoError(t, err)

	// Name should be computed from script+class
	assert.Equal(t, "worker-DO", req.Name)

	jsonBytes, err := json.Marshal(req)
	require.NoError(t, err)

	// Verify the JSON matches the expected API shape
	var raw map[string]interface{}
	err = json.Unmarshal(jsonBytes, &raw)
	require.NoError(t, err)

	// Must have these top-level keys
	for _, key := range []string{"name", "scheduling_policy", "configuration", "instances", "max_instances", "durable_objects", "rollout_active_grace_period"} {
		_, ok := raw[key]
		assert.True(t, ok, "JSON should have key %q", key)
	}

	// instances must be 0 (deprecated field)
	assert.Equal(t, float64(0), raw["instances"])

	// durable_objects.namespace_id must be set
	do := raw["durable_objects"].(map[string]interface{})
	assert.Equal(t, "ns-do-123", do["namespace_id"])
}

func TestContainerApplicationCustomInstanceType(t *testing.T) {
	ctx := context.Background()

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("test-account"),
		ScriptName:               types.StringValue("my-worker"),
		ClassName:                types.StringValue("MyDO"),
		Name:                     types.StringValue("my-app"),
		Image:                    types.StringValue("image:v1"),
		MaxInstances:             types.Int64Value(5),
		Vcpu:                     types.Float64Value(2.0),
		MemoryMib:                types.Int64Value(512),
		DiskSizeMb:               types.Int64Value(1024),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(0),
	}

	req, err := model.toCreateRequest(ctx, "ns-123")
	require.NoError(t, err)

	// Custom resources
	require.NotNil(t, req.Configuration.Vcpu)
	assert.Equal(t, 2.0, *req.Configuration.Vcpu)
	require.NotNil(t, req.Configuration.MemoryMib)
	assert.Equal(t, int64(512), *req.Configuration.MemoryMib)
	require.NotNil(t, req.Configuration.Disk)
	assert.Equal(t, int64(1024), req.Configuration.Disk.SizeMb)
	// InstanceType should be empty when custom resources are used
	assert.Equal(t, "", req.Configuration.InstanceType)

	// Verify JSON
	jsonBytes, err := json.Marshal(req)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(jsonBytes, &raw))
	config := raw["configuration"].(map[string]interface{})
	assert.Equal(t, 2.0, config["vcpu"])
	assert.Equal(t, float64(512), config["memory_mib"])
	disk := config["disk"].(map[string]interface{})
	assert.Equal(t, float64(1024), disk["size_mb"])
}

func TestContainerApplicationObservability(t *testing.T) {
	ctx := context.Background()

	t.Run("explicit observability disabled", func(t *testing.T) {
		model := &ContainerApplicationModel{
			AccountID:                types.StringValue("test-account"),
			ScriptName:               types.StringValue("worker"),
			ClassName:                types.StringValue("DO"),
			Name:                     types.StringValue("app"),
			Image:                    types.StringValue("image:v1"),
			MaxInstances:             types.Int64Value(1),
			InstanceType:             types.StringValue("lite"),
			SchedulingPolicy:         types.StringValue("default"),
			RolloutActiveGracePeriod: types.Int64Value(0),
			Observability: &ContainerApplicationObservabilityModel{
				LogsEnabled: types.BoolValue(false),
			},
		}

		req, err := model.toCreateRequest(ctx, "ns-1")
		require.NoError(t, err)

		require.NotNil(t, req.Configuration.Observability)
		require.NotNil(t, req.Configuration.Observability.Logs)
		assert.False(t, req.Configuration.Observability.Logs.Enabled)
	})

	t.Run("default observability when not specified", func(t *testing.T) {
		model := &ContainerApplicationModel{
			AccountID:                types.StringValue("test-account"),
			ScriptName:               types.StringValue("worker"),
			ClassName:                types.StringValue("DO"),
			Name:                     types.StringValue("app"),
			Image:                    types.StringValue("image:v1"),
			MaxInstances:             types.Int64Value(1),
			InstanceType:             types.StringValue("lite"),
			SchedulingPolicy:         types.StringValue("default"),
			RolloutActiveGracePeriod: types.Int64Value(0),
		}

		req, err := model.toCreateRequest(ctx, "ns-1")
		require.NoError(t, err)

		require.NotNil(t, req.Configuration.Observability)
		require.NotNil(t, req.Configuration.Observability.Logs)
		assert.True(t, req.Configuration.Observability.Logs.Enabled)
	})
}

func TestContainerApplicationAffinities(t *testing.T) {
	ctx := context.Background()

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("test-account"),
		ScriptName:               types.StringValue("worker"),
		ClassName:                types.StringValue("DO"),
		Name:                     types.StringValue("app"),
		Image:                    types.StringValue("image:v1"),
		MaxInstances:             types.Int64Value(5),
		InstanceType:             types.StringValue("lite"),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(0),
		Affinities: &ContainerApplicationAffinitiesModel{
			Colocation:         types.StringValue("best_effort"),
			HardwareGeneration: types.StringValue("metal"),
		},
	}

	req, err := model.toCreateRequest(ctx, "ns-1")
	require.NoError(t, err)

	require.NotNil(t, req.Affinities)
	assert.Equal(t, "best_effort", req.Affinities.Colocation)
	assert.Equal(t, "metal", req.Affinities.HardwareGeneration)

	// Verify JSON
	jsonBytes, err := json.Marshal(req)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(jsonBytes, &raw))
	affinities := raw["affinities"].(map[string]interface{})
	assert.Equal(t, "best_effort", affinities["colocation"])
	assert.Equal(t, "metal", affinities["hardware_generation"])
}

func TestContainerApplicationSSHConfig(t *testing.T) {
	ctx := context.Background()

	authKeys := []*ContainerApplicationAuthorizedKeyModel{
		{Name: types.StringValue("my-key"), PublicKey: types.StringValue("ssh-ed25519 AAAA...")},
	}
	caKeys := []*ContainerApplicationTrustedCAKeyModel{
		{Name: types.StringValue("ca-key"), PublicKey: types.StringValue("ssh-rsa BBBB...")},
	}

	model := &ContainerApplicationModel{
		AccountID:                types.StringValue("test-account"),
		ScriptName:               types.StringValue("worker"),
		ClassName:                types.StringValue("DO"),
		Name:                     types.StringValue("app"),
		Image:                    types.StringValue("image:v1"),
		MaxInstances:             types.Int64Value(1),
		InstanceType:             types.StringValue("lite"),
		SchedulingPolicy:         types.StringValue("default"),
		RolloutActiveGracePeriod: types.Int64Value(0),
		WranglerSSH: &ContainerApplicationWranglerSSHModel{
			Enabled: types.BoolValue(true),
			Port:    types.Int64Value(2222),
		},
		AuthorizedKeys:    &authKeys,
		TrustedUserCAKeys: &caKeys,
	}

	req, err := model.toCreateRequest(ctx, "ns-1")
	require.NoError(t, err)

	// SSH
	require.NotNil(t, req.Configuration.WranglerSSH)
	assert.True(t, req.Configuration.WranglerSSH.Enabled)
	assert.Equal(t, int64(2222), req.Configuration.WranglerSSH.Port)

	// Authorized keys
	require.Len(t, req.Configuration.AuthorizedKeys, 1)
	assert.Equal(t, "my-key", req.Configuration.AuthorizedKeys[0].Name)
	assert.Equal(t, "ssh-ed25519 AAAA...", req.Configuration.AuthorizedKeys[0].PublicKey)

	// Trusted CA keys
	require.Len(t, req.Configuration.TrustedUserCAKeys, 1)
	assert.Equal(t, "ca-key", req.Configuration.TrustedUserCAKeys[0].Name)
	assert.Equal(t, "ssh-rsa BBBB...", req.Configuration.TrustedUserCAKeys[0].PublicKey)

	// Verify JSON
	jsonBytes, err := json.Marshal(req)
	require.NoError(t, err)

	var raw map[string]interface{}
	require.NoError(t, json.Unmarshal(jsonBytes, &raw))
	config := raw["configuration"].(map[string]interface{})

	ssh := config["wrangler_ssh"].(map[string]interface{})
	assert.Equal(t, true, ssh["enabled"])
	assert.Equal(t, float64(2222), ssh["port"])

	keys := config["authorized_keys"].([]interface{})
	require.Len(t, keys, 1)
	key := keys[0].(map[string]interface{})
	assert.Equal(t, "my-key", key["name"])

	caKeysJSON := config["trusted_user_ca_keys"].([]interface{})
	require.Len(t, caKeysJSON, 1)
}

func TestContainerApplicationFromAPIApplicationFull(t *testing.T) {
	vcpu := 1.5
	memMib := int64(256)

	app := apiApplication{
		ID:               "app-full",
		Name:             "full-app",
		MaxInstances:     10,
		SchedulingPolicy: "regional",
		Configuration: apiConfiguration{
			Image:     "image:v2",
			Vcpu:      &vcpu,
			MemoryMib: &memMib,
			Disk:      &apiDisk{SizeMb: 2048},
			Observability: &apiObservability{
				Logs: &apiObservabilityLogs{Enabled: false},
			},
			WranglerSSH: &apiWranglerSSH{
				Enabled: true,
				Port:    2222,
			},
			AuthorizedKeys: []apiAuthorizedKey{
				{Name: "key1", PublicKey: "ssh-ed25519 AAA"},
			},
			TrustedUserCAKeys: []apiTrustedCAKey{
				{Name: "ca1", PublicKey: "ssh-rsa BBB"},
			},
		},
		Affinities: &apiAffinities{
			Colocation:         "best_effort",
			HardwareGeneration: "intel",
		},
		RolloutActiveGracePeriod: 60,
	}

	model := &ContainerApplicationModel{}
	model.fromAPIApplication(app)

	assert.Equal(t, "app-full", model.ID.ValueString())
	assert.Equal(t, "full-app", model.Name.ValueString())
	assert.Equal(t, "image:v2", model.Image.ValueString())
	assert.Equal(t, int64(10), model.MaxInstances.ValueInt64())
	assert.Equal(t, "regional", model.SchedulingPolicy.ValueString())
	assert.Equal(t, int64(60), model.RolloutActiveGracePeriod.ValueInt64())

	// Custom resources
	assert.Equal(t, 1.5, model.Vcpu.ValueFloat64())
	assert.Equal(t, int64(256), model.MemoryMib.ValueInt64())
	assert.Equal(t, int64(2048), model.DiskSizeMb.ValueInt64())

	// Affinities
	require.NotNil(t, model.Affinities)
	assert.Equal(t, "best_effort", model.Affinities.Colocation.ValueString())
	assert.Equal(t, "intel", model.Affinities.HardwareGeneration.ValueString())

	// Observability
	require.NotNil(t, model.Observability)
	assert.False(t, model.Observability.LogsEnabled.ValueBool())

	// SSH
	require.NotNil(t, model.WranglerSSH)
	assert.True(t, model.WranglerSSH.Enabled.ValueBool())
	assert.Equal(t, int64(2222), model.WranglerSSH.Port.ValueInt64())

	// Keys
	require.NotNil(t, model.AuthorizedKeys)
	require.Len(t, *model.AuthorizedKeys, 1)
	assert.Equal(t, "key1", (*model.AuthorizedKeys)[0].Name.ValueString())

	require.NotNil(t, model.TrustedUserCAKeys)
	require.Len(t, *model.TrustedUserCAKeys, 1)
	assert.Equal(t, "ca1", (*model.TrustedUserCAKeys)[0].Name.ValueString())
}

func TestContainerApplicationModifyWithAllFields(t *testing.T) {
	ctx := context.Background()

	authKeys := []*ContainerApplicationAuthorizedKeyModel{
		{Name: types.StringValue("k1"), PublicKey: types.StringValue("ssh-ed25519 AAA")},
		{Name: types.StringValue("k2"), PublicKey: types.StringValue("ssh-ed25519 BBB")},
	}

	model := &ContainerApplicationModel{
		Image:                    types.StringValue("image:v3"),
		Vcpu:                     types.Float64Value(4.0),
		MemoryMib:                types.Int64Value(1024),
		DiskSizeMb:               types.Int64Value(4096),
		MaxInstances:             types.Int64Value(100),
		SchedulingPolicy:         types.StringValue("moon"),
		RolloutActiveGracePeriod: types.Int64Value(120),
		Affinities: &ContainerApplicationAffinitiesModel{
			Colocation: types.StringValue("best_effort"),
		},
		Observability: &ContainerApplicationObservabilityModel{
			LogsEnabled: types.BoolValue(true),
		},
		WranglerSSH: &ContainerApplicationWranglerSSHModel{
			Enabled: types.BoolValue(true),
			Port:    types.Int64Value(3333),
		},
		AuthorizedKeys: &authKeys,
	}

	req, err := model.toModifyRequest(ctx)
	require.NoError(t, err)

	require.NotNil(t, req.Configuration)
	assert.Equal(t, "image:v3", req.Configuration.Image)
	require.NotNil(t, req.Configuration.Vcpu)
	assert.Equal(t, 4.0, *req.Configuration.Vcpu)
	require.NotNil(t, req.Configuration.MemoryMib)
	assert.Equal(t, int64(1024), *req.Configuration.MemoryMib)
	require.NotNil(t, req.Configuration.Disk)
	assert.Equal(t, int64(4096), req.Configuration.Disk.SizeMb)
	require.NotNil(t, req.Configuration.WranglerSSH)
	assert.True(t, req.Configuration.WranglerSSH.Enabled)
	assert.Equal(t, int64(3333), req.Configuration.WranglerSSH.Port)
	require.Len(t, req.Configuration.AuthorizedKeys, 2)
	require.NotNil(t, req.Affinities)
	assert.Equal(t, "best_effort", req.Affinities.Colocation)
	assert.Equal(t, int64(100), *req.MaxInstances)
	assert.Equal(t, "moon", *req.SchedulingPolicy)
}
