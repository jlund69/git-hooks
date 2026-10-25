## 2024-05-24 - Colorize CLI Output with terminal/color

**Learning:** The custom logger in `git-hooks` uses `github.com/wsxiaoys/terminal/color`. This library inherently parses string tags (e.g., `@g`, `@y`, `@r`) into ANSI color codes natively when printed.
**Action:** When adding color to CLI output, we can prepend these exact color tags to the message strings passed to the logger (e.g., `@y + message`), but we must remember that test environments capturing these logs will store the exact tag prefixes and thus assertions must match them verbatim.
