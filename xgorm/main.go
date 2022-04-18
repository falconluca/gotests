package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"hello/xgorm/model"
	"log"
)

type Store struct {
	engin *xorm.Engine
}

func NewStore() *Store {
	engine, err := xorm.NewEngine("mysql",
		"root:12345678@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return &Store{
		engin: engine,
	}
}

func (s *Store) SyncTables() {
	if err := s.engin.Sync2(
		new(model.User),
		new(model.Detail),
	); err != nil {
		log.Fatal("sync tables failed,", err)
	}
}

func (s *Store) Close() error {
	return s.engin.Close()
}

func (s *Store) ExecUser() {
	u := new(model.User)
	u.Name = "Sophia"
	u.Age = 20
	u.Salt = "RAD_CAT"
	u.Passwd = "XXX"
	// exec 不会自动生成 updated_at 和 created_at 字段，除非自己指定
	r, err := s.engin.Exec("insert into user(user_name, age, salt, passwd) "+
		"values(?, ?, ?, ?)", u.Name, u.Age, u.Salt, u.Passwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[Exec] insert user: ", u.Name, u.Age, u.Salt, u.Passwd)
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	insertId, err := r.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[Exec] rowsAffected: %v, insertId: %v", rowsAffected, insertId)
}

func (s *Store) InsertUsers() {
	u := &model.User{
		Name:   "Luca",
		Age:    24,
		Salt:   "asdasd",
		Passwd: "sadacqweqwe",
	}
	_, err := s.engin.Insert(u)
	if err != nil {
		// Error 1062: Duplicate entry 'Luca' for key 'user.UQE_user_user_name'
		log.Fatalln("insert user failed, ", err)
	}

	users := []*model.User{
		{
			Name:   "Alasd",
			Age:    23,
			Salt:   "dwe",
			Passwd: "asdasd",
		},
		{
			Name:   "Alasd23",
			Age:    222,
			Salt:   "dwe33",
			Passwd: "asdasd3",
		},
	}
	_, err = s.engin.Insert(users)
	if err != nil {
		log.Fatal("insert users failed,", err)
	}
}

func (s *Store) Get() {
	cols := []string{
		"user_name",
		"passwd",
		"salt",
	}
	colValues := make([]interface{}, len(cols))
	_, err := s.engin.Table(new(model.User)).
		Where("id = ?", 2).
		Get(&colValues)
	if err != nil {
		log.Fatal("get single user by cols failed,", err)
	}
	for i, colVal := range colValues {
		//val, ok := colVal.(string)
		fmt.Printf("i: %v => %v\n", i, colVal)
	}
}

func (s *Store) Exist() {
	// TODO
}

func (s *Store) Find() {
	name := "Luca"
	var userDetailList []model.UserDetail
	if err := s.engin.Table("user").
		Select("user.*, detail.user_id").
		Join("inner", "detail", "detail.user_id = user.id").
		Where("user.user_name = ?", name).
		Limit(10, 0).
		Find(&userDetailList); err != nil {
		log.Fatal("failed to find user by name,", err)
	}
	fmt.Println("len of userDetailList:", len(userDetailList))
}

func (s *Store) Distinct() {
	count, err := s.engin.Table("user").Distinct("user_name").Count()
	if err != nil {
		log.Fatal("distinct name failed", err)
	}
	fmt.Println("distinct name cnt:", count)

	var users []*model.User
	if err = s.engin.Distinct("user_name").Find(&users); err != nil {
		log.Fatal("distinct name and find failed,", err)
	}
	for _, user := range users {
		fmt.Println("name:", user.Name)
	}

	var userNames []string
	if err = s.engin.Table("user").
		Distinct("user_name").
		Find(&userNames); err != nil {
		log.Fatal("distinct name and find failed,", err)
	}
	for _, user := range userNames {
		fmt.Println("name:", user)
	}
}

func (s *Store) TestNoAutoCondition() {
	dbSession := s.engin.NewSession().NoAutoCondition(true)
	//dbSession := s.engin.NewSession()
	defer dbSession.Close()

	u := new(model.User)
	u.Name = "Luca"

	if _, err := dbSession.Get(u); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user: %+v\n", u)
}

func main() {
	store := NewStore()
	defer store.Close()

	store.SyncTables()

	//store.ExecUser()
	//store.InsertUsers()

	//store.Get()
	//store.Exist()
	//store.Find()
	//store.Distinct()

	store.TestNoAutoCondition()
}

//func joinQuery(engine *xorm.Engine) {
//	var users = make([]*User, 0)
//	//var users []*User
//	//fmt.Println(users)
//	name := "Luca"
//	if err := engine.Table(&Userinfo{}).
//		Join("left", "userdetail", "userinfo.detail_id = userdetail.id").
//		Where("name = ?", name).
//		Find(&users); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("len: %v ,user name: %s, gender: %v\n", len(users), users[0].Name, users[0].Gender)
//}

//func queryOrderByIdDesc(engine *xorm.Engine) {
//	users := []*model.User{}
//	if err := engine.Desc("id").Find(&users); err != nil {
//		log.Print(err)
//	}
//	for _, u := range users {
//		printUser(u)
//	}
//}
//
//func queryWhereAnd(engine *xorm.Engine) {
//	u := new(model.User)
//	if _, err := engine.Where("usr_name = ?", "Sophia").
//		And("open_id = ?", "asd234adsfc").
//		Get(u); err != nil {
//		log.Fatal(err)
//	}
//	printUser(u)
//}

//func printUser(u *model.User) {
//	log.Printf("UserInfo ==> name: %s, unionid: %s, id: %v\n", u.Name, u.UnionId, u.Id)
//}
//
//func insertUser(engine *xorm.Engine) {
//	user := new(model.User)
//	user.Name = "Luca"
//	user.OpenId = "openidluca"
//	user.UnionId = "unionidluca"
//	affected, err := engine.Insert(user)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("insert user success, affected: %v\n", affected)
//}
//
//func queryAlias(engine *xorm.Engine) {
//	u := new(model.User)
//	name := "Allen"
//
//	_, err := engine.Alias("u").Where("u.usr_name = ?", name).Get(u)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Print(u.Name)
//	log.Print(u.UnionId)
//	log.Print(u.Id)
//}
