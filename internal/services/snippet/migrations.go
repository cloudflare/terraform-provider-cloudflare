// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*SnippetResource)(nil)

func (r *SnippetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			// Use a minimal PriorSchema since we'll handle everything via RawState
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Optional: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Use RawState.JSON to handle all transformations since state could be in v4 or v5 format
				rawStateJSON := req.RawState.JSON
				fmt.Printf("DEBUG: RawState.JSON: %s\n", string(rawStateJSON))

				var rawState map[string]interface{}
				if err := json.Unmarshal(rawStateJSON, &rawState); err != nil {
					resp.Diagnostics.AddError("Failed to parse state JSON", err.Error())
					return
				}

				// Debug logging
				fmt.Printf("DEBUG: Raw state keys: %v\n", len(rawState))
				for k, v := range rawState {
					fmt.Printf("DEBUG: State key %s: type=%T value=%v\n", k, v, v)
				}

				// Initialize new state
				newState := SnippetModel{}

				// Initialize files as empty slice to avoid nil
				emptyFiles := make([]SnippetFile, 0)
				newState.Files = &emptyFiles

				// Handle zone_id (present in both v4 and v5)
				if zoneID, ok := rawState["zone_id"].(string); ok {
					newState.ZoneID = types.StringValue(zoneID)
				}

				// Handle name transformation (v4: name, v5: snippet_name)
				if snippetName, ok := rawState["snippet_name"].(string); ok {
					// v5 format
					newState.SnippetName = types.StringValue(snippetName)
				} else if name, ok := rawState["name"].(string); ok {
					// v4 format
					newState.SnippetName = types.StringValue(name)
				}

				// Handle metadata transformation
				if metadataRaw, ok := rawState["metadata"].(map[string]interface{}); ok {
					// v5 format - metadata is an object
					if mainModule, ok := metadataRaw["main_module"].(string); ok {
						newState.Metadata = &SnippetMetadataModel{
							MainModule: types.StringValue(mainModule),
						}
					}
				} else if mainModule, ok := rawState["main_module"].(string); ok {
					// v4 format - main_module is top-level
					newState.Metadata = &SnippetMetadataModel{
						MainModule: types.StringValue(mainModule),
					}
				}

				// Handle files transformation
				if filesRaw, exists := rawState["files"]; exists && filesRaw != nil {
					switch v := filesRaw.(type) {
					case []interface{}:
						// Already in v5 array format (from cmd/migrate or native v5)
						files := make([]SnippetFile, len(v))
						for i, fileRaw := range v {
							if fileMap, ok := fileRaw.(map[string]interface{}); ok {
								name := ""
								content := ""
								if nameVal, ok := fileMap["name"].(string); ok {
									name = nameVal
								}
								if contentVal, ok := fileMap["content"].(string); ok {
									content = contentVal
								}
								files[i] = NewSnippetsFileValueMust(name, content)
								fmt.Printf("DEBUG: Created file %d: name=%s, content=%d chars\n", i, name, len(content))
							}
						}
						newState.Files = &files
						fmt.Printf("DEBUG: Set newState.Files with %d files\n", len(files))
					}
				} else {
					// Check for v4 indexed format (files.#, files.0.name, etc.)
					// This path shouldn't be hit if cmd/migrate already transformed it
					if filesCountRaw, exists := rawState["files.#"]; exists {
						var filesCount int
						switch v := filesCountRaw.(type) {
						case string:
							fmt.Sscanf(v, "%d", &filesCount)
						case float64:
							filesCount = int(v)
						}

						if filesCount > 0 {
							files := make([]SnippetFile, filesCount)
							for i := 0; i < filesCount; i++ {
								name := ""
								content := ""

								// Try to get name from files.X.name
								nameKey := fmt.Sprintf("files.%d.name", i)
								if nameVal, ok := rawState[nameKey].(string); ok {
									name = nameVal
								}

								// Try to get content from files.X.content
								contentKey := fmt.Sprintf("files.%d.content", i)
								if contentVal, ok := rawState[contentKey].(string); ok {
									content = contentVal
								}

								files[i] = NewSnippetsFileValueMust(name, content)
							}
							newState.Files = &files
						}
					}
				}

				// Handle timestamps (computed fields)
				if createdOn, ok := rawState["created_on"].(string); ok && createdOn != "" {
					newState.CreatedOn = timetypes.NewRFC3339ValueMust(createdOn)
				}
				if modifiedOn, ok := rawState["modified_on"].(string); ok && modifiedOn != "" {
					newState.ModifiedOn = timetypes.NewRFC3339ValueMust(modifiedOn)
				}

				// Debug: print new state
				if newState.Files != nil {
					fmt.Printf("DEBUG: New state values: ZoneID=%v, SnippetName=%v, Files=%d\n",
						newState.ZoneID.ValueString(),
						newState.SnippetName.ValueString(),
						len(*newState.Files))
					for i, file := range *newState.Files {
						fmt.Printf("DEBUG: File %d in newState: %+v\n", i, file)
					}
				} else {
					fmt.Printf("DEBUG: New state values: ZoneID=%v, SnippetName=%v, Files=nil\n",
						newState.ZoneID.ValueString(),
						newState.SnippetName.ValueString())
				}

				// Set the upgraded state
				resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
				if resp.Diagnostics.HasError() {
					fmt.Printf("DEBUG: Error setting state: %v\n", resp.Diagnostics)
				} else {
					fmt.Printf("DEBUG: Successfully set upgraded state\n")
				}
			},
		},
	}
}

