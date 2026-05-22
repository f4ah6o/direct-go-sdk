# Track daab-go feature parity separately from direct-go coverage

## Background

The direct-go porting coverage metric only covers the low-level Direct4B SDK RPC baseline. It does not guarantee that daab-go, its CLI commands, bot framework behavior, daemon mode, examples, or webhook integration are fully compatible with the upstream daab JavaScript implementation.

Source context: direct-go-sdk porting coverage completion on 2026-05-22.

## Scope

- Build a daab-go parity checklist against upstream daab.
- Separate CLI parity, bot framework parity, webhook behavior, examples, and daemon behavior.
- Identify behavior already covered by tests versus behavior only smoke-tested manually.
- Keep direct-go RPC wrapper coverage out of the daab-go completion metric.

## Acceptance criteria

- [ ] daab-go has its own parity report or checklist.
- [ ] CLI commands have coverage for expected arguments, generated files, and failure modes.
- [ ] Bot framework message handling and response helpers are tested against representative events.
- [ ] Remaining upstream daab incompatibilities are tracked as explicit follow-up issues.
