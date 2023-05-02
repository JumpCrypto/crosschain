package aptos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CrosschainTestSuite struct {
	suite.Suite
	Ctx context.Context
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Ctx = context.Background()
}

func TestAptosTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}
