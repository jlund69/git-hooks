## 2024-05-24 - Command Injection in Git Exec
**Vulnerability:** Command injection vulnerability due to improperly sanitizing `args` in `gitExecWithDir` where `strings.Split(strings.Join(args, " "), " ")` was used.
**Learning:** This approach flattened all arguments and split by spaces, allowing attackers to inject additional Git arguments or shell commands if the input arguments contained spaces. Variadic functions taking `...string` should pass the elements directly to `exec.Command` rather than rebuilding a string.
**Prevention:** Avoid modifying the variadic `args` before passing it to `exec.Command("git", args...)`. Maintain argument separation to prevent command injection.
