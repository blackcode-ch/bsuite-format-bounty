# Formatting task

A standalone Go exercise: make a Unicode-aware truncation helper handle non-positive limits without panicking.

This package uses only the Go standard library. It needs no account, service, database, environment file, or private repository.

## Run

Use Go 1.26.2 or later.

```sh
go test ./...
go vet ./...
go run ./cmd/repro
```

The tests cover existing positive-limit behavior. The reproduction command currently panics for a zero limit. After the fix, it should print an empty quoted string and exit successfully.

## Intended behavior

- A limit of zero or less returns an empty string, for both empty and nonempty inputs.
- For a positive limit, return the input unchanged when it fits.
- Otherwise keep at most that many Unicode code points, including one trailing ellipsis.
- A limit of one returns the ellipsis when truncation is needed.
- Count Unicode code points (Go runes), not bytes or grapheme clusters.

## Submit

See the priced GitHub issue for the bounty and payment terms. Fork this repository, make a branch, and open a PR against main linking the issue. Include regression tests for zero and negative limits, including the minimum integer value, and preserve the existing positive-limit behavior. Run the commands above and link your public CI result.

Anyone may submit, including agent-operated contributors. The same acceptance tests and review apply. A comment or an early PR does not reserve the bounty; the first PR that passes review wins. Do not include personal data, credentials, or copied private material.

