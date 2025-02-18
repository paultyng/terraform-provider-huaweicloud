package cfw

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

var nonUpdatableParams = []string{"object_id"}

// @API CFW POST /v1/{project_id}/acl-rule
// @API CFW DELETE /v1/{project_id}/acl-rule/{id}
// @API CFW PUT /v1/{project_id}/acl-rule/{id}
// @API CFW GET /v1/{project_id}/acl-rules
// @API CFW PUT /v1/{project_id}/acl-rule/order/{id}
// @API CFW POST /v1/{project_id}/acl-rule/count
// @API CFW DELETE /v1/{project_id}/acl-rule/count
func ResourceAclRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceACLRuleCreate,
		UpdateContext: resourceACLRuleUpdate,
		ReadContext:   resourceACLRuleRead,
		DeleteContext: resourceACLRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceACLRuleImportState,
		},

		CustomizeDiff: config.FlexibleForceNew(nonUpdatableParams),

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The protected object ID.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The rule name.`,
			},
			"type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The rule type.`,
			},
			"action_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The action type.`,
			},
			"address_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The address type.`,
			},
			"sequence": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        ACLRuleOrderRuleAclDtoSchema(),
				Required:    true,
				Description: `The sequence configuration.`,
			},
			"status": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The rule status.`,
			},
			"long_connect_enable": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  `Whether to support persistent connections.`,
				ValidateFunc: validation.IntInSlice([]int{0, 1}),
			},
			"applications": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The application list.`,
			},
			"custom_services": {
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          ACLRuleServiceItemSchema(),
				Description:   `The custom service configuration.`,
				ConflictsWith: []string{"custom_service_groups", "predefined_service_groups"},
			},
			"custom_service_groups": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          ACLRuleServiceGroupSchema(),
				Description:   `The custom service group list.`,
				ConflictsWith: []string{"custom_services"},
			},
			"predefined_service_groups": {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				Elem:          ACLRuleServiceGroupSchema(),
				Description:   `The predefined service group list.`,
				ConflictsWith: []string{"custom_services"},
			},
			"source_addresses": {
				Type:          schema.TypeList,
				MinItems:      1,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   `The source IP address list.`,
				ConflictsWith: []string{"source_region_list", "source_address_groups", "source_predefined_groups"},
			},
			"source_region_list": {
				Type:          schema.TypeList,
				MinItems:      1,
				Optional:      true,
				Elem:          ACLRuleIpRegionDtoSchema(),
				Description:   `The source region list.`,
				ConflictsWith: []string{"source_addresses", "source_address_groups", "source_predefined_groups"},
			},
			"source_address_groups": {
				Type:          schema.TypeList,
				MinItems:      1,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   `The source address group list.`,
				ConflictsWith: []string{"source_addresses", "source_region_list"},
			},
			"source_predefined_groups": {
				Type:          schema.TypeList,
				MinItems:      1,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				Description:   `The source predefined address group list.`,
				ConflictsWith: []string{"source_addresses", "source_region_list"},
			},
			"source_address_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `The source address type.`,
			},
			"destination_addresses": {
				Type:        schema.TypeList,
				MinItems:    1,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The destination IP address list.`,
				ConflictsWith: []string{
					"destination_region_list", "destination_domain_address_name", "destination_domain_group_id",
					"destination_domain_group_name", "destination_domain_group_type", "destination_address_groups",
				},
			},
			"destination_region_list": {
				Type:        schema.TypeList,
				MinItems:    1,
				Optional:    true,
				Elem:        ACLRuleIpRegionDtoSchema(),
				Description: `The destination region list.`,
				ConflictsWith: []string{
					"destination_addresses", "destination_domain_address_name", "destination_domain_group_id",
					"destination_domain_group_name", "destination_domain_group_type", "destination_address_groups",
				},
			},
			"destination_domain_address_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The destination domain address name.`,
				ConflictsWith: []string{
					"destination_addresses", "destination_region_list", "destination_domain_group_id",
					"destination_domain_group_name", "destination_domain_group_type", "destination_address_groups",
				},
			},
			"destination_domain_group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The destination domain group ID.`,
				RequiredWith: []string{"destination_domain_group_name", "destination_domain_group_type"},
				ConflictsWith: []string{
					"destination_addresses", "destination_region_list", "destination_domain_address_name",
					"destination_address_groups",
				},
			},
			"destination_domain_group_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The destination domain group name.`,
				RequiredWith: []string{"destination_domain_group_id", "destination_domain_group_type"},
			},
			"destination_domain_group_type": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  `The destination domain group type.`,
				RequiredWith: []string{"destination_domain_group_id", "destination_domain_group_name"},
			},
			"destination_address_groups": {
				Type:        schema.TypeList,
				MinItems:    1,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The destination address group list.`,
				ConflictsWith: []string{
					"destination_addresses", "destination_region_list", "destination_domain_address_name",
					"destination_domain_group_id", "destination_domain_group_name", "destination_domain_group_type",
				},
			},
			"destination_address_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `The destination address type.`,
			},
			"long_connect_time_hour": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The persistent connection duration (hour).`,
			},
			"long_connect_time_minute": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The persistent connection duration (minute).`,
			},
			"long_connect_time_second": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The persistent Connection Duration (second).`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The rule description.`,
			},
			"direction": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: `The rule direction.`,
			},
			"rule_hit_count": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"0"}, true),
				Description:  `The number of times the ACL rule is hit.`,
			},
			"tags": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ValidateFunc: func(v interface{}, _ string) ([]string, []error) {
					if keys, ok := v.(map[string]interface{}); ok && len(keys) > 1 {
						return nil, []error{fmt.Errorf("tags can take at most one key-value pair")}
					}
					return nil, nil
				},
				Description: `The key/value pairs to associate with the ACL rule.`,
			},
			"enable_force_new": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
			},
		},
	}
}

func ACLRuleOrderRuleAclDtoSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"dest_rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The ID of the rule that the added rule will follow.`,
			},
			"top": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Whether to pin on top.`,
			},
			"bottom": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Whether to pin on bottom.`,
			},
		},
	}
	return &sc
}

func ACLRuleServiceItemSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocol": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `The protocol type.`,
			},
			"source_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The source port.`,
			},
			"dest_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The destination port.`,
			},
		},
	}
	return &sc
}

func ACLRuleServiceGroupSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"protocols": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: `The protocols used in the service groups.`,
			},
			"group_ids": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The IDs of the service groups.`,
			},
		},
	}
	return &sc
}

func ACLRuleIpRegionDtoSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The region ID.`,
			},
			"region_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The region type.",
			},
			"description_cn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Chinese description of the region.",
			},
			"description_en": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The English description of the region.",
			},
		},
	}
	return &sc
}

func resourceACLRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)

	// createACLRule: Create a CFW ACL Rule.
	var (
		createACLRuleHttpUrl = "v1/{project_id}/acl-rule"
		createACLRuleProduct = "cfw"
	)
	createACLRuleClient, err := conf.NewServiceClient(createACLRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating CFW client: %s", err)
	}

	createACLRulePath := createACLRuleClient.Endpoint + createACLRuleHttpUrl
	createACLRulePath = strings.ReplaceAll(createACLRulePath, "{project_id}", createACLRuleClient.ProjectID)

	createACLRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	createACLRuleOpt.JSONBody = utils.RemoveNil(buildCreateACLRuleBodyParams(d))
	createACLRuleResp, err := createACLRuleClient.Request("POST", createACLRulePath, &createACLRuleOpt)
	if err != nil {
		return diag.Errorf("error creating ACL rule: %s", err)
	}

	createACLRuleRespBody, err := utils.FlattenResponse(createACLRuleResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id := utils.PathSearch("data.rules[0].id", createACLRuleRespBody, "").(string)
	if id == "" {
		return diag.Errorf("error creating ACL rule: ID is not found in API response")
	}
	d.SetId(id)

	return resourceACLRuleRead(ctx, d, meta)
}

func buildCreateACLRuleBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"object_id": utils.ValueIgnoreEmpty(d.Get("object_id")),
		"type":      d.Get("type"),
		"rules":     []map[string]interface{}{buildACLRulesOpts(d, false)},
	}
	return bodyParams
}

func buildACLRulesOpts(d *schema.ResourceData, isUpdate bool) map[string]interface{} {
	applicationList := utils.ExpandToStringList(d.Get("applications").(*schema.Set).List())
	params := map[string]interface{}{
		"action_type":              d.Get("action_type"),
		"address_type":             d.Get("address_type"),
		"applications":             utils.ValueIgnoreEmpty(applicationList),
		"description":              utils.ValueIgnoreEmpty(d.Get("description")),
		"direction":                d.Get("direction"),
		"long_connect_enable":      d.Get("long_connect_enable"),
		"long_connect_time_hour":   utils.ValueIgnoreEmpty(d.Get("long_connect_time_hour")),
		"long_connect_time_minute": utils.ValueIgnoreEmpty(d.Get("long_connect_time_minute")),
		"long_connect_time_second": utils.ValueIgnoreEmpty(d.Get("long_connect_time_second")),
		"name":                     utils.ValueIgnoreEmpty(d.Get("name")),
		"status":                   d.Get("status"),
		"tag":                      buildACLRuleRequestBodyTagsVO(d.Get("tags").(map[string]interface{})),
	}

	if !isUpdate {
		params["sequence"] = buildACLRuleRequestBodyOrderRuleAclDto(d.Get("sequence"))
	}

	if isUpdate {
		params["type"] = d.Get("type")
	}

	customServices, customServicesOk := d.GetOk("custom_services")
	rawCustomServiceGroups, customServiceGroupsOk := d.GetOk("custom_service_groups")
	rawPredefinedServiceGroups, predefinedServiceGroupsOk := d.GetOk("predefined_service_groups")

	switch {
	case customServicesOk:
		params["service"] = buildACLRuleRequestBodyRuleCustomServices(customServices)
	case customServiceGroupsOk, predefinedServiceGroupsOk:
		customServiceGroups := rawCustomServiceGroups.([]interface{})
		predefinedServiceGroups := rawPredefinedServiceGroups.([]interface{})
		params["service"] = buildACLRuleRequestBodyRuleServiceGroups(customServiceGroups, predefinedServiceGroups)
	default:
		// Indicates setting any service.
		params["service"] = map[string]interface{}{
			"type":        0,
			"protocol":    -1,
			"source_port": "1-65535",
			"dest_port":   "1-65535",
		}
	}

	sourceAddresses, sourceAddressesOk := d.GetOk("source_addresses")
	sourceRegionList, sourceRegionListOk := d.GetOk("source_region_list")
	rawSourceAddressGroups, sourceAddressGroupOk := d.GetOk("source_address_groups")
	rawSourcePredefinedGroups, sourcePredefinedGroupsOk := d.GetOk("source_predefined_groups")
	sourceAddressType := d.Get("source_address_type").(int)

	switch {
	case sourceAddressesOk:
		params["source"] = buildACLRuleRequestBodyRuleAddresses(utils.ExpandToStringList(sourceAddresses.([]interface{})), sourceAddressType)
	case sourceRegionListOk:
		params["source"] = buildACLRuleRequestBodyIpRegionDto(sourceRegionList, sourceAddressType)
	case sourceAddressGroupOk, sourcePredefinedGroupsOk:
		sourceAddressGroups := rawSourceAddressGroups.([]interface{})
		sourcePredefinedGroups := rawSourcePredefinedGroups.([]interface{})
		params["source"] = buildACLRuleRequestBodyRuleAddressGroups(sourceAddressGroups, sourcePredefinedGroups, sourceAddressType)
	default:
		params["source"] = buildACLRuleRequestBodyRuleAnyAddress()
	}

	destinationAddresses, destinationAddressesOk := d.GetOk("destination_addresses")
	destinationRegionList, destinationRegionListOk := d.GetOk("destination_region_list")
	rawDestinationAddressGroups, destinationAddressGroupsOk := d.GetOk("destination_address_groups")
	destinationDomainAddressName, destinationDomainAddressNameOk := d.GetOk("destination_domain_address_name")
	destinationDomainGroupId, destinationDomainGroupIdOk := d.GetOk("destination_domain_group_id")
	destinationDomainGroupName := d.Get("destination_domain_group_name").(string)
	destinationDomainGroupType := d.Get("destination_domain_group_type").(int)
	destinationAddressType := d.Get("destination_address_type").(int)

	switch {
	case destinationAddressesOk:
		params["destination"] = buildACLRuleRequestBodyRuleAddresses(
			utils.ExpandToStringList(destinationAddresses.([]interface{})),
			destinationAddressType,
		)
	case destinationRegionListOk:
		params["destination"] = buildACLRuleRequestBodyIpRegionDto(destinationRegionList, destinationAddressType)
	case destinationDomainAddressNameOk:
		params["destination"] = buildACLRuleRequestBodyDomainAddress(destinationDomainAddressName, destinationAddressType)
	case destinationDomainGroupIdOk:
		params["destination"] = buildACLRuleRequestBodyDomainAddressGroup(
			destinationDomainGroupId, destinationDomainGroupName,
			destinationDomainGroupType, destinationAddressType,
		)
	case destinationAddressGroupsOk:
		destinationAddressGroups := rawDestinationAddressGroups.([]interface{})
		params["destination"] = buildACLRuleRequestBodyRuleAddressGroups(destinationAddressGroups, make([]interface{}, 0), destinationAddressType)
	default:
		params["destination"] = buildACLRuleRequestBodyRuleAnyAddress()
	}
	return params
}

func buildACLRuleRequestBodyTagsVO(tagmap map[string]interface{}) map[string]interface{} {
	tags := make(map[string]interface{})
	for k, v := range tagmap {
		tags["tag_key"] = k
		tags["tag_value"] = v
	}
	return tags
}

func buildACLRuleRequestBodyOrderRuleAclDto(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw := rawArray[0].(map[string]interface{})
		params := map[string]interface{}{
			"dest_rule_id": utils.ValueIgnoreEmpty(raw["dest_rule_id"]),
			"top":          raw["top"],
			"bottom":       raw["bottom"],
		}
		return params
	}
	return nil
}

func buildACLRuleRequestBodyRuleCustomServices(rawParams interface{}) map[string]interface{} {
	customServices := buildACLRuleRequestBodyRuleCustomService(rawParams)
	if customServices != nil {
		return map[string]interface{}{
			"custom_service": customServices,
			"type":           2,
		}
	}
	return nil
}

func buildACLRuleRequestBodyRuleCustomService(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		params := make([]map[string]interface{}, 0, len(rawArray))
		for _, rawService := range rawArray {
			service := rawService.(map[string]interface{})
			param := map[string]interface{}{
				"protocol":    service["protocol"],
				"source_port": service["source_port"],
				"dest_port":   service["dest_port"],
			}
			params = append(params, param)
		}
		return params
	}
	return nil
}

func buildACLRuleRequestBodyRuleServiceGroups(customServiceGroups, predefinedServiceGroups []interface{}) map[string]interface{} {
	var customServiceGroupsProtocols []interface{}
	var predefinedServiceGroupsProtocols []interface{}

	if len(customServiceGroups) == 0 && len(predefinedServiceGroups) == 0 {
		return nil
	}

	param := map[string]interface{}{
		"type": 2,
	}
	if len(customServiceGroups) > 0 {
		v := customServiceGroups[0].(map[string]interface{})
		param["service_group"] = utils.ExpandToStringList(v["group_ids"].([]interface{}))
		customServiceGroupsProtocols = v["protocols"].(*schema.Set).List()
	}
	if len(predefinedServiceGroups) > 0 {
		v := predefinedServiceGroups[0].(map[string]interface{})
		param["predefined_group"] = utils.ExpandToStringList(v["group_ids"].([]interface{}))
		predefinedServiceGroupsProtocols = v["protocols"].(*schema.Set).List()
	}
	param["protocols"] = mergeProtocols(customServiceGroupsProtocols, predefinedServiceGroupsProtocols)

	return param
}

func mergeProtocols(customProtocols, predefinedProtocols []interface{}) []int {
	protocolMap := make(map[int]struct{})

	for _, protocol := range customProtocols {
		protocolMap[protocol.(int)] = struct{}{}
	}

	for _, num := range predefinedProtocols {
		protocolMap[num.(int)] = struct{}{}
	}

	result := make([]int, 0, len(protocolMap))
	for num := range protocolMap {
		result = append(result, num)
	}

	return result
}

func buildACLRuleRequestBodyRuleAddresses(rawParams []string, addressType int) map[string]interface{} {
	if len(rawParams) == 0 {
		return nil
	}
	return map[string]interface{}{
		"ip_address":   rawParams,
		"type":         5,
		"address_type": addressType,
	}
}

func buildACLRuleRequestBodyIpRegionDto(rawParams interface{}, addressType int) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		params := make([]map[string]interface{}, 0, len(rawArray))
		for _, rawRegion := range rawArray {
			region := rawRegion.(map[string]interface{})
			param := map[string]interface{}{
				"region_id":      region["region_id"],
				"region_type":    region["region_type"],
				"description_cn": utils.ValueIgnoreEmpty(region["description_cn"]),
				"description_en": utils.ValueIgnoreEmpty(region["description_en"]),
			}
			params = append(params, param)
		}
		return map[string]interface{}{
			"type":        3,
			"region_list": params,
			"addressType": addressType,
		}
	}
	return nil
}

func buildACLRuleRequestBodyRuleAddressGroups(addressGroups, predefinedGroups []interface{}, addressType int) map[string]interface{} {
	if len(addressGroups) == 0 && len(predefinedGroups) == 0 {
		return nil
	}
	return map[string]interface{}{
		"address_group":    utils.ValueIgnoreEmpty(utils.ExpandToStringList(addressGroups)),
		"predefined_group": utils.ValueIgnoreEmpty(utils.ExpandToStringList(predefinedGroups)),
		"address_type":     addressType,
		"type":             5,
	}
}

func buildACLRuleRequestBodyDomainAddress(domainAddressName interface{}, addressType int) map[string]interface{} {
	if domainAddressName == nil {
		return nil
	}
	return map[string]interface{}{
		"domain_address_name": domainAddressName,
		"type":                2,
		"address_type":        addressType,
	}
}

func buildACLRuleRequestBodyDomainAddressGroup(domainGroupId, domainGroupName interface{}, domainGroupType, addressType int) map[string]interface{} {
	if domainGroupId == nil {
		return nil
	}
	return map[string]interface{}{
		"domain_set_id":   domainGroupId,
		"domain_set_name": domainGroupName,
		"type":            domainGroupType,
		"address_type":    addressType,
	}
}

func buildACLRuleRequestBodyRuleAnyAddress() map[string]interface{} {
	return map[string]interface{}{
		"type":    0,
		"address": "0.0.0.0/0",
	}
}

func resourceACLRuleRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)

	var mErr *multierror.Error

	// getACLRule: Query the CFW ACL rule detail
	getACLRuleProduct := "cfw"
	getACLRuleClient, err := conf.NewServiceClient(getACLRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating CFW client: %s", err)
	}

	objectID := d.Get("object_id").(string)
	rule, err := GetACLRule(getACLRuleClient, d.Id(), objectID)
	if err != nil {
		return common.CheckDeletedDiag(d, parseError(err), "error retrieving ACL rule")
	}

	count, err := getACLRuleHitCount(getACLRuleClient, d.Id())
	if err != nil {
		return diag.Errorf("error retrieving ACL rule hit count: %s", err)
	}

	ruleHitCount := ""
	if count != nil {
		if v, ok := count.(float64); ok {
			ruleHitCount = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}

	// the params 'sequence' and 'type 'not not returned
	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("action_type", utils.PathSearch("action_type", rule, nil)),
		d.Set("address_type", utils.PathSearch("address_type", rule, nil)),
		d.Set("applications", utils.PathSearch("applications", rule, nil)),
		d.Set("description", utils.PathSearch("description", rule, nil)),
		d.Set("direction", utils.PathSearch("direction", rule, nil)),
		d.Set("long_connect_enable", utils.PathSearch("long_connect_enable", rule, nil)),
		d.Set("long_connect_time_hour", utils.PathSearch("long_connect_time_hour", rule, nil)),
		d.Set("long_connect_time_minute", utils.PathSearch("long_connect_time_minute", rule, nil)),
		d.Set("long_connect_time_second", utils.PathSearch("long_connect_time_second", rule, nil)),
		d.Set("name", utils.PathSearch("name", rule, nil)),
		d.Set("source_addresses", utils.PathSearch("source.ip_address", rule, nil)),
		d.Set("source_region_list", flattenGetACLRuleResponseBodyRuleIpRegionDto(utils.PathSearch("source.region_list", rule, nil))),
		d.Set("source_address_groups", utils.PathSearch("source.address_group_names[?address_set_type==`0`].set_id", rule, nil)),
		d.Set("source_address_type", utils.PathSearch("source.address_type", rule, nil)),
		d.Set("destination_addresses", utils.PathSearch("destination.ip_address", rule, nil)),
		d.Set("destination_region_list", flattenGetACLRuleResponseBodyRuleIpRegionDto(utils.PathSearch("destination.region_list", rule, nil))),
		d.Set("destination_domain_address_name", utils.PathSearch("destination.domain_address_name", rule, nil)),
		d.Set("destination_address_groups", utils.PathSearch("destination.address_group_names[?address_set_type==`0`].set_id", rule, nil)),
		d.Set("destination_address_type", utils.PathSearch("destination.address_type", rule, nil)),
		d.Set("status", utils.PathSearch("status", rule, nil)),
		d.Set("tags", flattenGetACLRuleResponseBodyRuleTagsVO(rule)),
		d.Set("rule_hit_count", ruleHitCount),
	)

	customServices := utils.PathSearch("service.custom_service", rule, nil)
	customServiceGroups := utils.PathSearch("service.service_group_names[?service_set_type==`0`].set_id", rule, make([]interface{}, 0))
	if customServices != nil {
		mErr = multierror.Append(d.Set("custom_services", flattenGetACLRuleResponseBodyRuleCustomServices(rule)), mErr)
	}
	if len(customServiceGroups.([]interface{})) > 0 {
		mErr = multierror.Append(d.Set("custom_service_groups", flattenGetACLRuleResponseBodyRuleCustomServiceGroup(rule)), mErr)
	}

	domainSetID := utils.PathSearch("destination.domain_set_id", rule, nil)
	if domainSetID != nil {
		mErr = multierror.Append(
			mErr,
			d.Set("destination_domain_group_id", domainSetID.(string)),
			d.Set("destination_domain_group_name", utils.PathSearch("destination.domain_set_name", rule, nil)),
			d.Set("destination_domain_group_type", utils.PathSearch("destination.type", rule, nil)),
		)
	}

	return diag.FromErr(mErr.ErrorOrNil())
}

func GetACLRule(client *golangsdk.ServiceClient, id, objectID string) (interface{}, error) {
	httpUrl := "v1/{project_id}/acl-rules"
	path := client.Endpoint + httpUrl
	path = strings.ReplaceAll(path, "{project_id}", client.ProjectID)
	opt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	offset := 0
	for {
		path := fmt.Sprintf("%s?object_id=%s&limit=100&offset=%d", path, objectID, offset)
		resp, err := client.Request("GET", path, &opt)

		if err != nil {
			return nil, err
		}

		respBody, err := utils.FlattenResponse(resp)
		if err != nil {
			return nil, err
		}

		findRuleStr := fmt.Sprintf("data.records[?rule_id=='%s']|[0]", id)
		rule := utils.PathSearch(findRuleStr, respBody, nil)
		if rule != nil {
			return rule, nil
		}

		offset += 100
		total := utils.PathSearch("data.total", respBody, float64(0))
		if int(total.(float64)) <= offset {
			return nil, golangsdk.ErrDefault404{}
		}
	}
}

func getACLRuleHitCount(client *golangsdk.ServiceClient, id string) (interface{}, error) {
	getACLRuleHitCountHttpUrl := "v1/{project_id}/acl-rule/count"
	getACLRuleHitCountPath := client.Endpoint + getACLRuleHitCountHttpUrl
	getACLRuleHitCountPath = strings.ReplaceAll(getACLRuleHitCountPath, "{project_id}", client.ProjectID)

	getACLRuleHitCountOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         buildRuleHitCountBodyParams(id),
	}

	getACLRuleHitCountResp, err := client.Request("POST", getACLRuleHitCountPath, &getACLRuleHitCountOpt)
	if err != nil {
		return nil, err
	}

	getACLRuleHitCountRespBody, err := utils.FlattenResponse(getACLRuleHitCountResp)
	if err != nil {
		return nil, err
	}

	count := utils.PathSearch("data.records[0].rule_hit_count", getACLRuleHitCountRespBody, nil)
	if count == nil {
		return nil, fmt.Errorf("error parsing rule_hit_count from response= %#v", getACLRuleHitCountRespBody)
	}
	return count, nil
}

func buildRuleHitCountBodyParams(id string) map[string]interface{} {
	return map[string]interface{}{
		"rule_ids": []string{id},
	}
}

func flattenGetACLRuleResponseBodyRuleTagsVO(resp interface{}) map[string]interface{} {
	curJson := utils.PathSearch("tag", resp, nil)

	if curJson == nil {
		return nil
	}

	if tagMap, ok := curJson.(map[string]interface{}); ok {
		key, value := "", ""
		for k, v := range tagMap {
			switch k {
			case "tag_key":
				key = v.(string)
			case "tag_value":
				value = v.(string)
			}
		}
		return map[string]interface{}{key: value}
	}
	return nil
}

func flattenGetACLRuleResponseBodyRuleCustomServices(resp interface{}) []interface{} {
	curJson := utils.PathSearch("service.custom_service", resp, nil)

	if curJson == nil {
		return nil
	}

	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"protocol":    utils.PathSearch("protocol", v, nil),
			"source_port": utils.PathSearch("source_port", v, nil),
			"dest_port":   utils.PathSearch("dest_port", v, nil),
		})
	}
	return rst
}

func flattenGetACLRuleResponseBodyRuleCustomServiceGroup(resp interface{}) []interface{} {
	param := map[string]interface{}{
		"group_ids": utils.PathSearch("service.service_group_names[?service_set_type==`0`].set_id", resp, make([]interface{}, 0)),
		"protocols": flattenGetACLRuleResponseBodyRuleProtocols(resp),
	}
	return []interface{}{param}
}

func flattenGetACLRuleResponseBodyRuleProtocols(resp interface{}) []interface{} {
	curJson := utils.PathSearch("service.service_group_names[?service_set_type==`0`].protocols", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})

	if len(curArray) == 0 {
		return nil
	}

	protocolSet := make(map[int]struct{})
	for _, protocols := range curArray {
		for _, v := range protocols.([]interface{}) {
			protocol := int(v.(float64))
			protocolSet[protocol] = struct{}{}
		}
	}

	uniqueProtocols := make([]interface{}, 0, len(protocolSet))
	for protocol := range protocolSet {
		uniqueProtocols = append(uniqueProtocols, protocol)
	}
	return uniqueProtocols
}

func flattenGetACLRuleResponseBodyRuleIpRegionDto(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}

	curArray := resp.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"region_id":      utils.PathSearch("region_id", v, nil),
			"region_type":    utils.PathSearch("region_type", v, nil),
			"description_cn": utils.PathSearch("description_cn", v, nil),
			"description_en": utils.PathSearch("description_en", v, nil),
		})
	}
	return rst
}

func resourceACLRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)

	updateACLRulehasChanges := []string{
		"action_type",
		"address_type",
		"applications",
		"description",
		"direction",
		"long_connect_enable",
		"long_connect_time_hour",
		"long_connect_time_minute",
		"long_connect_time_second",
		"name",
		"custom_services",
		"custom_service_groups",
		"predefined_service_groups",
		"source_addresses",
		"source_region_list",
		"source_address_groups",
		"source_predefined_groups",
		"source_address_type",
		"destination_addresses",
		"destination_region_list",
		"destination_domain_address_name",
		"destination_domain_group_id",
		"destination_domain_group_name",
		"destination_domain_group_type",
		"destination_address_groups",
		"destination_address_type",
		"status",
		"type",
		"tags",
	}

	var (
		updateACLRuleHttpUrl         = "v1/{project_id}/acl-rule/{id}"
		updateACLRuleOrderHttpUrl    = "v1/{project_id}/acl-rule/order/{id}"
		updateACLRuleHitCountHttpUrl = "v1/{project_id}/acl-rule/count"
		updateACLRuleProduct         = "cfw"
	)
	updateACLRuleClient, err := conf.NewServiceClient(updateACLRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating CFW client: %s", err)
	}

	if d.HasChanges(updateACLRulehasChanges...) {
		// updateACLRule: Update the configuration of CFW ACL Rule
		updateACLRulePath := updateACLRuleClient.Endpoint + updateACLRuleHttpUrl
		updateACLRulePath = strings.ReplaceAll(updateACLRulePath, "{project_id}", updateACLRuleClient.ProjectID)
		updateACLRulePath = strings.ReplaceAll(updateACLRulePath, "{id}", d.Id())

		updateACLRuleOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
		}
		updateACLRuleOpt.JSONBody = utils.RemoveNil(buildACLRulesOpts(d, true))
		_, err = updateACLRuleClient.Request("PUT", updateACLRulePath, &updateACLRuleOpt)
		if err != nil {
			return diag.Errorf("error updating ACL rule: %s", err)
		}
	}

	if d.HasChange("sequence") {
		// updateACLRuleOrder: Update the order of CFW ACL Rule
		updateACLRuleOrderPath := updateACLRuleClient.Endpoint + updateACLRuleOrderHttpUrl
		updateACLRuleOrderPath = strings.ReplaceAll(updateACLRuleOrderPath, "{project_id}", updateACLRuleClient.ProjectID)
		updateACLRuleOrderPath = strings.ReplaceAll(updateACLRuleOrderPath, "{id}", d.Id())

		updateACLRuleOrderOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
		}
		updateACLRuleOrderOpt.JSONBody = utils.RemoveNil(buildACLRuleRequestBodyOrderRuleAclDto(d.Get("sequence")))
		_, err = updateACLRuleClient.Request("PUT", updateACLRuleOrderPath, &updateACLRuleOrderOpt)
		if err != nil {
			return diag.Errorf("error updating ACL rule order: %s", err)
		}
	}

	if d.HasChange("rule_hit_count") {
		updateRuleHitCountPath := updateACLRuleClient.Endpoint + updateACLRuleHitCountHttpUrl
		updateRuleHitCountPath = strings.ReplaceAll(updateRuleHitCountPath, "{project_id}", updateACLRuleClient.ProjectID)
		updateRuleHitCountOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			JSONBody:         buildRuleHitCountBodyParams(d.Id()),
		}
		_, err := updateACLRuleClient.Request("DELETE", updateRuleHitCountPath, &updateRuleHitCountOpt)
		if err != nil {
			return diag.Errorf("error updating ACL rule hit count: %s", err)
		}
	}

	return resourceACLRuleRead(ctx, d, meta)
}

func resourceACLRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)

	// deleteACLRule: Delete an existing CFW ACL Rule
	var (
		deleteACLRuleHttpUrl = "v1/{project_id}/acl-rule/{id}"
		deleteACLRuleProduct = "cfw"
	)
	deleteACLRuleClient, err := conf.NewServiceClient(deleteACLRuleProduct, region)
	if err != nil {
		return diag.Errorf("error creating CFW client: %s", err)
	}

	deleteACLRulePath := deleteACLRuleClient.Endpoint + deleteACLRuleHttpUrl
	deleteACLRulePath = strings.ReplaceAll(deleteACLRulePath, "{project_id}", deleteACLRuleClient.ProjectID)
	deleteACLRulePath = strings.ReplaceAll(deleteACLRulePath, "{id}", d.Id())

	deleteACLRuleOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}
	_, err = deleteACLRuleClient.Request("DELETE", deleteACLRulePath, &deleteACLRuleOpt)
	if err != nil {
		return common.CheckDeletedDiag(d, parseError(err), "error deleting ACL rule")
	}

	err = deleteACLRuleWaitingForCompleted(ctx, deleteACLRuleClient, d.Id(), d.Get("object_id").(string), d.Timeout(schema.TimeoutDelete))
	return common.CheckDeletedDiag(d, parseError(err), "error deleting ACL rule")
}

func deleteACLRuleWaitingForCompleted(ctx context.Context, client *golangsdk.ServiceClient, id, objectID string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETE"},
		Refresh: func() (interface{}, string, error) {
			rule, err := GetACLRule(client, id, objectID)
			if rule != nil {
				return rule, "PENDING", nil
			}
			return nil, "COMPLETE", err
		},
		Timeout:      timeout,
		Delay:        10 * time.Second,
		PollInterval: 10 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func resourceACLRuleImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format specified for import ID, must be <object_id>/<id>")
	}

	d.Set("object_id", parts[0])
	d.SetId(parts[1])

	return []*schema.ResourceData{d}, nil
}
