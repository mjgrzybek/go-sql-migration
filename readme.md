# SQL schema migration

Reading: https://awesome-go.com/#database

## Raw sql
Pros:
- sql queries can be unit tested (without Go code)
- no new dependencies

Cons:
- Requires sqlite3 specific knowledge to make it idiomatic way
- Logic is maintained by us

## gorm
https://github.com/go-gorm/gorm \
**30k** 🌟 \
Automatically adds additional columns (`created_at`, `deleted_at`, `deleted_at`). \
TODO: check if it's idiomatic approach and use regardless we decide on gorm.

Pros:
- widely used, respected and maintained
- go-native API (structs, queries)
- SQL details hidden from API user
    - unless specific cases must be handled
- clean and shiny, pleasure to work with 

Cons:
- Maybe an overkill for our simple, single table
- How to test it? SQL seems more straightforward and separated from go
  - e2e tests? 
    - may be risky because some "sql cases" can be unintentionally omitted

Questions:
- migration from SQL to ORM may be tricky (go code must match with existing SQL code 1:1)
- how schema version is stored? -> [gormigrate](##gormigrate)


## golang-migrate
https://github.com/golang-migrate/migrate \
**10k** ⭐ \
Creates additional `schema_migrations` table for migrations tracking. 

Pros:
- provides API for up- and down- migrations
- manages transactions automatically
  - it's aware of sqlite3 specific logic
- SQL queries can be tested without testing operator
- existing code doesn't have to be refactored
- testing tools [available](https://github.com/golang-migrate/migrate/blob/master/database/testing/testing.go)

Cons:
- raw SQL must be maintained
- raw SQL is used in our Go code

Questions:
- do we want to work with raw SQL in our codebase? ORM is more convenient
- `Unlike other migrate database drivers, the sqlite3 driver will automatically wrap each migration in an implicit transaction by default. Migrations must not contain explicit BEGIN or COMMIT statements. This behavior may change in a future major release.` \
[src](https://github.com/golang-migrate/migrate/tree/master/database/sqlite3)

## goose
https://github.com/pressly/goose \
**3k** ⭐ 

## sql-migrate
https://github.com/rubenv/sql-migrate \
**2.6k** ⭐ 

## gobufallo
https://github.com/gobuffalo/pop \
**1.2k** ⭐

## gormigrate
https://github.com/go-gormigrate/gormigrate \
**1k** ⭐ \
❗ Gormigrate is a minimalistic migration helper for Gorm. Gorm already has useful migrate functions, just misses proper schema versioning and migration rollback support.

