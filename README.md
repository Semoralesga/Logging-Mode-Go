# Logging-Mode-Go

Calculadora de dos digitos en GO, con el fin de ver el concepto de modo logging para despliegues en ambientes productivos.

La idea del modo logging en una feature es aplicar la feature en un ambiente productivo pero con la excepción de que no tiene que generar cambios en el sistema. En el caso en que todo el procedimiento funcionase de forma adecuada, loguear de que todo funcionó de forma exitosa. En el caso en que fallé, loguear de que la funcionalidad nueva no funcionó.

El motivo por el cual se hace el deploy de esta forma es para poder corroborar en un ambiente productivo de que una feature nueva en un sistema no va a generar problemas. Si bien uno puede probar en staging, hay veces que lo que diferencia entre ambos ambientes es la data que se maneja (por ejemplo, consultar data de otros sistemas y pedir permisos, entre otros).

Además, esta forma de deploy de la feature es ideal cuando se realizan cambios estructurales de gran magnitud en el sistema.
 
## Ejecución

- Verificar si se tiene instalado Go, de lo contrario seguir los pasos de la [guia](https://go.dev/doc/install).
- Una vez instalado Go nos dirigimos al directorio raiz ejecutamos el comando desde la terminal: ``` go run Calculadora.go```
- Veremos la siguiente informacion: 
    ``` 
    ************** Calculadora de 2 Digitos **************
  1. Sumar
  2. Restar
  * Cualquier otro valor finalizara la ejecución
    Digite el numero de la operación que desea hacer:
    ```
- Donde podras seleccionar la operación de tu preferencia, "1" para este ejemplo:
  ```
  Digite el numero de la operación que desea hacer: 1
  ``` 
- Seguido se solicitaran los numeros a operar separados por coma, ej. 10,30 y finalmente retornara el resultado de la suma, junto con el resultado de la división en modo logging 
  ```
  Digite el numero de la operación que desea hacer: 1
  Digite los numeros a operar separados por coma(1,2): 10,30
  La suma es: 40
  2023/03/07 15:08:48  [Logging] La división es: 0.33333334
  ```
- En este caso el modo logging nos sirve para detectar posibles errores en un feature nuevo sin alterar la logica o funcionalidad de los features actuales, por ejemplo cuando dividimos por 0.
  ```
  ************** Calculadora de 2 Digitos **************
  1. Sumar
  2. Restar
  * Cualquier otro valor finalizara la ejecución
  Digite el numero de la operación que desea hacer: 1
  Digite los numeros a operar separados por coma(1,2): 10,0
  >> Error: El divisor no puede ser igual a 0
  >> 2023/03/07 15:41:05  [Logging] La división es: +Inf
  La suma es: 10
     ```
  Podemos observar que a pesar que se registra el error en la división, la ejecución del programa continua retornado el valor de la suma.