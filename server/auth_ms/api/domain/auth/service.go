package auth

import (
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	Client *redis.Client
	Url string
}

func (r Repository) Create() []byte {
	// query := "/find?username=root"
	req, _ := http.NewRequest("GET","http://localhost:1000/api/v1/user/find?username=root", nil)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	return b
} 