# direct-js Source Code

This directory contains the deminified source code from [lisb/direct-js](https://github.com/lisb/direct-js).

## File Information

- **Source**: https://github.com/lisb/direct-js
- **File**: lib/direct-node.min.js
- **Last synced**: 2025-12-10
- **Commit SHA**: ab1d017f3c0b3481fe40914ae64e6ebcfd7ec161

## Processing

The minified code is automatically:
1. Downloaded from the upstream repository
2. Formatted using `prettier` for readability
3. Committed to this repository

## Purpose

This deminified source is used as a reference for porting functionality to the Go implementation
in this repository. The formatted JavaScript code makes it easier to:

- Understand the original implementation
- Reference API method signatures
- Port features accurately to Go
- Compare behavior between implementations

## Automatic Updates

This source is automatically synced daily via GitHub Actions workflow.
The workflow:
- Checks for updates to lisb/direct-js daily at 00:00 UTC
- Downloads and deminifies new versions when detected
- Runs coverage analysis to update porting status
- Commits changes automatically

See `.github/workflows/sync-direct-js.yml` for implementation details.

## Manual Sync

To manually trigger a sync:

```bash
# Using GitHub CLI
gh workflow run sync-direct-js.yml

# Or via GitHub UI
# Go to Actions → Sync direct-js Source → Run workflow
```

## Files

- `direct-node.js` - Deminified JavaScript source code
- `.last-sync-sha` - SHA of the last synced commit (used to detect updates)
- `README.md` - This file
