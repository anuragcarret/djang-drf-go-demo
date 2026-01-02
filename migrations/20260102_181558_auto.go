package migrations

import (
	"github.com/anuragcarret/djang-drf-go/orm/migrations"
)

func init() {
	migrations.GlobalRegistry.Register("demo", &migrations.Migration{
		ID: "20260102_181558",
		Operations: []migrations.Operation{
			
			
			&migrations.CreateTable{
				Name: "go_groups",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "TEXT UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_permissions",
				Fields: map[string]string{
					"codename": "TEXT UNIQUE NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_outstanding_tokens",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"exp": "TEXT NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"jti": "TEXT UNIQUE NOT NULL",
					"token": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"user_id": "INTEGER NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "accounts",
				Fields: map[string]string{
					"avatar": "TEXT",
					"bio": "TEXT",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"date_joined": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"email": "TEXT UNIQUE NOT NULL",
					"first_name": "TEXT NOT NULL",
					"followers": "TEXT",
					"id": "SERIAL PRIMARY KEY",
					"is_active": "BOOLEAN NOT NULL DEFAULT true",
					"is_staff": "BOOLEAN NOT NULL DEFAULT false",
					"is_superuser": "BOOLEAN NOT NULL DEFAULT false",
					"last_login": "TIMESTAMP WITH TIME ZONE",
					"last_name": "TEXT NOT NULL",
					"location": "TEXT",
					"password": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"username": "TEXT UNIQUE NOT NULL",
					
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
					"author_id": "INTEGER NOT NULL",
					"comments": "TEXT NOT NULL",
					"content": "TEXT NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"likes": "TEXT NOT NULL",
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
			
			
			
			&migrations.CreateTable{
				Name: "comments",
				Fields: map[string]string{
					"author_id": "INTEGER NOT NULL",
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"post_id": "INTEGER NOT NULL",
					"text": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_blacklisted_tokens",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"token": "TEXT NOT NULL",
					"token_id": "INTEGER UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "blog_categories",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"name": "TEXT UNIQUE NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					
				},
			},
			
			
			
			&migrations.CreateTable{
				Name: "go_users",
				Fields: map[string]string{
					"created_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"date_joined": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"email": "TEXT UNIQUE NOT NULL",
					"first_name": "TEXT NOT NULL",
					"id": "SERIAL PRIMARY KEY",
					"is_active": "BOOLEAN NOT NULL DEFAULT true",
					"is_staff": "BOOLEAN NOT NULL DEFAULT false",
					"is_superuser": "BOOLEAN NOT NULL DEFAULT false",
					"last_login": "TIMESTAMP WITH TIME ZONE",
					"last_name": "TEXT NOT NULL",
					"password": "TEXT NOT NULL",
					"updated_at": "TIMESTAMP WITH TIME ZONE NOT NULL",
					"username": "TEXT UNIQUE NOT NULL",
					
				},
			},
			
			
		},
	})
}
