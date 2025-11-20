# makeqr — Roadmap

This roadmap summarizes the planned features and the GitHub issues created for them.

## High Priority

- #2 Support alternative input types (text, vCard, WiFi, file, STDIN)
- #3 Add support for SVG/JPEG output and stdout writing
- #4 Add foreground/background color, margin, and scale options
- #5 Add clipboard copy and open-in-default-viewer options
- #6 Add GenerateBytes API and support writing to stdout
- #7 Add structured logging and verbosity flags (--verbose, --quiet)
- #8 Add cross-platform CI matrix and automatic releases
- #9 Add golden file tests for new formats and features (SVG, logos, colors)
- #10 Improve README with examples and 'using as library' guidance
- #11 Add --version and improved --help support
- #12 Add input validation and size limits for CLI flags

## Medium Priority

- #13 Add batch generation: CSV/JSON input and templated outputs
- #14 Add logo overlay/center image support for PNG/SVG
- #15 Add text labels/captions under or over the QR code
- #16 Add advanced payload helpers (WiFi, vCard, geo, calendar event)
- #17 Add CLI output metadata (file size, content type, checksum)
- #18 Add an interactive TUI mode for previewing and saving QR codes
- #19 Add QR layout templates and multi-QR printable pages (PDF)
- #20 Add terminal preview (ASCII or base64) for quick viewing
- #21 Add security hardening: sanitize and rate-limit inputs for server mode
- #22 Add benchmarks and coverage reporting to CI
- #23 Add a Dockerfile and GitHub Action for reusable demo and action
- #24 Add HTTP server mode to serve generated QR codes
- #25 Create a GitHub Action to generate QR images for README or releases

## Low Priority / Advanced

- #26 Support arbitrary binary payloads (file input)
- #27 Add structured append support (split payload across multiple QR codes)
- #28 Add optional cryptographic signing and validation of payloads
- #29 Add optional dynamic short-link generation + analytics endpoints (SaaS integration)
- #30 Add a cross-platform GUI (desktop) for interactive QR generation
- #31 Publish as Homebrew formula or package for easy install
- #32 Add shell auto-completions (zsh, bash, fish)
- #33 Create a VS Code extension to generate QR from selected text
- #34 Add vulnerability scanning and dependency checks in CI
- #35 Add machine-readable JSON output flag (--json) with metadata
- #36 Add i18n/localization support for CLI help and messages
- #37 Add automated dependency maintenance and go mod tidy in CI
- #38 Create a project Roadmap and label/issue templates

