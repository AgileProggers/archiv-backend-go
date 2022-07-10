package helpers

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)


func BuildPredicates(columns []string, query map[string][]string) (func(selector *sql.Selector), error) {
    for key, value := range query {
        if len(value) == 0 {
            continue
        }
        contains := false
        for _, column := range columns {
            if column == value[0] {
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

