## 3.35.1

This is a security rebuild of the v3 provider following v3.35.0, the final
regular v3 release.

DEPENDENCIES:

* provider: build with Go 1.27.3, which distrusts a set of intermediate certificate authorities that have been revoked for CA compromise. The distrust list is compiled into the binary, so affected chains are rejected regardless of the host's system trust store.

Upgrade to v3.35.1 to receive the rebuilt provider. The module's Go 1.20 directive
and dependencies are unchanged.

Windows ARM32 binaries are no longer available because the release compiler does
not support that target. Windows ARM64, AMD64, and 386 binaries remain available.
