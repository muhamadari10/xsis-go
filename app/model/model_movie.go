package model

import (
	"movie-app/app/model/rpc/movie"
	"time"
)

const (
	TableMovie = "movie"
)

type (
	MovieReq struct {
		Add       *movie.AddMovieReq
		Updated   *movie.AddMovieReq
		Detail    *movie.DetailMovieReq
		MovieList []*movie.MovieData
	}
	MovieRes struct {
		Add         *movie.AddMovieRes
		Updated     *movie.AddMovieRes
		Detail      *movie.AddMovieRes
		MovieDetail *movie.MovieData
		MovieById   movie.MovieData
		Remove      *movie.RemoveMovieRes
	}
)

func (m *MovieReq) InsertMovieAdd() string {
	return "INSERT INTO %s" +
		"(title, description, rating, image, created_at, updated_at)" +
		"VALUES($1, $2, $3, $4, $5, $6)"
}

func (m *MovieReq) UpdatedMovie() string {
	return "UPDATE %s SET title=$1, description=$2, rating=$3, image=$4, updated_at=$5  where id = $6"
}

func (m *MovieReq) DeleteMoveById() string {
	return "delete from %s where id = $1"
}

func (m *MovieReq) CountMovie() string {
	return "select count(*) as movie from %s where id = $1"
}

func (m *MovieReq) GetMovieById() string {
	return "Select id, title, description, rating, image, created_at, updated_at from %s where id = $1"
}

func (m *MovieReq) GetMovie() string {
	return "select id, title, description, rating, image, created_at, updated_at from %s "
}

func (m *MovieReq) SetAddMovie(req *movie.AddMovieReq, now time.Time) (data []any) {
	return append(data, req.Title,
		req.Description,
		req.Rating,
		req.Image, now, now)
}

func (m *MovieReq) SetUpdateMovie(req *movie.AddMovieReq, now time.Time) (data []any) {
	return append(data, req.Title,
		req.Description,
		req.Rating,
		req.Image, now, req.Id)
}

func (m *MovieRes) SetListMovie() (data []any) {
	m.MovieDetail = &movie.MovieData{}
	return append(data,
		&m.MovieDetail.Id,
		&m.MovieDetail.Title,
		&m.MovieDetail.Description,
		&m.MovieDetail.Rating,
		&m.MovieDetail.Image,
		&m.MovieDetail.CreatedAt,
		&m.MovieDetail.UpdatedAt,
	)
}
