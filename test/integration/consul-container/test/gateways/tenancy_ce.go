//go:build !consulent
// +build !consulent

package gateways

import (
	"testing"

	"github.com/shulutkov/yellow-pages/api"
)

func getOrCreateNamespace(_ *testing.T, _ *api.Client) string {
	return ""
}
