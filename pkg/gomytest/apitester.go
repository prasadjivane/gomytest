// pkg/gomytest/gomytest.go

package gomytest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TestGET performs a GET request to the specified URL and prints response data and execution time
func TestGET(url string) error {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("GET Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
	fmt.Printf("Execution Time: %s\n", time.Since(start))
	return nil
}

// TestPOST performs a POST request to the specified URL with the given data and prints response data and execution time
func TestPOST(url string, data []byte) error {
	start := time.Now()
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("POST Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
	fmt.Printf("Execution Time: %s\n", time.Since(start))
	return nil
}
func TestPUT(url string, data []byte) error {
	start := time.Now()
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("PUT Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
	fmt.Printf("Execution Time: %s\n", time.Since(start))
	return nil
}

// TestDELETE performs a DELETE request to the specified URL and prints response data and execution time
func TestDELETE(url string) error {
	start := time.Now()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("DELETE Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
	fmt.Printf("Execution Time: %s\n", time.Since(start))
	return nil
}
