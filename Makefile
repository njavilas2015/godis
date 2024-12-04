BINARY_NAME=godis

build:
	@echo "Compilando el proyecto..."
	go build -o $(BINARY_NAME) .

run:
	@echo "Ejecutando la aplicación..."
	./$(BINARY_NAME)

test:
	@echo "Ejecutando las pruebas..."
	go test ./...

clean:
	@echo "Limpiando los archivos generados..."
	rm -f $(BINARY_NAME)

tidy:
	@echo "Limpiando y actualizando las dependencias..."
	go mod tidy

docs:
	@echo "Generando la documentación..."
	go doc

vendor:
	@echo "Actualizando las dependencias..."
	go mod vendor

lint:
	@echo "Ejecutando el linter..."
	golint .

all: build test clean
