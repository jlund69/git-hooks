## 2024-05-24 - CLI Color Feedback Pattern
**Learning:** Adding color tags (like `@g` for green, `@y` for yellow) to CLI output generation provides immediate visual feedback on system state and improves scannability for users without altering underlying configuration strings.
**Action:** Use color-tag prepending directly in logger calls (e.g., `logger.Infoln("@g" + msg)`) for status messages to enhance CLI UX, and ensure tests account for these exact string prefixes.
