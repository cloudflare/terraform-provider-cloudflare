package worker_version

import "github.com/cloudflare/cloudflare-go/v7/workers"

// PlacementFromSettingsForTest exports placementFromSettings for use in tests.
var PlacementFromSettingsForTest = placementFromSettings

// ScriptPlacementGetResponse is an alias for the SDK type to avoid import verbosity in tests.
type ScriptPlacementGetResponse = workers.ScriptScriptAndVersionSettingGetResponsePlacement

// ScriptPlacementModeMode is an alias for the SDK mode type.
type ScriptPlacementModeMode = workers.ScriptScriptAndVersionSettingGetResponsePlacementModeMode
