package tutor

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

type Tutor struct {
	Email    string     `json:"email"`
	Subjects []Subjects `json:"subjects"`
}

type Subjects struct {
	Name string `json:"name"`
}

type DefaultResponse struct {
	Value            []Tutor  `json:"value"`
	IsSuccess        bool     `json:"IsSuccess"`
	ErrorModel       string   `json:"errorModel"`
	ValidationErrors []string `json:"validationErrors"`
}

type Actions interface {
	GetAll() []Tutor
	GetById(id int) Tutor
}

func (t *Tutor) GetAll() []Tutor {
	var response DefaultResponse
	resp, err := http.Get("http://localhost:5273/tutors/all")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range response.Value {
		fmt.Println(t)
	}
	//tutors = []Tutor{
	//	{Email: "test@gmail.com", Id: 1, Specialises: []string{"english", "russian"}},
	//	{Email: "test@yandex.ru", Id: 2, Specialises: []string{"c#", "java"}},
	//}
	var tutors []Tutor
	return tutors
}

func (t *Tutor) GetById(id int) Tutor {
	return Tutor{"email", []Subjects{{"rus"}}}
}
