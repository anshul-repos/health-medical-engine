package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "patient-service/proto/patientpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPatientServiceServer
}

func (s *server) GetPatient(ctx context.Context, req *pb.PatientRequest) (*pb.PatientResponse, error) {
	log.Println("PatientService: GetPatient called for", req.PatientId)
	return &pb.PatientResponse{
		PatientId: req.PatientId,
		Name:      "Anshul Tiwari",
		Age:       30,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterPatientServiceServer(s, &server{})
	fmt.Println("PatientService listening on :50051")
	log.Fatal(s.Serve(lis))
}
