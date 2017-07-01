package dbkit

import (
	"crypto/tls"
	"fmt"
	"git.gumpcome.com/go_kit/logiccode"
	"git.gumpcome.com/go_kit/strkit"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
)

const (
	SOCKET_TIME_OUT_MS = 10000 //操作超时,默认10s
	CONN_TIME_OUT_MS   = 10000 //连接超时,默认10s
	MAX_POOL_SIZE      = 100   //连接池大小
	MAX_RETRIES        = 5     //连接失败后,重试次数
)

var (
	mongoInited   bool //是否已初始化
	globalSession *mgo.Session
	connUrl       string //完整连接URL
	mongoDBName   string
)

type MongoSearch struct {
	Collection string
	Key        string
	Value      interface{}
}

// 非SSL协议初始K化数据库
// @connUrl 连接字符串
// @dbName  数据库名称
func InitMongoDB(connUrl string, dbName string) {
	if connUrl == "" || dbName == "" {
		panic("conn url or db name is empty!")
	}
	if mongoInited {
		return
	}
	fullUrl := setConnUrlOptions(connUrl)
	mySession, err := mgo.Dial(fullUrl)
	if err != nil {
		panic(err)
	}
	mySession.SetMode(mgo.Monotonic, true)
	mySession.SetPoolLimit(100)
	globalSession = mySession
	connUrl = fullUrl
	mongoDBName = dbName
	mongoInited = true
}

// SSL协议初始化数据库
// @connUrl 连接字符串
// @dbName  数据库名称
func InitMongoDBWithSSL(connUrl string, dbName string) {
	if connUrl == "" || dbName == "" {
		panic("conn url or db name is empty!")
	}
	if mongoInited {
		return
	}
	fullUrl := setConnUrlOptions(connUrl)
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	dialInfo, err := mgo.ParseURL(fullUrl)
	if err != nil {
		panic(err)
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	mySession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	err = mySession.Ping()
	if err != nil {
		panic(err)
	}
	mySession.SetMode(mgo.Monotonic, true)
	mySession.SetPoolLimit(100)
	globalSession = mySession
	connUrl = fullUrl
	mongoDBName = dbName
	mongoInited = true
}

// 查询出与search匹配的结果,如果没有就添加,否则会覆盖原有记录
// 注意:如果Mongo搜索到多条与search匹配的记录,只会更新最新插入的一条记录。
func MongoUpsertDoc(search *MongoSearch, doc interface{}) (*mgo.ChangeInfo, bool, error) {
	session, err := getSession()
	if err != nil {
		return &mgo.ChangeInfo{}, false, err
	}
	defer session.Clone()
	if search == nil || doc == nil || search.Collection == "" || search.Key == "" || search.Value == nil {
		return &mgo.ChangeInfo{}, false, logiccode.MongoParamsErrorCode()
	}
	changeInfo, err := session.DB(mongoDBName).C(search.Collection).Upsert(bson.M{search.Key: search.Value}, doc)
	if err != nil {
		return &mgo.ChangeInfo{}, false, logiccode.MongoUpsertErrorCode(err)
	}
	return changeInfo, true, nil
}

//插入记录
func MongoInsert(colelection string, data interface{}) error {
	session, err := getSession()
	if err != nil {
		return err
	}
	defer session.Clone()
	c := session.DB(mongoDBName).C(colelection)
	err = c.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

// 查找单个记录
func MongoFindDoc(collection string, fun func(*mgo.Collection)) error {
	session, err := getSession()
	if err != nil {
		return err
	}
	defer session.Clone()
	c := session.DB(mongoDBName).C(collection)
	fun(c)
	return nil
}

// 删除所有记录
func MongoRemoveAllDoc(search *MongoSearch) (bool, error) {
	session, err := getSession()
	if err != nil {
		return false, err
	}
	defer session.Clone()
	if search == nil || search.Collection == "" || search.Key == "" || search.Value == "" {
		return false, logiccode.MongoParamsErrorCode()
	}
	_, err = session.DB(mongoDBName).C(search.Collection).RemoveAll(bson.M{search.Key: search.Value})
	if err != nil {
		return false, logiccode.MongoRemoveErrorCode(err)
	}
	return true, nil
}

func Find(search *MongoSearch) (*mgo.Query, error) {
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	data := session.DB(mongoDBName).C(search.Collection).Find(bson.M{search.Key: search.Value})
	return data, nil
}

// 设置连接字符串后缀可选项
func setConnUrlOptions(connUlr string) string {
	opts := make([]string, 0)
	opts = append(opts, connUlr)
	opts = append(opts, "?")
	opts = append(opts, "authMechanism=MONGODB-CR")
	opts = append(opts, "&maxPoolSize=100")
	//opts = append(opts, "&connectTimeoutMS=10000") //10s连接超时
	//opts = append(opts, "&socketTimeoutMS=10000")  //10s操作超时
	return strkit.StrJoin(opts...)
}

func getSession() (*mgo.Session, error) {
	if !mongoInited || globalSession == nil {
		return nil, logiccode.MongoSessionErrorCode()
	}
	isSessionOk := true
	err := globalSession.Ping()
	if err != nil {
		isSessionOk = false
		globalSession.Refresh()
		for i := 0; i < MAX_RETRIES; i++ {
			err = globalSession.Ping()
			if err == nil {
				isSessionOk = true
				beego.Info("Reconnect to mongodb successful.")
				break
			} else {
				beego.Error(fmt.Sprintf("Reconnect to mongodb fail:%v"), i)
			}
		}
	}
	if isSessionOk {
		return globalSession.Clone(), nil
	}
	return nil, logiccode.MongoSessionCloneErrorCode()
}

//
//func InitMongoDBWithUrl(url string) {
//	if inited {
//		return
//	}
//}
//
//func InitMongoDBWithUrlSSL(url string) {
//	if inited {
//		return
//	}
//}
//
//func (this *MongoDB) New(user, password, host, admin, dbname string) *MongoDB {
//	if user == "" || password == "" || host == "" || admin == "" || dbname == "" {
//		panic("please complete params")
//	}
//	if this.PoolLimit <= 0 {
//		this.PoolLimit = DEFAUL_POOL_LIMIT
//	}
//	if this.init {
//		return this
//	}
//	this.init = true
//	this.User = user
//	this.Password = password
//	this.Host = host
//	this.Admin = admin
//	this.Db = dbname
//
//	if this.MaxReties <= 0 {
//		this.MaxReties = MAX_RETRIES
//	}
//	session, err := mgo.Dial(fmt.Sprintf(MONGO_CR, this.User, this.Password, this.Host, this.Admin))
//	if err != nil {
//		panic(err)
//	}
//	session.SetPoolLimit(this.PoolLimit)
//	err = session.Ping()
//	if err != nil {
//		panic(err)
//	}
//	session.SetMode(mgo.Monotonic, true)
//
//	this.db = session.DB(this.Db)
//	return this
//}
//
//func (this *MongoDB) GetMongo() *mgo.Database {
//	if this.init == false {
//		panic("please call New before")
//	}
//	err := this.db.Session.Ping()
//	if err != nil {
//		beego.Debug("%v", "Lost connection to db!")
//		this.db.Session.Refresh()
//		for i := 0; i < MAX_RETRIES; i++ {
//			err = this.db.Session.Ping()
//			if err == nil {
//				beego.Debug("%v", "Reconnect to db successful.")
//				break
//			} else {
//				beego.Error("Reconnect to db faild:%v", i)
//			}
//		}
//	}
//	return this.db
//}
