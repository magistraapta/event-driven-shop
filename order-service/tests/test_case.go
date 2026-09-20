package tests

import (
	"github.com/goravel/framework/testing"

	"order-service/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
