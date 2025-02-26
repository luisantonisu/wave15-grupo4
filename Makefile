# Variables
TEST_DIR := ./internal/handler/... ./internal/service/...
COVERAGE_FILE := coverage_temp.out
COVERAGE_FINAL := coverage.out
COVERAGE_HTML := coverage.html
EXCLUDE_FILES := _mock.go

.PHONY: test coverage html-coverage total-coverage lint

all:

# Ejecutar los tests
test:
	go test -v $(TEST_DIR)

# Generar reporte de coverage
coverage:
	go test -coverprofile=$(COVERAGE_FILE) $(TEST_DIR)
	grep -v "$(EXCLUDE_FILES)" $(COVERAGE_FILE) > $(COVERAGE_FINAL)
	rm -r $(COVERAGE_FILE)

# Generar reporte HTML de coverage
report: coverage
	go tool cover -html=$(COVERAGE_FINAL) -o $(COVERAGE_HTML)
	@echo "Coverage HTML report generated: $(COVERAGE_HTML)"ˇ
	open $(COVERAGE_HTML)
	

# Obtener información del coverage total
total-coverage: coverage
	go tool cover -func=$(COVERAGE_FINAL)

# Ejecutar el linter
lint:
	staticcheck ./...

# Limpiar archivos generados
clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_FINAL) $(COVERAGE_HTML)