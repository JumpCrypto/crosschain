package config

import (
	"errors"
	"os"
	"testing"

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
