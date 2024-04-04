package cmd

import (
	"fmt"
	"net/http"
	"sync"
)

func tense(url string, n int64) error {
	var wg sync.WaitGroup
	wg.Add(int(n))

	for i := int64(0); i < n; i++ {
		go func() {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()
			// Process the response as needed
			// ...

			// Print the response
			//fmt.Println("Response:", resp)
			if resp.StatusCode != 200 {
				fmt.Println("server crached")
				fmt.Println("Error:", resp.Status)
				fmt.Println("status code:", resp.StatusCode)
			}

			
			}()
		}
		
		wg.Wait()
		
	fmt.Println("server okay!")
	return nil
}