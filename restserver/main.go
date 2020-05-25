package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	fileMongo "calvinechols.com/vault/database/mongo"
	"calvinechols.com/vault/env"
	"calvinechols.com/vault/file"
	"calvinechols.com/vault/filestore/disk"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Start is a method that starts the server.
func main() {
	var mongoConnectionString = env.Get("MONGODB_DATA_CONNECTION", "mongodb://root:password@localhost:27017")
	options := options.Client().ApplyURI(mongoConnectionString)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatalf("mongo connection failed: %v", err)
		return
	}
	userHandler, err := buildFileHandler(client)
	if err != nil {
		log.Fatalf("user handler creation failed: %v", err)
		return
	}
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/putfile", userHandler.PutFile).Methods("POST", "OPTIONS")

	http.Handle("/", accessControl(mux))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func buildFileHandler(client *mongo.Client) (file.Handler, error) {
	maxFileSize, err := strconv.ParseInt(env.Get("VAULT_MAX_FILE_UPLOAD_SIZE", "10000000000"), 10, 64)
	if err != nil {
		return nil, err
	}
	savePath := env.Get("VAULT_DISK_SAVE_PATH", "../../vault_files")
	diskRepo := disk.NewFileStoreDiskRepo(savePath)
	fileRepo := fileMongo.NewFileMongoRepo(client)
	fileService := file.NewFileService(fileRepo, diskRepo)
	fileHandler := file.NewFileHandler(fileService, maxFileSize)
	return fileHandler, nil
}

// CORS control
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// func getMongoConnection(connectionString string) *mongo.Client
