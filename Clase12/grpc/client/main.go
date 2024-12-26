package main

import (
	"context"
	"fmt"
	"log"

	pb "main/grpcClient"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

var ctx = context.Background()

type Data struct {
	Carnet   int32
	Nombre   string
	Curso    string
	Nota     int32
	Semestre string
	Year     int32
}

func insertData(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}
	var carnetInt int32
	if carnetFloat, ok := data["carnet"].(float64); ok {
		carnetInt = int32(carnetFloat)
	}

	var yearInt int32
	if yearFloat, ok := data["year"].(float64); ok {
		yearInt = int32(yearFloat)
	}

	var notaInt int32
	if notaFloat, ok := data["nota"].(float64); ok {
		notaInt = int32(notaFloat)
	}

	nota := Data{
		Carnet:   carnetInt,
		Nombre:   data["nombre"].(string),
		Curso:    data["curso"].(string),
		Nota:     notaInt,
		Semestre: data["semestre"].(string),
		Year:     yearInt,
	}

	go sendMongoServer(nota)

	return nil
}

func sendMongoServer(data Data) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalln("No se pudo conectar al servidor de Mongo", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error al cerrar la conexión: %v", err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Carnet:   data.Carnet,
		Nombre:   data.Nombre,
		Curso:    data.Curso,
		Nota:     data.Nota,
		Semestre: data.Semestre,
		Year:     data.Year,
	})
	if err != nil {
		log.Fatalf("Error al enviar la información: %v", err)
	}

	fmt.Println("Respuesta del servidor de Mongo: ", ret.GetInfo())
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
