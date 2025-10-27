package handler

import (
	"context"
	"encoding/json"
	serviceregistry "gateway/service-registry"
	"net/http"
	patientpb "patient-service/proto/patientpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	patientID := r.URL.Query().Get("id")
	if patientID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	conn, err := grpc.NewClient(serviceregistry.ServiceRegistry["patient"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := patientpb.NewPatientServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.GetPatient(ctx, &patientpb.PatientRequest{PatientId: patientID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func DoctorHandler(w http.ResponseWriter, r *http.Request) {

}
