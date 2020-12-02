package config

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

var contextFunctions = map[string]function.Function{
	"env": function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name:         "key",
				Type:         cty.String,
				AllowNull:    false,
				AllowUnknown: false,
			},
			{
				Name:      "default",
				Type:      cty.String,
				AllowNull: true,
			},
		},
		VarParam: nil,
		Type:     function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var key = args[0].AsString()
			var def = args[1]

			if val := os.Getenv(key); val != "" {
				return cty.StringVal(val), nil
			}

			return def, nil
		},
	}),
}

// CreateContext creates a context to be used when writing the configuration file.
func CreateContext() *hcl.EvalContext {
	return &hcl.EvalContext{
		Functions: contextFunctions,
	}
}
