## 2026-11-02 - Add Semantic Colors to CLI Status Output
**Learning:** In text-based interfaces like CLIs, color plays a critical role in conveying the state of a system quickly. A simple block of text can be hard to parse, but explicitly green/yellow/red text immediately indicates success, warning, or error states.
**Action:** When printing status indicators in CLI applications, explicitly prepend semantic color tags (e.g., using a library like `wsxiaoys/terminal/color`) to make the interface more scannable and accessible.
