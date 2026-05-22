# Add live Direct4B integration coverage

## Background

Mock-server tests verify RPC method names, argument order, response decoding, and error propagation. They do not prove that every implemented wrapper succeeds against the real Direct4B service with real account permissions and current production behavior.

Source context: direct-go-sdk porting coverage completion on 2026-05-22.

## Scope

- Define a credential-safe live integration test plan.
- Use runtime secret injection rather than printing or storing tokens.
- Start with read-only or low-risk RPCs before mutating Direct4B resources.
- Mark destructive or account-control flows as opt-in.

## Acceptance criteria

- [ ] Integration tests can run only when explicit environment variables are present.
- [ ] No secret values are printed, committed, or required in test fixtures.
- [ ] A documented subset of baseline RPCs is validated against live Direct4B.
- [ ] Unsupported or permission-dependent live cases are listed separately from SDK defects.
