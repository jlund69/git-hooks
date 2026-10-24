## 2024-05-24 - [Adding color tags to CLI output]
**Learning:** Using terminal color libraries (like `github.com/wsxiaoys/terminal/color`) makes status messages in CLI output much easier to scan. By prepending color tags directly to output messages instead of configuration dictionaries, the separation of logic is maintained while significantly improving UX.
**Action:** Always consider the visual hierarchy and scannability of text-based terminal outputs. Use appropriate colors (e.g. green for success, yellow/red for warnings/errors) to draw attention and guide the user intuitively.
