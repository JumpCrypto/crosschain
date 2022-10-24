package crosschain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

const MockTime = 12345

type CrosschainTestSuite struct {
	suite.Suite
	Ctx context.Context
}

func (s *CrosschainTestSuite) SetupTest() {
	s.Ctx = context.Background()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CrosschainTestSuite))
}
