# Sistema asincronico simple utilizando GO y RabbitMQ

Este proyecto simula un sistema donde se envian tareas a una cola de mensajes, la cual es consumida por un cliente para visualizar sus nombres respectivos.

Utiliza Go y RabbitMQ para la comunicación asincrónica entre producer y consumer, **sin gRPC ni Docker**.

### Características:

- Producer: Envía tareas a una cola de RabbitMQ.

- Consumer: Recibe las tareas desde RabbitMQ.


### Estructura del proyecto:
```
simple-homework-system/
├── producer/
│   └── main.go
└── consumer/
    └── main.go
```

# Instrucciones para ejecutar el proyecto

1. Clonar el repositorio:
```
git clone https://github.com/tu-usuario/simple-homework-system.git
cd simple-homework-system
```
2. Ejecutar el consumer:
```
cd consumer
go run main.go
```
3. Ejecutar el producer:
```
cd producer
go run main.go
```

Esto enviará una tarea a RabbitMQ, que será recibida y procesada por el consumidor.
