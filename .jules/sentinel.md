## 2024-05-18 - [Command Injection Risk]
**Vulnerability:** `gitExecWithDir` function splits arguments by space and joins them, potentially leading to command injection if spaces are provided inside individual arguments. Passing `args ...string` as separate items is best practice to avoid injection.
**Learning:** In Go, splitting an already formed string to `exec.Command` ignores OS-level argument escaping and quoting rules.
**Prevention:** Pass `args ...string` without doing `strings.Split(strings.Join(args, " "), " ")`.
