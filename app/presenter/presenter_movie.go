package presenter

import (
	"movie-app/app/model"
	"movie-app/app/model/rpc/movie"
	"movie-app/app/repository"
	"movie-app/app/usecase"
	env "movie-app/config"

	"github.com/gin-gonic/gin"
)

func Add(env *env.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.MovieReq{Add: &movie.AddMovieReq{}}
		res := &model.MovieRes{Add: &movie.AddMovieRes{}}
		rO := repository.NewMovieRepository(c, env, req, res)
		uO := usecase.NewMovieUseCase(rO)
		uO.AddMovie()
	}
}

func Detail(env *env.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.MovieReq{Add: &movie.AddMovieReq{}}
		res := &model.MovieRes{Add: &movie.AddMovieRes{}}
		rO := repository.NewMovieRepository(c, env, req, res)
		uO := usecase.NewMovieUseCase(rO)
		uO.DetailMovie()
	}
}

func List(env *env.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.MovieReq{}
		res := &model.MovieRes{Add: &movie.AddMovieRes{}}
		rO := repository.NewMovieRepository(c, env, req, res)
		uO := usecase.NewMovieUseCase(rO)
		uO.ListMovie()
	}
}

func Remove(env *env.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.MovieReq{}
		res := &model.MovieRes{}
		rO := repository.NewMovieRepository(c, env, req, res)
		uO := usecase.NewMovieUseCase(rO)
		uO.RemoveMovie()
	}
}

func Updated(env *env.Dependency) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &model.MovieReq{Updated: &movie.AddMovieReq{}}
		res := &model.MovieRes{Updated: &movie.AddMovieRes{}}
		rO := repository.NewMovieRepository(c, env, req, res)
		uO := usecase.NewMovieUseCase(rO)
		uO.UpdatedMovie()
	}
}
