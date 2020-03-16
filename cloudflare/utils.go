package cloudflare

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

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
