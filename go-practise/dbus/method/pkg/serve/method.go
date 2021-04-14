package serve

import (
	"crypto/md5"
	"fmt"

	"pkg.deepin.io/lib/dbusutil"

	"github.com/godbus/dbus"
)

/*
	方法
*/
func (o *Object) GetMd5sum(msg string) (string, *dbus.Error) {
	var err error
	hash := md5.Sum([]byte(msg))
	md5sum := fmt.Sprintf("%x", hash)
	if len(md5sum) == 0 {
		err = fmt.Errorf("Hash calculation failed：%s\n", msg)
	}
	return md5sum, dbusutil.ToError(err)
}

func (o *Object) SendSignal(message string, object string) *dbus.Error {
	var err error
	if len(message) == 0 || len(object) == 0 {
		err = fmt.Errorf("The incoming content is incorrect：%s %s\n", message, object)
		goto end
	}
	err = o.emitSignal(GetService().conn, message, object)
end:
	return dbusutil.ToError(err)
}

/*
	信号
*/
func (o *Object) emitSignal(srv *dbusutil.Service, messgae string, object string) error {
	return srv.Emit(o, messgae, object)
}
