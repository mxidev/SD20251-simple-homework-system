package main

import (
	"log"
	"time"

	"fmt"

	"github.com/streadway/amqp"
)

func main() {

	// Conexion Basica
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf(" Error al conectar : %v", err)
	}
	defer conn.Close()

	// Apertura de canal
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf(" Error al abrir el canal: %v", err)
	}
	defer channel.Close()

	// Declaracion de la cola de mensajes
	queue, err := channel.QueueDeclare(
		"tareas", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		log.Fatalf(" Error al declarar la cola de mensajes: %v", err)
	}
	defer conn.Close()

	// "Publicacion" de tareas
	for i := 1; i <= 15; i++ {
		body := "Tarea número " + fmt.Sprint(i)
		channel.Publish(
			"",         // exchange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		log.Printf("🟢 Enviado: %s", body)
		time.Sleep(2 * time.Second)
	}
}
