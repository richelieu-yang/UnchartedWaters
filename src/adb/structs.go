package adb

import (
	"context"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

type Adb struct {
	// Address 安卓设备的地址
	/*
		（1）基本语法：adb connect <设备IP地址>:<端口号>
		（2）默认端口号是 5555，如果使用默认端口，可以省略端口号
			e.g. adb connect 192.168.1.100 <=> adb connect 192.168.1.100:5555
		（3）mac上，BlueStacks Air默认是 "127.0.0.1:5555"
	*/
	Address string
}

func (adb Adb) CheckEnv() error {
	path, err := cmdKit.LookPath("adb")
	if err != nil {
		return errorKit.Wrapf(err, "fail to look path of adb")
	}
	console.Infof("adb path: %s", path)

	// adb 版本号
	{
		str, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "version")
		if err != nil {
			return errorKit.Wrapf(err, "fail to run 'adb version'")
		}
		console.Infof("adb version:\n%s", str)
	}

	// adb devices
	{
		str, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "devices")
		if err != nil {
			return errorKit.Wrapf(err, "fail to run 'adb devices'")
		}
		console.Infof("adb devices:\n%s", str)
	}

	return nil
}
