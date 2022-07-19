package query

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

func BuildPredicate(columns []string, query map[string][]string) (func(selector *sql.Selector), error) {
	for key, value := range query {
		if len(value) == 0 {
			continue
		}

		if !ContainsColumn(columns, key) {
			return nil, fmt.Errorf("invalid column %s", key)
		}
	}

	return func(selector *sql.Selector) {
		for key, value := range query {
			if len(value) == 0 {
				continue
			}
			selector.Where(sql.EQ(key, value[0]))
		}
	}, nil
}

func ContainsColumn(modelColumns []string, column string) bool {
	contains := false
	for _, modelColumn := range modelColumns {
		if column == modelColumn {
			contains = true
			break
		}
	}

	return contains

}
