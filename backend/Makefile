db-schema:
	go run -mod=mod entgo.io/ent/cmd/ent new $(db-schema)

gen:
	go generate .

run:
	ENV=dev go run ./main.go