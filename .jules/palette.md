## 2024-05-24 - CLI Empty States
**Learning:** Empty states in CLI output are critical for user confidence. Without them, a user might think a command failed silently or missed something. The existing `list` command output a header without any contents if no hooks were configured.
**Action:** Always provide explicit textual feedback when iterating over an empty collection or when a command succeeds but has no data to display (e.g., "(No items found)").
