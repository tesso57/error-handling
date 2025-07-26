 package main

 import (
   "log"
   "net/http"

   "github.com/tesso57/error-handling-sample/internal/presentation"
 )

 func main() {
   handler := presentation.NewRouter()
   log.Println("Server listening on :8080")
   if err := http.ListenAndServe(":8080", handler); err != nil {
     log.Fatalf("failed to start server: %v", err)
   }
 }