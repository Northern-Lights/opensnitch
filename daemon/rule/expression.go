package rule

import (
	"fmt"

	"github.com/Northern-Lights/os-rules-engine/network"
	"github.com/Northern-Lights/os-rules-engine/rules"
)

// An Expression can evaluate whether a connection matches
type Expression interface {
	Evaluate(*network.Connection) bool
}

// ExpressionSerializer is an Expression that can be serialized into a
// protocol buffer Expression
type ExpressionSerializer interface {
	Expression
	Serializer
}

// A Serializer can be serialized into a protocol buffer Expression
type Serializer interface {
	Serialize() *rules.Expression
}

// An ExpressionDeserializer deserializes a protobuf-based expression into an ExpressionSerializer
type ExpressionDeserializer func(*rules.Expression) (ExpressionSerializer, error)

// DeserializeExpression should be set to an implementation of ExpressionDeserializer
var DeserializeExpression ExpressionDeserializer = func(_ *rules.Expression) (_ ExpressionSerializer, err error) {
	err = fmt.Errorf(`rule: DeserializeExpression not set`)
	return
}
