# Test Runner Specification

## Overview
The test runner is a flexible tool that validates circuit simulator implementations by testing core functionality through HTTP requests to your backend service.

## Configuration
Create a `test-config.json` file to specify how the test runner should interact with your API:

```json
{
  "server_url": "http://localhost:8080",
  "endpoints": {
    "create_circuit": "/your-circuit-endpoint",
    "evaluate_circuit": "/your-evaluation-endpoint",
    "create_component": "/your-component-endpoint"
  },
  "request_format": "your-preferred-format",
  "timeout_ms": 5000
}
```

## Test Categories

### Basic Logic Gates
- AND, OR, NOT, XOR gate functionality
- Multi-input gates
- Chained gate combinations

### Circuit Evaluation
- Simple circuits with known outputs
- Complex circuits with multiple paths
- Edge cases and error conditions

### Component System
- Component creation and reuse
- Nested components
- Component library functionality

### Performance
- Evaluation speed benchmarks
- Memory usage validation
- Concurrent request handling

## Usage
```bash
./test-runner --config test-config.json --suite all
./test-runner --config test-config.json --suite basic
./test-runner --config test-config.json --test specific-test-name
```

## Output
Test results are provided in multiple formats:
- Console output with pass/fail status
- JSON report for automated processing
- Performance metrics and timing data

## Adaptation
The test runner adapts to your API design by:
- Using your specified endpoints and request formats
- Parsing responses according to your schema
- Validating core functionality regardless of implementation details