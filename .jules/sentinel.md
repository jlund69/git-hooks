## 2023-10-06 - Prevent Argument Injection in gitExec Calls
**Vulnerability:** Argument injection via split-string subcommands in `gitExecWithDir` and global `exec.Command` string-joining.
**Learning:** `exec.Command` handles argument separating correctly when elements are slices. If arguments are passed dynamically as a single string and then split (e.g., `strings.Split(strings.Join(args, " "), " ")`), any string containing a space acts as multiple arguments instead of a single quoted one.
**Prevention:** Store commands and parameters as slices (e.g., `[]string{"config", "--global", "init.templatedir"}`). Avoid reconstructing slices from space-joined strings. Let `exec.Command` handle string formatting internally.
