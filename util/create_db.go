package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func postAnswers() {
	values := map[string]string{"question_text": "John Doe", "answer_id": "gardener"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/answers", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])

}
