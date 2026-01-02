package migrations

import (
	"github.com/anuragcarret/djang-drf-go/orm/migrations"
)

func init() {
	migrations.GlobalRegistry.Register("demo", &migrations.Migration{
		ID: "20260102_230322",
		Operations: []migrations.Operation{
			
			
			&migrations.RemoveField{
				TableName: "accounts",
				FieldName: "followers",
			},
			
			
			
			&migrations.AlterField{
				TableName: "complex_data",
				FieldName: "status",
				FieldType: "VARCHAR(20) NOT NULL DEFAULT 'active'",
			},
			
			
			
			&migrations.AddField{
				TableName: "posts",
				FieldName: "title",
				FieldType: "VARCHAR(100)",
			},
			
			
		},
	})
}
