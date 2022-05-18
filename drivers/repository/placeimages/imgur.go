package placeimages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (Repo *placeimageRepo) Upload(image string) (string, error) {
	url := "https://api.imgur.com/3/image"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", image)
	// _ = writer.WriteField("type", "base64")

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", "Client-ID f4a9a61acd375d4")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	resJSON := map[string]interface{}{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(resJSON)
	fmt.Println(string(body))
	return resJSON["data"].(map[string]interface{})["link"].(string), nil
}


