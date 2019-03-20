// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetAppEngineFirewallRuleCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := replaceVars(d, config, "//appengine.googleapis.com/apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetAppEngineFirewallRuleApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "google.appengine.FirewallRule",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1/rest",
				DiscoveryName:        "FirewallRule",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetAppEngineFirewallRuleApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_range"); !isEmptyValue(reflect.ValueOf(sourceRangeProp)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	return obj, nil
}

func expandAppEngineFirewallRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleSourceRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRulePriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
