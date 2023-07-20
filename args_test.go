package sqlz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgs(t *testing.T) {
	var args Args
	s := `
		1:` + args.Add("ADD") + `
		2:` + args.Add("TRUE", true) + `
		3:` + args.Add("FALSE", false) + `
		4:` + args.Addf(`one=$?`, "ONE") + `
		5:` + args.Addf(`two=$?,$?`, "TWO") + `
		6:` + args.Addf(`true=$?`, "TRUE", true) + `
		7:` + args.Addf(`false=$?`, "FALSE", false) + `
	`
	require.Equal(t, `
		1:$1
		2:$2
		3:
		4:one=$3
		5:two=$4,$4
		6:true=$5
		7:
	`, s)
	require.Equal(t, Args{"ADD", "TRUE", "ONE", "TWO", "TRUE"}, args)
}
