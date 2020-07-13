package queryutils

import "github.com/go-pg/pg/v9/orm"

func WithRelations(q *orm.Query, relations ...string) *orm.Query {
	for i := range relations {
		q = q.Relation(relations[i])
	}

	return q
}
