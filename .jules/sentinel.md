## 2024-05-18 - Argument Injection in gitExecWithDir
**Vulnerability:** Command arguments were being unsafely tokenized via `strings.Split(strings.Join(args, " "), " ")` before being passed to `exec.Command`.
**Learning:** This pattern destroys argument boundaries and creates an argument injection vulnerability where user-supplied inputs or file paths with spaces are split into separate command options.
**Prevention:** Always pass arguments natively as separate string slice elements (variadic arguments) to `exec.Command` without intermediate string joining and splitting.
