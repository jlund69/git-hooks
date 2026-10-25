## 2024-05-24 - Fix command injection vulnerability in gitExec

**Vulnerability:** The `gitExecWithDir` function in `util.go` was taking a variadic `args ...string`, joining them with spaces, and then splitting them by spaces before passing them to `exec.Command("git", args...)`. This allows for command injection if an argument contains spaces or if arguments are carefully crafted, as any input containing spaces would be evaluated as separate arguments to `git`, bypassing the intended argument boundaries.

**Learning:** This existed because of an attempt to easily pass multiple arguments as a single string (e.g. `gitExec("config --global init.templatedir")`), which creates a pattern where argument parsing is done insecurely instead of passing each argument explicitly as a separate string element. The codebase heavily relied on `GIT` config map having single strings with multiple arguments instead of string slices.

**Prevention:** Always pass arguments to `exec.Command` and wrapper functions as explicit separate strings (string slices) rather than joining and re-splitting. Update configurations (like the `GIT` map) to store string slices representing individual arguments.
