package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/asadbek21coder/test-project/service-1/storage"
	"github.com/jackc/pgx/v4/pgxpool"

	pb "github.com/asadbek21coder/test-project/service-1/genproto/service_1"
)

type service_1_Repo struct {
	db *pgxpool.Pool
}

type Response struct {
	Meta Meta    `json:"meta"`
	Data []Posts `json:"data"`
}

type Posts struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Meta struct {
	Pagination Pagination
}

type Pagination struct {
	Total int   `json:"total"`
	Pages int   `json:"page"`
	Page  int   `json:"pages"`
	Limit int   `json:"limit"`
	Links Links `json:"links"`
}
type Links struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

func NewService_1_Repo(db *pgxpool.Pool) storage.Service_1_I {
	return &service_1_Repo{
		db: db,
	}
}

var errorResponse error

var waitgroup sync.WaitGroup

func (r *service_1_Repo) GetAll(ctx context.Context) (*pb.Status, error) {
	for i := 1; i < 50; i = i + 10 {
		go getter(r, ctx, i, i+9)
		waitgroup.Add(1)
	}
	waitgroup.Wait()

	if errorResponse != nil {
		return nil, errorResponse
	} else {
		return &pb.Status{
			Status: string("OK"),
		}, nil
	}

}
func getter(r *service_1_Repo, ctx context.Context, start, end int) {
	defer waitgroup.Done()

	for i := start; i <= end; i++ {
		url := "https://gorest.co.in/public/v1/posts?page=" + fmt.Sprint(i)
		if i == 1 {
			url = url[:len(url)-7]
		}
		response, err := http.Get(url)
		if err != nil {
			errorResponse = err
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			errorResponse = err
		}
		var data Response
		json.Unmarshal(responseData, &data)
		for _, v := range data.Data {
			query := `INSERT into posts (id,user_id, title,body) VALUES ($1,$2,$3,$4)`

			_, err := r.db.Exec(ctx, query, v.Id, v.UserId, v.Title, v.Body)
			if err != nil {
				errorResponse = err
			}

		}

	}
}
