package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"movie-app/app/model"
	"movie-app/app/repository"
	"movie-app/app/viewmodel"
)

type (
	movieUseCase struct {
		repo repository.MovieRepoInterface
		err  error
	}
	MovieUseCaseInterface interface {
		AddMovie()
		DetailMovie()
		ListMovie()
		RemoveMovie()
		UpdatedMovie()
	}
)

func (m *movieUseCase) AddMovie() {
	m.repo.BindAdd()                          // bind data
	m.err = m.repo.ValidateForm(validateForm) // validate
	m.repo.AddMovie(validateForm)
	viewmodel.VMMovieAdd(m.repo.ResponseData(m.err))
}

func (m *movieUseCase) UpdatedMovie() {
	m.repo.BindAdd()                                // bind data
	m.err = m.repo.ValidateFormUpdate(validateForm) // validate
	m.err = m.repo.CountMovie(m.repo.BindUpdateId, m.validateMovie)
	m.repo.UpdateMovie()
	viewmodel.VMMovieAdd(m.repo.ResponseUpdated(m.err))
}

func (m *movieUseCase) ListMovie() {
	m.err = m.repo.GetMovieList(m.dbScan)
	viewmodel.VMMovieList(m.repo.ResponseList(m.err))
}

func (m *movieUseCase) DetailMovie() {
	m.repo.GetMovieDetail(m.repo.BindId)
	viewmodel.VMMovieAdd(m.repo.ResponseData(m.err))
}

func (m *movieUseCase) RemoveMovie() {
	m.err = m.repo.CountMovie(m.repo.BindId, m.validateMovie)
	m.repo.RemoveById(m.repo.BindId)
	viewmodel.VMMovieRemove(m.repo.ResponseRemove(m.err))
}

func (m *movieUseCase) validateMovie(countMovie int8) error {
	if countMovie < 1 {
		return errors.New(model.NotFound.GetMessage())
	}
	return nil
}

func (e *movieUseCase) dbScan(rows *sql.Rows, data func() []any, append func()) (err error) {
	for rows.Next() {
		if err = rows.Scan(data()...); err != nil {
			fmt.Println(err)
			return
		}

		append()
	}
	return
}

func validateForm(title, desription string, rating float64, types string, id int64) (error, []string) {
	var stringMsg []string
	tamplate := "%s harus di isi"

	if types == "UPDATED" && id == 0 {
		stringMsg = append(stringMsg, fmt.Sprintf(tamplate, "Id"))
	}

	if title == "" {
		stringMsg = append(stringMsg, fmt.Sprintf(tamplate, "Title"))
	}

	if desription == "" {
		stringMsg = append(stringMsg, fmt.Sprintf(tamplate, "Description"))
	}

	if rating == 0 {
		stringMsg = append(stringMsg, fmt.Sprintf(tamplate, "Rating"))
	}

	if title == "" || desription == "" || rating <= 0 {
		return errors.New("Invalid Input"), stringMsg
	}

	return nil, stringMsg
}

func NewMovieUseCase(repo repository.MovieRepoInterface) MovieUseCaseInterface {
	return &movieUseCase{repo: repo}
}
