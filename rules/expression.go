package rules

import (
	"github.com/evilsocket/opensnitch/network"
)

// An Evaluator can evaluate whether a connection matches
type Evaluator interface {
	Evaluate(*network.Connection) bool
}

// EvaluatorSerializer is an Evaluator that can be serialized into a
// protocol buffer Expression
type EvaluatorSerializer interface {
	Evaluator
	Serialize() *Expression
}

// An ExpressionDeserializer deserializes a protobuf-based expression into an EvaluatorSerializer
type ExpressionDeserializer func(*Expression) (EvaluatorSerializer, error)
