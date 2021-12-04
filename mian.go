package main

import (
	"database/sql"
	"errors"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"os"
)

var Customerdata1 struct{
	ID string `db:"id"`
	Password int `db:"password"`
}

var Customerdata2 struct{
	ID string `db:"id"`
	Questions string `db:"questions"`
	Answers string `db:"answers"`

}

var db *sql.DB//Db数据库连接池

func initDB() (err error){
	dsn := "root:123456@(127.0.0.1:3306)/customerData"
	db,err = sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("err: %v\n",err)
		return
	}
	err = db.Ping()
	if err != nil{
		fmt.Printf("open %s faild, err:%v\n",dsn,err)
		return
	}
	return
}//连接数据库

func main(){
	initDB()
	err := initDB()
	if err !=nil{
		fmt.Print("initDB failed,err:%v\n",err)
	}
	menu()
	defer db.Close()
}

/*
var scanner *bufio.Scanner
func getInput() string {
	scanner.Scan()
	return scanner.Text()
}//用函数接收界面开始时输入的值
*/

/*
func getOptionSelect(title string, sList []string) int {
	if len(sList) == 0 {
		return 0
	}
	fmt.Println(title)
	for i, s := range sList {
		fmt.Printf("%d.%s\n", i+1, s)
	}
	for {
		content := getInput()
		choice, err := strconv.Atoi(content)
		if err != nil || choice <= 0 || choice > len(sList) {
			fmt.Println("输入不合法，请重新输入")
			continue
		}
		return choice
	}
}//判断输入的数字是否合法
*/

var errUnknown error = errors.New("unknown error")//设置一个值，接收未知的错误，防止panic；

var choice int

func menu() {
	fmt.Println("欢迎来到世界树")
	t := []string{"1注册","2登录","3设置密保信息","4忘记密码"}
	fmt.Println(t)
	fmt.Print("请输入您的选择")
	fmt.Scanln(&choice)
	switch choice {
		case 1:
			Resign()
		case 2:
			Login()
		case 3:
			Settings()
	    case 4:
			Verify()
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}//主菜单3

func Resign(){
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&Customerdata1.ID)
	fmt.Print("请输入您的密码：")
	fmt.Scanln(&Customerdata1.Password)

	r := "insert into customerdata1(id,password) values (?,?)"

	re,err := db.Exec(r,Customerdata1.ID,Customerdata1.Password)

	if err != nil{
		fmt.Printf("注册失败，\n err:%v",err)
		return
	}
	newID,err :=re.LastInsertId()

	if err != nil{
		fmt.Printf("get lastinsert id failed,err:%v\n",err)
	}
	fmt.Printf("insert success, the id is %d.\n", newID)
	fmt.Println("恭喜您，创建成功！")


}//注册时调用的函数

func Settings(){
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&Customerdata2.ID)
	fmt.Print("请输入您的密保答案：")
	fmt.Scanln(&Customerdata2.Questions)
	fmt.Print("请输入您的密保答案：")
	fmt.Scanln(&Customerdata2.Answers)

	str := "insert into C(id,questions,answers) values (?,?,?)"

	re,err := db.Exec(str,Customerdata2.ID,Customerdata2.Questions,Customerdata2.Answers)
	if err != nil{
		fmt.Printf("insert failed,err:%v\n",err)
		return
	}
	new,err := re.LastInsertId()
	if err != nil{
		fmt.Printf("set 密保 failed,err:%v\n",err)
	}
	fmt.Printf("insert success, the id is %d.\n", new)
	fmt.Println("恭喜您，密保设置成功")

}//设置密保调用的函数

func Login(){

	fmt.Print("请输入您的账号：")
	fmt.Scanln(&Customerdata1.ID)
	fmt.Print("请输入您的密码：")
	fmt.Scanln(&Customerdata1.Password)

	str := "select id,password from customerdata1 where id =?"
	/*
	单行查询db.QueryRow()执行一次查询，并且期望返回最多一行数据（Row）。
	QueryRow总是返回非nil值，直到返回值被Scan方法调用，才会返回被延迟错误。（如：没找到结果）
    */
	var U struct{
		ID string `db:"id"`
		Password int `db:"password"`
	}
	err :=db.QueryRow(str,Customerdata1.ID).Scan(&U.ID,&U.Password)
	if err !=nil{
		fmt.Println("账号错误")
	}else if U.Password == Customerdata1.Password{
		fmt.Println("恭喜您登陆成功")
	}else{
		fmt.Println("密码错误")
		os.Exit(0)
		}
	}//登陆时调用的函数

func Verify(){
	fmt.Print("请输入您的账号：")
	fmt.Scanln(&Customerdata2.ID)

	var U struct{
		ID string `db:"id"`
		Questions string `db:"questions"`
		Answers string `db:"answers"`
	}

	str := "select id from customerdata2 where id = ?" //查询custmerdata2中的id信息
	err :=db.QueryRow(str,Customerdata2.ID).Scan(&U.ID)//查询输入的id与 customerdata2 中的id是否一致

	if err != nil {
		fmt.Println("账号错误")
	} else if Customerdata2.ID == U.ID {
		str2 := "select questions from customerdata2 where id = ?"
		err := db.QueryRow(str2,Customerdata2.ID).Scan(&U.Questions)
		if err != nil {
			fmt.Println("未设置密保")
		} else {
			fmt.Println(U.Questions)
		}
	}
	fmt.Print("请输入您的密保答案：")
	fmt.Scanln(&Customerdata2.Answers)

	 str3 := "select customer2_answers from where id = ?"
	 err2 := db.QueryRow(str3,Customerdata2.ID).Scan(&U.Answers)

	 if err2 != nil{
		 fmt.Println("查询密保答案这里出现error！")
	 }else if Customerdata2.Answers == U.Answers{
		 fmt.Println("恭喜您，回答正确，请输入新的密码！")
		 fmt.Scanln(&Customerdata1.Password)
		 Updata()
	}
}//忘记密码，调用的验证函数

func Updata(){

		sqlStr := "update customerdata1 set password =? where id = ?"
		ret, err := db.Exec(sqlStr, Customerdata1.Password, Customerdata2.ID)
		if err != nil {
			fmt.Printf("修改失败！ err:%v\n", err)
			return
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			fmt.Printf("操作行数处理失败  err:%v\n", err)
			return
		}
		fmt.Printf("密码修改成功！  affected rows:%d\n", n)
	}
