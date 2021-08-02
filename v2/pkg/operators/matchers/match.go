package matchers

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/projectdiscovery/nuclei/v2/pkg/operators/common/context"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolstate"
)

// MatchStatusCode matches a status code check against a corpus
func (m *Matcher) MatchStatusCode(statusCode int) *context.Context {
	// Iterate over all the status codes accepted as valid
	//
	// Status codes don't support AND conditions.
	for _, status := range m.Status {
		// Continue if the status codes don't match
		if statusCode != status {
			continue
		}
		// Return on the first match.
		return m.contextWithReason(true, strconv.Itoa(status))
	}
	return m.contextWithReason(false, "")
}

// MatchSize matches a size check against a corpus
func (m *Matcher) MatchSize(length int) *context.Context {
	// Iterate over all the sizes accepted as valid
	//
	// Sizes codes don't support AND conditions.
	for _, size := range m.Size {
		// Continue if the size doesn't match
		if length != size {
			continue
		}
		// Return on the first match.
		return m.contextWithReason(true, strconv.Itoa(size))
	}
	return m.contextWithReason(false, "")
}

// MatchWords matches a word check against a corpus.
func (m *Matcher) MatchWords(corpus string) *context.Context {
	ctx := context.NewContext()

	// Iterate over all the words accepted as valid
	for i, word := range m.Words {
		// Continue if the word doesn't match
		if !strings.Contains(corpus, word) {
			// If we are in an AND request and a match failed,
			// return false as the AND condition fails on any single mismatch.
			if m.condition == ANDCondition {
				return ctx.AddSubContext(m.contextWithReason(false, word)).SetResult(false, "AND condition did not match")
			}
			// Continue with the flow since its an OR Condition.
			continue
		}
		ctx.AddSubContext(m.contextWithReason(true, word))

		// If the condition was an OR, return on the first match.
		if m.condition == ORCondition {
			return ctx.SetResult(true, "OR Condition matched")
		}

		// If we are at the end of the words, return with true
		if len(m.Words)-1 == i {
			return ctx.SetResult(true, "AND Condition matched")
		}
	}
	return m.contextWithReason(false, "")
}

// MatchRegex matches a regex check against a corpus
func (m *Matcher) MatchRegex(corpus string) bool {
	// Iterate over all the regexes accepted as valid
	for i, regex := range m.regexCompiled {
		// Continue if the regex doesn't match
		if !regex.MatchString(corpus) {
			// If we are in an AND request and a match failed,
			// return false as the AND condition fails on any single mismatch.
			if m.condition == ANDCondition {
				return false
			}
			// Continue with the flow since its an OR Condition.
			continue
		}

		// If the condition was an OR, return on the first match.
		if m.condition == ORCondition {
			return true
		}

		// If we are at the end of the regex, return with true
		if len(m.regexCompiled)-1 == i {
			return true
		}
	}
	return false
}

// MatchBinary matches a binary check against a corpus
func (m *Matcher) MatchBinary(corpus string) bool {
	// Iterate over all the words accepted as valid
	for i, binary := range m.Binary {
		// Continue if the word doesn't match
		hexa, _ := hex.DecodeString(binary)
		if !strings.Contains(corpus, string(hexa)) {
			// If we are in an AND request and a match failed,
			// return false as the AND condition fails on any single mismatch.
			if m.condition == ANDCondition {
				return false
			}
			// Continue with the flow since its an OR Condition.
			continue
		}

		// If the condition was an OR, return on the first match.
		if m.condition == ORCondition {
			return true
		}

		// If we are at the end of the words, return with true
		if len(m.Binary)-1 == i {
			return true
		}
	}
	return false
}

// MatchDSL matches on a generic map result
func (m *Matcher) MatchDSL(data map[string]interface{}) bool {
	// Iterate over all the expressions accepted as valid
	for i, expression := range m.dslCompiled {
		result, err := expression.Evaluate(data)
		if err != nil {
			continue
		}

		var bResult bool
		bResult, ok := result.(bool)

		// Continue if the regex doesn't match
		if !ok || !bResult {
			// If we are in an AND request and a match failed,
			// return false as the AND condition fails on any single mismatch.
			if m.condition == ANDCondition {
				return false
			}
			// Continue with the flow since its an OR Condition.
			continue
		}

		// If the condition was an OR, return on the first match.
		if m.condition == ORCondition {
			return true
		}

		// If we are at the end of the dsl, return with true
		if len(m.dslCompiled)-1 == i {
			return true
		}
	}
	return false
}

// contextWithReason returns a condensed context reason structure for a matcher
func (m *Matcher) contextWithReason(matched bool, value string) *context.Context {
	if !protocolstate.IsDebug() && !matched {
		return context.NewContext().SetResult(false, "")
	}
	builder := &strings.Builder{}

	if m.Name != "" {
		builder.WriteString(m.Name)
		builder.WriteString(" ")
	}
	builder.WriteString(m.Type)
	builder.WriteString(" matcher ")

	ctx := context.NewContext()
	if matched {
		builder.WriteString(value)
		builder.WriteString(" matched")
	} else {
		builder.WriteString(" did not match")
	}
	ctx.SetResult(matched, builder.String())
	return ctx
}
