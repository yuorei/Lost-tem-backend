package handlers

import (
	"context"
	"log"
	"lost-item/cloud/googlecloud"
	"lost-item/database"
	"lost-item/database/postgresd"
	"lost-item/model"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	imgupload "github.com/olahol/go-imageupload"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db    database.DBConn
	cloud googlecloud.GCloud
}

func (h *Handler) Init() {
	var err error
	if h.db, err = postgresd.NewPostgresd(); err != nil {
		log.Fatalf("Database connection failed")
	}
	ctx := context.Background()
	x, err := googlecloud.NewGoogleCloud(ctx)
	if err != nil {
		log.Fatalf("Cloud initialization failure")
	}
	h.cloud = *x
}

func (h Handler) Search(c *gin.Context) {
	search_query := c.Param("q")
	search_result, err := h.db.SearchItemsFor(search_query)

	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func (h Handler) ItemList(c *gin.Context) {
	search_query := model.AreaSearchQuery{}
	err := c.Bind(&search_query)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	search_result, err := h.db.SearchItemsArea(search_query.Location1, search_query.Location2)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, search_result)
}

func (h Handler) ItemDetail(c *gin.Context) {
	item_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	item_detail, err := h.db.ItemDetail(uint64(item_id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, item_detail)
}

func (h Handler) RegisterItem(c *gin.Context) {

}

func (h Handler) DeleteItem(c *gin.Context) {

}

func (h Handler) Parse(c *gin.Context) {
	img, err := imgupload.Process(c.Request, "file")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	filename := uuid.String()

	if img.ContentType == "image/png" {
		filename += ".png"
	} else if img.ContentType == "image/jpeg" {
		filename += ".jpg"
	} else {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.cloud.UploadImage(img.Data, filename); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	objects, err := h.cloud.ObjectRecognition(filename)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var img_info model.ImageInfo
	img_info.ImageName = filename
	img_info.Tags = objects

	c.JSON(http.StatusOK, img_info)
}
