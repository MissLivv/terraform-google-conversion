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

func GetMonitoringAlertPolicyCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := replaceVars(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetMonitoringAlertPolicyApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "google.monitoring.AlertPolicy",
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "AlertPolicy",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetMonitoringAlertPolicyApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringAlertPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	combinerProp, err := expandMonitoringAlertPolicyCombiner(d.Get("combiner"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("combiner"); !isEmptyValue(reflect.ValueOf(combinerProp)) && (ok || !reflect.DeepEqual(v, combinerProp)) {
		obj["combiner"] = combinerProp
	}
	enabledProp, err := expandMonitoringAlertPolicyEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}
	conditionsProp, err := expandMonitoringAlertPolicyConditions(d.Get("conditions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("conditions"); !isEmptyValue(reflect.ValueOf(conditionsProp)) && (ok || !reflect.DeepEqual(v, conditionsProp)) {
		obj["conditions"] = conditionsProp
	}
	notificationChannelsProp, err := expandMonitoringAlertPolicyNotificationChannels(d.Get("notification_channels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_channels"); !isEmptyValue(reflect.ValueOf(notificationChannelsProp)) && (ok || !reflect.DeepEqual(v, notificationChannelsProp)) {
		obj["notificationChannels"] = notificationChannelsProp
	}
	labelsProp, err := expandMonitoringAlertPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	documentationProp, err := expandMonitoringAlertPolicyDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("documentation"); !isEmptyValue(reflect.ValueOf(documentationProp)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}

	return obj, nil
}

func expandMonitoringAlertPolicyDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyCombiner(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedConditionAbsent, err := expandMonitoringAlertPolicyConditionsConditionAbsent(original["condition_absent"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConditionAbsent); val.IsValid() && !isEmptyValue(val) {
			transformed["conditionAbsent"] = transformedConditionAbsent
		}

		transformedName, err := expandMonitoringAlertPolicyConditionsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedConditionThreshold, err := expandMonitoringAlertPolicyConditionsConditionThreshold(original["condition_threshold"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConditionThreshold); val.IsValid() && !isEmptyValue(val) {
			transformed["conditionThreshold"] = transformedConditionThreshold
		}

		transformedDisplayName, err := expandMonitoringAlertPolicyConditionsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !isEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAggregations, err := expandMonitoringAlertPolicyConditionsConditionAbsentAggregations(original["aggregations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAggregations); val.IsValid() && !isEmptyValue(val) {
		transformed["aggregations"] = transformedAggregations
	}

	transformedTrigger, err := expandMonitoringAlertPolicyConditionsConditionAbsentTrigger(original["trigger"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrigger); val.IsValid() && !isEmptyValue(val) {
		transformed["trigger"] = transformedTrigger
	}

	transformedDuration, err := expandMonitoringAlertPolicyConditionsConditionAbsentDuration(original["duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["duration"] = transformedDuration
	}

	transformedFilter, err := expandMonitoringAlertPolicyConditionsConditionAbsentFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !isEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentAggregations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPerSeriesAligner, err := expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAligner(original["per_series_aligner"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPerSeriesAligner); val.IsValid() && !isEmptyValue(val) {
			transformed["perSeriesAligner"] = transformedPerSeriesAligner
		}

		transformedGroupByFields, err := expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsGroupByFields(original["group_by_fields"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByFields); val.IsValid() && !isEmptyValue(val) {
			transformed["groupByFields"] = transformedGroupByFields
		}

		transformedAlignmentPeriod, err := expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsAlignmentPeriod(original["alignment_period"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlignmentPeriod); val.IsValid() && !isEmptyValue(val) {
			transformed["alignmentPeriod"] = transformedAlignmentPeriod
		}

		transformedCrossSeriesReducer, err := expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducer(original["cross_series_reducer"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCrossSeriesReducer); val.IsValid() && !isEmptyValue(val) {
			transformed["crossSeriesReducer"] = transformedCrossSeriesReducer
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAligner(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsGroupByFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsAlignmentPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducer(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentTrigger(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPercent, err := expandMonitoringAlertPolicyConditionsConditionAbsentTriggerPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	transformedCount, err := expandMonitoringAlertPolicyConditionsConditionAbsentTriggerCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !isEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	return transformed, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentTriggerPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentTriggerCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionAbsentFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedThresholdValue, err := expandMonitoringAlertPolicyConditionsConditionThresholdThresholdValue(original["threshold_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThresholdValue); val.IsValid() && !isEmptyValue(val) {
		transformed["thresholdValue"] = transformedThresholdValue
	}

	transformedDenominatorFilter, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorFilter(original["denominator_filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDenominatorFilter); val.IsValid() && !isEmptyValue(val) {
		transformed["denominatorFilter"] = transformedDenominatorFilter
	}

	transformedDenominatorAggregations, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(original["denominator_aggregations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDenominatorAggregations); val.IsValid() && !isEmptyValue(val) {
		transformed["denominatorAggregations"] = transformedDenominatorAggregations
	}

	transformedDuration, err := expandMonitoringAlertPolicyConditionsConditionThresholdDuration(original["duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !isEmptyValue(val) {
		transformed["duration"] = transformedDuration
	}

	transformedComparison, err := expandMonitoringAlertPolicyConditionsConditionThresholdComparison(original["comparison"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedComparison); val.IsValid() && !isEmptyValue(val) {
		transformed["comparison"] = transformedComparison
	}

	transformedTrigger, err := expandMonitoringAlertPolicyConditionsConditionThresholdTrigger(original["trigger"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrigger); val.IsValid() && !isEmptyValue(val) {
		transformed["trigger"] = transformedTrigger
	}

	transformedAggregations, err := expandMonitoringAlertPolicyConditionsConditionThresholdAggregations(original["aggregations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAggregations); val.IsValid() && !isEmptyValue(val) {
		transformed["aggregations"] = transformedAggregations
	}

	transformedFilter, err := expandMonitoringAlertPolicyConditionsConditionThresholdFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !isEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdThresholdValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPerSeriesAligner, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAligner(original["per_series_aligner"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPerSeriesAligner); val.IsValid() && !isEmptyValue(val) {
			transformed["perSeriesAligner"] = transformedPerSeriesAligner
		}

		transformedGroupByFields, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsGroupByFields(original["group_by_fields"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByFields); val.IsValid() && !isEmptyValue(val) {
			transformed["groupByFields"] = transformedGroupByFields
		}

		transformedAlignmentPeriod, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsAlignmentPeriod(original["alignment_period"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlignmentPeriod); val.IsValid() && !isEmptyValue(val) {
			transformed["alignmentPeriod"] = transformedAlignmentPeriod
		}

		transformedCrossSeriesReducer, err := expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducer(original["cross_series_reducer"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCrossSeriesReducer); val.IsValid() && !isEmptyValue(val) {
			transformed["crossSeriesReducer"] = transformedCrossSeriesReducer
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAligner(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsGroupByFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsAlignmentPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducer(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdComparison(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdTrigger(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPercent, err := expandMonitoringAlertPolicyConditionsConditionThresholdTriggerPercent(original["percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPercent); val.IsValid() && !isEmptyValue(val) {
		transformed["percent"] = transformedPercent
	}

	transformedCount, err := expandMonitoringAlertPolicyConditionsConditionThresholdTriggerCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !isEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	return transformed, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdTriggerPercent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdTriggerCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdAggregations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPerSeriesAligner, err := expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAligner(original["per_series_aligner"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPerSeriesAligner); val.IsValid() && !isEmptyValue(val) {
			transformed["perSeriesAligner"] = transformedPerSeriesAligner
		}

		transformedGroupByFields, err := expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsGroupByFields(original["group_by_fields"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByFields); val.IsValid() && !isEmptyValue(val) {
			transformed["groupByFields"] = transformedGroupByFields
		}

		transformedAlignmentPeriod, err := expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsAlignmentPeriod(original["alignment_period"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlignmentPeriod); val.IsValid() && !isEmptyValue(val) {
			transformed["alignmentPeriod"] = transformedAlignmentPeriod
		}

		transformedCrossSeriesReducer, err := expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducer(original["cross_series_reducer"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCrossSeriesReducer); val.IsValid() && !isEmptyValue(val) {
			transformed["crossSeriesReducer"] = transformedCrossSeriesReducer
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAligner(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsGroupByFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsAlignmentPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducer(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsConditionThresholdFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyConditionsDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyNotificationChannels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyDocumentation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedContent, err := expandMonitoringAlertPolicyDocumentationContent(original["content"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
		transformed["content"] = transformedContent
	}

	transformedMimeType, err := expandMonitoringAlertPolicyDocumentationMimeType(original["mime_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMimeType); val.IsValid() && !isEmptyValue(val) {
		transformed["mimeType"] = transformedMimeType
	}

	return transformed, nil
}

func expandMonitoringAlertPolicyDocumentationContent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringAlertPolicyDocumentationMimeType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
