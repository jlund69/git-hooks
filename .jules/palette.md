## 2024-11-01 - Add Colors to CLI Output

**Learning:** When developing a CLI tool, using colors can drastically improve UX. The `git-hooks` codebase uses `github.com/wsxiaoys/terminal/color` which allows styling output.
**Action:** Enhance user experience by outputting colorized text! Adding the tags `@g`, `@y`, `@r`, and `@c` to messages logged via `logger.Info`, `logger.Warn`, and `logger.Error` natively colorizes the output.
