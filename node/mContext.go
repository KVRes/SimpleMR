package node

import "github.com/KVRes/SimpleMR/types"

type MapContext struct {
	m types.Intermediate
}

func NewMapContext() *MapContext {
	return &MapContext{
		m: make(map[string][]any),
	}
}

func (c *MapContext) Emit(key string, value any) {
	c.m[key] = append(c.m[key], value)
}

func (c *MapContext) Get(key string) []any {
	return c.m[key]
}

func (c *MapContext) Clear() {
	c.m = make(map[string][]any)
}
