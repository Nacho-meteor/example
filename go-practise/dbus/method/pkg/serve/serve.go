package serve

import (
	"fmt"

	"pkg.deepin.io/lib/dbusutil"
)

const (
	dbusName = "com.practise.Example"  //名称
	dbusPath = "/com/practise/Example" //地址
	dbusIFC  = "com.practise.Example"  //接口名
)

// 实体化服务对象
type Object struct {
	methods *struct {
		GetMd5sum  func() `in:"msg" out:"md5"`
		SendSignal func() `in:"message,object"`
	}
	signal *struct {
		Send struct {
			message string
			object  string
		}
	}
}

// dbus对象
type Service struct {
	conn   *dbusutil.Service
	Object *Object
}

var dbusObj *Service

// 获取初始化的dbus对象,不存在就新建
func GetService() *Service {
	if dbusObj != nil {
		return dbusObj
	}
	dbusObj, err := newService()
	if err != nil {
		panic(err)
	}
	return dbusObj
}

// 新建
func newService() (*Service, error) {
	var dbusService *Service
	srv, err := dbusutil.NewSystemService()
	if err != nil {
		return nil, fmt.Errorf("new system service is error:%s\n", err)
	}
	obj := &Object{}
	dbusService = &Service{conn: srv, Object: obj}
	return dbusService, nil
}

// 外部调用
func (srv *Service) Init() error {
	return srv.initDBus()
}

// 外调
func (srv *Service) initDBus() error {
	err := srv.conn.Export(dbusPath, GetService().Object)
	if err != nil {
		return err
	}
	return srv.conn.RequestName(dbusName)
}

// 获取 dbus对象 ifc名称
func (o *Object) GetInterfaceName() string {
	return dbusIFC
}

// 阻塞
func (srv *Service) Loop() {
	fmt.Println("dbus success")

	srv.conn.Wait()
}
