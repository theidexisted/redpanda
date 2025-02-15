// Copyright 2021 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package admin

import "net/http"

// FeatureState enumerates the possible states of a feature.
type FeatureState string

const (
	FeatureStateActive      FeatureState = "active"
	FeatureStatePreparing   FeatureState = "preparing"
	FeatureStateAvailable   FeatureState = "available"
	FeatureStateUnavailable FeatureState = "unavailable"
	FeatureStateDisabled    FeatureState = "disabled"
)

// Feature contains information on the state of a feature.
type Feature struct {
	Name      string       `json:"name"`
	State     FeatureState `json:"state"`
	WasActive bool         `json:"was_active"`
}

// FeaturesResponse contains information on the features available on a Redpanda cluster.
type FeaturesResponse struct {
	ClusterVersion int       `json:"cluster_version"`
	Features       []Feature `json:"features"`
}

// GetFeatures returns information about the available features.
func (a *AdminAPI) GetFeatures() (FeaturesResponse, error) {
	var features FeaturesResponse
	return features, a.sendAny(
		http.MethodGet,
		"/v1/features",
		nil,
		&features)
}
