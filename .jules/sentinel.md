## 2024-05-18 - Argument Injection in Command Execution
**Vulnerability:** Argument injection via `strings.Split(strings.Join(args, " "), " ")` in the `gitExecWithDir` function.
**Learning:** Joining and then splitting variadic arguments on space destroys the boundaries of arguments containing spaces. It allows paths containing spaces or malicious inputs to inject unexpected options/arguments into git command executions.
**Prevention:** Always pass string slices directly to `exec.Command` variadic arguments and maintain them throughout any wrapper layers without converting back to space-delimited strings.
