## 2024-05-18 - Fix Command Injection Vulnerability in gitExecWithDir
**Vulnerability:** A command injection vulnerability existed because arguments provided to `gitExec` functions were concatenated into a single string separated by spaces, which was then split and passed to `exec.Command`. This could allow attackers to inject arbitrary arguments and potentially commands by carefully crafting input strings.
**Learning:** Concatenating arguments into a single string for shell execution is extremely dangerous because it bypasses the argument separation provided by `exec.Command`.
**Prevention:** Always pass arguments to `exec.Command` as separate elements in a slice. Avoid joining arguments and relying on manual string splitting, especially with input originating from external sources.
