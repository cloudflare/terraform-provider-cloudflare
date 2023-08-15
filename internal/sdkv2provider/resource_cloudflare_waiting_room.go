package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWaitingRoom() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareWaitingRoomCreate,
		ReadContext:   resourceCloudflareWaitingRoomRead,
		UpdateContext: resourceCloudflareWaitingRoomUpdate,
		DeleteContext: resourceCloudflareWaitingRoomDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWaitingRoomImport,
		},

		Schema: resourceCloudflareWaitingRoomSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
		Description: "Provides a Cloudflare Waiting Room resource.",
	}
}

func buildWaitingRoom(d *schema.ResourceData) cloudflare.WaitingRoom {
	routes := d.Get("additional_routes").([]interface{})
	additional_routes := []*cloudflare.WaitingRoomRoute{}
	if routes != nil {
		for _, route := range routes {
			r := route.(map[string]interface{})
			additional_routes = append(additional_routes, &cloudflare.WaitingRoomRoute{
				Host: r["host"].(string),
				Path: r["path"].(string),
			})
		}
	}

	return cloudflare.WaitingRoom{
		Name:                    d.Get("name").(string),
		Description:             d.Get("description").(string),
		Suspended:               d.Get("suspended").(bool),
		Host:                    d.Get("host").(string),
		Path:                    d.Get("path").(string),
		TotalActiveUsers:        d.Get("total_active_users").(int),
		NewUsersPerMinute:       d.Get("new_users_per_minute").(int),
		CustomPageHTML:          d.Get("custom_page_html").(string),
		QueueingMethod:          d.Get("queueing_method").(string),
		DefaultTemplateLanguage: d.Get("default_template_language").(string),
		SessionDuration:         d.Get("session_duration").(int),
		JsonResponseEnabled:     d.Get("json_response_enabled").(bool),
		QueueAll:                d.Get("queue_all").(bool),
		DisableSessionRenewal:   d.Get("disable_session_renewal").(bool),
		CookieSuffix:            d.Get("cookie_suffix").(string),
		AdditionalRoutes:        additional_routes,
		QueueingStatusCode:      d.Get("queueing_status_code").(int),
	}
}

func resourceCloudflareWaitingRoomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	newWaitingRoom := buildWaitingRoom(d)

	waitingRoom, err := client.CreateWaitingRoom(ctx, zoneID, newWaitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error creating waiting room %q: %w", name, err))
	}

	d.SetId(waitingRoom.ID)

	return resourceCloudflareWaitingRoomRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoom, err := client.WaitingRoom(ctx, zoneID, waitingRoomID)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing waiting room from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error getting waiting room %q: %w", name, err))
	}
	d.SetId(waitingRoom.ID)
	d.Set("name", waitingRoom.Name)
	d.Set("description", waitingRoom.Description)
	d.Set("suspended", waitingRoom.Suspended)
	d.Set("host", waitingRoom.Host)
	d.Set("path", waitingRoom.Path)
	d.Set("queue_all", waitingRoom.QueueAll)
	d.Set("new_users_per_minute", waitingRoom.NewUsersPerMinute)
	d.Set("total_active_users", waitingRoom.TotalActiveUsers)
	d.Set("session_duration", waitingRoom.SessionDuration)
	d.Set("disable_session_renewal", waitingRoom.DisableSessionRenewal)
	d.Set("queueing_method", waitingRoom.QueueingMethod)
	d.Set("custom_page_html", waitingRoom.CustomPageHTML)
	d.Set("default_template_language", waitingRoom.DefaultTemplateLanguage)
	d.Set("json_response_enabled", waitingRoom.JsonResponseEnabled)
	d.Set("cookie_suffix", waitingRoom.CookieSuffix)
	d.Set("additional_routes", flattenWaitingRoomAdditionalRoutes(waitingRoom.AdditionalRoutes))
	d.Set("queueing_status_code", waitingRoom.QueueingStatusCode)
	return nil
}

func resourceCloudflareWaitingRoomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	waitingRoom := buildWaitingRoom(d)

	_, err := client.ChangeWaitingRoom(ctx, zoneID, waitingRoomID, waitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error updating waiting room %q: %w", name, err))
	}

	return resourceCloudflareWaitingRoomRead(ctx, d, meta)
}

func resourceCloudflareWaitingRoomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteWaitingRoom(ctx, zoneID, waitingRoomID)

	if err != nil {
		name := d.Get("name").(string)
		return diag.FromErr(fmt.Errorf("error deleting waiting room %q: %w", name, err))
	}

	return nil
}

func resourceCloudflareWaitingRoomImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var waitingRoomID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		waitingRoomID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/waitingRoomID\" for import", d.Id())
	}

	waitingRoom, err := client.WaitingRoom(ctx, zoneID, waitingRoomID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Waiting room %s", waitingRoomID)
	}

	d.SetId(waitingRoom.ID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)

	resourceCloudflareWaitingRoomRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}

func flattenWaitingRoomAdditionalRoutes(routes []*cloudflare.WaitingRoomRoute) []interface{} {
	flattened := []interface{}{}
	if routes == nil {
		return flattened
	}
	for _, r := range routes {
		flattened = append(flattened, map[string]interface{}{
			"host": r.Host,
			"path": r.Path,
		})
	}
	return flattened
}
