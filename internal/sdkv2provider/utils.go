package sdkv2provider

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandInterfaceToStringList(list interface{}) []string {
	ifaceList := list.([]interface{})
	vs := make([]string, 0, len(ifaceList))
	for _, v := range ifaceList {
		vs = append(vs, v.(string))
	}
	return vs
}

func expandStringListToSet(list []string) *schema.Set {
	values := schema.NewSet(schema.HashString, []interface{}{})
	for _, h := range list {
		values.Add(h)
	}
	return values
}

func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}

func flattenIntList(list []int) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}

func IntIdentity(i interface{}) int {
	return i.(int)
}

func HashByMapKey(key string) func(v interface{}) int {
	return func(v interface{}) int {
		m := v.(map[string]interface{})
		return schema.HashString(m[key])
	}
}

// stringChecksum takes a string and returns the checksum of the string.
func stringChecksum(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func stringListChecksum(s []string) string {
	sort.Strings(s)
	return stringChecksum(strings.Join(s, ""))
}

// Returns true if string value exists in string slice.
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func sliceContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func itemExistsInSlice(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		log.Print("[DEBUG] invalid data type detected")
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

// findIndex returns the smallest index i at which x == a[i],
// or (0, false) if there is no such index.
func findIndex(a []interface{}, x interface{}) (int, bool) {
	for i, n := range a {
		if x == n {
			return i, true
		}
	}
	return 0, false
}

func boolFromString(status string) bool {
	if status == "on" {
		return true
	}
	return false
}

func stringFromBool(status bool) string {
	if status {
		return "on"
	}
	return "off"
}

func initIdentifier(d *schema.ResourceData) (*cloudflare.ResourceContainer, error) {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	if accountID == "" && zoneID == "" {
		return nil, fmt.Errorf(`error determining resource: "zone_id" or "account_id" required`)
	}

	if accountID != "" {
		return cloudflare.AccountIdentifier(accountID), nil
	}

	return cloudflare.ZoneIdentifier(zoneID), nil
}

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func hashCodeString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// Strings hashes a list of strings to a unique hashcode.
func hashCodeStrings(strings []string) string {
	var buf bytes.Buffer

	for _, s := range strings {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}

	return fmt.Sprintf("%d", hashCodeString(buf.String()))
}

// GetRawValue receives a path spec (similar to ResourceData.GetOK) and a raw Value, and iterates
// through each key given for the underlying map or slice represented. Calls to this function attempt
// to protect against incorrect or mismatching data, but will return a "dangerous" NilVal value if
// path resolution was impossible.
func getRawValue(key string, value cty.Value) cty.Value {
	parts := strings.SplitN(key, ".", 2)
	index, indexErr := strconv.ParseInt(parts[0], 10, 32)

	if value.IsNull() {
		return value
	} else if indexErr == nil && value.CanIterateElements() && value.LengthInt() > int(index) {
		value = value.AsValueSlice()[index]
	} else if indexErr != nil && (value.Type().IsObjectType() || value.Type().IsMapType()) {
		if v, ok := value.AsValueMap()[parts[0]]; ok {
			value = v
		} else {
			return cty.NilVal
		}
	} else {
		return cty.NilVal
	}

	if len(parts) == 2 && parts[1] != "" {
		return getRawValue(parts[1], value)
	}

	return value
}

// renderAvailableDocumentationValuesStringSlice takes a slice of strings and
// formats it for documentation output use.
//
// Example: ["foo", "bar", "baz"] -> `foo`, `bar`, `baz`.
func renderAvailableDocumentationValuesStringSlice(s []string) string {
	output := ""
	if s != nil && len(s) > 0 {
		values := make([]string, len(s))
		for i, c := range s {
			if c == "" {
				values[i] = "`\"\"`"
			} else {
				values[i] = fmt.Sprintf("`%s`", c)
			}

		}
		output = fmt.Sprintf("Available values: %s", strings.Join(values, ", "))
	}
	return output
}

// renderAvailableDocumentationValuesIntSlice takes a slice of ints and
// formats it for documentation output use.
//
// Example: [1, 2, 3] -> `1`, `2`, `3`.
func renderAvailableDocumentationValuesIntSlice(s []int) string {
	output := ""
	if s != nil && len(s) > 0 {
		values := make([]string, len(s))
		for i, c := range s {
			values[i] = fmt.Sprintf("`%d`", c)
		}
		output = fmt.Sprintf("Available values: %s", strings.Join(values, ", "))
	}
	return output
}
