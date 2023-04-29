package graph

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ArxivInsanity/backend-service/src/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Params struct {
	Authors     string `form:"authors"`
	MinYear     string `form:"minYear"`
	MaxYear     string `form:"maxYear"`
	MinCitation string `form:"minCitation"`
}

// Graph godoc
// @Summary Retrieve graph for a paper
// @Schemes
// @Description Endpoint to get the generated graph for a paper
// @Tags Graph
// @Accept json
// @Produce json
// @Param id path string true "The paper id"
// @Param authors query []string false "The authors"
// @Param minYear query string false "The min year"
// @Param maxYear query string false "The max year"
// @Param minCitation query string false "The minimum number of citations"
// @Success 200 {object} models.Response
// @Security Bearer
// @Router /api/graph/{id} [get]
func GetGraph(c *gin.Context) {
	client := &http.Client{}
	var params Params
	if err := c.BindQuery(&params); err != nil {
		log.Error().Msg("Something went wrong parsing query params " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Something went wrong parsing query params " + err.Error(), nil})
		return
	}

	var req *http.Request
	if params.Authors != "" || params.MinYear != "" || params.MaxYear != "" || params.MinCitation != "" {
		req, _ = http.NewRequest("GET", BASE_URL+GET_GRAPH_WITH_FILTER+"/"+c.Param("id"), nil)
		q := req.URL.Query()
		q.Add("authors", params.Authors)
		q.Add("minYear", params.MinYear)
		q.Add("maxYear", params.MaxYear)
		q.Add("minCitation", params.MinCitation)
		req.URL.RawQuery = q.Encode()
	} else {
		req, _ = http.NewRequest("GET", BASE_URL+GET_GRAPH+"/"+c.Param("id"), nil)
	}
	log.Info().Msg("Calling graph service endpoint " + req.URL.RequestURI())
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Something went wrong when calling graph service: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Something went wrong when calling graph service " + err.Error(), nil})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("Error reading response from graph service: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error reading response from graph service:  " + err.Error(), nil})
		return
	}
	graph := map[string]any{}
	err = json.Unmarshal(body, &graph)
	if err != nil {
		log.Error().Msg("Error unmarshaling response from graph service: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error unmarshaling response from graph service: " + err.Error(), nil})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Response{http.StatusOK, "Retrieved graph", graph})
}
