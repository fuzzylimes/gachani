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

func RankingHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	res, err := GetRankings(query)
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

func GetRankings(q url.Values) (string, error) {
	query := &mal.RankingQuery{
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
		Limit:       50,
		RankingType: mal.RankingByPopularity,
	}

	if rankingType := q.Get("ranking_type"); rankingType != "" && mal.RankTypeQueries.IsValid(rankingType) {
		query.RankingType = mal.RankingType(rankingType)
	}
	if offset := q.Get("offset"); offset != "" {
		if offsetInt, err := strconv.Atoi(offset); err != nil {
			return "", errors.New("invalid_query")
		} else {
			query.Offset = offsetInt
		}
	}

	c := mal.NewClient(os.Getenv("MAL_API_KEY"))
	res, err := c.GetRanking(query)

	if err != nil {
		return "", errors.New("client_error")
	}

	if j, err := res.JSON(); err != nil {
		return "", errors.New("decode_error")
	} else {
		return j, nil
	}

}
