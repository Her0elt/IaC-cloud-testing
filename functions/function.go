package function

import (
	"encoding/json"
	"errors"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:message`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := vars["id"]
	res, _ := json.Marshal(albums)
	if ok {
		album, err := getAlbumsById(id[0])
		if err != nil {
			msg := ErrorMessage{Message: err.Error()}
			res, _ = json.Marshal(msg)
		} else {
			res, _ = json.Marshal(album)
		}
		w.Write(res)
		return
	}
	w.Write(res)
}

func getAlbumsById(id string) (*album, error) {
	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, errors.New("No album with that Id")
}
