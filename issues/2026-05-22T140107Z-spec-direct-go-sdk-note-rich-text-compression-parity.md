# Implement note rich text and compression parity

## Background

The note-management baseline wrappers intentionally preserve compressed note content exactly as returned. The current direct-go implementation does not guarantee direct-js parity for gzip compression, decompression, rich-text note creation, or rich-text note update behavior.

Source context: direct-go-sdk note-management porting completion on 2026-05-22.

## Scope

- Study direct-js note create/update rich-text payload construction.
- Identify the exact compression format and encoding boundaries used by Direct4B.
- Add direct-go helpers only after the raw wire behavior is understood.
- Preserve existing raw note APIs for callers that need exact payload passthrough.

## Acceptance criteria

- [ ] Note content compression and decompression behavior is documented with fixture evidence.
- [ ] `create_note` and `update_note` parity work is scoped separately from raw note read/delete/lock APIs.
- [ ] Tests cover compressed content round trips without requiring live credentials.
- [ ] Public APIs make raw passthrough versus high-level rich-text behavior explicit.
