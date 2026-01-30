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

func (a *Adb) CheckEnv() error {
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

	return nil
}

func (a *Adb) Initialize() error {
	// pkill -f HD-Adb
	// Richelieu: 此处返回的 err 不用管
	_, _ = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "HD-Adb")

	// pkill -f adb
	// Richelieu: 此处返回的 err 不用管
	_, _ = cmdKit.RunCombinedlyToString(context.TODO(), "pkill", "-f", "adb")

	// adb kill-server
	_, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "kill-server")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb kill-server'")
	}

	// adb start-server
	_, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "start-server")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb start-server'")
	}

	// adb connect {a.Address}
	_, err = cmdKit.RunCombinedlyToString(context.TODO(), "adb", "connect", a.Address)
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb connect %s'", a.Address)
	}

	// adb devices
	devices, err := cmdKit.RunCombinedlyToString(context.TODO(), "adb", "devices")
	if err != nil {
		return errorKit.Wrapf(err, "fail to run 'adb devices'")
	}
	console.Infof("adb devices:\n%s", devices)

	return nil
}
