package paper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ArxivInsanity/backend-service/src/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AutoCompleteDetails struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	AuthorsYear string `json:"authorsYear"`
}
type AutoCompleteResult struct {
	Matches []AutoCompleteDetails `json:"matches"`
}

// Papers godoc
// @Summary Autocomplete suggestions for paper search
// @Schemes
// @Description Endpoint to retrieve autocomplete suggestions for papers that the user types in the search
// @Tags Papers
// @Accept json
// @Produce json
// @Param query query string true "The query to search for papers"
// @Security Bearer
// @Success 200 {object} models.Response
// @Router /api/papers/autocomplete [get]
func GetPaperAutocompleteSuggestions(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", BASE_URL+AUTOCOMPLETE_ENDPOINT, nil)
	req.Header.Add("x-api-key", os.Getenv(SS_API_KEY))
	q := req.URL.Query()
	q.Add("query", c.Query("query"))
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Something went wrong when calling autocomplete to semantic scholar api: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Something went wrong when calling autocomplete to semantic scholar api: " + err.Error(), nil})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("Error reading response from semantic scholar for autocomplete endpoint: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error reading response from semantic scholar for autocomplete endpoint: " + err.Error(), nil})
		return
	}
	var respContent AutoCompleteResult
	err = json.Unmarshal(body, &respContent)
	if err != nil {
		log.Error().Msg("Error unmarshaling response from semantic scholar for autocomplete endpoint: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error unmarshaling response from semantic scholar for autocomplete endpoint: " + err.Error(), nil})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Response{http.StatusOK, "Queried successfully", respContent.Matches})
}

type Author struct {
	AuthorId string `json:"authorId"`
	Name     string `json:"name"`
}
type Reference struct {
	PaperId string `json:"paperId"`
	Title   string `json:"title"`
}
type PaperDetails struct {
	Url        string      `json:"url"`
	Year       int         `json:"year"`
	Authors    []Author    `json:"authors"`
	Abstract   string      `json:"abstract"`
	References []Reference `json:"references"`
}

// Papers godoc
// @Summary Retrieves paper details
// @Schemes
// @Description Endpoint to retrieve paper details given a paper id
// @Tags Papers
// @Accept json
// @Produce json
// @Param id path string true "The paper id"
// @Success 200 {object} models.Response
// @Security Bearer
// @Router /api/papers/{id} [get]
func GetPaperDetails(c *gin.Context) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", BASE_URL+PAPER_DETAILS_ENDPOINT+"/"+c.Param("id"), nil)
	req.Header.Add("x-api-key", os.Getenv(SS_API_KEY))
	q := req.URL.Query()
	q.Add("fields", "url,year,authors,abstract,references")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Something went wrong when calling paper details to semantic scholar api: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Something went wrong when calling paper details to semantic scholar api: " + err.Error(), nil})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("Error reading response from semantic scholar for paper details endpoint: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error reading response from semantic scholar for paper details endpoint: " + err.Error(), nil})
		return
	}
	var paperDetails PaperDetails
	err = json.Unmarshal(body, &paperDetails)
	if err != nil {
		log.Error().Msg("Error unmarshaling response from semantic scholar for paper details endpoint: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, models.Response{http.StatusInternalServerError, "Error unmarshaling response from semantic scholar for paper details endpoint: " + err.Error(), nil})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Response{http.StatusOK, "Queried successfully", paperDetails})
}
