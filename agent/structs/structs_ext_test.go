// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package structs_test

import (
	"testing"

	"github.com/arenadata/consul/agent/structs"
	"github.com/arenadata/consul/proto/private/pbpeering"
	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeProto(t *testing.T) {
	arg := pbpeering.SecretsWriteRequest{
		PeerID: "bbd26d02-a831-47b6-8a20-d16a99f56def",
	}
	buf, err := structs.EncodeProto(structs.PeeringSecretsWriteType, &arg)
	require.NoError(t, err)

	var out pbpeering.SecretsWriteRequest

	err = structs.DecodeProto(buf[1:], &out)
	require.NoError(t, err)
	require.Equal(t, arg.PeerID, out.PeerID)
}
