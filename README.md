# Estructuras de Datos: Stack, Queue y HashTable

Este proyecto implementa y prueba tres estructuras de datos fundamentales: **Stack**, **Queue** y **HashTable**.

## Descripción

Se desarrollaron las estructuras de datos **Stack**, **Queue** y **HashTable** en el paquete `ds`, y se realizaron pruebas para verificar su correcto funcionamiento.

## Estructura del Proyecto

```plaintext
/project-root
│── ds/
│   ├── stack/
│   │   ├── stack.go
│   │   ├── stack_test.go
│   ├── queue/
│   │   ├── queue.go
│   │   ├── queue_test.go
│   ├── hashtable/
│   │   ├── hashtable.go
│   │   ├── hashtable_test.go
│── main.go
│── README.md
```

## Pruebas Realizadas

### **1. Stack (Pila)**

**Operaciones probadas:**
* `Push()`: Agrega elementos a la pila.
* `Pop()`: Extrae elementos de la pila.
* `IsEmpty()`: Verifica si la pila está vacía.
* `ShowStack()`: Muestra el estado actual de la pila.

**Validación:** Se comprueba que sigue el principio **LIFO (Last In, First Out)**.

### **2. Queue (Cola)**

**Operaciones probadas:**
* `Push()`: Agrega elementos a la cola.
* `Dequeue()`: Extrae elementos de la cola.
* `IsEmpty()`: Verifica si la cola está vacía.
* `ShowQueue()`: Muestra el estado actual de la cola.

**Validación:** Se comprueba que sigue el principio **FIFO (First In, First Out)**.

### **3. HashTable (Tabla Hash)**

**Operaciones probadas:**
* `Set()`: Inserta pares clave-valor.
* `Delete()`: Elimina un elemento por su clave.
* `Print()`: Muestra el contenido de la tabla hash.

**Validación:** Se verifica la correcta inserción, eliminación y recuperación de valores.

## Cómo Ejecutar

1. Clona el repositorio.
2. Navega al directorio raíz del proyecto.
3. Ejecuta:

```sh
go run main.go
```

## Tecnologías Usadas

* **Go** (Golang)

## Licencia

Este proyecto está bajo la Licencia MIT.
