package db

import "fmt"

type DbError struct {
	message string
}

type NonExistingDbError struct {
	message string
}

type NonPermissionsDbError struct {
	message string
}

func (c *DbError) Error() string {
	return fmt.Sprintf("db error: %s", c.message)
}

func (c *NonExistingDbError) Error() string {
	return fmt.Sprintf("db does not exist : %s", c.message)
}

func (c *NonPermissionsDbError) Error() string {
	return fmt.Sprintf("db does not exist : %s", c.message)
}
