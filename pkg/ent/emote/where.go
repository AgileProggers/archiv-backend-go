// Code generated by entc, DO NOT EDIT.

package emote

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Emote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Emote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Emote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Emote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Emote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.Provider) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Emote) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Emote) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Emote) predicate.Emote {
	return predicate.Emote(func(s *sql.Selector) {
		p(s.Not())
	})
}
