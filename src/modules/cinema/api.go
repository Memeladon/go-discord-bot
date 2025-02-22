package cinema

import (
	"encoding/json"
	"net/http"
	"os"
)

const KinopoiskApiHost = "https://api.kinopoisk.dev"
const PosterFallback = "https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Question_in_a_question_in_a_question_in_a_question.gif/1024px-Question_in_a_question_in_a_question_in_a_question.gif"

var cache = make(map[string]*KinopoiskMovie)

type Poster struct {
	Url        string `json:"url"`
	PreviewUrl string `json:"previewUrl"`
}

type Genre struct {
	Name string `json:"name"`
}

type KinopoiskMovie struct {
	Id              uint    `json:"id"`
	Name            string  `json:"name"`
	AlternativeName string  `json:"alternativeName"`
	Description     string  `json:"description"`
	Year            uint    `json:"year"`
	MovieLength     uint    `json:"movieLength"`
	Genres          []Genre `json:"genres"`
	Poster          Poster  `json:"poster"`
}

func (m *KinopoiskMovie) Init(id string) error {
	//API имеет ограничение поэтому тоже кэшируем (на всякий случай)
	if cachedMovie, found := cache[id]; found {
		*m = *cachedMovie
		return nil
	}

	req, _ := http.NewRequest("GET", KinopoiskApiHost+"/v1.4/movie/"+id, nil)
	req.Header.Set("X-API-KEY", os.Getenv("KINOPOISK_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(m)
	if err != nil {
		return err
	}

	if m.Poster.Url == "" {
		m.Poster.Url = PosterFallback
	}

	if m.Poster.PreviewUrl == "" {
		m.Poster.PreviewUrl = PosterFallback
	}

	if m.Name == "" {
		m.Name = m.AlternativeName
	}

	cache[id] = m
	return nil
}
