all:
	@echo Please specify the target. >&2
	@exit 1

# graphql schemaからgqlgenのコードを生成する
gen:
	go run github.com/99designs/gqlgen generate
	go run cmd/gqlgen/add-directive-tag.go

