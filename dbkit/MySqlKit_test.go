package dbkit

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"git.gumpcome.com/go_kit/timekit"
	"testing"
)

//INSERT INTO user(name,age,email,gender,height,interests) VALUES (?,?,?,?,?,?)
func TestCreateMysqlInsertSQL(t *testing.T) {
	tableName := "user"
	data := make(map[string]interface{})
	data["name"] = "小刘"
	data["age"] = 18
	data["email"] = "xiaoliu@gumpcome.com"
	data["gender"] = 1
	data["height"] = 180
	data["interests"] = "游泳,爬山"

	sql, params := CreateMysqlInsertSQL(tableName, data)
	fmt.Println(sql)
	fmt.Println(params)
}

func TestInitMysql(t *testing.T) {
	logger := logs.NewLogger(1000)
	logger.SetLogger(logs.AdapterConsole)
	dbUserName := "root"
	dbUserPwd := "123456"
	dbHost := "127.0.0.1:3306"
	dbName := "godb"
	dbMaxIdle := 10
	dbMaxActive := 20
	InitMysql(dbUserName, dbUserPwd, dbHost, dbName, dbMaxIdle, dbMaxActive, logger)
}

func TestSaveInMysql(t *testing.T) {
	tableName := "user"
	data := make(map[string]interface{})
	data["name"] = "大龙"
	data["age"] = 18
	data["email"] = "xiaoliu@gumpcome.com"
	data["gender"] = 1
	data["height"] = 180
	data["interests"] = "游泳,爬山"
	unixTime, createTime, _ := timekit.GetNowTimeMsAndDate(timekit.DateFormat_YYYY_MM_DD_HH_MM_SS)
	data["createtime"] = createTime
	data["unixtime"] = unixTime

	result, id, err := SaveInMysql(GetMysqlCon(), tableName, data)
	if !result && err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Printf("保存记录返回的主键结果 id=%d\n", id)
}

func TestUpdateByIdInMysql(t *testing.T) {
	tableName := "user"
	data := make(map[string]interface{})
	data["id"] = 12
	data["name"] = "小刘3"
	data["age"] = 18
	data["email"] = "xiaoliu@gumpcome.com"
	data["gender"] = 1
	data["height"] = 180
	data["interests"] = "游泳1,爬山1"
	unixTime, createTime, _ := timekit.GetNowTimeMsAndDate(timekit.DateFormat_YYYY_MM_DD_HH_MM_SS)
	data["createtime"] = createTime
	data["unixtime"] = unixTime

	result, err := UpdateByIdInMysql(GetMysqlCon(), tableName, data)
	if !result && err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Printf("根据ID更新记录返回结果 result=%t\n", result)
}

func TestUpdateInMysql(t *testing.T) {
	sql := `UPDATE user SET name = ? , age = ? , email = ? , gender = ? , interests = ? , createtime = ? , unixtime = ? , height = ?  WHERE id = ?`
	unixTime, createTime, _ := timekit.GetNowTimeMsAndDate(timekit.DateFormat_YYYY_MM_DD_HH_MM_SS)
	result, err := UpdateInMysql(GetMysqlCon(), sql, "大龙", 20, "dalong@gumpcome.com", 1, "潜水,旅游", createTime, unixTime, 176, 13)
	if !result && err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Printf("更新记录返回结果 result=%t\n", result)
}

func TestDeleteByIdInMysql(t *testing.T) {
	result, err := DeleteByIdInMysql(GetMysqlCon(), "user", 14)
	if !result && err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Printf("根据ID删除记录返回结果 result=%t\n", result)
}

func TestDeleteInMysql(t *testing.T) {
	sql := `DELETE FROM user WHERE name = ?`
	result, err := DeleteInMysql(GetMysqlCon(), sql, "小刘2")
	if !result && err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Printf("记录返回结果 result=%t\n", result)
}

func TestFindInMysql(t *testing.T) {
	sql := `SELECT * FROM user`
	result, err := FindInMysql(GetMysqlCon(), sql)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Println(result)
}

func TestFindFirstInMysql(t *testing.T) {
	sql := `SELECT * FROM user LIMIT 1`
	result, err := FindFirstInMysql(GetMysqlCon(), sql)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Println(result)
}

func TestFindFirstInMysql2(t *testing.T) {
	sql := `SELECT COUNT(*) AS count FROM user WHERE name LIKE ? ORDER BY id`
	result, err := FindFirstInMysql(GetMysqlCon(), sql, "%大龙%")
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Println(result)
}

func TestPaginateInMysql(t *testing.T) {
	selectSql := `SELECT name AS user_name, age As user_age`
	sqlExceptSelect := `FROM user WHERE name LIKE ?`
	result, err := PaginateInMysql(GetMysqlCon(), 1, 3, selectSql, sqlExceptSelect, "%大龙%")
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	fmt.Println(result)
}
