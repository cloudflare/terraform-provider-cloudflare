package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareWaitingRoom() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareWaitingRoomCreate,
		Read:   resourceCloudflareWaitingRoomRead,
		Update: resourceCloudflareWaitingRoomUpdate,
		Delete: resourceCloudflareWaitingRoomDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareWaitingRoomImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(i interface{}) string {
					return strings.ToLower(i.(string))
				},
			},

			"host": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(i interface{}) string {
					return strings.ToLower(i.(string))
				},
			},

			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"total_active_users": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"new_users_per_minute": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"custom_page_html": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"queue_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_session_renewal": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"suspended": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"session_duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
			Update: schema.DefaultTimeout(30 * time.Second),
		},
	}
}

func buildWaitingRoom(d *schema.ResourceData) cloudflare.WaitingRoom {
	return cloudflare.WaitingRoom{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		Suspended:             d.Get("suspended").(bool),
		Host:                  d.Get("host").(string),
		Path:                  d.Get("path").(string),
		TotalActiveUsers:      d.Get("total_active_users").(int),
		NewUsersPerMinute:     d.Get("new_users_per_minute").(int),
		CustomPageHtml:        d.Get("custom_page_html").(string),
		SessionDuration:       d.Get("session_duration").(int),
		QueueAll:              d.Get("queue_all").(bool),
		DisableSessionRenewal: d.Get("disable_session_renewal").(bool),
	}
}

func resourceCloudflareWaitingRoomReadFromWaitingRoomId(d *schema.ResourceData, meta interface{}, zoneID, waitingRoomID string) error {
	client := meta.(*cloudflare.API)
	waitingRoom, err := client.WaitingRoom(context.Background(), zoneID, waitingRoomID)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[WARN] Removing waiting room from state because it's not found in API")
			d.SetId("")
			return nil
		}
		name := d.Get("name").(string)
		return fmt.Errorf("Error getting waiting room %q: %s", name, err)
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
	d.Set("custom_page_html", waitingRoom.CustomPageHtml)
	return nil
}

func resourceCloudflareWaitingRoomCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	newWaitingRoom := buildWaitingRoom(d)

	waitingRoom, err := client.CreateWaitingRoom(context.Background(), zoneID, newWaitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("Error creating waiting room %q: %s", name, err)
	}

	d.SetId(waitingRoom.ID)

	return resourceCloudflareWaitingRoomRead(d, meta)
}

func resourceCloudflareWaitingRoomRead(d *schema.ResourceData, meta interface{}) error {
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	return resourceCloudflareWaitingRoomReadFromWaitingRoomId(d, meta, zoneID, waitingRoomID)
}

func resourceCloudflareWaitingRoomUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	waitingRoom := buildWaitingRoom(d)

	err := client.UpdateWaitingRoom(context.Background(), zoneID, waitingRoomID, waitingRoom)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("Error updating waiting room %q: %s", name, err)
	}

	return resourceCloudflareWaitingRoomRead(d, meta)
}

func resourceCloudflareWaitingRoomDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	waitingRoomID := d.Id()
	zoneID := d.Get("zone_id").(string)

	err := client.DeleteWaitingRoom(context.Background(), zoneID, waitingRoomID)

	if err != nil {
		name := d.Get("name").(string)
		return fmt.Errorf("Error updating waiting room %q: %s", name, err)
	}

	return nil
}

func resourceCloudflareWaitingRoomImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneID string
	var waitingRoomID string
	if len(idAttr) == 2 {
		zoneID = idAttr[0]
		waitingRoomID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/waitingRoomID\" for import", d.Id())
	}

	err := resourceCloudflareWaitingRoomReadFromWaitingRoomId(d, meta, zoneID, waitingRoomID)
	if err != nil {
		return nil, fmt.Errorf("Unable to find waiting room %s/%s: %s", zoneID, waitingRoomID, err)
	}

	if d.Id() != waitingRoomID {
		return nil, fmt.Errorf("Unable to find waiting room %s/%s: %s", zoneID, waitingRoomID, err)
	}
	d.Set("zone_id", zoneID)

	return []*schema.ResourceData{d}, nil
}
