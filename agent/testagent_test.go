// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package agent

import (
	"testing"

	"github.com/arenadata/consul/agent/config"
	"github.com/hashicorp/hcl"
)

// TestDefaultConfig triggers a data race in the HCL parser.
func TestDefaultConfig(t *testing.T) {
	for i := 0; i < 500; i++ {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			var c config.Config
			data := config.DefaultSource().(config.FileSource).Data
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
			hcl.Decode(&c, data)
		})
	}
}
