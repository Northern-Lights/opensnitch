package rule

import (
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
	Serialize() *rules.Expression
}

// An ExpressionDeserializer deserializes a protobuf-based expression into an ExpressionSerializer
type ExpressionDeserializer func(*rules.Expression) (ExpressionSerializer, error)
