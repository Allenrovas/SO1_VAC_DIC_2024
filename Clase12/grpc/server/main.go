package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "main/grpcServer"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var mongoClient *mongo.Client
var mongoCollection *mongo.Collection

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	// serverPort is the port where the server will listen
	serverPort = ":3001"
)

type Data struct {
	Carnet   int32  `bson:"carnet"`
	Nombre   string `bson:"nombre"`
	Curso    string `bson:"curso"`
	Nota     int32  `bson:"nota"`
	Semestre string `bson:"semestre"`
	Year     int32  `bson:"year"`
}

func mongoConnect() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("MONGO_COLLECTION")

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mongoCollection = mongoClient.Database(dbName).Collection(collectionName)

	fmt.Println("Connected to MongoDB")
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibi nota de carnet: ", in.GetCarnet())
	data := Data{
		Carnet:   in.GetCarnet(),
		Nombre:   in.GetNombre(),
		Curso:    in.GetCurso(),
		Nota:     in.GetNota(),
		Semestre: in.GetSemestre(),
		Year:     in.GetYear(),
	}
	insertMongo(data)
	return &pb.ReplyInfo{Info: "Nota recibida"}, nil
}

func insertMongo(data Data) {
	_, err := mongoCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted in MongoDB")
}

func main() {
	mongoConnect()
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
