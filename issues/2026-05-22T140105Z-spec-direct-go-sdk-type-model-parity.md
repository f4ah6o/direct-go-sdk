# Define direct-go SDK type model parity

## Background

The direct-go porting coverage reached the current baseline RPC method count, but that only proves wrapper presence for the tracked RPC names. It does not guarantee that every public Go type matches direct-js semantics or that every nested result field has a stable typed model.

Source context: direct-go-sdk porting coverage completion on 2026-05-22.

## Scope

- Audit direct-js request and response shapes for implemented baseline RPCs.
- Compare the shapes against direct-go public structs and raw map wrappers.
- Decide where raw payloads should stay raw and where stable typed models should be added.
- Document any intentional API differences.

## Acceptance criteria

- [ ] Each baseline RPC has an explicit type-model status: typed, raw-by-design, or missing.
- [ ] Any raw-by-design APIs explain why a typed model is deferred.
- [ ] Public type changes are covered by mock-server tests.
- [ ] Compatibility notes are added to direct-go documentation or coverage reporting.
