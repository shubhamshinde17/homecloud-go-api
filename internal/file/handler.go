package file

import (
	"encoding/json"
	"net/http"

	"github.com/homecloud/go-api/internal/common"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (service *Service) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Method not Allowed",
		))
		return
	}

	// Parse multipart form (max 10 MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid multipart form", http.StatusBadRequest)
		return
	}

	// // Get uploaded file
	// file, header, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, "File is required", http.StatusBadRequest)
	// 	return
	// }
	// defer file.Close()

	// // Get authenticated username from context/middleware
	// username, ok := r.Context().Value("username").(string)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	// // Call service layer
	// uploadLog, err := h.fileService.SaveFile(file, header, username)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// response := ResponseDTO[UploadLog, any]{
	// 	Success: true,
	// 	Message: "File uploaded successfully",
	// 	Data:    &uploadLog,
	// 	Error:   nil,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(response)
}
