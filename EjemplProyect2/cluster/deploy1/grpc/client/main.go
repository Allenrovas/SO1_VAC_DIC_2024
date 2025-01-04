package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	pb "main/grpcClient"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

var ctx = context.Background()

// Contador para alternar entre gRPC y Rust
var requestCounter int32 = 0

// Dirección del servidor Rust
const rustServerURL = "http://localhost:4000/insert" // Cambia el puerto si es necesario

// Estructura para el nuevo formato de JSON
type AssignmentData struct {
	Curso    string `json:"curso"`
	Facultad string `json:"facultad"`
	Carrera  string `json:"carrera"`
	Region   string `json:"region"`
}

func insertData(c *fiber.Ctx) error {
	// Parsear el cuerpo de la solicitud al nuevo formato de JSON
	var data AssignmentData
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Alternar entre gRPC y Rust
	if atomic.AddInt32(&requestCounter, 1)%2 == 0 {
		go sendRustServer(data)
	} else {
		go sendMongoServer(data)
	}

	return c.JSON(fiber.Map{"message": "Solicitud procesada correctamente"})
}

func sendMongoServer(data AssignmentData) {
	// Conectar al servidor gRPC
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("No se pudo conectar al servidor gRPC", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error al cerrar la conexión: %v", err)
		}
	}(conn)

	// Enviar los datos al servidor gRPC
	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Curso:    data.Curso,
		Facultad: data.Facultad,
		Carrera:  data.Carrera,
		Region:   data.Region,
	})
	if err != nil {
		log.Fatalf("Error al enviar la información al servidor gRPC: %v", err)
	}

	fmt.Println("Respuesta del servidor gRPC: ", ret.GetInfo())
}

func sendRustServer(data AssignmentData) {
	// Convertir la estructura AssignmentData a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error al convertir los datos a JSON: %v", err)
	}

	// Enviar la solicitud al servidor Rust
	resp, err := http.Post(rustServerURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error al enviar la información al servidor Rust: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: el servidor Rust devolvió el estado %d", resp.StatusCode)
	} else {
		fmt.Println("Datos enviados al servidor Rust correctamente")
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	app.Post("/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
