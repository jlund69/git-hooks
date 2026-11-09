## 2024-05-14 - Injecting color tags for CLI terminal output using github.com/wsxiaoys/terminal/color
**Learning:** Terminal outputs can be formatted using color tags like `@c`, `@g`, `@y`, `@r` and `@w` by prepending them to the output string. The custom logger natively parses these before printing them out, which enhances the visual hierarchy and UX of CLI apps.
**Action:** When adding visual cues to terminal interfaces, always modify the printed string via prefix (e.g., `"@g" + msg`) rather than messing with the underlying logger settings, but be aware this impacts assertions in tests.
