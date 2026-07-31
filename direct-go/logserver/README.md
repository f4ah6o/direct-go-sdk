# Direct debug log server

The Direct SDK emits metadata-only diagnostics by default. Safe diagnostics may
include method names, event types, byte counts, durations, collection sizes,
and stable one-way hashes of user, talk, domain, or message identifiers.
Message bodies, raw protocol maps, access tokens, authorization headers, and
secret values are not emitted by the SDK's normal or verbose diagnostic paths.

Payload tracing is an unsafe troubleshooting mode. It must be enabled
explicitly with `debuglog.EnableUnsafePayloadTracing()` (or by setting
`DIRECT_DEBUG_UNSAFE_PAYLOADS=true` before process startup), and raw output must
use `UnsafePrintf`, `UnsafeVerbose`, or `UnsafeWriter`. The logger emits a
warning when that mode is enabled. Do not use it with production accounts or
with a log server reachable by other users.

The HTTP server has no authentication, authorization, or encryption of its own.
It accepts log entries at `POST /log` and exposes collected entries through
`GET /logs` and `GET /stream`. Bind it to localhost during development or put
it behind an authenticated, TLS-terminating proxy. The server trusts clients
to use the safe logger; application logs, webhook payloads, MCP tool results,
and data returned to user callbacks are outside this guarantee and must be
handled by the application.
