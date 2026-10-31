## 2023-10-31 - Fix command injection in `gitExec`
**Vulnerability:** The `gitExec` and `gitExecWithDir` functions were vulnerable to command and argument injection. They took a variadic slice of strings `args ...string`, concatenated them into a single string, split them by space, and passed them to `exec.Command`.
**Learning:** Concatenating arguments with space and then splitting them can lead to argument injection if any of the individual strings contain spaces. It also bypasses the natural argument separation provided by `exec.Command` which correctly escapes arguments.
**Prevention:** Pass arguments to `exec.Command` exactly as they are provided without any intermediate string concatenation and splitting.
