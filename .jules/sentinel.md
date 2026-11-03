## 2024-05-18 - Argument Injection in Command Execution
**Vulnerability:** Argument injection risk via space-joined command arguments in `gitExecWithDir`.
**Learning:** `gitExecWithDir` formerly accepted variadic string arguments, joined them with spaces, and split them by spaces again before passing to `exec.Command`. This makes any argument containing spaces break into multiple arguments, which allows malicious argument injection.
**Prevention:** Pass arguments to `exec.Command` as separate strings and use variadic slices `[]string` consistently for configurations containing complex commands. Avoid string manipulation (join and split) on commands.
