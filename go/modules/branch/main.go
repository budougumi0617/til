package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/driver"
	database3 "github.com/golang-migrate/migrate/v3/database"
	database4 "github.com/golang-migrate/migrate/v4/database"
)

func main() {
	fmt.Println(database3.GenerateAdvisoryLockId("mod test"))
	fmt.Println(database4.GenerateAdvisoryLockId("mod test"))
	fmt.Println(driver.GetDriver("mod test"))
}
