// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ca

import (
	"fmt"
	"os"
	"testing"

	"github.com/shulutkov/yellow-pages/agent/structs"
	"github.com/stretchr/testify/require"
)

func TestVaultCAProvider_JwtAuthClient(t *testing.T) {
	tokenF, err := os.CreateTemp("", "token-path")
	require.NoError(t, err)
	defer func() { os.Remove(tokenF.Name()) }()
	_, err = tokenF.WriteString("test-token")
	require.NoError(t, err)
	err = tokenF.Close()
	require.NoError(t, err)

	cases := map[string]struct {
		authMethod *structs.VaultAuthMethod
		expData    map[string]any
		expErr     error
	}{
		"base-case": {
			authMethod: &structs.VaultAuthMethod{
				Type: "jwt",
				Params: map[string]any{
					"role": "test-role",
					"path": tokenF.Name(),
				},
			},
			expData: map[string]any{
				"role": "test-role",
				"jwt":  "test-token",
			},
		},
		"no-role": {
			authMethod: &structs.VaultAuthMethod{
				Type:   "jwt",
				Params: map[string]any{},
			},
			expErr: fmt.Errorf("missing 'role' value"),
		},
		"no-path": {
			authMethod: &structs.VaultAuthMethod{
				Type: "jwt",
				Params: map[string]any{
					"role": "test-role",
				},
			},
			expErr: fmt.Errorf("missing 'path' value"),
		},
		"no-path-but-jwt": {
			authMethod: &structs.VaultAuthMethod{
				Type: "jwt",
				Params: map[string]any{
					"role": "test-role",
					"jwt":  "test-jwt",
				},
			},
			expData: map[string]any{
				"role": "test-role",
				"jwt":  "test-jwt",
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			auth, err := NewJwtAuthClient(c.authMethod)
			if c.expErr != nil {
				require.EqualError(t, c.expErr, err.Error())
				return
			}
			require.NoError(t, err)
			if auth.LoginDataGen != nil {
				data, err := auth.LoginDataGen(c.authMethod)
				require.NoError(t, err)
				require.Equal(t, c.expData, data)
			}
		})
	}
}

func TestVaultCAProvider_K8sAuthClient(t *testing.T) {
	tokenF, err := os.CreateTemp("", "token-path")
	require.NoError(t, err)
	defer func() { os.Remove(tokenF.Name()) }()
	_, err = tokenF.WriteString("test-token")
	require.NoError(t, err)
	err = tokenF.Close()
	require.NoError(t, err)

	cases := map[string]struct {
		authMethod *structs.VaultAuthMethod
		expData    map[string]any
		expErr     error
	}{
		"base-case": {
			authMethod: &structs.VaultAuthMethod{
				Type: "kubernetes",
				Params: map[string]any{
					"role":       "test-role",
					"token_path": tokenF.Name(),
				},
			},
			expData: map[string]any{
				"role": "test-role",
				"jwt":  "test-token",
			},
		},
		"legacy-case": {
			authMethod: &structs.VaultAuthMethod{
				Type: "kubernetes",
				Params: map[string]any{
					"role": "test-role",
					"jwt":  "test-token",
				},
			},
			expData: map[string]any{
				"role": "test-role",
				"jwt":  "test-token",
			},
		},
		"no-role": {
			authMethod: &structs.VaultAuthMethod{
				Type:   "kubernetes",
				Params: map[string]any{},
			},
			expErr: fmt.Errorf("missing 'role' value"),
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			auth, err := NewK8sAuthClient(c.authMethod)
			if c.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, c.expErr, err.Error())
				return
			}
			require.NoError(t, err)
			if auth.LoginDataGen != nil {
				data, err := auth.LoginDataGen(c.authMethod)
				require.NoError(t, err)
				require.Equal(t, c.expData, data)
			}
		})
	}
}

func TestVaultCAProvider_AppRoleAuthClient(t *testing.T) {
	roleID, secretID := "test_role_id", "test_secret_id"

	roleFd, err := os.CreateTemp("", "role")
	require.NoError(t, err)
	_, err = roleFd.WriteString(roleID)
	require.NoError(t, err)
	err = roleFd.Close()
	require.NoError(t, err)

	secretFd, err := os.CreateTemp("", "secret")
	require.NoError(t, err)
	_, err = secretFd.WriteString(secretID)
	require.NoError(t, err)
	err = secretFd.Close()
	require.NoError(t, err)

	roleIdPath := roleFd.Name()
	secretIdPath := secretFd.Name()

	defer func() {
		os.Remove(secretFd.Name())
		os.Remove(roleFd.Name())
	}()

	cases := map[string]struct {
		authMethod *structs.VaultAuthMethod
		expData    map[string]any
		expErr     error
	}{
		"base-case": {
			authMethod: &structs.VaultAuthMethod{
				Type: "approle",
				Params: map[string]any{
					"role_id_file_path":   roleIdPath,
					"secret_id_file_path": secretIdPath,
				},
			},
			expData: map[string]any{
				"role_id":   roleID,
				"secret_id": secretID,
			},
		},
		"optional-secret-left-out": {
			authMethod: &structs.VaultAuthMethod{
				Type: "approle",
				Params: map[string]any{
					"role_id_file_path": roleIdPath,
				},
			},
			expData: map[string]any{
				"role_id": roleID,
			},
		},
		"missing-role-id-file-path": {
			authMethod: &structs.VaultAuthMethod{
				Type:   "approle",
				Params: map[string]any{},
			},
			expErr: fmt.Errorf("missing '%s' value", "role_id_file_path"),
		},
		"legacy-direct-values": {
			authMethod: &structs.VaultAuthMethod{
				Type: "approle",
				Params: map[string]any{
					"role_id":   "test-role",
					"secret_id": "test-secret",
				},
			},
			expData: map[string]any{
				"role_id":   "test-role",
				"secret_id": "test-secret",
			},
		},
	}

	for k, c := range cases {
		t.Run(k, func(t *testing.T) {
			auth, err := NewAppRoleAuthClient(c.authMethod)
			if c.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, c.expErr, err.Error())
				return
			}
			require.NoError(t, err)
			if auth.LoginDataGen != nil {
				data, err := auth.LoginDataGen(c.authMethod)
				require.NoError(t, err)
				require.Equal(t, c.expData, data)
			}
		})
	}
}
