## 2024-05-15 - Command Injection via Incorrect Argument Joining
**Vulnerability:** Found a vulnerability in `util.go`'s `gitExecWithDir` where the variadic arguments `args` were unsafely joined by spaces and then split again before being passed to `exec.Command`. This could lead to argument or command injection if an argument originally contained a space (e.g., in a directory path or user input).
**Learning:** Joining and splitting arguments defeats the protection that `exec.Command` provides when separating commands from their arguments.
**Prevention:** Always pass arguments to `exec.Command` as distinct elements in the slice without joining them.
