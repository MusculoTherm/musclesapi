package api

import (
	"fmt"
	"github.com/MusculoTherm/musclesapi/models"
	"github.com/MusculoTherm/musclesapi/utils"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
)

func V0_API_Upload_Artwork(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		failFormat := models.Response{Success: false, Debug: "An Error Occured Decoding the Request Body", Message: "Failed Uploading the Image"}
		JSONResponse(w, failFormat, 400)
		return
	}

	contentType := handler.Header.Get("Content-Type")
	if contentType == "" {
		failFormat := models.Response{Success: false, Debug: "The `Content-Type` header is missing on the file", Message: "Failed Uploading the Image"}
		JSONResponse(w, failFormat, 400)
		return
	}

	extension, err := utils.CheckImageContentType(contentType)
	if err != nil {
		failFormat := models.Response{Success: false, Debug: err.Error(), Message: "Failed Uploading the Image"}
		JSONResponse(w, failFormat, 400)
		return
	}

	fileId := generateFileId()
	path := fmt.Sprintf("./uploads/%s%s", fileId, extension)
	defer file.Close()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("An error occured saving an uploaded file:", err)
		failFormat := models.Response{Success: false, Debug: "Internal Server Error", Message: "Failed Uploading the Image"}
		JSONResponse(w, failFormat, 500)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	resp := models.Artwork{
		Url: fmt.Sprintf("%suploads/%s%s", models.GlobalConfig.FullyQualifiedHost, fileId, extension),
	}
	success := models.Response{Success: true, Data: resp, Message: "Successfully Uploaded the Image"}
	JSONResponse(w, success, 200)
}

func generateFileId() string {
	return uuid.NewV4().String()
}
