## 2023-10-18 - Command Injection Risk in Command Argument Splitting

**Vulnerability:** The `gitExecWithDir` function originally used `strings.Split(strings.Join(args, " "), " ")` to parse command line arguments before executing them using `exec.Command`. This could lead to a command injection risk where spaces in arguments, intended to be parsed as single values (e.g. paths with spaces), are incorrectly split.
**Learning:** Joining arguments with spaces and splitting them blindly undermines the protection provided by using a slice of strings in `exec.Command`, which passes arguments directly to the underlying OS and handles escaping automatically.
**Prevention:** Instead of manipulating arguments with string functions, store and construct multi-word arguments like commands using string slices (e.g., `[]string{"config", "--get", "hooks.global"}`). If dynamic building of arrays is required, use `append()` and `append(...slice)`.
