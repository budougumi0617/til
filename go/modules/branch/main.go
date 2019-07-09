package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v3/database"
)

func main() {
	fmt.Println(database.GenerateAdvisoryLockId("mod test"))
}
