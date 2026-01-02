package migrations

import (
	"github.com/anuragcarret/djang-drf-go/orm/migrations"
)

func init() {
	migrations.GlobalRegistry.Register("demo", &migrations.Migration{
		ID: "20260102_223940",
		Operations: []migrations.Operation{
			
			
			&migrations.CreateTable{
				Name: "go_permissions",
				Fields: map[string]string{
					"codename": "VARCHAR(100) UNIQUE NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "VARCHAR(255) NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_blacklisted_tokens",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"token": "TEXT NOT NULL",
					"token_id": "BIGINT UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "blog_categories",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "VARCHAR(50) UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "comments",
				Fields: map[string]string{
					"author_id": "BIGINT NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"post_id": "BIGINT NOT NULL",
					"text": "VARCHAR(500) NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
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
			
			
			
			&migrations.CreateTable{
				Name: "go_users",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"date_joined": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"email": "VARCHAR(254) UNIQUE NOT NULL",
					"first_name": "VARCHAR(30) NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"is_active": "BOOLEAN NOT NULL DEFAULT true",
					"is_staff": "BOOLEAN NOT NULL DEFAULT false",
					"is_superuser": "BOOLEAN NOT NULL DEFAULT false",
					"last_login": "TEXT",
					"last_name": "VARCHAR(150) NOT NULL",
					"password": "VARCHAR(128) NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"username": "VARCHAR(150) UNIQUE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_groups",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "VARCHAR(150) UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_outstanding_tokens",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"exp": "BIGINT NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"jti": "TEXT UNIQUE NOT NULL",
					"token": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"user_id": "BIGINT NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "accounts",
				Fields: map[string]string{
					"avatar": "TEXT",
					"bio": "TEXT",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"date_joined": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"email": "VARCHAR(254) UNIQUE NOT NULL",
					"first_name": "VARCHAR(30) NOT NULL",
					"followers": "TEXT[]",
					"id": "SERIAL PRIMARY KEY",
					"is_active": "BOOLEAN NOT NULL DEFAULT true",
					"is_staff": "BOOLEAN NOT NULL DEFAULT false",
					"is_superuser": "BOOLEAN NOT NULL DEFAULT false",
					"last_login": "TEXT",
					"last_name": "VARCHAR(150) NOT NULL",
					"location": "VARCHAR(100)",
					"password": "VARCHAR(128) NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"username": "VARCHAR(150) UNIQUE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "account_follows",
				Fields: map[string]string{
					"follower_id": "INTEGER NOT NULL",
					"following_id": "INTEGER NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "posts",
				Fields: map[string]string{
					"author_id": "BIGINT NOT NULL",
					"comments": "TEXT[] NOT NULL",
					"content": "VARCHAR(280) NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"likes": "TEXT[] NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "post_likes",
				Fields: map[string]string{
					"account_id": "INTEGER NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"post_id": "INTEGER NOT NULL",
					
				},
			},
			
			
		},
	})
}
