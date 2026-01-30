package main

import (
	"github.com/richelieu-yang/UnchartedWaters/src/adb"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

func main() {
	a := &adb.Adb{
		Address: "127.0.0.1:5555",
	}
	if err := a.CheckEnv(); err != nil {
		console.Fatalf("a.CheckEnv() failed: %s", err)
		return
	}

}
