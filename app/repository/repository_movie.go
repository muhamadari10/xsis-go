package repository

import (
	"database/sql"
	"fmt"
	"movie-app/app/model"
	"movie-app/app/model/rpc/movie"
	env "movie-app/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	movieRepo struct {
		ctx      *gin.Context
		env      *env.Dependency
		req      *model.MovieReq
		res      *model.MovieRes
		err      error
		rows     *sql.Rows
		errMulti []string
	}
	MovieRepoInterface interface {
		BindAdd(...any) error
		ValidateForm(...any) error
		AddMovie(...any) error
		ResponseData(error) (*gin.Context, model.StatusData, *movie.AddMovieRes, error)
		ValidateFormUpdate(...any) error
		BindUpdateId() string
		UpdateMovie(...any) error
		ResponseUpdated(error) (*gin.Context, model.StatusData, *movie.AddMovieRes, error)
		BindId() string
		GetMovieDetail(...any) error
		GetMovieList(...any) error
		ResponseList(error) (*gin.Context, model.StatusData, *movie.ListMovieRes, error)
		CountMovie(...any) error
		RemoveById(...any) error
		ResponseRemove(error) (*gin.Context, model.StatusData, *movie.RemoveMovieRes, error)
	}
)

func (m *movieRepo) BindAdd(...any) error {
	return m.ctx.BindJSON(&m.req.Add)
}

func (m *movieRepo) ValidateForm(fn ...any) error {
	m.err, m.errMulti = fn[0].(func(title, desription string, rating float64, types string, id int64) (error, []string))(m.req.Add.Title, m.req.Add.Description, m.req.Add.Rating, "ADD", 0)
	return m.err
}

func (m *movieRepo) AddMovie(fn ...any) error {
	q := fmt.Sprintf(m.req.InsertMovieAdd(), model.TableMovie)
	now := time.Now()
	_, err := m.env.DB.ExecContext(m.ctx, q, m.req.SetAddMovie(m.req.Add, now)...)

	if err != nil {
		return err
	}

	m.res.Add.Data = &movie.MovieData{
		Title:       m.req.Add.Title,
		Description: m.req.Add.Description,
		Image:       m.req.Add.Image,
		Rating:      m.req.Add.Rating,
		CreatedAt:   now.String(),
	}

	return err
}

func (m *movieRepo) ResponseData(err error) (*gin.Context, model.StatusData, *movie.AddMovieRes, error) {
	res := &movie.AddMovieRes{
		Code:    int64(model.Success.GetCode()),
		Message: model.Success.GetMessage(),
		Data:    &movie.MovieData{},
	}

	if err != nil {
		res.Code = int64(model.InvInput.GetCode())
		res.Message = fmt.Sprintf(model.InvInput.GetMessage(), err.Error())
		res.ErrorData = m.errMulti
		return m.ctx, model.InvInput, res, err
	}

	res.Data = m.res.Add.Data
	return m.ctx, model.Success, res, err
}

func (m *movieRepo) ValidateFormUpdate(fn ...any) error {
	m.err, m.errMulti = fn[0].(func(title, desription string, rating float64, types string, id int64) (error, []string))(m.req.Add.Title, m.req.Add.Description, m.req.Add.Rating, "UPDATED", m.req.Add.Id)
	return m.err
}

func (m *movieRepo) BindUpdateId() string {
	id := strconv.Itoa(int(m.req.Add.Id))
	return id
}

func (m *movieRepo) UpdateMovie(fn ...any) error {
	q := fmt.Sprintf(m.req.UpdatedMovie(), model.TableMovie)
	now := time.Now()
	_, err := m.env.DB.ExecContext(m.ctx, q, m.req.SetUpdateMovie(m.req.Add, now)...)
	if err != nil {
		return err
	}

	m.res.Updated.Data = &movie.MovieData{
		Title:       m.req.Add.Title,
		Description: m.req.Add.Description,
		Image:       m.req.Add.Image,
		Rating:      m.req.Add.Rating,
		CreatedAt:   now.String(),
		UpdatedAt:   now.String(),
	}

	return err
}

func (m *movieRepo) ResponseUpdated(err error) (*gin.Context, model.StatusData, *movie.AddMovieRes, error) {
	res := &movie.AddMovieRes{
		Code:    int64(model.Success.GetCode()),
		Message: model.Success.GetMessage(),
		Data:    &movie.MovieData{},
	}

	if err != nil {
		res.Code = int64(model.InvInput.GetCode())
		res.Message = fmt.Sprintf(model.InvInput.GetMessage(), err.Error())
		res.ErrorData = m.errMulti
		return m.ctx, model.InvInput, res, err
	}

	res.Data = m.res.Updated.Data
	return m.ctx, model.Success, res, err
}

func (m *movieRepo) BindId() string {
	return m.ctx.Param("id")
}

func (m *movieRepo) GetMovieDetail(fn ...any) error {
	id := fn[0].(func() string)()
	q := fmt.Sprintf(m.req.GetMovieById(), model.TableMovie)

	err := m.env.DB.QueryRowContext(m.ctx, q, id).Scan()

	if err != nil {
		return err
	}

	m.res.Add.Data = &m.res.MovieById
	return nil
}

func (m *movieRepo) CountMovie(fn ...any) error {
	var CountMovie int8
	id := fn[0].(func() string)()

	q := fmt.Sprintf(m.req.CountMovie(), model.TableMovie)
	err := m.env.DB.QueryRowContext(m.ctx, q, id).Scan(&CountMovie)

	if err != nil {
		return err
	}

	return fn[1].(func(CountMovie int8) error)(CountMovie)
}

func (m *movieRepo) RemoveById(fn ...any) error {
	id := fn[0].(func() string)()
	q := fmt.Sprintf(m.req.DeleteMoveById(), model.TableMovie)
	_, err := m.env.DB.ExecContext(m.ctx, q, id)

	if err != nil {
		return err
	}
	return nil
}

func (m *movieRepo) ResponseRemove(err error) (*gin.Context, model.StatusData, *movie.RemoveMovieRes, error) {
	res := &movie.RemoveMovieRes{
		Code:    int64(model.Success.GetCode()),
		Message: fmt.Sprintf(model.Success.GetMessage(), " Remove By ID"),
	}

	if err != nil {
		res.Code = int64(model.NotFound.GetCode())
		res.Message = err.Error()
		return m.ctx, model.NotFound, res, err
	}

	return m.ctx, model.Success, res, err
}

func (m *movieRepo) GetMovieList(fn ...any) error {

	q := fmt.Sprintf(m.req.GetMovie(), model.TableMovie)

	if m.rows, m.err = m.env.DB.QueryContext(m.ctx, q); m.err != nil {
		return m.err
	}

	defer func() { _ = m.rows.Close() }()
	return fn[0].(func(*sql.Rows, func() []any, func()) error)(m.rows, m.res.SetListMovie, m.appendMovie)
}

func (m *movieRepo) ResponseList(err error) (*gin.Context, model.StatusData, *movie.ListMovieRes, error) {
	res := &movie.ListMovieRes{
		Code:    int64(model.Success.GetCode()),
		Message: model.Success.GetMessage(),
		Data:    []*movie.MovieData{},
	}

	if err != nil {
		res.Code = int64(model.InvInput.GetCode())
		res.Message = fmt.Sprintf(model.InvInput.GetMessage(), err.Error())
		return m.ctx, model.InvInput, res, err
	}

	res.Data = m.req.MovieList

	return m.ctx, model.Success, res, err
}

func (m *movieRepo) appendMovie() {
	m.req.MovieList = append(m.req.MovieList, m.res.MovieDetail)
}

func NewMovieRepository(ctx *gin.Context, env *env.Dependency, reqData *model.MovieReq, resData *model.MovieRes) MovieRepoInterface {
	return &movieRepo{ctx: ctx, env: env, err: nil, req: reqData, res: resData}
}
