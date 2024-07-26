package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareManagedHeaders() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareManagedHeadersSchema(),
		CreateContext: resourceCloudflareManagedHeadersCreate,
		ReadContext:   resourceCloudflareManagedHeadersRead,
		UpdateContext: resourceCloudflareManagedHeadersUpdate,
		DeleteContext: resourceCloudflareManagedHeadersDelete,
		SchemaVersion: 0,
		Description: heredoc.Doc(`
			The [Cloudflare Managed Headers](https://developers.cloudflare.com/rules/transform/managed-transforms/)
			allows you to add or remove some predefined headers to one's
			requests or origin responses.
		`),
	}
}

func resourceCloudflareManagedHeadersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	d.SetId(zoneID)
	return resourceCloudflareManagedHeadersUpdate(ctx, d, meta)
}

func resourceCloudflareManagedHeadersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	headers, err := client.ListZoneManagedHeaders(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListManagedHeadersParams{
		Status: "enabled",
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading managed headers: %w", err))
	}

	if err := d.Set("managed_request_headers", buildResourceFromManagedHeaders(headers.ManagedRequestHeaders)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("managed_response_headers", buildResourceFromManagedHeaders(headers.ManagedResponseHeaders)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func buildResourceFromManagedHeaders(headers []cloudflare.ManagedHeader) interface{} {
	headersState := []map[string]interface{}{}
	for _, header := range headers {
		headersState = append(headersState, map[string]interface{}{
			"id":      header.ID,
			"enabled": header.Enabled,
		})
	}

	return headersState
}

func resourceCloudflareManagedHeadersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	mh, err := buildManagedHeadersFromResource(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error building managed headers from resource: %w", err))
	}

	managedHeadersPresentInResource := map[string]bool{}
	for _, header := range mh.ManagedRequestHeaders {
		managedHeadersPresentInResource[header.ID] = true
	}
	for _, header := range mh.ManagedResponseHeaders {
		managedHeadersPresentInResource[header.ID] = true
	}

	headers, err := client.ListZoneManagedHeaders(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListManagedHeadersParams{
		Status: "enabled",
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading managed headers: %w", err))
	}

	for _, header := range headers.ManagedRequestHeaders {
		if !managedHeadersPresentInResource[header.ID] {
			header.Enabled = false
			mh.ManagedRequestHeaders = append(mh.ManagedRequestHeaders, header)
		}
	}
	for _, header := range headers.ManagedResponseHeaders {
		if !managedHeadersPresentInResource[header.ID] {
			header.Enabled = false
			mh.ManagedResponseHeaders = append(mh.ManagedResponseHeaders, header)
		}
	}

	if _, err := client.UpdateZoneManagedHeaders(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateManagedHeadersParams{
		ManagedHeaders: mh,
	}); err != nil {
		return diag.FromErr(fmt.Errorf("error updating managed headers: %w", err))
	}
	return resourceCloudflareManagedHeadersRead(ctx, d, meta)
}

// receives the resource config and builds a managed headers struct.
func buildManagedHeadersFromResource(d *schema.ResourceData) (cloudflare.ManagedHeaders, error) {
	requestHeaders, ok := d.Get("managed_request_headers").(*schema.Set)
	if !ok {
		return cloudflare.ManagedHeaders{}, errors.New("unable to create interface array type assertion")
	}
	reqHeaders, err := buildManagedHeadersListFromResource(requestHeaders)
	if err != nil {
		return cloudflare.ManagedHeaders{}, err
	}

	responseHeaders, ok := d.Get("managed_response_headers").(*schema.Set)
	if !ok {
		return cloudflare.ManagedHeaders{}, errors.New("unable to create interface array type assertion")
	}
	respHeaders, err := buildManagedHeadersListFromResource(responseHeaders)
	if err != nil {
		return cloudflare.ManagedHeaders{}, err
	}

	return cloudflare.ManagedHeaders{
		ManagedRequestHeaders:  reqHeaders,
		ManagedResponseHeaders: respHeaders,
	}, nil
}

func buildManagedHeadersListFromResource(resource *schema.Set) ([]cloudflare.ManagedHeader, error) {
	headers := make([]cloudflare.ManagedHeader, 0, len(resource.List()))
	for _, header := range resource.List() {
		h, ok := header.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to create interface map type assertion for managed header")
		}
		id, ok := h["id"].(string)
		if !ok {
			return nil, errors.New("unable to create string type assertion for managed header ID")
		}
		enabled, ok := h["enabled"].(bool)
		if !ok {
			return nil, errors.New("unable to create bool type assertion for managed header enabled")
		}

		if enabled {
			headers = append(headers, cloudflare.ManagedHeader{
				ID:      id,
				Enabled: enabled,
			})
		}
	}
	return headers, nil
}

func resourceCloudflareManagedHeadersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	headers, err := client.ListZoneManagedHeaders(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListManagedHeadersParams{
		Status: "enabled",
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading managed headers: %w", err))
	}

	requestHeaders := make([]cloudflare.ManagedHeader, 0, len(headers.ManagedRequestHeaders))
	for _, header := range headers.ManagedRequestHeaders {
		header.Enabled = false
		requestHeaders = append(requestHeaders, header)
	}
	responseHeaders := make([]cloudflare.ManagedHeader, 0, len(headers.ManagedResponseHeaders))
	for _, header := range headers.ManagedResponseHeaders {
		header.Enabled = false
		responseHeaders = append(responseHeaders, header)
	}

	if _, err := client.UpdateZoneManagedHeaders(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateManagedHeadersParams{
		ManagedHeaders: cloudflare.ManagedHeaders{
			ManagedRequestHeaders:  requestHeaders,
			ManagedResponseHeaders: responseHeaders,
		},
	}); err != nil {
		return diag.FromErr(fmt.Errorf("error deleting managed headers with ID %q: %w", d.Id(), err))
	}

	return nil
}
