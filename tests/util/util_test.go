package util_test

import (
	"fmt"
	"testing"

	"coursesheduling/lib/config"
	"coursesheduling/lib/util"
)

func TestSetValue(t *testing.T) {
	var config config.Configuration
	util.SetDefaults(&config)
	fmt.Println("config",config)
}
