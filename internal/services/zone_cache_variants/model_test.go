package zone_cache_variants_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_variants"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func variants(values ...string) *[]types.String {
	out := make([]types.String, len(values))
	for i, v := range values {
		out[i] = types.StringValue(v)
	}
	return &out
}

// The variants PATCH endpoint replaces the whole stored value, so an update
// must send every planned key and leave a removed key out instead of sending
// it as null.
func TestMarshalJSONForUpdateSendsFullValue(t *testing.T) {
	t.Parallel()
	state := zone_cache_variants.ZoneCacheVariantsModel{
		ID:     types.StringValue("zone"),
		ZoneID: types.StringValue("zone"),
		Value: &zone_cache_variants.ZoneCacheVariantsValueModel{
			AVIF: variants("image/avif", "image/webp"),
			GIF:  variants("image/gif", "image/webp"),
			WebP: variants("image/webp"),
		},
	}

	tests := []struct {
		name  string
		value *zone_cache_variants.ZoneCacheVariantsValueModel
		want  map[string][]string
	}{
		{
			name: "removed key",
			value: &zone_cache_variants.ZoneCacheVariantsValueModel{
				AVIF: variants("image/avif", "image/webp"),
				WebP: variants("image/webp"),
			},
			want: map[string][]string{
				"avif": {"image/avif", "image/webp"},
				"webp": {"image/webp"},
			},
		},
		{
			name: "one key edited",
			value: &zone_cache_variants.ZoneCacheVariantsValueModel{
				AVIF: variants("image/avif"),
				GIF:  variants("image/gif", "image/webp"),
				WebP: variants("image/webp"),
			},
			want: map[string][]string{
				"avif": {"image/avif"},
				"gif":  {"image/gif", "image/webp"},
				"webp": {"image/webp"},
			},
		},
	}

	// Subtests stay sequential: apijson's cached encoders share mutable state,
	// so concurrent marshals race under -race.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := zone_cache_variants.ZoneCacheVariantsModel{
				ID:     state.ID,
				ZoneID: state.ZoneID,
				Value:  tt.value,
			}
			data, err := plan.MarshalJSONForUpdate(state)
			if err != nil {
				t.Fatalf("MarshalJSONForUpdate: %v", err)
			}

			var body struct {
				Value map[string]json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(data, &body); err != nil {
				t.Fatalf("unmarshal %s: %v", data, err)
			}
			got := make(map[string][]string, len(body.Value))
			for k, raw := range body.Value {
				var v []string
				if err := json.Unmarshal(raw, &v); err != nil || v == nil {
					t.Fatalf("value.%s=%s, want a string array (body %s)", k, raw, data)
				}
				got[k] = v
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("value = %v, want the full planned value %v (body %s)", got, tt.want, data)
			}
		})
	}
}
