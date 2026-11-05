## 2024-05-18 - [CLI Output Accessibility]
**Learning:** This app uses `github.com/wsxiaoys/terminal/color` to inherently parse string tags into ANSI color codes. Prepending specific tags (`@r`, `@y`, `@g`) to messages before output improves scannability for users while isolating styling from base configuration. This was well received.
**Action:** Use string tags like `@r` and `@g` rather than raw ANSI codes when styling CLI output. Ensure they're prepended in the function layer, not the message config map, to preserve testability and separation of concerns.
