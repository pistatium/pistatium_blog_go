package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pistatium/pistatium_blog_go/repos"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const ImageMaxSize = 680

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const AdminEntriesPerPage = 30

func (s *Server) GetAdminEntries(gc *gin.Context) {
	ctx := gc.Request.Context()

	page, _ := strconv.Atoi(gc.DefaultQuery("page", "0"))
	offset := 0
	if page != 0 {
		offset = page * AdminEntriesPerPage
	}

	entries, err := s.Entries.GetEntries(ctx, offset, AdminEntriesPerPage, false)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, &repos.Entries{Entries: entries})
}

func (s *Server) PostEntry(gc *gin.Context) {

	ctx := gc.Request.Context()
	var entry repos.Entry
	if err := gc.ShouldBindJSON(&entry); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.Entries.UpdateEntry(ctx, entry.Id, entry)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gc.JSON(http.StatusOK, entry)
}

func (s *Server) UploadPhoto(gc *gin.Context) {
	ctx := gc.Request.Context()

	file, header, err := gc.Request.FormFile("file")
	filename := header.Filename
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	img, err := generatePhoto(bs, filename)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = s.Photos.PutPhoto(ctx, img)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, &map[string]string{"status": "uploaded", "path": fmt.Sprintf("/photo/show/%d.jpg", img.Id)})
}

func (s *Server) IsLogin(gc *gin.Context) {
	gc.JSON(http.StatusOK, &map[string]string{"status": "ok",})
}

func (s *Server) AdminLogin(gc *gin.Context) {
	ctx := gc.Request.Context()
	var user LoginForm
	if err := gc.ShouldBindJSON(&user); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := s.Admin.GetValidUser(ctx, user.Username, user.Password)
	if err != nil {
		if err == err.(*repos.LoginError) {
			gc.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		} else {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	SetLogin(gc, u.Username)
	gc.JSON(http.StatusOK, &map[string]string{"status": "logged in",})
}

func generatePhoto(bs []byte, filename string) (*repos.Photo, error) {
	now := time.Now()
	format, err := guessImageFormat(bytes.NewBuffer(bs))
	var contentType string
	if err != nil {
		contentType = ""
	} else {
		contentType = "image/" + format
	}

	// リサイズ
	if format == "jpeg" || format == "png" {
		img, _, err := image.Decode(bytes.NewBuffer(bs))
		if err != nil {
			return nil, err
		}

		w := float32(img.Bounds().Dx())
		h := float32(img.Bounds().Dy())

		if w > ImageMaxSize {
			h = float32(ImageMaxSize) / w * h
			w = float32(ImageMaxSize)
		}
		dst := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
		draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

		var buf bytes.Buffer
		switch format {
		case "png":
			err = png.Encode(&buf, dst)
			if err != nil {
				return nil, err
			}
		case "jpeg":
			err = jpeg.Encode(&buf, dst, nil)
			if err != nil {
				return nil, err
			}
		}
		bs = buf.Bytes()
	}

	p := repos.Photo{
		Datetime:    &now,
		Image:       bs,
		Title:       filename,
		ContentType: contentType,
	}
	return &p, nil
}

func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	return
}
