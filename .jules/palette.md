## 2024-03-24 - CLI Output Styling
**Learning:** Adding color to CLI applications should separate the styling from the message configuration to ensure proper rendering and keep the core configurations unpolluted.
**Action:** Use specific color prefixes for different states (`@g` for success, `@r` for error, `@y` for warning, `@c` for informational headers) in the logging output, keeping the source strings intact.
