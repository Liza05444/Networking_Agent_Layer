package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"agent/models"
	"agent/segment"
	"agent/minio"
	"agent/sender"
)

func writeJSON(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(models.Response{Status: code, Message: msg})
}

//	@Summary	Загрузка запрошенного с транспортного уровня изображения из MinIO, деление на сегменты и вызов метода для их отправки на транспортный уровень
//	@Description	Метод Process принимает запрос с транспортного уровня, загружает изображение из MinIO, делит его на сегменты и для каждого сегмента вызывает отправку на транспортный уровень.
//	@Tags		agent
//	@Accept		json
//	@Produce	json
//	@Param		body	body		models.Request	true	"Название документа и номер страницы"	
//	@Success	200	{object}	models.Response
//	@Failure	400	{object}	models.Response
//	@Failure	405	{object}	models.Response
//	@Failure	500	{object}	models.Response
//	@Router		/process [post]
func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	var req models.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, "Ошибка парсинга JSON: "+err.Error())
		return
	}

	objectName := fmt.Sprintf("%s/%d.png", req.DocumentID, req.PageID)

	imageData, err := minio.DownloadImage(objectName)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Ошибка загрузки изображения: "+err.Error())
		return
	}

	segments := segment.Split(imageData)
	totalSegments := len(segments)

	log.Printf("Документ %s, страница %d: %d байт, %d сегментов",
		req.DocumentID, req.PageID, len(imageData), totalSegments)

	for i, segmentData := range segments {
		seg := models.Segment{
			DocumentID:    req.DocumentID,
			PageID:        req.PageID,
			SegmentID:     i + 1,
			TotalSegments: totalSegments,
			Payload:       base64.StdEncoding.EncodeToString(segmentData),
		}

		if err := sender.Send(seg); err != nil {
			log.Printf("Ошибка отправки сегмента %d: %v", i+1, err)
			writeJSON(w, http.StatusInternalServerError, "Ошибка отправки сегмента")
			return
		}

		log.Printf("Сегмент %d/%d отправлен", i+1, totalSegments)
	}

	writeJSON(w, http.StatusOK, fmt.Sprintf("Отправлено %d сегментов", totalSegments))
}
