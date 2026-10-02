## 2024-05-18 - Argument Injection in Command Execution
**Vulnerability:** Argument injection vulnerability in `gitExecWithDir` where arguments were joined by spaces and then split by spaces before being passed to `exec.Command`.
**Learning:** This approach causes arguments containing spaces to be improperly split into multiple arguments, allowing attackers to inject arbitrary arguments into the command (e.g., git options like `--upload-pack` in clone).
**Prevention:** Always pass arguments to `exec.Command` as separate strings in a slice, rather than joining and splitting them. Avoid using `strings.Join` and `strings.Split` for command arguments.
