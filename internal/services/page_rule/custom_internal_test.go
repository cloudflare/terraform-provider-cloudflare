package page_rule

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func strs(vs ...string) []basetypes.StringValue {
	if vs == nil {
		return nil
	}
	out := make([]basetypes.StringValue, 0, len(vs))
	for _, v := range vs {
		out = append(out, types.StringValue(v))
	}
	return out
}

// The API models query_string.include/exclude as a oneOf: the bare string "*"
// means all parameters, an array is a list of literal parameter names. Sending
// ["*"] verbatim lands in the array branch and matches nothing.
func TestEncodeQueryStringField(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []basetypes.StringValue
		want any
	}{
		{"wildcard becomes bare string", strs("*"), "*"},
		{"named parameters stay a list", strs("foo", "bar"), []string{"foo", "bar"}},
		{"empty stays an empty list", strs(), []string{}},
		{"nil stays an empty list", nil, []string{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := encodeQueryStringField(tc.in)
			gotJSON, _ := json.Marshal(got)
			wantJSON, _ := json.Marshal(tc.want)
			if string(gotJSON) != string(wantJSON) {
				t.Errorf("encodeQueryStringField(%v) = %s, want %s", tc.in, gotJSON, wantJSON)
			}
		})
	}
}

// Regression: the string form previously failed a []interface{} type assertion
// and was discarded without error, so a wildcard set outside Terraform read back
// as empty and the next apply overwrote it.
func TestDecodeQueryStringField(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   any
		want []basetypes.StringValue
	}{
		{"bare string wildcard becomes a list", "*", strs("*")},
		{"array wildcard is preserved", []interface{}{"*"}, strs("*")},
		{"named parameters are preserved", []interface{}{"foo", "bar"}, strs("foo", "bar")},
		{"empty array decodes to nil", []interface{}{}, nil},
		{"unexpected string decodes to nil", "unexpected", nil},
		{"unexpected type decodes to nil", 42, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := decodeQueryStringField(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("decodeQueryStringField(%v) = %v, want %v", tc.in, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("element %d = %v, want %v", i, got[i], tc.want[i])
				}
			}
		})
	}
}

// "*" is a whole-value sentinel, not a list element, so mixing it with named
// parameters has no meaning on the API. v4 rejected it outright.
func TestValidateQueryStringField(t *testing.T) {
	for _, tc := range []struct {
		name    string
		in      []basetypes.StringValue
		wantErr bool
	}{
		{"wildcard alone is valid", strs("*"), false},
		{"named parameters are valid", strs("foo", "bar"), false},
		{"empty is valid", nil, false},
		{"wildcard first is rejected", strs("*", "foo"), true},
		{"wildcard last is rejected", strs("foo", "*"), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := validateQueryStringField("include", tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("validateQueryStringField(%v) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			}
		})
	}
}

// End-to-end import path (UnmarshalPageRuleModel is reached only from
// ImportState) against payloads the live API actually returns, verified by
// creating page rules with each form and reading them back.
func TestUnmarshalPageRuleQueryStringWildcard(t *testing.T) {
	payload := func(qs string) []byte {
		return []byte(fmt.Sprintf(`{"result":{
			"id":"abc","zone_id":"z","priority":1,"status":"disabled",
			"created_on":"2026-01-01T00:00:00Z","modified_on":"2026-01-01T00:00:00Z",
			"targets":[{"target":"url","constraint":{"operator":"matches","value":"x.example.com/*"}}],
			"actions":[{"id":"cache_key_fields","value":{
				"host":{"resolved":false},
				"user":{"geo":true,"device_type":true,"lang":false},
				"query_string":%s}}]}}`, qs))
	}

	for _, tc := range []struct {
		name                     string
		apiQueryString           string
		wantInclude, wantExclude []basetypes.StringValue
	}{
		{
			name:           "include string wildcard",
			apiQueryString: `{"include":"*","exclude":[]}`,
			wantInclude:    strs("*"),
			wantExclude:    nil,
		},
		{
			name:           "include array",
			apiQueryString: `{"include":["foo"],"exclude":[]}`,
			wantInclude:    strs("foo"),
			wantExclude:    nil,
		},
		{
			name:           "exclude string wildcard",
			apiQueryString: `{"include":[],"exclude":"*"}`,
			wantInclude:    nil,
			wantExclude:    strs("*"),
		},
		{
			name:           "exclude array",
			apiQueryString: `{"include":[],"exclude":["foo"]}`,
			wantInclude:    nil,
			wantExclude:    strs("foo"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m, err := UnmarshalPageRuleModel(payload(tc.apiQueryString))
			if err != nil {
				t.Fatalf("UnmarshalPageRuleModel: %v", err)
			}

			var ckf PageRuleActionsCacheKeyFieldsModel
			if diags := m.Actions.CacheKeyFields.As(t.Context(), &ckf, basetypes.ObjectAsOptions{}); diags.HasError() {
				t.Fatalf("decoding cache_key_fields: %v", diags)
			}
			var qs PageRuleActionsCacheKeyFieldsQueryStringModel
			if diags := ckf.QueryString.As(t.Context(), &qs, basetypes.ObjectAsOptions{}); diags.HasError() {
				t.Fatalf("decoding query_string: %v", diags)
			}

			for _, f := range []struct {
				field     string
				got, want []basetypes.StringValue
			}{
				{"include", qs.Include, tc.wantInclude},
				{"exclude", qs.Exclude, tc.wantExclude},
			} {
				if len(f.got) != len(f.want) {
					t.Fatalf("%s = %v, want %v", f.field, f.got, f.want)
				}
				for i := range f.got {
					if f.got[i] != f.want[i] {
						t.Errorf("%s[%d] = %v, want %v", f.field, i, f.got[i], f.want[i])
					}
				}
			}
		})
	}
}
