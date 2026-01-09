package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_concurryCmdRunfunc_defaults(t *testing.T) {

	type (
		Args struct {
			SourceTypeName string
			OutputFileName string
			OutputTypeName string
		}
		Expected struct {
			OutputFileName string
			OutputTypeName string
		}
		Test struct {
			Name     string
			Args     Args
			Expected Expected
		}
	)

	tests := []Test{
		{
			Name: "source: SourceType",
			Args: Args{
				SourceTypeName: "SourceType",
			},
			Expected: Expected{
				OutputFileName: "source_type_with_context.go",
				OutputTypeName: "SourceTypeWithContext",
			},
		},
		{
			Name: "source: someType",
			Args: Args{
				SourceTypeName: "someType",
			},
			Expected: Expected{
				OutputFileName: "some_type_with_context.go",
				OutputTypeName: "someTypeWithContext",
			},
		},
		{
			Name: "source: ABC",
			Args: Args{
				SourceTypeName: "ABC",
			},
			Expected: Expected{
				OutputFileName: "abc_with_context.go",
				OutputTypeName: "ABCWithContext",
			},
		},
		{
			Name: "source: IDThing",
			Args: Args{
				SourceTypeName: "IDThing",
			},
			Expected: Expected{
				OutputFileName: "id_thing_with_context.go",
				OutputTypeName: "IDThingWithContext",
			},
		},
		{
			Name: "source: CrAzyyyWOORD123a1b2n3n4A8S7DFGH",
			Args: Args{
				SourceTypeName: "CrAzyyyWOORD123a1b2n3n4A8S7DFGH",
			},
			Expected: Expected{
				OutputFileName: "cr_azyyy_woord_123_a_1_b_2_n_3_n_4_a_8_s_7_dfgh_with_context.go",
				OutputTypeName: "CrAzyyyWOORD123a1b2n3n4A8S7DFGHWithContext",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sourceTypeName = test.Args.SourceTypeName
			outputFileName = test.Args.OutputFileName
			outputTypeName = test.Args.OutputTypeName

			concurryCmdRunfunc(nil, nil)

			assert.Equal(t, test.Expected.OutputFileName, normalizedOutputFileName, "unexpected output file name generated")
			assert.Equal(t, test.Expected.OutputTypeName, normalizedOutputTypeName, "unexpected output type name generated")
		})
	}
}
