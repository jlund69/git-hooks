## 2024-05-15 - Command Injection via gitExec arguments

**Vulnerability:** The function `gitExecWithDir` and the calls to it join variadic string arguments using `strings.Join` and then split them using `strings.Split`, or callers pass single string containing spaces that gets split into separate arguments implicitly when passed to `exec.Command`. This can lead to command and argument injection if user input is passed this way.
**Learning:** The memory mentions: "When calling `gitExec` or `exec.Command` in this codebase, always pass arguments as separate strings (e.g., using a slice or variadic arguments) rather than joining them into a single space-separated string to prevent argument and command injection vulnerabilities."
**Prevention:** Always pass the strings straight through directly to `exec.Command` without joining and splitting.
