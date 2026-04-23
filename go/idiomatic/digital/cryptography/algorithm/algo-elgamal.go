package algorithm

// Capabilities: Encrypt (not used for signing)
// Key size:
// 	Minimum: 768 bits (GnuPG does not allow lower).
// 	Maximum: Often up to 4096 bits depending on build, though GnuPG prompts don’t explicitly cap it.
// Notes:
// 	Appears as “DSA and ElGamal” when generating a classic PGP‑style key.
// 	Slower, produces large ciphertexts, but still supported.
