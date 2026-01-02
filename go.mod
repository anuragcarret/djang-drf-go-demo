module demo

go 1.25.5

replace github.com/anuragcarret/djang-drf-go => ../django_drf_go

require github.com/anuragcarret/djang-drf-go v0.0.0-00010101000000-000000000000

require (
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
