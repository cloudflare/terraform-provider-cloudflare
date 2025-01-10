package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/cloudflare/cloudflare-go/v4/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if req != nil {
			LogRequest(ctx, req)
		}

		resp, err := next(req)

		if resp != nil {
			LogResponse(ctx, resp)
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) error {
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

		// Log the body
		lines = append(lines, ">\n", string(bodyBytes), "\n")
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
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

	lines = append(lines, "<\n", string(bodyBytes), "\n")

	// Log the body
	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}
