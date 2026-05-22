# Verify edge case and notification parity

## Background

RPC wrapper coverage does not guarantee behavior parity for edge cases such as permission errors, pagination boundaries, retry behavior, context cancellation, websocket reconnects, or server notification event ordering.

Source context: direct-go-sdk porting coverage completion on 2026-05-22.

## Scope

- Compare direct-js and direct-go behavior for pagination markers, limits, and empty pages.
- Exercise permission-denied, not-found, validation, and server error responses.
- Review retry and reconnect behavior around notifications.
- Check event payload handling for baseline message and notification flows.

## Acceptance criteria

- [ ] Pagination boundary behavior is tested for representative APIs.
- [ ] RPC error mapping is documented and tested for common server error shapes.
- [ ] Notification and reconnect behavior has a parity checklist against direct-js.
- [ ] Any behavioral differences are documented as intentional or tracked for repair.
