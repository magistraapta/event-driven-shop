package tests

import (
	"github.com/goravel/framework/testing"

	"inventory-service/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
