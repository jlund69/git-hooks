## 2024-11-07 - [Colorizing CLI Outputs]
**Learning:** In CLI apps, plain text output can be hard to parse visually. Using color tags natively supported by the logger (e.g., `@g` for success, `@y` for warnings, `@r` for errors) gives immediate visual feedback on the state, improving scanability.
**Action:** Always prepend appropriate color tags in the output generation logic (instead of config files) to keep separation of concerns while improving user experience.
