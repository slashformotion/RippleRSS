watch:
	echo "db/query.sql" | entr -s "echo rebuild queries; sqlc generate"
