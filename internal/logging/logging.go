package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v3/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		LogRequest(ctx, req)

		resp, err := next(req)

		LogResponse(ctx, resp)

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) error {
	lines := []string{"\n== Request ==", fmt.Sprintf("url: %s %s", req.Method, req.URL.Path)}

	// Log headers
	lines = append(lines, "== Headers ==")
	for name, values := range req.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("- %s: %s", name, value))
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
		lines = append(lines, fmt.Sprintf("\n== Body ==\n%s\n", string(bodyBytes)))
	}

	tflog.Warn(ctx, strings.Join(lines, "\n"))

	return nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
	// Log the status code
	lines := []string{"\n== Response ==", fmt.Sprintf("status: %s", resp.Status)}

	// Log headers
	lines = append(lines, "== Headers ==")
	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("- %s: %s", name, value))
		}
	}

	// Read the body without mutating the original response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Restore the original body to the response so it can be read again
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, fmt.Sprintf("\n== Body ==\n%s\n", string(bodyBytes)))

	// Log the body
	tflog.Warn(ctx, strings.Join(lines, "\n"))

	return nil
}
