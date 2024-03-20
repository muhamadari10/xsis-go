package viewmodel

import (
	"movie-app/app/model"
	"movie-app/app/model/rpc/movie"

	"github.com/gin-gonic/gin"
)

func VMMovieAdd(c *gin.Context, status model.StatusData, res *movie.AddMovieRes, err error) {
	c.JSON(model.StatusData(status).GetCode(), res)
}

func VMMovieList(c *gin.Context, status model.StatusData, res *movie.ListMovieRes, err error) {
	c.JSON(model.StatusData(status).GetCode(), res)
}

func VMMovieRemove(c *gin.Context, status model.StatusData, res *movie.RemoveMovieRes, err error) {
	c.JSON(model.StatusData(status).GetCode(), res)
}
