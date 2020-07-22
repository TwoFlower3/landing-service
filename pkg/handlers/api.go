package handlers

import (
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/twoflower3/interview-service/pkg/datastore"
	model "github.com/twoflower3/interview-service/pkg/models/rest"
	msg "github.com/twoflower3/interview-service/pkg/msgbuilder"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

const reg = `.+\.pdf`

// Send ...
// @Summary отправка резюме
// @Description отправка резюме кандидата на почту
// @Accept  json
// @Param body body model.Resume true "resume"
// @Success 200 {string} string ""
// @Failure 400 {object} string "error"
// @Failure 500 {object} string "error"
// @Router /send [post]
func Send(c *gin.Context) {

	attachment := msg.Attachment{}

	forms, err := c.MultipartForm()
	if err != nil {
		log.Errorf("get multipart error: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resume := getResume(forms)

	if resume.File.Content != nil {
		check, err := regexp.MatchString(reg, resume.File.Filename)
		if err != nil {
			log.Errorf("validate error: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "#0"})
			return
		} else if !check {
			log.Error("file not valid")
			c.JSON(http.StatusBadRequest, gin.H{"error": "file not valid"})
			return
		}

		file, err := resume.File.Content.Open()
		if err != nil {
			log.Error("file not open")
			c.JSON(http.StatusBadRequest, gin.H{"error": "#4"})
			return
		}
		defer file.Close()

		attachment = msg.NewAttachment(resume.File.Filename, file)
	}

	message, err := msg.GetMessage(
		resume.Name,
		resume.Email,
		resume.Number,
		resume.Project,
		attachment,
	)
	if err != nil {
		log.Errorf("get message error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "#2"})
		return
	}

	if err := datastore.ClientSMTP.SendMessage(message); err != nil {
		log.Errorf("send message error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "#3"})
	}

	c.Status(http.StatusOK)
}

func getResume(forms *multipart.Form) *model.Resume {
	var resume model.Resume

	for key, value := range forms.Value {
		switch key {
		case "name":
			resume.Name = value[0]
		case "number":
			resume.Number = value[0]
		case "email":
			resume.Email = value[0]
		case "project":
			resume.Project = value[0]
		}
	}

	for key, file := range forms.File {
		resume.File.Filename = key
		resume.File.Content = file[0]
	}

	log.Debugf("resume: %+v", resume)

	return &resume
}
