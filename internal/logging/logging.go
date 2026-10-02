package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/cloudflare/cloudflare-go/v7/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context, sensitiveBodyFieldNames ...string) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if req != nil {
			if err := LogRequest(ctx, req, sensitiveBodyFieldNames...); err != nil {
				return nil, err
			}
		}

		resp, err := next(req)

		if resp != nil {
			if err := LogResponse(ctx, resp, sensitiveBodyFieldNames...); err != nil {
				return nil, err
			}
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request, sensitiveBodyFieldNames ...string) error {
	sensitiveHeaderNames := []string{"x-auth-email", "x-auth-key", "x-auth-user-service-key", "authorization"}

	lines := []string{fmt.Sprintf("\n%s %s %s", req.Method, req.URL.Path, req.Proto)}

	// Log headers
	for name, values := range req.Header {
		for _, value := range values {

			if slices.Contains(sensitiveHeaderNames, strings.ToLower(name)) {
				value = "[redacted]"
			}

			lines = append(lines, fmt.Sprintf("> %s: %s", strings.ToLower(name), value))
		}
	}

	if req.Body != nil {
		// Read the body without mutating the original response
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		// Restore the original body to the response so it can be read again
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		lines = append(lines, ">\n", string(redactSensitiveBodyFields(bodyBytes, sensitiveBodyFieldNames...)), "\n")
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

func redactSensitiveBodyFields(body []byte, sensitiveBodyFieldNames ...string) []byte {
	if len(sensitiveBodyFieldNames) == 0 {
		return body
	}

	var value any
	if err := json.Unmarshal(body, &value); err != nil {
		return body
	}

	sensitiveFields := make(map[string]struct{}, len(sensitiveBodyFieldNames))
	for _, fieldName := range sensitiveBodyFieldNames {
		sensitiveFields[strings.ToLower(fieldName)] = struct{}{}
	}
	redactSensitiveJSONFields(value, sensitiveFields)
	redacted, err := json.Marshal(value)
	if err != nil {
		return body
	}
	return redacted
}

func redactSensitiveJSONFields(value any, sensitiveFields map[string]struct{}) {
	switch value := value.(type) {
	case map[string]any:
		for key, child := range value {
			if _, sensitive := sensitiveFields[strings.ToLower(key)]; sensitive {
				value[key] = "[redacted]"
				continue
			}
			redactSensitiveJSONFields(child, sensitiveFields)
		}
	case []any:
		for _, child := range value {
			redactSensitiveJSONFields(child, sensitiveFields)
		}
	}
}

func LogResponse(ctx context.Context, resp *http.Response, sensitiveBodyFieldNames ...string) error {
	// Log the status code
	lines := []string{fmt.Sprintf("\n< %s %s", resp.Proto, resp.Status)}

	// Log headers
	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("< %s: %s", strings.ToLower(name), value))
		}
	}

	// Read the body without mutating the original response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Restore the original body to the response so it can be read again
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, "<\n", string(redactSensitiveBodyFields(bodyBytes, sensitiveBodyFieldNames...)), "\n")

	// Log the body
	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}
