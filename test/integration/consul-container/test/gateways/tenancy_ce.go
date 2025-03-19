//go:build !consulent
// +build !consulent

package gateways

import (
	"testing"

	"github.com/arenadata/consul/api"
)

func getOrCreateNamespace(_ *testing.T, _ *api.Client) string {
	return ""
}
