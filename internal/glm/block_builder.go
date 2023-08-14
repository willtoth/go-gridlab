package glm

import (
	"fmt"
	"strings"
)

type BlockType string

const (
	ObjectType BlockType = "object"
	ClockType  BlockType = "clock"
	ModuleType BlockType = "module"
)

type BlockBuilder struct {
	// Key value pairs
	options   []string
	blockType BlockType
	name      string
}

func NewBlockBuilder(name string, blockType BlockType) *BlockBuilder {
	return &BlockBuilder{
		name:      name,
		blockType: blockType,
		options:   make([]string, 0),
	}
}

// Add an option to a block and return self
func (b *BlockBuilder) AddOption(name, value string) *BlockBuilder {
	b.options = append(b.options, fmt.Sprintf("%s %s", name, value))
	return b
}

// Add a struct to a block with glm tags
func (b *BlockBuilder) AddStruct(s interface{}) *BlockBuilder {
	fields, _ := parseGLMTags(s)
	b.options = append(b.options, fields...)
	return b
}

// Output an object as "object name { options }" or "object name;"
func (b *BlockBuilder) Build() string {
	if len(b.options) == 0 {
		return fmt.Sprintf("%s %s;", b.blockType, b.name)
	} else {
		var builder strings.Builder

		builder.WriteString(fmt.Sprintf("%s %s {\n", b.blockType, b.name))
		for _, value := range b.options {
			builder.WriteString(fmt.Sprintf("\t%s\n", value))
		}
		builder.WriteString("}")
		return builder.String()
	}
}
