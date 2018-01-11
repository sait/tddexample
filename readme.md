[![Build Status](https://travis-ci.org/sait/golang-tdd-example.svg?branch=master)](https://travis-ci.org/sait/golang-tdd-example)


# Ejemplo de TDD
 Ejemplo de como desarrollar una app guiado por pruebas

## Tecnologia

- Go
- Ginkgo: test framework 
- Gomega: libreria de asersiones


## Metodologia

1. Definir las especificaciones o historias de usuario
2. Codificar la estructura de las pruebas con ginkgo y gomega
3. Ejecutar pruebas ** si no pasan se corrigen
4. Crear estructura de la applicacion/funcionalidad (modelos y metodos) y se simulan los datos (mocking) ** lo minimo para funcionar
5. Codificar las pruebas acorde a la nueva estructura de la applicacion/funcionalidad
6. Implementar la aplicacion y ejecutar pruebas



## Ejecutar pruebas

```
#cambiar de directorio
cd test/

# poner parametro "-v" para ver mas detalle en las pruebas
go test -v

```