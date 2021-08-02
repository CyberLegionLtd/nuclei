// Package context provides context describing helpers
// over Operator based Matchers And Extractors.
package context

import (
	"strconv"
	"strings"
)

// Context is a structure that provides context around operators
//
//
type Context struct {
	chains   []*Context // sub-contexts for this context
	result   bool
	reason   string
	metadata map[string]struct{}
}

// NewContext returns a new context structure
func NewContext() *Context {
	return &Context{}
}

// GetResult returns the result boolean for a context
func (c *Context) GetResult() bool {
	return c.result
}

// SetResult sets the final result for a context structure with a reason string
func (c *Context) SetResult(result bool, reason string) *Context {
	c.result = result
	c.reason = reason
	return c
}

// SetMetadata adds metadata to the context
func (c *Context) SetMetadata(metadata map[string]struct{}) *Context {
	c.metadata = metadata
	return c
}

// AddSubContext adds a subcontext to the context
func (c *Context) AddSubContext(subcontext *Context) *Context {
	c.chains = append(c.chains, subcontext)
	return c
}

// String returns the string representation of a context
func (c *Context) String() string {
	builder := &strings.Builder{}
	builder.Grow(len(c.reason))
	builder.WriteString(strconv.FormatBool(c.result))
	builder.WriteString("| ")
	builder.WriteString(c.reason)

	for _, subctx := range c.chains {
		builder.WriteString("\n\t")
		builder.WriteString(subctx.String())
	}
	return builder.String()
}
