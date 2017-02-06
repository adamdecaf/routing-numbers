# routing-numbers

A simple [routing transit number](https://en.wikipedia.org/wiki/Routing_transit_number) validation library. The algorithm is described there.

### Usage

Include `github.com/adamdecaf/routing-numbers` in your imports and call out to `routing.Valid(string)` for a `bool` back. You can call `routing.Check(string)` instead for an `error` type returned.

### Tests

The test file of routing numbers is downloaded from: https://www.frbservices.org/EPaymentsDirectory/fpddir.txt

### Contributions

1. Submit a PR with passing tests
1. It'll likely get merged

### License

Apache 2
