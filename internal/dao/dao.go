package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// dao DaoStruct interface
type Dao interface {
   Close()
   Ping(ctx context.Context) (err error)
}

// DaoStruct DaoStruct.
type DaoStruct struct {
	db          *sql.DB
	//redis       *redis.Pool
	//redisExpire int32
	mc          *memcache.Memcache
	mcExpire    int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a DaoStruct and return.
func New() (*DaoStruct) {
	var (
		dc struct {
			SqlConfig *sql.Config
		}
		//rc struct {
		//	Demo       *redis.Config
		//	DemoExpire xtime.Duration
		//}
		mc struct {
			MemCache       *memcache.Config
			MemCacheExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	//checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	checkErr(paladin.Get("memcache.toml").UnmarshalTOML(&mc))
	return &DaoStruct{
		// mysql
		db: sql.NewMySQL(dc.SqlConfig),
		// redis
		//redis:       redis.NewPool(rc.Demo),
		//redisExpire: int32(time.Duration(rc.DemoExpire) / time.Second),
		// memcache
		mc:       memcache.New(mc.MemCache),
		mcExpire: int32(time.Duration(mc.MemCacheExpire) / time.Second),
	}
}

// Close close the resource.
func (d *DaoStruct) Close() {
	d.mc.Close()
	//d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *DaoStruct) Ping(ctx context.Context) (err error) {
	if err = d.pingMC(ctx); err != nil {
		return
	}
	//if err = d.pingRedis(ctx); err != nil {
	//	return
	//}
	return d.db.Ping(ctx)
}

func (d *DaoStruct) pingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

//func (d *DaoStruct) pingRedis(ctx context.Context) (err error) {
//	conn := d.redis.Get(ctx)
//	defer conn.Close()
//	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
//		log.Error("conn.Set(PING) error(%v)", err)
//	}
//	return
//}
