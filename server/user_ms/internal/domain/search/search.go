package search

type Search struct {
	ID string `dynamodbav:"id" bson:"_id,omitempty" json:"id"`
}