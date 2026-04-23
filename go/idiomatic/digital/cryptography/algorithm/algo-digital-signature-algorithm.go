package algorithm

// Capabilities: Sign, Certify
// Key size:
// 	Minimum: 512 bits (DSA1), though GnuPG enforces minimum 768 bits for keys broadly.
// 	Maximum: 1024 bits for original DSA, larger for DSA2 support (2048+ depending on SHA variant).
// GnuPG doc: “The size of a DSA key must be between 512 and 1024 bits.”
// Notes:
// 	DSA is sign‑only and is normally paired with ElGamal for encryption.
// 	Older and less common for new keys today.
