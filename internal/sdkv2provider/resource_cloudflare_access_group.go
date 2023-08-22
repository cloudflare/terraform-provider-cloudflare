package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessGroup() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAccessGroupSchema(),
		CreateContext: resourceCloudflareAccessGroupCreate,
		ReadContext:   resourceCloudflareAccessGroupRead,
		UpdateContext: resourceCloudflareAccessGroupUpdate,
		DeleteContext: resourceCloudflareAccessGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAccessGroupImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Access Group resource. Access Groups are used
			in conjunction with Access Policies to restrict access to a
			particular resource based on group membership.
		`),
	}
}

func resourceCloudflareAccessGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessGroup, err := client.GetAccessGroup(ctx, identifier, d.Id())

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Access Group %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Access Group %q: %w", d.Id(), err))
	}

	d.Set("name", accessGroup.Name)

	if err := d.Set("require", TransformAccessGroupForSchema(ctx, accessGroup.Require)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set require attribute: %w", err))
	}

	if err := d.Set("exclude", TransformAccessGroupForSchema(ctx, accessGroup.Exclude)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set exclude attribute: %w", err))
	}

	if err := d.Set("include", TransformAccessGroupForSchema(ctx, accessGroup.Include)); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set include attribute: %w", err))
	}

	return nil
}

func resourceCloudflareAccessGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	newAccessGroup := cloudflare.CreateAccessGroupParams{
		Name: d.Get("name").(string),
	}

	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		if value != nil {
			newAccessGroup.Exclude = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		if value != nil {
			newAccessGroup.Require = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		if value != nil {
			newAccessGroup.Include = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Access Group from struct: %+v", newAccessGroup))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessGroup, err := client.CreateAccessGroup(ctx, identifier, newAccessGroup)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Access Group for ID %q: %w", accessGroup.ID, err))
	}

	d.SetId(accessGroup.ID)

	return resourceCloudflareAccessGroupRead(ctx, d, meta)
}

func resourceCloudflareAccessGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	updatedAccessGroup := cloudflare.UpdateAccessGroupParams{
		Name: d.Get("name").(string),
		ID:   d.Id(),
	}

	exclude := d.Get("exclude").([]interface{})
	for _, value := range exclude {
		if value != nil {
			updatedAccessGroup.Exclude = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	require := d.Get("require").([]interface{})
	for _, value := range require {
		if value != nil {
			updatedAccessGroup.Require = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	include := d.Get("include").([]interface{})
	for _, value := range include {
		if value != nil {
			updatedAccessGroup.Include = BuildAccessGroupCondition(value.(map[string]interface{}))
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Access Group from struct: %+v", updatedAccessGroup))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	accessGroup, err := client.UpdateAccessGroup(ctx, identifier, updatedAccessGroup)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Access Group for ID %q: %w", d.Id(), err))
	}

	if accessGroup.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Access Group ID in update response; resource was empty"))
	}

	return resourceCloudflareAccessGroupRead(ctx, d, meta)
}

func resourceCloudflareAccessGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Access Group using ID: %s", d.Id()))

	identifier, err := initIdentifier(d)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.DeleteAccessGroup(ctx, identifier, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Access Group for ID %q: %w", d.Id(), err))
	}

	resourceCloudflareAccessGroupRead(ctx, d, meta)

	return nil
}

func resourceCloudflareAccessGroupImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/accessGroupID\"", d.Id())
	}

	accountID, accessGroupID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Access Group: accountID %q, accessGroupID %q", accountID, accessGroupID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(accessGroupID)

	resourceCloudflareAccessGroupRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

// BuildAccessGroupCondition iterates the provided `map` of values and
// generates the required (repetitive) structs.
//
// Returns the intended combination structure of Access Groups to build the
// desired policy.
func BuildAccessGroupCondition(options map[string]interface{}) []interface{} {
	var group []interface{}
	for accessGroupType, values := range options {
		if accessGroupType == "everyone" {
			if values == true {
				group = append(group, cloudflare.AccessGroupEveryone{})
			}
		} else if accessGroupType == "any_valid_service_token" {
			if values == true {
				group = append(group, cloudflare.AccessGroupAnyValidServiceToken{})
			}
		} else if accessGroupType == "certificate" {
			if values == true {
				group = append(group, cloudflare.AccessGroupCertificate{})
			}
		} else if accessGroupType == "common_name" {
			if values != "" {
				group = append(group, cloudflare.AccessGroupCertificateCommonName{CommonName: struct {
					CommonName string `json:"common_name"`
				}{CommonName: values.(string)}})
			}
		} else if accessGroupType == "auth_method" {
			if values != "" {
				group = append(group, cloudflare.AccessGroupAuthMethod{AuthMethod: struct {
					AuthMethod string `json:"auth_method"`
				}{AuthMethod: values.(string)}})
			}
		} else if accessGroupType == "gsuite" {
			for _, v := range values.([]interface{}) {
				gsuiteCfg := v.(map[string]interface{})
				for _, email := range gsuiteCfg["email"].([]interface{}) {
					group = append(group, cloudflare.AccessGroupGSuite{Gsuite: struct {
						Email              string `json:"email"`
						IdentityProviderID string `json:"identity_provider_id"`
					}{
						Email:              email.(string),
						IdentityProviderID: gsuiteCfg["identity_provider_id"].(string),
					}})
				}
			}
		} else if accessGroupType == "github" {
			for _, v := range values.([]interface{}) {
				githubCfg := v.(map[string]interface{})
				if len(githubCfg["teams"].([]interface{})) > 0 {
					for _, team := range githubCfg["teams"].([]interface{}) {
						group = append(group, cloudflare.AccessGroupGitHub{GitHubOrganization: struct {
							Name               string `json:"name"`
							Team               string `json:"team,omitempty"`
							IdentityProviderID string `json:"identity_provider_id"`
						}{
							Name:               githubCfg["name"].(string),
							Team:               team.(string),
							IdentityProviderID: githubCfg["identity_provider_id"].(string),
						}})
					}
				} else {
					group = append(group, cloudflare.AccessGroupGitHub{GitHubOrganization: struct {
						Name               string `json:"name"`
						Team               string `json:"team,omitempty"`
						IdentityProviderID string `json:"identity_provider_id"`
					}{
						Name:               githubCfg["name"].(string),
						IdentityProviderID: githubCfg["identity_provider_id"].(string),
					}})
				}
			}
		} else if accessGroupType == "azure" {
			for _, v := range values.([]interface{}) {
				azureCfg := v.(map[string]interface{})
				for _, id := range azureCfg["id"].([]interface{}) {
					group = append(group, cloudflare.AccessGroupAzure{AzureAD: struct {
						ID                 string `json:"id"`
						IdentityProviderID string `json:"identity_provider_id"`
					}{
						ID:                 id.(string),
						IdentityProviderID: azureCfg["identity_provider_id"].(string),
					}})
				}
			}
		} else if accessGroupType == "okta" {
			for _, v := range values.([]interface{}) {
				oktaCfg := v.(map[string]interface{})
				for _, name := range oktaCfg["name"].([]interface{}) {
					group = append(group, cloudflare.AccessGroupOkta{Okta: struct {
						Name               string `json:"name"`
						IdentityProviderID string `json:"identity_provider_id"`
					}{
						Name:               name.(string),
						IdentityProviderID: oktaCfg["identity_provider_id"].(string),
					}})
				}
			}
		} else if accessGroupType == "saml" {
			for _, v := range values.([]interface{}) {
				samlCfg := v.(map[string]interface{})
				group = append(group, cloudflare.AccessGroupSAML{Saml: struct {
					AttributeName      string `json:"attribute_name"`
					AttributeValue     string `json:"attribute_value"`
					IdentityProviderID string `json:"identity_provider_id"`
				}{
					AttributeName:      samlCfg["attribute_name"].(string),
					AttributeValue:     samlCfg["attribute_value"].(string),
					IdentityProviderID: samlCfg["identity_provider_id"].(string),
				}})
			}
		} else if accessGroupType == "external_evaluation" {
			for _, v := range values.([]interface{}) {
				eeCfg := v.(map[string]interface{})
				group = append(group, cloudflare.AccessGroupExternalEvaluation{ExternalEvaluation: struct {
					EvaluateURL string `json:"evaluate_url"`
					KeysURL     string `json:"keys_url"`
				}{
					EvaluateURL: eeCfg["evaluate_url"].(string),
					KeysURL:     eeCfg["keys_url"].(string),
				}})
			}
		} else if accessGroupType == "auth_context" {
			for _, v := range values.([]interface{}) {
				ctxCfg := v.(map[string]interface{})
				group = append(group, cloudflare.AccessGroupAzureAuthContext{AuthContext: struct {
					ID                 string `json:"id"`
					IdentityProviderID string `json:"identity_provider_id"`
					ACID               string `json:"ac_id"`
				}{
					ID:                 ctxCfg["id"].(string),
					IdentityProviderID: ctxCfg["identity_provider_id"].(string),
					ACID:               ctxCfg["ac_id"].(string),
				}})
			}
		} else {
			for _, value := range values.([]interface{}) {
				switch accessGroupType {
				case "email":
					group = append(group, cloudflare.AccessGroupEmail{Email: struct {
						Email string `json:"email"`
					}{Email: value.(string)}})
				case "email_domain":
					group = append(group, cloudflare.AccessGroupEmailDomain{EmailDomain: struct {
						Domain string `json:"domain"`
					}{Domain: value.(string)}})
				case "ip":
					group = append(group, cloudflare.AccessGroupIP{IP: struct {
						IP string `json:"ip"`
					}{IP: value.(string)}})
				case "ip_list":
					group = append(group, cloudflare.AccessGroupIPList{IPList: struct {
						ID string `json:"id"`
					}{ID: value.(string)}})
				case "service_token":
					group = append(group, cloudflare.AccessGroupServiceToken{ServiceToken: struct {
						ID string `json:"token_id"`
					}{ID: value.(string)}})
				case "group":
					group = append(group, cloudflare.AccessGroupAccessGroup{Group: struct {
						ID string `json:"id"`
					}{ID: value.(string)}})
				case "geo":
					group = append(group, cloudflare.AccessGroupGeo{Geo: struct {
						CountryCode string `json:"country_code"`
					}{CountryCode: value.(string)}})
				case "login_method":
					group = append(group, cloudflare.AccessGroupLoginMethod{LoginMethod: struct {
						ID string `json:"id"`
					}{ID: value.(string)}})
				case "device_posture":
					group = append(group, cloudflare.AccessGroupDevicePosture{DevicePosture: struct {
						ID string `json:"integration_uid"`
					}{ID: value.(string)}})
				}
			}
		}
	}

	return group
}

// TransformAccessGroupForSchema takes the incoming `accessGroup` from the API
// response and converts it to a usable schema for the conditions.
func TransformAccessGroupForSchema(ctx context.Context, accessGroup []interface{}) []map[string]interface{} {
	data := []map[string]interface{}{}

	if len(accessGroup) == 0 {
		return data
	}

	everyone := false
	anyValidServiceToken := false
	certificate := false
	emails := []string{}
	emailDomains := []string{}
	ips := []string{}
	ipList := []string{}
	serviceTokens := []string{}
	groups := []string{}
	commonName := ""
	authMethod := ""
	geos := []string{}
	loginMethod := []string{}
	oktaID := ""
	oktaGroups := []string{}
	gsuiteID := ""
	gsuiteEmails := []string{}
	githubName := ""
	githubTeams := []string{}
	githubID := ""
	azureGroups := []map[string]interface{}{}
	samlGroups := []map[string]string{}
	externalEvaluationURL := ""
	externalEvaluationKeysURL := ""
	devicePostureRuleIDs := []string{}
	authCtxID := ""
	authCtxIDPID := ""
	authCtxACID := ""

	for _, group := range accessGroup {
		for groupKey, groupValue := range group.(map[string]interface{}) {
			switch groupKey {
			case "everyone":
				everyone = true
			case "any_valid_service_token":
				anyValidServiceToken = true
			case "certificate":
				certificate = true
			case "email":
				for _, email := range groupValue.(map[string]interface{}) {
					emails = append(emails, email.(string))
				}
			case "email_domain":
				for _, domain := range groupValue.(map[string]interface{}) {
					emailDomains = append(emailDomains, domain.(string))
				}
			case "ip":
				for _, ip := range groupValue.(map[string]interface{}) {
					ips = append(ips, ip.(string))
				}
			case "ip_list":
				for _, ipListID := range groupValue.(map[string]interface{}) {
					ipList = append(ipList, ipListID.(string))
				}
			case "service_token":
				for _, serviceToken := range groupValue.(map[string]interface{}) {
					serviceTokens = append(serviceTokens, serviceToken.(string))
				}
			case "common_name":
				for _, name := range groupValue.(map[string]interface{}) {
					commonName = name.(string)
				}
			case "auth_method":
				for _, method := range groupValue.(map[string]interface{}) {
					authMethod = method.(string)
				}
			case "geo":
				for _, geo := range groupValue.(map[string]interface{}) {
					geos = append(geos, geo.(string))
				}
			case "login_method":
				for _, method := range groupValue.(map[string]interface{}) {
					loginMethod = append(loginMethod, method.(string))
				}
			case "okta":
				oktaCfg := groupValue.(map[string]interface{})
				oktaID = oktaCfg["identity_provider_id"].(string)
				oktaGroups = append(oktaGroups, oktaCfg["name"].(string))
			case "gsuite":
				gsuiteCfg := groupValue.(map[string]interface{})
				gsuiteID = gsuiteCfg["identity_provider_id"].(string)
				gsuiteEmails = append(gsuiteEmails, gsuiteCfg["email"].(string))
			case "github-organization":
				githubCfg := groupValue.(map[string]interface{})
				githubID = githubCfg["identity_provider_id"].(string)
				githubName = githubCfg["name"].(string)
				if v, ok := githubCfg["team"]; ok {
					githubTeams = append(githubTeams, v.(string))
				}
			case "azureAD":
				azureCfg := groupValue.(map[string]interface{})
				azureIdPID := azureCfg["identity_provider_id"].(string)
				azureGroupID := azureCfg["id"].(string)

				var azureGroup map[string]interface{}
				for _, ag := range azureGroups {
					if ag["identity_provider_id"] == azureIdPID {
						azureGroup = ag
						break
					}
				}

				if len(azureGroup) == 0 {
					azureGroups = append(azureGroups, map[string]interface{}{"identity_provider_id": azureIdPID, "id": []string{azureGroupID}})
				} else {
					azureGroup["id"] = append(azureGroup["id"].([]string), azureGroupID)
				}
			case "saml":
				samlCfg := groupValue.(map[string]interface{})
				samlAttrName := samlCfg["attribute_name"].(string)
				samlAttrValue := samlCfg["attribute_value"].(string)
				s := map[string]string{"attribute_name": samlAttrName, "attribute_value": samlAttrValue}
				if idpID, ok := samlCfg["identity_provider_id"].(string); ok {
					s["identity_provider_id"] = idpID
				}
				samlGroups = append(samlGroups, s)
			case "external_evaluation":
				eeCfg := groupValue.(map[string]interface{})
				externalEvaluationURL = eeCfg["evaluate_url"].(string)
				externalEvaluationKeysURL = eeCfg["keys_url"].(string)
			case "group":
				for _, group := range groupValue.(map[string]interface{}) {
					groups = append(groups, group.(string))
				}
			case "device_posture":
				for _, dprID := range groupValue.(map[string]interface{}) {
					devicePostureRuleIDs = append(devicePostureRuleIDs, dprID.(string))
				}
			case "auth_context":
				ctxCfg := groupValue.(map[string]interface{})
				authCtxIDPID = ctxCfg["identity_provider_id"].(string)
				authCtxID = ctxCfg["id"].(string)
				authCtxACID = ctxCfg["ac_id"].(string)
			default:
				tflog.Debug(ctx, fmt.Sprintf("Access Group key %q not transformed", groupKey))
			}
		}
	}

	groupMap := map[string]interface{}{}

	if everyone {
		groupMap["everyone"] = everyone
	}

	if anyValidServiceToken {
		groupMap["any_valid_service_token"] = anyValidServiceToken
	}

	if certificate {
		groupMap["certificate"] = certificate
	}

	if len(emails) > 0 {
		groupMap["email"] = emails
	}

	if len(emailDomains) > 0 {
		groupMap["email_domain"] = emailDomains
	}

	if len(ips) > 0 {
		groupMap["ip"] = ips
	}

	if len(ipList) > 0 {
		groupMap["ip_list"] = ipList
	}

	if len(serviceTokens) > 0 {
		groupMap["service_token"] = serviceTokens
	}

	if commonName != "" {
		groupMap["common_name"] = commonName
	}

	if authMethod != "" {
		groupMap["auth_method"] = authMethod
	}

	if len(geos) > 0 {
		groupMap["geo"] = geos
	}

	if len(loginMethod) > 0 {
		groupMap["login_method"] = loginMethod
	}

	if len(oktaGroups) > 0 && oktaID != "" {
		groupMap["okta"] = []interface{}{
			map[string]interface{}{
				"identity_provider_id": oktaID,
				"name":                 oktaGroups,
			},
		}
	}

	if len(gsuiteEmails) > 0 && gsuiteID != "" {
		groupMap["gsuite"] = []interface{}{
			map[string]interface{}{
				"identity_provider_id": gsuiteID,
				"email":                gsuiteEmails,
			},
		}
	}

	if githubID != "" && githubName != "" {
		groupMap["github"] = []interface{}{
			map[string]interface{}{
				"name":                 githubName,
				"teams":                githubTeams,
				"identity_provider_id": githubID,
			},
		}
	}

	if len(azureGroups) > 0 {
		groupMap["azure"] = azureGroups
	}

	if len(samlGroups) > 0 {
		groupMap["saml"] = samlGroups
	}

	if externalEvaluationURL != "" && externalEvaluationKeysURL != "" {
		groupMap["external_evaluation"] = []interface{}{
			map[string]interface{}{
				"evaluate_url": externalEvaluationURL,
				"keys_url":     externalEvaluationKeysURL,
			},
		}
	}

	if authCtxID != "" && authCtxACID != "" && authCtxIDPID != "" {
		groupMap["auth_context"] = []interface{}{
			map[string]interface{}{
				"id":                   authCtxID,
				"ac_id":                authCtxACID,
				"identity_provider_id": authCtxIDPID,
			},
		}
	}

	if len(groups) > 0 {
		groupMap["group"] = groups
	}

	if len(devicePostureRuleIDs) > 0 {
		groupMap["device_posture"] = devicePostureRuleIDs
	}

	data = append(data, groupMap)

	return data
}
