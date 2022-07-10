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
		contains := false
		for _, column := range columns {
			if column == key {
				contains = true
				break
			}
		}
		if !contains {
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

// func BuildPredicate(columns []string, query map[string][]string) (func(selector *sql.Selector), error) {
//     for _, value := range query {
// 		hasColumn := ContainsColumn(columns, value[0])
// 		if !hasColumn {
// 			return nil, nil;
// 		}
//     }

//     return func(selector *sql.Selector) {
//         for key, value := range query {
//             if len(value) == 0 {
//                 continue
//             }
//             selector.Where(sql.EQ(key, value[0]))
//         }
//     }, nil
// }

func ContainsColumn(columns []string, column string) bool {
	if len(column) == 0 {
		return false
	}

	contains := false

	for _, value := range columns {
		if value == column {
			contains = true
			break
		}
	}

	if !contains {
		fmt.Errorf("invalid column %s", column)
		return false
	}

	return contains
}
