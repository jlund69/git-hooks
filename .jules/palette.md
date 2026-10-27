## 2024-10-27 - Color-coded CLI list output
**Learning:** Adding visual hierarchy via color codes (e.g., green for installed, yellow for not installed, red for errors, cyan for headers) drastically improves scanability in terminal applications where large blocks of text are presented.
**Action:** Consistently use the terminal color library to apply a semantic color palette to CLI output (using inline tags like @c, @g, @r, @y) instead of relying solely on spacing or uncolored text, while making sure test assertions reflect the literal formatted strings with the color tags included.
