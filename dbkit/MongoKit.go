package dbkit

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2"
)

const (
	MAX_RETRIES       = 5
	DEFAUL_POOL_LIMIT = 10
	MONGO_CR          = "mongodb://%s:%s@%s/%s?authMechanism=MONGODB-CR"
)

type MongoDB struct {
	User      string
	Password  string
	Host      string
	Admin     string
	Db        string
	db        *mgo.Database
	MaxReties int
	PoolLimit int
	*logs.BeeLogger
	init bool
}

func (this *MongoDB) New(user, password, host, admin, dbname string) *MongoDB {
	if user == "" || password == "" || host == "" || admin == "" || dbname == "" {
		panic("please complete params")
	}
	if this.BeeLogger == nil {
		panic("please init log")
	}
	if this.PoolLimit <= 0 {
		this.PoolLimit = DEFAUL_POOL_LIMIT
	}
	if this.init {
		return this
	}
	this.init = true
	this.User = user
	this.Password = password
	this.Host = host
	this.Admin = admin
	this.Db = dbname

	if this.MaxReties <= 0 {
		this.MaxReties = MAX_RETRIES
	}
	session, err := mgo.Dial(fmt.Sprintf(MONGO_CR, this.User, this.Password, this.Host, this.Admin))
	if err != nil {
		panic(err)
	}
	session.SetPoolLimit(this.PoolLimit)
	err = session.Ping()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	this.db = session.DB(this.Db)
	return this
}

func (this *MongoDB) GetMongo() *mgo.Database {
	if this.init == false {
		panic("please call New before")
	}
	err := this.db.Session.Ping()
	if err != nil {
		this.Debug("%v", "Lost connection to db!")
		this.db.Session.Refresh()
		for i := 0; i < MAX_RETRIES; i++ {
			err = this.db.Session.Ping()
			if err == nil {
				this.Debug("%v", "Reconnect to db successful.")
				break
			} else {
				this.Error("Reconnect to db faild:%v", i)
			}
		}
	}
	return this.db
}
