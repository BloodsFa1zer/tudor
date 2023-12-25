swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs

sql:
	sqlc generate 