package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type UpdatePostpayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func updatePost(postID int, p UpdatePostpayload, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("http://localhost:5000/v1/posts/%d", postID)

	b, _ := json.Marshal(p)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request: ", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Update response status:", resp.Status)
}

func main() {
	var wg sync.WaitGroup

	postID := 13

	// simulate user A and B updating the same post concurrently
	wg.Add(2)
	content := "NEW CONTENT FROM USER B"
	title := "NEW TITLE FROM USER A"

	go updatePost(postID, UpdatePostpayload{Title: &title}, &wg)
	go updatePost(postID, UpdatePostpayload{Content: &content}, &wg)
	wg.Wait()
}
