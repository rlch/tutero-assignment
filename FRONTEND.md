# Frontend Development Guide

Flutter web application skeleton with Ferry GraphQL client.

## Skeleton Overview

Basic Flutter web app structure with Ferry GraphQL integration configured. Provides foundation for building circuit editor interface.

> [!NOTE]
> **What's Provided:**
> - Flutter app with Material 3 theme
> - Ferry GraphQL client setup with code generation
> - Basic GraphQL operations (circuits query)
> - Sample UI component showing GraphQL integration
> - Task-based build automation

> [!WARNING]
> **What's Missing:**
> - Circuit editor UI components
> - Canvas-based drag-and-drop interface  
> - Component palette for logic gates
> - Circuit visualization and interaction logic

## Project Structure

```
frontend/
├── lib/
│   ├── main.dart                    # App entry point
│   ├── graphql_client.dart         # Ferry client setup
│   ├── test_query.dart             # Sample GraphQL integration
│   ├── src/graphql/                # GraphQL operations
│   │   ├── circuits.graphql        # Circuit queries
│   │   ├── mutations.graphql       # Circuit mutations
│   │   └── __generated__/          # Generated Ferry code
│   └── schema.graphql              # Backend schema copy
├── pubspec.yaml                    # Dependencies
├── build.yaml                     # Ferry config
└── Taskfile.yml                   # Build automation
```

## Development Setup

### Prerequisites
- Flutter 3.8.1+
- Chrome browser
- Task runner

### Commands
```bash
flutter pub get      # Install dependencies
task dev            # Run Flutter web app
task generate       # Generate GraphQL code from schema
```

## Ferry GraphQL Integration

### Client Setup
The GraphQL client is configured in `lib/graphql_client.dart`:
```dart
final link = HttpLink("http://localhost:8080/query");
final cache = Cache(possibleTypes: possibleTypesMap);
final client = Client(link: link, cache: cache);
```

### Code Generation
Ferry automatically generates type-safe Dart code from GraphQL operations:
- Edit `.graphql` files in `lib/src/graphql/`
- Run `task generate` to update generated code
- Import generated request classes in your widgets

### Sample Usage
See `lib/test_query.dart` for example of using Ferry's Operation widget:
```dart
Operation(
  client: client,
  operationRequest: GGetCircuitsReq(),
  builder: (context, response, error) {
    // Handle loading, error, and data states
  },
)
```

## Development Tips

> [!TIP]
> ### Canvas-based UI
> For circuit editor, consider:
> - **CustomPainter** for drawing circuits and connections
> - **GestureDetector** for drag-and-drop interactions
> - **Transform** widgets for zoom/pan functionality
> - **Stack** layout for layering components

### State Management
- Start with StatefulWidget for local component state
- Consider Provider or Riverpod for complex app state
- Leverage Ferry's built-in caching for GraphQL data

> [!IMPORTANT]
> ### Flutter Web Considerations
> - Use `flutter run -d chrome --web-renderer html` for debugging
> - Test with different screen sizes and browsers
> - Be mindful of web-specific constraints (no native file system, etc.)

> [!TIP]
> ### Performance
> - Use `RepaintBoundary` around canvas components
> - Implement efficient hit-testing for interactive elements
> - Consider virtualization for large circuits

## Common Patterns

### GraphQL Mutations
```dart
// Execute mutation
final request = GCreateCircuitReq((b) => b..vars.title = 'My Circuit');
client.request(request).listen((response) {
  // Handle response
});
```

### Error Handling
```dart
if (response?.hasErrors == true) {
  // Show user-friendly error message
  final error = response!.graphqlErrors!.first.message;
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(content: Text('Error: $error')),
  );
}
```

### Loading States
```dart
if (response?.loading == true) {
  return Center(child: CircularProgressIndicator());
}
```

## Troubleshooting

> [!CAUTION]
> ### Code Generation Issues
> ```bash
> # Clear generated files and rebuild
> rm -rf lib/__generated__ lib/src/graphql/__generated__
> task generate
> ```

> [!CAUTION]
> ### CORS Errors
> - Ensure backend server is running on port 8080
> - Check browser network tab for failed requests
> - Verify GraphQL endpoint URL in `graphql_client.dart`

> [!WARNING]
> ### Hot Reload Not Working
> ```bash
> # Full restart after generated code changes
> flutter clean
> flutter pub get
> flutter run -d chrome
> ```

> [!NOTE]
> ### Build Errors
> Common issues:
> - Missing imports after code generation
> - Outdated pubspec.lock file
> - Flutter SDK version compatibility

## References

- [Flutter Web Documentation](https://docs.flutter.dev/platform-integration/web)
- [Ferry GraphQL Client](https://ferrygraphql.com/)
- [Material 3 Design System](https://m3.material.io/)
- [CustomPainter Tutorial](https://docs.flutter.dev/cookbook/animation/draw-a-circle)

---

Focus on implementing the circuit editor canvas and connecting it to the GraphQL mutations for a complete circuit building experience.