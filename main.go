package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func vote(w http.ResponseWriter, r *http.Request) {
	/*
		Lecture données envoyées
	*/
	body, _ := io.ReadAll(r.Body)
	/*
		Anciennes données
	*/
	var data []map[string]any

	file, err := os.ReadFile("data.json")

	if err == nil {

		json.Unmarshal(file, &data)
	}

	/*
		Nouvelle donnée
	*/

	var newData map[string]interface{}

	json.Unmarshal(body, &newData)

	/*
		Ajout
	*/

	data = append(data, newData)

	/*
		Sauvegarde
	*/

	finalData, _ := json.MarshalIndent(
		data,
		"",
		"  ",
	)

	os.WriteFile(
		"data.json",
		finalData,
		0644,
	)

	/*
		Réponse
	*/

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.Write([]byte(`{
		"success":true,
		"message":"Vote reçu"
	}`))
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./Public/index.html")
}

func FolderPublic(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(
		"/Public/",
		http.FileServer(http.Dir("./Public")),
	).ServeHTTP(w, r)

}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc(
		"/Public/",
		FolderPublic,
	)

	http.HandleFunc(
		"/vote",
		vote,
	)

	port := os.Getenv("PORT")

	http.ListenAndServe(
		":" + port,
		nil,
	)
}
