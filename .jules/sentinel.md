## 2024-05-17 - Command Injection via `strings.Join`

**Vulnerability:** The `gitExecWithDir` and `gitExec` functions took a variadic list of strings, joined them with spaces, and then split them again by spaces before passing them to `exec.Command`. This could allow command injection or argument injection if a user-supplied string contained spaces. Furthermore, `GIT` map values in `config.go` contained spaces, necessitating a change to `map[string][]string`.

**Learning:** Shell-like parsing of commands should be avoided. The `exec.Command` function in Go is designed to accept individual arguments safely, bypassing the shell. Joining and splitting strings reintroduces shell-like parsing vulnerabilities.

**Prevention:** Always pass arguments to `exec.Command` as individual strings in a slice or variadic arguments, without any intermediate joining and splitting. Ensure configuration values containing commands and arguments are stored as slices of strings rather than single space-separated strings.
