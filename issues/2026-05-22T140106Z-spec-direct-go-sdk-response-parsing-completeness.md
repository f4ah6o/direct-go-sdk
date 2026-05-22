# Verify direct-go SDK response parsing completeness

## Background

The porting coverage tool counts RPC wrappers, not exhaustive parsing behavior. A wrapper can exist and still fail to parse a legitimate Direct4B payload if optional fields, array variants, numeric types, or nested map shapes differ from the mock fixtures.

Source context: direct-go-sdk porting coverage completion on 2026-05-22.

## Scope

- Expand representative response fixtures beyond the happy-path mock payloads.
- Include optional fields, empty arrays, omitted fields, mixed integer widths, and nested maps.
- Check APIs that currently return conservative raw payload structs.
- Separate real parsing bugs from intentionally raw passthrough APIs.

## Acceptance criteria

- [ ] High-risk baseline RPCs have response parsing tests with realistic variants.
- [ ] Parsing failures produce actionable errors or preserve raw payloads as documented.
- [ ] Coverage reporting clearly distinguishes wrapper coverage from parsing completeness.
- [ ] Any discovered parser defects are fixed or tracked as separate implementation issues.
