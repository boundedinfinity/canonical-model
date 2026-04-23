package algorithm

// Capabilities: Sign, Certify, Encrypt
// GPG relevance: Default & most universally compatible.
// Key size:
// 		Minimum: 1024 bits (GnuPG enforces ≥768, but RSA options begin at 1024)
// 		Maximum: 4096 bits (commonly presented default max in GnuPG prompts)
// Notes:
// 		RSA is offered as: RSA, RSA‑E (encrypt‑only), RSA‑S (sign‑only).
// 		“RSA and RSA” is the recommended multi‑purpose default for new keys
