package gachani

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	mal "github.com/fuzzylimes/malgomate"
)

func SeasonHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	res, err := GetSeasons(query)
	if err != nil {
		switch e := err.Error(); e {
		case "invalid_query":
			http.Error(w, e, http.StatusBadRequest)
		case "client_error":
			http.Error(w, e, http.StatusInternalServerError)
		case "decode_error":
			http.Error(w, e, http.StatusInternalServerError)
		default:
			http.Error(w, e, http.StatusInternalServerError)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, res)
}

func GetSeasons(q url.Values) (string, error) {
	query := &mal.SeasonalQuery{
		Fields: []mal.QueryField{
			mal.FieldID,
			mal.FieldTitle,
			mal.FieldMainPicture,
			mal.FieldSynopsis,
			mal.FieldStartDate,
			mal.FieldEndDate,
			mal.FieldMean,
			mal.FieldGenres,
			mal.FieldNsfw,
			mal.FieldNumEpisodes,
			mal.FieldAverageEpisodeDuration,
			mal.FieldAlternativeTitles,
		},
		Limit: 10,
	}

	if season := q.Get("season"); season != "" && mal.SeasonTypeQueries.IsValid(season) {
		query.Season = mal.Season(season)
	}

	if year := q.Get("year"); year != "" {
		if yearInt, err := strconv.Atoi(year); err != nil {
			return "", errors.New("invalid_query")
		} else {
			query.Year = yearInt
		}
	}

	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetSeason(query)

	if err != nil {
		return "", errors.New("client_error")
	}

	if j, err := res.JSON(); err != nil {
		return "", errors.New("decode_error")
	} else {
		return j, nil
	}

}
