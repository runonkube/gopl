# GOPL 4.11 — GitHub issues CLI

**Extracted to its own repository:** https://github.com/runonkube/issues-cli

The solution to GOPL exercise 4.11 (a command-line tool for creating, reading,
updating, and closing GitHub issues) grew beyond a book exercise into a small
standalone CLI. It now lives as its own Go module with:

- Subcommand dispatch (`create` / `list` / `show` / `update` / `close`)
- `$EDITOR` integration via `--edit` (VISUAL → EDITOR → vi/notepad fallback)
- HTTP client parameterisable at construction for `httptest`-based tests
- Partial updates via `omitempty`
- GitHub Actions CI (test + `golangci-lint` + `govulncheck`)
- Apache 2.0 license

Local path (development): `~/Documents/projects/issues-cli/`

Extraction date: 2026-07-13.
