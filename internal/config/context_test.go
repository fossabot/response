package config_test

import (
	"os"
	"testing"

	"github.com/responserms/response/internal/config"
	"github.com/stretchr/testify/suite"
	"github.com/zclconf/go-cty/cty"
)

type ContextSuite struct {
	suite.Suite
}

var expectedFunctionsCount = 1

func (t *ContextSuite) TestThatCreateContextReturnsFunctions() {
	resp := config.CreateContext()

	t.Assert().Equal(len(resp.Functions), 1)
}

func (t *ContextSuite) TestThatEnvFunctionReturnsValidEnvValue() {
	ctx := config.CreateContext()
	os.Setenv("TEST_ENV_WORKS", "YES")

	env := ctx.Functions["env"]

	// should properly return YES becasue TEST_ENV_WORKS was set above to "YES"
	yes, err := env.Call([]cty.Value{cty.StringVal("TEST_ENV_WORKS"), cty.StringVal("DEFAULT")})
	if err != nil {
		t.Fail("env() on yes returned an error!")
	}

	t.Assert().Equal(yes.AsString(), "YES")

	// should properly return DEFAULT because the DOES_NOT_EXIST environment variable does not exist. heh.
	def, err := env.Call([]cty.Value{cty.StringVal("DOES_NOT_EXIST"), cty.StringVal("DEFAULT")})
	if err != nil {
		t.Fail("env() on default returned an error")
	}

	t.Assert().Equal(def.AsString(), "DEFAULT")
}

func TestContextSuite(t *testing.T) {
	suite.Run(t, &ContextSuite{})
}
