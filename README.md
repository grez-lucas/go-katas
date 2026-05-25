# go-katas

TDD practice in Go. One module, one package per kata. Each kata is driven
entirely by tests — write a failing test, make it pass, refactor, repeat.

## Running

```sh
go test ./...            # run every kata
go test ./fizzbuzz/      # run one kata
go test ./... -v         # verbose, see each test case
go test ./... -run Add   # run tests matching a name
```

Handy while practicing a single kata:

```sh
go test ./stringcalculator/ -v        # rerun on demand
```

## The TDD loop

1. **Red** — write the smallest test for the next bit of behavior; watch it fail.
2. **Green** — write the least code that makes it pass.
3. **Refactor** — clean up with the tests as your safety net.

## Katas

| Kata               | Status        | Notes                                              |
| ------------------ | ------------- | -------------------------------------------------- |
| `fizzbuzz`         | ✅ reference   | Fully worked example of the structure/conventions. |
| `stringcalculator` | 🔴 your turn  | Classic kata. Rules are in `string_calculator.go`. |

`stringcalculator` ships with one intentionally failing test — that's your
starting point. `go test ./...` will be red until you implement it; that's TDD
working as intended.

## Adding a new kata

```sh
mkdir bowling
# create bowling/bowling.go        -> package bowling
# create bowling/bowling_test.go   -> package bowling
go test ./bowling/
```

Ideas to try next: Bowling Game, Roman Numerals, Prime Factors, Bank Account,
Gilded Rose, Mars Rover.

## Syncing across machines

This repo lives on GitHub. On the other machine:

```sh
git clone https://github.com/grez-lucas/go-katas.git
```

Then the usual `git pull` before a session and `git push` after.
