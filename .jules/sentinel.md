## 2026-10-21 - Command Injection via args.Split
**Vulnerability:** Argument splitting via `strings.Split(strings.Join(args, " "), " ")` allows argument and command injection.
**Learning:** Never parse user input this way. Pass array explicitly instead.
**Prevention:** Always use slices or variadic arguments and pass directly to `exec.Command`.
