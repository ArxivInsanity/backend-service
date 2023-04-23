package project

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ArxivInsanity/backend-service/src/auth"
	"github.com/ArxivInsanity/backend-service/src/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectDetails struct {
	Name           string `json:"name"`
	LastModifiedAt string `json:"lastModifiedAt"`
}

type ProjectHandler struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

type CreateProjectDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// Project godoc
// @Summary Endpoint for listing all the projects for the user
// @Schemes
// @Description Returns a list of objects that contain details of all the projects created by the user
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} []ProjectDetails
// @Router /api/projects [get]
func (ph *ProjectHandler) GetAllProjects(c *gin.Context) {
	curr, err := ph.Collection.Find(ph.Ctx, bson.D{{auth.USER, c.GetString(auth.USER)}})
	if err != nil {
		log.Error().Msg("Error fetching documents : " + err.Error())
		return
	}
	res := []models.ProjectDoc{}
	for curr.Next(ph.Ctx) {
		var pd models.ProjectDoc
		err := curr.Decode(&pd)
		if err != nil {
			log.Error().Msg("Error decoding the document : " + err.Error())
		}
		log.Info().Msg("Found project: " + fmt.Sprint(pd.Name, pd.LastModifiedAt))
		res = append(res, pd)
	}
	c.IndentedJSON(http.StatusOK, res)
}

// Project godoc
// @Summary Endpoint for creating a new project
// @Schemes
// @Description Returns a status json that describes if the project was created successfully or not
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body CreateProjectDetails true "Project Details"
// @Success 200 {object} models.Response
// @Router /api/projects [post]
func (ph *ProjectHandler) CreateProject(c *gin.Context) {

	var projectDoc models.ProjectDoc
	err := c.ShouldBindJSON(&projectDoc)

	if err != nil {
		log.Error().Msg("Something went wrong binding the post body: " + err.Error())
		return
	}
	projectDoc.User = c.GetString(auth.USER)
	projectDoc.LastModifiedAt = time.Now()
	projectDoc.SeedPapers = make([]models.Paper, 0)
	projectDoc.ReadingList = make([]models.Paper, 0)
	projectDocument := models.ProjectDocument(&projectDoc).GetFullDocument()

	_, err = ph.Collection.InsertOne(ph.Ctx, projectDocument)
	if err != nil {
		log.Error().Msg("Something went wrong when inserting the document")
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := models.Response{http.StatusOK, "Success", projectDocument}
	c.IndentedJSON(http.StatusOK, response)
}

// Project godoc
// @Summary Endpoint for deleting an existing project
// @Schemes
// @Description Returns a status json that describes if the project was deleted successfully.
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param name path string true "Project name"
// @Success 200 {object} models.Response
// @Router /api/projects/{name} [delete]
func (ph *ProjectHandler) DeleteProject(c *gin.Context) {

	name := c.Param("name")

	deleteResult, err := ph.Collection.DeleteOne(ph.Ctx, bson.D{{"name", name}, {auth.USER, c.GetString(auth.USER)}})
	if err != nil {
		log.Error().Msg("Something went wrong when deleting the document")
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := models.Response{http.StatusOK, "Success", fmt.Sprint("Deleted ", deleteResult.DeletedCount, " documents")}
	c.IndentedJSON(http.StatusOK, response)
}

// Project godoc
// @Summary Endpoint for updating an existing project details
// @Schemes
// @Description Returns a status json that describes if the project was updated successfully.
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param name path string true "Existing project name"
// @Param body body CreateProjectDetails true "Project Details"
// @Success 200 {object} models.Response
// @Router /api/projects/{name} [put]
func (ph *ProjectHandler) UpdateProject(c *gin.Context) {

	oldName := c.Param("name")
	var doc models.ProjectDoc
	err := c.ShouldBindJSON(&doc)
	doc.User = c.GetString(auth.USER)
	doc.LastModifiedAt = time.Now()
	projectDocument := models.ProjectDocument(&doc).GetBaseDocument()

	updateResult, err := ph.Collection.UpdateOne(ph.Ctx, bson.D{{"name", oldName}, {auth.USER, c.GetString(auth.USER)}}, bson.D{{"$set", projectDocument}})
	if err != nil {
		log.Error().Msg("Something went wrong when updating the document")
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := models.Response{http.StatusOK, "Success", fmt.Sprint("Updated ", updateResult.ModifiedCount, " documents")}
	c.IndentedJSON(http.StatusOK, response)
}

// Project godoc
// @Summary Endpoint for getting all the seed papers saved for the project
// @Schemes
// @Description Returns a list of seed papers that have been saved for this project
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param name path string true "Project name"
// @Success 200 {object} []string
// @Router /api/projects/{name}/seedPapers [get]
func (ph *ProjectHandler) GetSeedPapers(c *gin.Context) {
	projectName := c.Param("name")
	res := ph.Collection.FindOne(ph.Ctx, bson.D{{"name", projectName}, {auth.USER, c.GetString(auth.USER)}})
	var doc models.ProjectDoc
	res.Decode(&doc)
	c.IndentedJSON(http.StatusOK, doc.SeedPapers)
}

// Project godoc
// @Summary Endpoint to add a seed paper to a project
// @Schemes
// @Description Adding a seed paper to the list of existing seed papers in the project
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param name path string true "Project name"
// @Param body body models.Paper true "Seed paper details"
// @Success 200 {object} models.Response
// @Router /api/projects/{name}/seedPapers [put]
func (ph *ProjectHandler) AddSeedPaper(c *gin.Context) {
	projectName := c.Param("name")
	var paperDetails models.Paper
	err := c.ShouldBindJSON(&paperDetails)
	var doc models.ProjectDoc
	res := ph.Collection.FindOne(ph.Ctx, bson.D{{"name", projectName}, {auth.USER, c.GetString(auth.USER)}})
	res.Decode(&doc)
	doc.SeedPapers = append(doc.SeedPapers, paperDetails)
	seedPaperUpdate := models.ProjectDocument(&doc).GetSeedPaperDocument()
	updateResult, err := ph.Collection.UpdateOne(ph.Ctx, bson.D{{"name", projectName}, {auth.USER, c.GetString(auth.USER)}}, bson.D{{"$set", seedPaperUpdate}})
	if err != nil {
		log.Error().Msg("Something went wrong when updating the document")
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := models.Response{http.StatusOK, "Success", fmt.Sprint("Updated ", updateResult.ModifiedCount, " documents")}
	c.IndentedJSON(http.StatusOK, response)
}

func indexOf(element models.Paper, data []models.Paper) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// Project godoc
// @Summary Endpoint removing a seed paper from project
// @Schemes
// @Description Deleting a seed paper from the list of existing seed papers in the project
// @Tags Project
// @Accept json
// @Produce json
// @Security Bearer
// @Param name path string true "Project name"
// @Param body body models.Paper true "Seed paper details"
// @Success 200 {object} models.Response
// @Router /api/projects/{name}/seedPapers [delete]
func (ph *ProjectHandler) DeleteSeedPaper(c *gin.Context) {
	projectName := c.Param("name")
	var paperDetails models.Paper
	err := c.ShouldBindJSON(&paperDetails)
	var doc models.ProjectDoc
	res := ph.Collection.FindOne(ph.Ctx, bson.D{{"name", projectName}, {auth.USER, c.GetString(auth.USER)}})
	res.Decode(&doc)
	index := indexOf(paperDetails, doc.SeedPapers)
	if index == -1 {
		c.IndentedJSON(http.StatusBadRequest, models.Response{http.StatusBadRequest, "Error", fmt.Sprint("This seed paper doesn't exist - ", paperDetails)})
		return
	}
	doc.SeedPapers = append(doc.SeedPapers[:index], doc.SeedPapers[index+1:]...)
	seedPaperUpdate := models.ProjectDocument(&doc).GetSeedPaperDocument()
	updateResult, err := ph.Collection.UpdateOne(ph.Ctx, bson.D{{"name", projectName}, {auth.USER, c.GetString(auth.USER)}}, bson.D{{"$set", seedPaperUpdate}})
	if err != nil {
		log.Error().Msg("Something went wrong when updating the document")
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	response := models.Response{http.StatusOK, "Success", fmt.Sprint("Updated ", updateResult.ModifiedCount, " documents")}
	c.IndentedJSON(http.StatusOK, response)
}
