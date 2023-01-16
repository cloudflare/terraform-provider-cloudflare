package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessMutualTLSCertificate() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessMutualTLSCertificateSchema(),
		CreateContext: resourceCloudflareAccessMutualTLSCertificateCreate,
		ReadContext:   resourceCloudflareAccessMutualTLSCertificateRead,
		UpdateContext: resourceCloudflareAccessMutualTLSCertificateUpdate,
		DeleteContext: resourceCloudflareAccessMutualTLSCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessMutualTLSCertificateImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Mutual TLS Certificate resource.
			Mutual TLS authentication ensures that the traffic is secure and
			trusted in both directions between a client and server and can be
			 used with Access to only allows requests from devices with a
			 corresponding client certificate.
		`),
	}
}

func resourceCloudflareAccessMutualTLSCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	newAccessMutualTLSCertificate := cloudflare.AccessMutualTLSCertificate{
		Name:        d.Get("name").(string),
		Certificate: d.Get("certificate").(string),
	}
	newAccessMutualTLSCertificate.AssociatedHostnames = expandInterfaceToStringList(d.Get("associated_hostnames"))

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Mutual TLS certificate from struct: %+v", newAccessMutualTLSCertificate))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var accessMutualTLSCert cloudflare.AccessMutualTLSCertificate
	if identifier.Type == AccountType {
		accessMutualTLSCert, err = client.CreateAccessMutualTLSCertificate(ctx, identifier.Value, newAccessMutualTLSCertificate)
	} else {
		accessMutualTLSCert, err = client.CreateZoneAccessMutualTLSCertificate(ctx, identifier.Value, newAccessMutualTLSCertificate)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Mutual TLS Certificate for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	d.SetId(accessMutualTLSCert.ID)

	return resourceCloudflareAccessMutualTLSCertificateRead(ctx, d, meta)
}

func resourceCloudflareAccessMutualTLSCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var accessMutualTLSCert cloudflare.AccessMutualTLSCertificate
	if identifier.Type == AccountType {
		accessMutualTLSCert, err = client.AccessMutualTLSCertificate(ctx, identifier.Value, d.Id())
	} else {
		accessMutualTLSCert, err = client.ZoneAccessMutualTLSCertificate(ctx, identifier.Value, d.Id())
	}

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Mutal TLS Certificate %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Access Mutual TLS Certificate %q: %w", d.Id(), err))
	}

	d.Set("name", accessMutualTLSCert.Name)
	d.Set("associated_hostnames", accessMutualTLSCert.AssociatedHostnames)
	d.Set("fingerprint", accessMutualTLSCert.Fingerprint)

	return nil
}

func resourceCloudflareAccessMutualTLSCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	updatedAccessMutualTLSCert := cloudflare.AccessMutualTLSCertificate{
		ID:   d.Id(),
		Name: d.Get("name").(string),
	}
	updatedAccessMutualTLSCert.AssociatedHostnames = expandInterfaceToStringList(d.Get("associated_hostnames"))

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Mutal TLS Certificate from struct: %+v", updatedAccessMutualTLSCert))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if identifier.Type == AccountType {
		_, err = client.UpdateAccessMutualTLSCertificate(ctx, identifier.Value, d.Id(), updatedAccessMutualTLSCert)
	} else {
		_, err = client.UpdateZoneAccessMutualTLSCertificate(ctx, identifier.Value, d.Id(), updatedAccessMutualTLSCert)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Mutual TLS Certificate for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	return resourceCloudflareAccessMutualTLSCertificateRead(ctx, d, meta)
}

func resourceCloudflareAccessMutualTLSCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	certID := d.Id()

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Mutual TLS Certificate using ID: %s", certID))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	// To actually delete the certificate, it cannot have any hostnames associated
	// with it so here we perform an update (to remove them) before we continue on
	// with wiping the certificate itself.
	deletedCertificate := cloudflare.AccessMutualTLSCertificate{
		ID:                  d.Id(),
		Name:                d.Get("name").(string),
		AssociatedHostnames: []string{},
	}

	if identifier.Type == AccountType {
		_, err = client.UpdateAccessMutualTLSCertificate(ctx, identifier.Value, d.Id(), deletedCertificate)
	} else {
		_, err = client.UpdateZoneAccessMutualTLSCertificate(ctx, identifier.Value, d.Id(), deletedCertificate)
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Mutual TLS Certificate for %s %q: %w", identifier.Type, identifier.Value, err))
	}

	retryErr := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		if identifier.Type == AccountType {
			err = client.DeleteAccessMutualTLSCertificate(ctx, identifier.Value, certID)
		} else {
			err = client.DeleteZoneAccessMutualTLSCertificate(ctx, identifier.Value, certID)
		}

		if err != nil {
			if strings.Contains(err.Error(), "access.api.error.certificate_has_active_associations") {
				return resource.RetryableError(fmt.Errorf("certificate associations are not yet removed"))
			} else {
				return resource.NonRetryableError(fmt.Errorf("error deleting Access Mutual TLS Certificate for %s %q: %w", identifier.Type, identifier.Value, err))
			}
		}

		d.SetId("")

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	return nil
}

func resourceCloudflareAccessMutualTLSCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"account/accountID/accessMutualTLSCertificateID\" or \"zone/zoneID/accessMutualTLSCertificateID\"", d.Id())
	}

	identifierType, identifierID, accessMutualTLSCertificateID := attributes[0], attributes[1], attributes[2]

	if AccessIdentifierType(identifierType) != AccountType && AccessIdentifierType(identifierType) != ZoneType {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"account/accountID/accessMutualTLSCertificateID\" or \"zone/zoneID/accessMutualTLSCertificateID\"", d.Id())
	}

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Mutual TLS Certificate: id %s for %s %s", accessMutualTLSCertificateID, identifierType, identifierID))

	//lintignore:R001
	d.Set(fmt.Sprintf("%s_id", identifierType), identifierID)
	d.SetId(accessMutualTLSCertificateID)

	resourceCloudflareAccessMutualTLSCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
