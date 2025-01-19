package models

type User struct {
	ID        uint   `bson:"_id,ompitempty"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
}
