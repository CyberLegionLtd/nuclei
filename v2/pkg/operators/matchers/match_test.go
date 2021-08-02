package matchers

import (
	"fmt"
	"testing"

	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestANDCondition(t *testing.T) {
	protocolstate.Init(&types.Options{Debug: true})

	m := &Matcher{Name: "wordMatcher", condition: ANDCondition, Words: []string{"a", "b"}}

	matched := m.MatchWords("cd")
	//	require.True(t, matched.GetResult(), "Could not match valid AND condition")
	fmt.Printf("%v\n", matched)

	matched = m.MatchWords("b")
	require.False(t, matched.GetResult(), "Could match invalid AND condition")
}

func TestORCondition(t *testing.T) {
	m := &Matcher{condition: ORCondition, Words: []string{"a", "b"}}

	matched := m.MatchWords("a b")
	require.True(t, matched.GetResult(), "Could not match valid OR condition")

	matched = m.MatchWords("b")
	require.True(t, matched.GetResult(), "Could not match valid OR condition")

	matched = m.MatchWords("c")
	require.False(t, matched.GetResult(), "Could match invalid OR condition")
}

func TestHexEncoding(t *testing.T) {
	m := &Matcher{Encoding: "hex", Type: "word", Part: "body", Words: []string{"50494e47"}}
	err := m.CompileMatchers()
	require.Nil(t, err, "could not compile matcher")

	matched := m.MatchWords("PING")
	require.True(t, matched.GetResult(), "Could not match valid Hex condition")
}
