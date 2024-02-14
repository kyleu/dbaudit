<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Releasing

DB Audit uses `goreleaser` to create build artifacts.

You can release to GitHub using `./bin/build/release.sh`, or test the release by running `./bin/build/release-test.sh`.

Your releases are available at https://github.com/kyleu/dbaudit/releases

### Checksums

All release binaries are checksummed, available in `checksums.txt` in the root of the release

### Changelog

A changelog will be created based on the commit history, including all authors and messages

### Docker Images

Multiple Docker images will be created. The main image is `ghcr.io/kyleu/dbaudit/x.x.x`, and a debug image is provided at `ghcr.io/kyleu/dbaudit/x.x.x-debug` that includes `delve` for debugging

### Homebrew

Packages for macOS and Linux will be pushed to Homebrew at `kyleu/homebrew-kyleu`

### Signing

Release binaries and the checksum file are signed using `gpg`

### Source Code

The source code will be bundled in the release, available as `dbaudit_x.x.x_source.zip`

### Universal Binaries

A universal macOS app will be created, containing the complete app for x86-64 and ARM

