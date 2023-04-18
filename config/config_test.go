package config

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	vault "github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/suite"
)

type CrosschainTestSuite struct {
	suite.Suite
}

func (s *CrosschainTestSuite) SetupTest() {
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}

func (s *CrosschainTestSuite) TestRequireConfig() {
	require := s.Require()
	xcConfig := RequireConfig("crosschain")
	require.NotNil(xcConfig)
	require.NotNil(xcConfig["chains"])
}

func (s *CrosschainTestSuite) TestRequireConfigErr() {
	require := s.Require()
	xcConfig := RequireConfig("crosschainINVALID")
	require.Equal(xcConfig, map[string]interface{}{})
}

func (s *CrosschainTestSuite) TestGetSecretEnv() {
	require := s.Require()
	os.Setenv("XCTEST", "mysecret")
	secret, err := GetSecret("env:XCTEST")
	os.Unsetenv("XCTEST")
	require.Equal("mysecret", secret)
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestGetSecretFile() {
	require := s.Require()
	secret, err := GetSecret("file:../LICENSE")
	require.Contains(secret, "Apache License")
	require.Nil(err)
}

func (s *CrosschainTestSuite) TestGetSecretFileHomeErrFileNotFound() {
	require := s.Require()
	secret, err := GetSecret("file:~/config-in-home")
	require.Equal("", secret)
	require.Error(err)
}

func (s *CrosschainTestSuite) TestGetSecretErrFileNotFound() {
	require := s.Require()
	secret, err := GetSecret("file:../LICENSEinvalid")
	require.Equal("", secret)
	require.Error(err)
}

func (s *CrosschainTestSuite) TestGetSecretErrNoColon() {
	require := s.Require()
	secret, err := GetSecret("invalid")
	require.Equal("", secret)
	require.Error(errors.New("invalid secret source for: ***"), err)
}

func (s *CrosschainTestSuite) TestGetSecretErrInvalidType() {
	require := s.Require()
	secret, err := GetSecret("invalid:value")
	require.Equal("", secret)
	require.Error(errors.New("invalid secret source for: ***"), err)
}

type MockedVaultLoaded struct {
	data map[string]interface{}
}

var _ VaultLoader = &MockedVaultLoaded{}

func (l *MockedVaultLoaded) LoadSecretData(path string) (*vault.Secret, error) {
	data, ok := l.data[path]
	if !ok {
		return &vault.Secret{}, errors.New("path not found")
	}
	return &vault.Secret{
		Data: data.(map[string]interface{}),
	}, nil
}

func (s *CrosschainTestSuite) TestGetSecretVault() {
	require := s.Require()
	NewVaultClient = func(cfg *vault.Config) (VaultLoader, error) {
		vaultRes := `{
			"path1/to": {
				"data": {
					"secret": "mysecret"
				}
			},
			"path2/to": {
				"data": {
					"secret2": "mysecret2"
				}
			}
		}`
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(vaultRes), &data)
		require.NoError(err)

		return &MockedVaultLoaded{
			data: data,
		}, nil
	}

	_, err := GetSecret("vault:wrong_args")
	require.ErrorContains(err, "vault secret has 2 comma separated arguments")
	_, err = GetSecret("vault:wrong_args,aaa,bbb")
	require.ErrorContains(err, "vault secret has 2 comma separated arguments")

	_, err = GetSecret("vault:url,aaa")
	require.ErrorContains(err, "malformed vault secret")

	_, err = GetSecret("vault:url,aaa/secret")
	require.EqualError(err, "path not found")

	secret, err := GetSecret("vault:https://example.com,path1/to/secret")
	require.NoError(err)
	require.Equal("mysecret", secret)

	secret, err = GetSecret("vault:https://example.com,path2/to/secret2")
	require.NoError(err)
	require.Equal("mysecret2", secret)

	secret, err = GetSecret("vault:https://example.com,path2/to/secret_none")
	require.NoError(err)
	require.Equal("", secret)
}
