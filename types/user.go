package types

// bson -> convert ID to _id that can be used in database
// json -> convert id to id that can be used in http api call
// omitempty -> it will omit id from the api call which is empty or we dont want to display it according to business requirement
type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
