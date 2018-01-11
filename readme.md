[![Build Status](https://travis-ci.org/sait/tddexample.svg?branch=master)](https://travis-ci.org/sait/tddexample)
[![Build status](https://ci.appveyor.com/api/projects/status/w8bjo1a5vr43frvn?svg=true)](https://ci.appveyor.com/project/xerardoo/tddexample)


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


###Documentacion Relacionada

*Test Driven Development en Go con Ginkgo y Gomega* ~
https://www.youtube.com/watch?v=5PMuFfBjpuQ

*Ginkgo and Gomega: BDD-style Testing in Go with Onsi Fakhouri* ~
1. https://www.youtube.com/watch?v=rGHu8IvGzNM
2. https://www.youtube.com/watch?v=xn6Erpr2p0o
3. https://www.youtube.com/watch?v=6XbEyZYNp4g

*Ginkgo y Gomega* ~
https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo

*Configurar AppVeyor CI* ~
https://blog.markvincze.com/setting-up-an-appveyor-pipeline-for-golang/

*Ejemplo AppVeyor.yml* ~
https://www.appveyor.com/docs/appveyor-yml/

*Webinar sobre Integración Continua - atSistemas* ~
https://www.youtube.com/watch?v=uizji8-I5_w