## 2024-05-14 - Fix CLI output color spacing
**Learning:** Added color tags properly to logger output by prefixing the color tag (e.g. '@c') onto the string directly, instead of adding it as a separate argument. This prevents unwanted spaces injected by terminal/color print functions.
**Action:** Always combine the color tag string with the first argument to avoid unintended spaces when passing multiple arguments to color.Print.
