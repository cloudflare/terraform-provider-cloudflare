package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareLogpushOwnershipChallenge() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareLogpushOwnershipChallengeSchema(),
		Create: resourceCloudflareLogpushOwnershipChallengeCreate,
		Update: resourceCloudflareLogpushOwnershipChallengeCreate,
		Read:   resourceCloudflareLogpushOwnershipChallengeNoop,
		Delete: resourceCloudflareLogpushOwnershipChallengeNoop,
	}
}

func resourceCloudflareLogpushOwnershipChallengeNoop(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudflareLogpushOwnershipChallengeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	destinationConf := d.Get("destination_conf").(string)
	identifier, err := initIdentifier(d)
	if err != nil {
		return err
	}

	var challenge *cloudflare.LogpushGetOwnershipChallenge
	if identifier.Type == AccountType {
		challenge, err = client.GetAccountLogpushOwnershipChallenge(context.Background(), identifier.Value, destinationConf)
	} else {
		challenge, err = client.GetZoneLogpushOwnershipChallenge(context.Background(), identifier.Value, destinationConf)
	}

	if err != nil {
		return fmt.Errorf("error requesting ownership challenge for %s: %w", identifier, err)
	}

	// The ownership challenge doesn't have a unique identifier so we generate it
	// here from the filename which will be unique.
	d.SetId(stringChecksum(challenge.Filename))
	d.Set("ownership_challenge_filename", challenge.Filename)

	log.Printf("[INFO] Created Cloudflare Logpush Ownership Challenge for %s: %s", identifier, d.Id())

	return nil
}
