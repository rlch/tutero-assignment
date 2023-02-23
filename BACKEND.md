# Backend Development Guide

Go GraphQL API skeleton using gqlgen for the Circuit Builder application.

## Skeleton Overview

Complete GraphQL schema and server setup provided. Includes service interfaces and entity models, but core business logic implementation is required.

> [!NOTE]
> **What's Provided:**
> - Complete GraphQL schema definition with all types and operations
> - Basic Go server with CORS middleware for web client support
> - Service interfaces defining all required operations
> - Entity models for circuits, nodes, edges, and evaluation
> - gqlgen configuration and code generation setup
> - Task-based build automation with hot reload (Air)

> [!WARNING]
> **What's Missing:**
> - Service layer implementation (core business logic)
> - Database integration and persistence layer
> - Circuit evaluation algorithm
> - GraphQL resolver implementations

## Technology Choice

> [!IMPORTANT]
> **Recommended**: Use the provided Go + GraphQL skeleton for rapid development and focus on solving complex circuit logic.

> [!TIP]
> **Alternative**: Implement with your preferred stack:
> - **Languages**: Node.js, Python, Rust, Java, C#, etc.
> - **APIs**: REST, gRPC, WebSocket, etc.
> - **Databases**: PostgreSQL, MySQL, SQLite, MongoDB, etc.
> 
> The GraphQL schema can serve as your API specification regardless of implementation technology.

## Project Structure

```
backend/
├── cmd/main.go              # Server entry point with CORS setup
├── graph/                   # GraphQL layer
│   ├── schema.graphql      # Complete schema definition
│   ├── schema.resolvers.go # Resolver stubs (to implement)
│   ├── resolver.go         # Resolver root configuration
│   └── generated.go        # Generated GraphQL code
├── internal/               # Business logic (to implement)
│   ├── entity/            # Domain models (provided)
│   └── service/           # Service interfaces (to implement)
├── gqlgen.yml             # gqlgen configuration
├── Taskfile.yml           # Build automation
└── .air.toml             # Hot reload configuration
```

## Development Setup

### Prerequisites
- Go 1.21+
- Task runner: `go install github.com/go-task/task/v3/cmd/task@latest`

### Commands
```bash
go mod download     # Install dependencies
task dev           # Start server with hot reload on :8080
task generate      # Regenerate GraphQL code after schema changes
task build         # Build production binary
```

## GraphQL Schema

The complete schema is defined in `graph/schema.graphql` with:

**Core Types:**
- `Circuit` - Container with nodes and edges
- `Node` interface - Base for all node types (Input, Output, AND, OR, NOT, Circuit)
- `Edge` - Connections between nodes  
- `EvaluationResult` - Circuit simulation results

**Operations:**
- **Queries**: `circuits`, `circuit(id)`, `evaluateCircuit`
- **Mutations**: Create circuits, nodes, and edges

## Service Layer Architecture

Service interfaces are defined in `internal/service/contract.go`:

```go
type CircuitService interface {
    // Circuit operations
    CreateCircuit(title string) (*entity.Circuit, error)
    GetCircuit(id string) (*entity.Circuit, error)
    GetAllCircuits() ([]*entity.Circuit, error)

    // Node operations  
    CreateInputNode(circuitID string, title string) (*entity.InputNode, error)
    // ... other node creation methods

    // Edge operations
    CreateEdge(circuitID string, sourceNodeID string, targetNodeID string) (*entity.Edge, error)

    // Evaluation (core challenge)
    EvaluateCircuit(circuit *entity.Circuit, inputs []*entity.InputNodeValue) (*entity.EvaluationResult, error)
}
```

> [!WARNING]
> **Current Status**: Interface defined, implementation required in `internal/service/impl.go`.

## Development Workflow

### Schema-First Development
1. Schema is already complete in `graph/schema.graphql`
2. Run `task generate` to update generated code
3. Implement resolver methods in `graph/schema.resolvers.go`
4. Implement business logic in service layer
5. Test with GraphQL playground at http://localhost:8080/

### Code Generation
```bash
task generate
# Runs: go run -mod=mod github.com/99designs/gqlgen generate
```

This updates:
- `graph/generated.go` - Core GraphQL server
- `graph/model/models_gen.go` - Type definitions
- Resolver stub methods

### Hot Reload Development
```bash
task dev  # Uses Air for automatic reloading
```

Air watches `**/*.go` and `graph/**/*.graphql` files for changes.

## Implementation Tips

> [!CAUTION]
> ### Circuit Evaluation Algorithm
> The core challenge is implementing boolean logic evaluation:
> 
> ```go
> func (s *circuitServiceImpl) EvaluateCircuit(circuit *entity.Circuit, inputs []*entity.InputNodeValue) (*entity.EvaluationResult, error) {
>     // 1. Build dependency graph from nodes and edges
>     // 2. Apply input values to InputNodes
>     // 3. Propagate values through logic gates (AND, OR, NOT)
>     // 4. Handle CircuitNode references (nested circuits)
>     // 5. Detect cycles and handle gracefully
>     // 6. Return output values from OutputNodes
> }
> ```
> 
> **Key Considerations:**
> - Topological sorting for evaluation order
> - Cycle detection in circuit graphs
> - Recursive evaluation for nested CircuitNodes
> - Error handling for invalid circuits

> [!TIP]
> ### Database Integration
> Recommended approach:
> - Use PostgreSQL or SQLite for persistence
> - Implement repository pattern in `internal/repository/`
> - Use database/sql or an ORM like GORM
> - Add database connection to service implementations

### Error Handling
```go
func (r *queryResolver) Circuit(ctx context.Context, id string) (*entity.Circuit, error) {
    circuit, err := r.CircuitService.GetCircuit(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get circuit %s: %w", id, err)
    }
    return circuit, nil
}
```

> [!TIP]
> ### Testing Strategy
> - Unit tests for service layer business logic
> - Integration tests for GraphQL resolvers
> - Use testify for assertions and mocking
> - Test circuit evaluation with complex scenarios

## Database Schema Example

```sql
CREATE TABLE circuits (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE nodes (
    id UUID PRIMARY KEY,
    circuit_id UUID REFERENCES circuits(id),
    node_type TEXT NOT NULL, -- 'input', 'output', 'and', 'or', 'not', 'circuit'
    title TEXT,              -- For input/output nodes
    referenced_circuit_id UUID REFERENCES circuits(id) -- For circuit nodes
);

CREATE TABLE edges (
    id UUID PRIMARY KEY,
    circuit_id UUID REFERENCES circuits(id),
    source_node_id UUID REFERENCES nodes(id),
    target_node_id UUID REFERENCES nodes(id)
);
```

## Configuration

### Environment Variables
```bash
PORT=8080                    # Server port
DATABASE_URL=postgres://...  # Database connection
LOG_LEVEL=info              # Logging level
```

### gqlgen Configuration
The `gqlgen.yml` file is pre-configured for:
- Schema location (`graph/schema.graphql`)
- Generated code output paths
- Resolver layout and naming

## Troubleshooting

> [!CAUTION]
> ### Code Generation Issues
> ```bash
> # Clear generated files and rebuild
> rm -rf graph/generated.go graph/model/
> task generate
> ```

> [!WARNING]
> ### Import/Dependency Issues
> ```bash
> go mod tidy
> go mod download
> ```

> [!CAUTION]
> ### CORS Problems
> The server includes CORS middleware in `cmd/main.go`. Verify:
> - Middleware is applied to `/query` endpoint  
> - Allowed origins, methods, and headers are correct
> - Browser network tab shows successful OPTIONS requests

> [!TIP]
> ### Performance Considerations
> - Use DataLoader pattern to prevent N+1 queries
> - Cache frequently accessed circuits
> - Optimize database queries with proper indexing
> - Consider circuit evaluation caching for complex circuits

## References

- [gqlgen Documentation](https://gqlgen.com/) - GraphQL server generation
- [GraphQL Specification](https://spec.graphql.org/) - API standard
- [Air Hot Reload](https://github.com/cosmtrek/air) - Development tool
- [Go Project Layout](https://github.com/golang-standards/project-layout) - Structure patterns

---

Focus on implementing the circuit evaluation algorithm and service layer for a working GraphQL API.