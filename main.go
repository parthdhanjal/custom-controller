package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
)

func handleMutate(w http.ResponseWriter, r *http.Request) {
	input := &admissionv1.AdmissionReview{}
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		sendErr(w, fmt.Errorf("Could not unmarshal review: %v", err))
		return
	}
}

func sendErr(w http.ResponseWriter, err error) {
	out, err := json.Marshal(map[string]string{"Err": err.Error()})
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(out)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", handleMutate)
	srv := &http.Server{Addr: ":443", Handler: mux}
	log.Fatal(srv.ListenAndServeTLS("/certs/webhook.crt", "/certs/webhook-key.pem"))
}
