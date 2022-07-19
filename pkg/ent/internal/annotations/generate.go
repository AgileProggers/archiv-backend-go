package annotations

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

func Generate(model ent.Interface) []schema.Annotation {
	fields := model.Fields()

	if fields == nil {
		return nil
	}

	var structTag = map[string]string{}

	for _, f := range fields {
		structTag[f.Descriptor().Name] = `json:"` + f.Descriptor().Name + `"`
	}
	return []schema.Annotation{
		field.Annotation{
			StructTag: structTag,
		},
	}
}
