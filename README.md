# Sistema asincronico simple utilizando GO y RabbitMQ

Este proyecto simula un sistema donde se envian tareas a una cola de mensajes, la cual es consumida por un cliente para visualizar sus nombres respectivos (Tarea 1, Tarea 2, ...).

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

Es posible que requiera instalar rabbit en su sistema (Linux):
```
sudo apt install rabbitmq-server
```

1. Clonar el repositorio:
```
git clone https://github.com/mxidev/simple-homework-system.git
cd simple-homework-system
```
2. Inicializar modulo Go
```
go mod init <mod_name>
go mod tidy
```
3. Instalacion de dependencia para utilizar RabbitMQ
```
go get github.com/streadway/amqp
```
4. Ejecutar el consumer:
```
cd consumer
go run main.go
```
5. Ejecutar el producer:
```
cd producer
go run main.go
```

Esto enviará una tarea a RabbitMQ, que será recibida y procesada por el consumidor.
