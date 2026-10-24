## 2024-10-24 - Argument Injection in Command Execution
**Vulnerability:** Argument and command injection vulnerability in `gitExecWithDir` function due to using `strings.Join(args, " ")` and then `strings.Split(..., " ")` on input arguments which could allow a malicious user to inject arbitrary shell arguments into the `git` command execution.
**Learning:** Joining arguments into a single string and resplitting by spaces defeats the security benefits of using a slice for variadic arguments in `exec.Command`, leading to injection flaws.
**Prevention:** Pass arguments as separate strings (using variadic functions and slices) directly to `exec.Command` and refrain from joining and splitting argument lists before execution.
