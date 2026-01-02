package migrations

import (
	"github.com/anuragcarret/djang-drf-go/orm/migrations"
)

func init() {
	migrations.GlobalRegistry.Register("demo", &migrations.Migration{
		ID: "20260102_183244",
		Operations: []migrations.Operation{
			
			
			&migrations.CreateTable{
				Name: "complex_data",
				Fields: map[string]string{
					"activation_date": "DATE",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"external_id": "UUID UNIQUE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"is_processed": "BOOLEAN NOT NULL DEFAULT false",
					"metadata": "JSONB NOT NULL",
					"score": "NUMERIC NOT NULL",
					"status": "VARCHAR(20) NOT NULL DEFAULT 'active'",
					"tags": "TEXT[] NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
		},
	})
}
