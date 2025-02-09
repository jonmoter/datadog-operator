// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2021 Datadog, Inc.

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		version    string
		minVersion string
		expected   bool
	}{
		{
			version:    "7.27.0",
			minVersion: "7.26.0-rc.5",
			expected:   true,
		},
		{
			version:    "7.29.0-rc.5",
			minVersion: "7.27.0-0",
			expected:   true,
		},
		{
			version:    "7.30.0",
			minVersion: "7.27.1-0",
			expected:   true,
		},
		{
			version:    "7.25.0",
			minVersion: "7.27.1-0",
			expected:   false,
		},
		{
			version:    "6.27.0",
			minVersion: "6.28.0-0",
			expected:   false,
		},
		{
			version:    "6.28.1",
			minVersion: "6.28.0-0",
			expected:   true,
		},
		{
			version:    "6.27.0",
			minVersion: "6.28.0",
			expected:   false,
		},
		{
			version:    "foobar",
			minVersion: "7.28.0",
			expected:   false,
		},
		{
			version:    "1.23.4",
			minVersion: "1.22-0",
			expected:   true,
		},
	}
	for _, test := range testCases {
		t.Run(test.version, func(t *testing.T) {
			result := IsAboveMinVersion(test.version, test.minVersion)
			assert.Equal(t, test.expected, result)
		})
	}
}
