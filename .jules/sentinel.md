## 2024-05-18 - [Command Injection via `strings.Join` and `strings.Split`]
**Vulnerability:** Command injection vulnerability in `gitExecWithDir` caused by insecure argument splitting: `args = strings.Split(strings.Join(args, " "), " ")`.
**Learning:** Joining and then splitting `exec.Command` arguments on spaces bypasses Go's built-in argument escaping, allowing malicious input to be injected as additional commands or arguments.
**Prevention:** Pass arguments to `exec.Command` exactly as they are constructed, using slice unpacking (`args...`) without attempting to reconstruct the command as a single string.
