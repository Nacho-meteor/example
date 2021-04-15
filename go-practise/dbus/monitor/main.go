package main

import (
	"flag"
	"fmt"
	dbus "go-lib/dbus1"
	"io/ioutil"
	"log"
	"strings"
)

var (
	uuid  = flag.String("uuid", "", "The value of uuid")
	allow = flag.Int("allow", -1, "Allow request")
	app   = flag.String("app", "", "Control application")
	ask   = flag.Bool("ask", false, "Whether to ask every time the message is processed")
	count = -1
	pid   = ""
)

func main() {
	flag.Parse()

	if len(*uuid) == 0 || len(*app) == 0 {
		flag.Usage()
		return
	}

	sysBus, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(err)
	}

	obj := sysBus.Object("com.deepin.daemon.ResourceManager", "/com/deepin/daemon/ResourceManager")

	err = sysBus.BusObject().AddMatchSignal("com.deepin.daemon.ResourceManager", "Notification",
		dbus.WithMatchObjectPath("/com/deepin/daemon/ResourceManager")).Err
	if err != nil {
		log.Fatal(err)
	}

	signalCh := make(chan *dbus.Signal, 10)
	sysBus.Signal(signalCh)

	fmt.Println(*app)
	go func() {
		log.Println("Start")
		for {
			select {
			case sig := <-signalCh:
				// log.Printf("sig: %#v\n", sig)
				if sig.Path == "/com/deepin/daemon/ResourceManager" &&
					sig.Name == "com.deepin.daemon.ResourceManager.Notification" {
					var subscrber []string
					var msg string
					err = dbus.Store(sig.Body, &subscrber, &msg)
					if err != nil {
						log.Println("WARN:", err)
					}
					if ischeckApp(msg, *app) {
						switch *allow {
						case 1:
							err = obj.Call("com.deepin.daemon.ResourceManager.AllowRequest", 0, *uuid, msg, true).Err
						case -1:
							if !isLive() {
								pid = ""
								count = -1
							}
							if count == 1 {
								err = obj.Call("com.deepin.daemon.ResourceManager.AllowRequest", 0, *uuid, msg, true).Err
							} else {
								fmt.Println("等待", msg, *app)
								if !*ask {
									count = 1
								}
							}
						case 0:
							err = obj.Call("com.deepin.daemon.ResourceManager.AllowRequest", 0, *uuid, msg, false).Err
						}
					} else {
						err = obj.Call("com.deepin.daemon.ResourceManager.AllowRequest", 0, *uuid, msg, true).Err
					}
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()
	select {}
}

func ischeckApp(msg string, app string) bool {
	items := strings.Split(msg, "-")
	if len(items) != 3 {
		return false
	}
	cmdPath := fmt.Sprintf("/proc/%s/cmdline", items[2])
	date, err := ioutil.ReadFile(cmdPath)
	if err != nil {
		return false
	}
	fmt.Println("对比：", string(date), app, items[2])
	if strings.Contains(string(date), app) {
		if !*ask {
			if count == -1 {
				pid = items[2]
			} else {
				count = 1
			}
		}
		return true
	}
	return false
}

func isLive() bool {
	cmdPath := fmt.Sprintf("/proc/%s/cmdline", pid)
	_, err := ioutil.ReadFile(cmdPath)
	if err != nil {
		return false
	}
	return true
}
