package service

import "backend/internal/entity"

type CircuitService interface {
	// Circuit operations
	CreateCircuit(title string) (*entity.Circuit, error)
	GetCircuit(id string) (*entity.Circuit, error)
	GetAllCircuits() ([]*entity.Circuit, error)

	// Node operations
	CreateInputNode(circuitID string, title string) (*entity.InputNode, error)
	CreateOutputNode(circuitID string, title string) (*entity.OutputNode, error)
	CreateAndNode(circuitID string) (*entity.AndNode, error)
	CreateOrNode(circuitID string) (*entity.OrNode, error)
	CreateNotNode(circuitID string) (*entity.NotNode, error)
	CreateCircuitNode(circuitID string, referencedCircuitID string) (*entity.CircuitNode, error)

	// Edge operations
	CreateEdge(circuitID string, sourceNodeID string, targetNodeID string) (*entity.Edge, error)

	// Evaluation operations
	EvaluateCircuit(circuit *entity.Circuit, inputs []*entity.InputNodeValue) (*entity.EvaluationResult, error)
}
