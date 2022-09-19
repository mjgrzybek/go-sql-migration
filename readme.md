# SQL schema migration

## Raw sql
Pros:
- sql queries can be unit tested (without go code)
- no new dependencies

Cons:
- Requires sqlite3 specific knowledge to make it right
- Logic is maintained by us

## gorm
https://github.com/go-gorm/gorm \
30k ⭐

Pros:
- go-native API (structs, queries)
- SQL details hidden from API user
    - unless specific cases must be handled

Cons:
- Maybe an overkill for our single table
- How to test it? SQL seems more straightforward and separated from go
  - e2e tests?


## golang-migrate
https://github.com/golang-migrate/migrate \
10k ⭐

Pros:
- provides API for up- and down- migrations
- manages transactions automatically
- SQL queries can be tested without testing operator