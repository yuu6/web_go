package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"

)

const(
	ip = "docker.for.mac.host.internal"
	port = "3306"
	userName = "root"
	password = "123456"
	dbName = "person"
)

type Person struct {
	Id int32
	Name string
}


//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB()  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

//func InsertUser() (bool){
//	//开启事务
//	tx, err := DB.Begin()
//	if err != nil{
//		fmt.Println("tx fail")
//		return false
//	}
//	//准备sql语句
//	stmt, err := tx.Prepare("INSERT INTO person.student VALUES (?, ?)")
//	if err != nil{
//		fmt.Println("Prepare fail")
//		return false
//	}
//
//	temp:= make(map[int]string)
//	temp[1] = "zhangsan"
//	temp[2] = "lisi"
//	temp[3] = "wangwu"
//
//	for k, v := range temp { //每行数据是放在values里面，现在把它挪到row里
//		//将参数传递到sql语句中并且执行
//		log.Println(k,v)
//		_, err := stmt.Exec(k, v)
//		if err != nil{
//			fmt.Println("Exec fail")
//			return false
//		}
//	}
//	log.Println("插入完成")
//	//将事务提交
//	tx.Commit()
//	//获得上一个插入自增的id
//	//fmt.Println(res.LastInsertId())
//	return true
//}

func SelectAllUser() ([]Person){
	// 首先初始化
	InitDB()
	//执行查询语句
	rows, err := DB.Query("SELECT * from person.student")
	if err != nil{
		fmt.Println("查询出错了")
	}
	list := []Person{}
	//循环读取结果
	for rows.Next(){
		user:=Person{}
		//将每一行的结果都赋值到一个user对象中

		err = rows.Scan(&user.Id, &user.Name)

		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		list = append(list, user)
	}
	log.Println(list)
	log.Println("it is over")
	return list
}


func SelectOneUser(id string) (Person){
	InitDB()
	//执行查询语句
	stmt, err := DB.Prepare("SELECT * from person.student where id = ?")
	rows, err := stmt.Query( id)

	if err != nil{
		fmt.Println("查询出错了")
	}
	list := []Person{}
	//循环读取结果
	for rows.Next(){
		user:=Person{}
		//将每一行的结果都赋值到一个user对象中

		err = rows.Scan(&user.Id, &user.Name)

		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		list = append(list, user)
	}
	log.Println(list)
	log.Println("it is over")
	return list[0]
}


func main(){
	InitDB()
	SelectAllUser()
}