package transaction

import "gopkg.in/mgo.v2/bson"

type Transaction struct {
	Id   bson.ObjectId `form:"id" bson:"_id,omitempty"`
	Data string        `form:"data" bson:"data"`
}
