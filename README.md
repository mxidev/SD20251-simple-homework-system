# Sistema asincronico simple utilizando GO y RabbitMQ

Este proyecto simula un sistema donde se envian tareas a una cola de mensajes, la cual es consumida por un cliente para visualizar sus nombres respectivos.

Utiliza Go y RabbitMQ para la comunicación asincrónica entre producer y consumer, **sin gRPC ni Docker**.

### Características:

- Producer: Envía tareas a una cola de RabbitMQ.

- Consumer: Recibe las tareas desde RabbitMQ.
