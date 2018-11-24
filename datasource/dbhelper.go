package datasource

import (
	"fmt"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yz124/superstar/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 主库，单例
//可以高并发访问该函数
func InstanceMaster() *xorm.Engine {
	//临界区的双重检测
	//快速检测
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	//慢检测
	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	//构建一个engine
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	//设置ORM数据库时区
	engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//设置基于内存的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher) //全局SQL缓存，可以设置针对某个struct的缓存
	//将新创建的engine赋值给masterEngine
	masterEngine = engine
	return engine
}

// 从库，单例
func InstanceSlave() *xorm.Engine {
	//临界区的双重检测
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	engine, err := xorm.NewEngine(conf.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		log.Fatal("dbhelper", "DbInstanceMaster", err)
		return nil
	}
	//设置xorm时区
	engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return engine
}
