module github.com/budougumi0617/til/go/modules/branch

go 1.12

require (
	github.com/golang-migrate/migrate v1.3.2
	github.com/golang-migrate/migrate/v3 v3.5.4

	github.com/golang-migrate/migrate/v4 v4.4.0
	gopkg.in/mattes/migrate.v1 v1.3.2 // indirect
)

replace github.com/golang-migrate/migrate/v3 v3.5.4 => github.com/golang-migrate/migrate v3.5.4+incompatible
