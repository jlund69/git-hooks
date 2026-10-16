## 2024-06-25 - Using Colors for CLI Status Indication
**Learning:** CLI status messages (e.g., error, warning, success, info) can be significantly improved by using color tags. This provides immediate visual feedback to the user without them needing to read the entire string to determine the context.
**Action:** Use color encoding (`@r`, `@y`, `@g`, etc.) consistently when logging output for system status updates in CLI applications. Ensure tests that assert log outputs match exact format strings including these color tags.
