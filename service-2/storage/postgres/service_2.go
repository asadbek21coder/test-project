package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/asadbek21coder/test-project/service-2/pkg/helper"
	"github.com/asadbek21coder/test-project/service-2/storage"
	"github.com/jackc/pgx/v4/pgxpool"

	pb "github.com/asadbek21coder/test-project/service-2/genproto/service_2"
)

type service_2_Repo struct {
	db *pgxpool.Pool
}

func NewService_2_Repo(db *pgxpool.Pool) storage.Service_2_I {
	return &service_2_Repo{
		db: db,
	}
}

func (r *service_2_Repo) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	var (
		resp   pb.GetAllResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND title ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM posts WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				user_id,
				title,
				body
			FROM posts
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var posts pb.Post

		err = rows.Scan(
			&posts.Id,
			&posts.UserId,
			&posts.Title,
			&posts.Body,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning profession err: %w", err)
		}

		resp.Posts = append(resp.Posts, &posts)
	}

	return &resp, nil

}

func (r *service_2_Repo) GetById(ctx context.Context, entity *pb.Id) (*pb.Post, error) {
	var resp pb.Post

	query := `SELECT * FROM posts WHERE id=$1`

	row, err := r.db.Query(
		ctx,
		query,
		entity.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting post by given id err: %w", err)
	}

	for row.Next() {
		err = row.Scan(
			&resp.Id,
			&resp.UserId,
			&resp.Title,
			&resp.Body,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning Post err: %w", err)
		}

	}

	return &resp, nil
}

func (r *service_2_Repo) Update(ctx context.Context, entity *pb.Post) (*pb.Post, error) {
	row := r.db.QueryRow(ctx, "select exists (select id from posts where id=$1) as output", entity.Id)
	var exists bool
	if row.Scan(&exists); !exists {
		return nil, fmt.Errorf("checking DB: %w", errors.New("no such id"))
	}

	query := `UPDATE posts SET user_id=$1, title=$2, body=$3 WHERE id=$4 returning *`
	_, err := r.db.Exec(ctx, query, entity.UserId, entity.Title, entity.Body, entity.Id)
	if err != nil {
		return nil, fmt.Errorf("error while updating posts err: %w", err)
	}

	return entity, nil
}

func (r *service_2_Repo) Delete(ctx context.Context, entity *pb.Id) (*pb.Id, error) {
	row := r.db.QueryRow(ctx, "select exists (select id from posts where id=$1) as output", entity.Id)
	var exists bool
	if row.Scan(&exists); !exists {
		return nil, fmt.Errorf("checking DB: %w", errors.New("no such id"))
	}

	query := `DELETE FROM posts WHERE id=$1`

	_, err := r.db.Exec(
		ctx,
		query,
		entity.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("error while deleting post err: %w", err)
	}

	return entity, nil
}
