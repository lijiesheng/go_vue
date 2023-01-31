package relate_tables

//// 一对一关系 [属于]
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

// 一对一 [包含关系]
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

// 一对多关系
// 一个 User2 对应多个 Article
type User2 struct {
	Id       int
	Name     string
	Age      int
	Addr     string
	Articles []Article `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article struct {
	Id      int
	Title   string
	Content string
	Desc    string
	// 外键
	UId int
}

// 多对多
