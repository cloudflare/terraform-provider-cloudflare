package cloudflare

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	errors "github.com/pkg/errors"
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

// Returns true if string value exists in string slice
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
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

type CloudflareAPIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CloudflareAPIErrorResponse struct {
	Errors []CloudflareAPIError `json:"errors"`
}

func cloudflareErrorIsOneOfCodes(err error, codes []int) bool {
	errorMsg := errors.Cause(err).Error()

	// We will parse the error message only if it's an error 400, in which
	// case we need to verify the kind of error.
	r := regexp.MustCompile(`^HTTP status 400: content "(.*)"$`)
	submatchs := r.FindStringSubmatch(errorMsg)
	if submatchs != nil {
		jsonData := strings.Replace(submatchs[1], "\\\"", "\"", -1)
		log.Printf("[DEBUG][cloudflareErrorIsCode] error matching status 400, content: %#v", jsonData)

		var cfer CloudflareAPIErrorResponse
		unmarshalErr := json.Unmarshal([]byte(jsonData), &cfer)

		// We check that there is only one error and that its code
		// matches what we expected
		if unmarshalErr == nil && len(cfer.Errors) == 1 {
			for _, code := range codes {
				if cfer.Errors[0].Code == code {
					return true
				}
			}
		}
	}

	return false
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

func getAccountIDFromZoneID(d *schema.ResourceData, client *cloudflare.API) (string, error) {
	accountID := d.Get("account_id").(string)
	if accountID == "" {
		zoneID := d.Get("zone_id").(string)
		zone, err := client.ZoneDetails(zoneID)
		if err != nil {
			return "", fmt.Errorf("error retrieving zone for zone_id %q: %s", zoneID, err)
		}
		accountID = zone.Account.ID
	}

	d.Set("account_id", accountID)
	return accountID, nil
}
