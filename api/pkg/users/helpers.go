package users

import (
	"db/pkg/utils"
	"errors"
	"io"
	"net/http"
)

func DBApiWrapper(endpoint string) ([]byte, int, error) {
	// Call the database API
	req, err := http.NewRequest("GET", utils.AppConfig.DBApiURL+endpoint, nil)
	if err != nil {
		return nil, 501, errors.New("error creating request")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 501, errors.New("error sending request")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 501, errors.New("error reading response")
	}
	return body, res.StatusCode, nil

}
