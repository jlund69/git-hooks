## 2024-05-24 - [Command Injection via String Splitting]
**Vulnerability:** `exec.Command` was executed with arguments joined and then split by spaces (`strings.Split(strings.Join(args, " "), " ")`), leading to potential command/argument injection.
**Learning:** This allowed space-separated inputs in variables like `homeTemplate` to be treated as additional commands/arguments. It was caused by attempting to generalize command execution without considering injection.
**Prevention:** Always use multiple string arguments or slices for `exec.Command` rather than formatting a single string. Maintain argument separation. Also be careful not to append to a global variable slice to avoid mutability side-effects.
