# khqr-learn

A from-scratch Go project for learning the Bakong KHQR standard before
tackling the real payment feature at work. Every function is stubbed
with `panic("not implemented")` — your job is to fill them in, in order.

## Why this order

Each milestone builds on the last, and each one has a way to verify
you got it right *before* moving on. Don't skip the verification step —
that's where the actual learning happens.

## Milestones

1. **TLV encoder** (`internal/khqr/tlv.go`)
   Implement `tlv()` and `nestedTLV()`. Test by hand-encoding a field
   from the NBC "KHQR Content Guideline v1.4" PDF and comparing byte-for-byte.
   Docs: https://api-bakong.nbc.gov.kh/document

2. **CRC16-CCITT** (`internal/khqr/crc16.go`)
   Implement the checksum. Verify against a known CRC-16/CCITT-FALSE
   test vector (search that exact term — it's a standard, not
   Bakong-specific, so plenty of reference values exist).

3. **KHQR generator** (`internal/khqr/generator.go`)
   Combine 1 + 2 into `GenerateIndividual()`. Verify by comparing your
   output against:
   - The worked example in the NBC guideline PDF
   - github.com/chhunneng/bakong-khqr generating the same inputs
   Then do `GenerateMerchant()`.

4. **MD5 hashing** (`internal/khqr/hash.go`)
   Already implemented — it's a one-liner. Read it, understand it,
   move on.

5. **Bakong sandbox client** (`internal/bakong/client.go`)
   Register for a dev token: https://api-bakong.nbc.gov.kh/register/
   Implement `CheckTransactionByMD5` and `CheckAccountByID` against
   https://sit-api-bakong.nbc.gov.kh

6. **Smart polling** (`internal/bakong/poll.go`)
   Implement backoff polling so you're not hammering the API every
   second. This same "don't hammer a resource" instinct applies to
   your Quartz-based Favorite Schedule jobs at work.

7. **Cross-check** (no file — this is a verification step)
   Same inputs into your generator and into `chhunneng/bakong-khqr`
   (Go) should produce identical QR strings. If they match, your
   implementation is spec-correct.

## Running it

```bash
go run ./cmd
```

Will panic until you implement the functions it calls — that's expected,
work through the milestones in order.

## Reference docs

- KHQR Content Guideline v1.4: https://api-bakong.nbc.gov.kh/document
- QR Payment Integration PDF: same location
- Official Go SDK (for cross-checking, not copying):
  https://github.com/chhunneng/bakong-khqr
- Sandbox base URL: https://sit-api-bakong.nbc.gov.kh
- Production base URL: https://api-bakong.nbc.gov.kh
