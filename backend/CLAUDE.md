# Backend Development Notes

## Project Structure
- **Entities**: `internal/entity/` - Domain models (Circuit, Node types, Edge, evaluation types)
- **Services**: `internal/service/` - Business logic interface and stub implementation
- **GraphQL**: `graph/` - Schema, resolvers, and generated code
- **Main**: `cmd/main.go` - Application entry point

## Key Implementation Points

### Service Layer
- **Contract**: `internal/service/contract.go` - Interface defining all operations
- **Implementation**: `internal/service/impl.go` - Stub implementation returning "not implemented" errors
- **TODO**: Implement all service methods with actual business logic

### Circuit Evaluation
- Support AND, OR, NOT gates as primitives
- Other gates (XOR, NAND, NOR) can be built from primitives
- Handle nested circuits (CircuitNode type)
- Input validation and error handling

### GraphQL Schema
- **Node Interface**: All gate types implement Node interface
- **InputNodeValue**: For providing input values during evaluation
- **Simplified API**: No wrapper input types for simple operations

## Development Commands
```bash
task dev        # Hot reload development
task generate   # Regenerate GraphQL code after schema changes  
task build      # Build application
task test       # Run tests
```

## Database Integration
- Add database layer under `internal/database/` or `internal/repository/`
- Inject database dependencies into service implementation
- Support persistent storage of circuits and components

## Testing Strategy
- Unit tests for service layer business logic
- Integration tests for GraphQL resolvers
- Test circuit evaluation with known input/output pairs

## Air Configuration
- Watches `.go` and `.graphql` files
- Excludes test files and generated models
- Builds to `tmp/main` with hot reload