package relate_tables

//// 一对一关系
//type User struct {
//	Id   int
//	Name string
//	Age  int
//	Addr string
//}
//
//type UserProfile struct {
//	Id    int
//	Pic   string
//	CPic  string
//	Phone string
//	User  User `gorm:"ForeignKey:UId;AssociationForeignKey:Id"` // 关联关系
//	//UserID int  // 默认关联字段为Id
//	UId int // uid
//}

// 一对一
// UserProfile 包含一个 User, 外键在User模型，外键字段为:PId

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
	PId  int
}

type UserProfile struct {
	Id    int
	Pic   string
	CPic  string
	Phone string
	User  User `gorm:"ForeignKey:PId;AssociationForeignKey:Id"` // 关联关系

}
