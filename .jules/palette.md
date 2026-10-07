## 2024-05-24 - CLI Output Styling
**Learning:** In terminal applications, text output with colors and emojis improves user experience by providing clear visual hierarchy and feedback. However, custom logger implementations may capture the exact arguments (including color tags) passed to them in test mode. Adding styling strings changes the structure of arguments logged and can break assertions in unit tests.
**Action:** Always update tests that make exact string matches on log outputs when adding color tags or emojis to ensure tests remain reliable.
