package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"agent/config"
	"agent/models"
)

//	@Summary	Отправка сегмента на транспортный уровень
//	@Description	Метод Send отправляет сегмент на транспортный уровень.
//	@Tags		agent
//	@Accept		json
//	@Produce	json
//	@Param		body	body	models.Segment	true	"Сегмент"
//	@Success	200	{object}	models.Response
//	@Failure	500	{object}	models.Response
//	@Router		/send [post]
func Send(segment models.Segment) error {
	jsonData, err := json.Marshal(segment)
	if err != nil {
		return err
	}

	resp, err := http.Post(config.TransportURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Транспортный уровень вернул статус %d", resp.StatusCode)
	}

	return nil
}
