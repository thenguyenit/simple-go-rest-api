package auth

type User struct {
	ID       int    `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Roles    string `bson:"roles"`
	Status   int    `bson:"status"`
}
