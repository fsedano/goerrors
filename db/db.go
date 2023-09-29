package db

import (
	"fmt"
	"log"
)

func DbOpen(option string) error {
	log.Printf("Opening db %s", option)

	switch option {
	case "nonexistingdb":
		return fmt.Errorf("things failed: %w", &NonExistingDbError{message: "We failed because db did not exist"})
	case "notpermissionsdb":
		return fmt.Errorf("things failed: %w", &NonPermissionsDbError{message: "We failed because permissions were not there"})
	default:
		return fmt.Errorf("things failed: %w", &DbError{message: "We failed in a non specific way"})
	}

}
