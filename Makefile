.PHONY: help db-generate clean

help:
	@echo "Database Commands"
	@echo ""
	@echo "Code Generation:"
	@echo "  make db-generate      Generate Go code from SQL (sqlc)"
	@echo ""
	@echo "Maintenance:"
	@echo "  make clean            Clean generated SQL code"

db-generate:
	@echo "Generating Go code from SQL queries..."
	sqlc generate
	@echo "Database code generated!"

clean:
	@echo "Cleaning generated SQL code..."
	rm -rf generated/*.go
	@echo "Clean complete!"
