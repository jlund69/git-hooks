## 2024-05-24 - [Argument Injection in Command Execution]
**Vulnerability:** `util.go:gitExecWithDir` was re-splitting variadic arguments by joining them with spaces and splitting them again (`args = strings.Split(strings.Join(args, " "), " ")`). This allowed argument injection if a directory path, git URL, or branch name contained spaces.
**Learning:** Variadic arguments passed to `exec.Command` must be passed natively as distinct strings.
**Prevention:** Never join and split argument arrays for shell commands. Always pass `args...` directly to `exec.Command`. For maps holding command bases like `GIT` in `config.go`, use `map[string][]string` rather than `map[string]string` with spaces.
