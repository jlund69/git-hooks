## 2023-10-10 - CLI Color Formatting UX
**Learning:** Using `fmt.Println`-style terminal loggers like `github.com/wsxiaoys/terminal/color` can introduce unwanted spaces if color tags are passed as separate arguments (e.g. `color.Print("@c", msg)` results in `@c msg`).
**Action:** Always prepend color tags directly to the message string (e.g. `"@c" + msg`) to ensure clean spacing and formatting in CLI outputs.
