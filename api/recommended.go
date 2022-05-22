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

func RecommendedHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	res, err := getRecommended(query)
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

func getRecommended(q url.Values) (string, error) {
	query := &mal.DetailsQuery{
		Fields: []mal.DetailField{
			mal.DetailRecommendations.SubFields(&mal.DetailFields{
				mal.DetailID,
				mal.DetailTitle,
				mal.DetailMainPicture,
				mal.DetailSynopsis,
				mal.DetailStartDate,
				mal.DetailEndDate,
				mal.DetailMean,
				mal.DetailMediaType,
				mal.DetailGenres,
				mal.DetailNumEpisodes,
				mal.DetailAverageEpisodeDuration,
				mal.DetailAlternativeTitles,
			}),
		},
	}

	if id := q.Get("id"); id != "" {
		if idInt, err := strconv.Atoi(id); err != nil {
			return "", errors.New("invalid_query")
		} else {
			query.Id = idInt
		}
	} else {
		return "", errors.New("invalid_query")
	}

	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetDetails(query)

	if err != nil {
		return "", errors.New("client_error")
	}

	if j, err := res.JSON(); err != nil {
		return "", errors.New("decode_error")
	} else {
		return j, nil
	}

}
