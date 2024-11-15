package customerror_test

import (
	"errors"
	"testing"

	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/suite"
	"github.com/swallowstalker/online-book-store/modules/bookstore/customerror"
)

type CustomErrorTestSuite struct {
	suite.Suite
}

func TestWrapperRepo(t *testing.T) {
	suite.Run(t, new(CustomErrorTestSuite))
}

func (s *CustomErrorTestSuite) TestIsErrNotFound() {
	s.Run("error is not gogox error", func() {
		result := customerror.IsErrNotFound(errors.New("not gogox"))
		s.Assert().False(result)
	})

	s.Run("error is gogox and not code not found", func() {
		result := customerror.IsErrNotFound(errorx.ErrInternal("internal"))
		s.Assert().False(result)
	})

	s.Run("error is gogox and code not found", func() {
		result := customerror.IsErrNotFound(errorx.ErrNotFound("not found"))
		s.Assert().True(result)
	})
}
