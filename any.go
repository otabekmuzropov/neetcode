package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const gotenbergURL = "http://localhost:3000"

// ConvertHTMLToPDF converts HTML content to PDF using Gotenberg
func ConvertHTMLToPDF(htmlContent string, outputPath string) error {
	// Create a buffer to write multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the HTML file
	part, err := writer.CreateFormFile("files", "index.html")
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = part.Write([]byte(htmlContent))
	if err != nil {
		return fmt.Errorf("failed to write HTML content: %w", err)
	}

	// Optional: Add paper size (default is A4)
	// writer.WriteField("paperWidth", "8.5")
	// writer.WriteField("paperHeight", "11")

	// Optional: Add margins (in inches)
	// writer.WriteField("marginTop", "0.5")
	// writer.WriteField("marginBottom", "0.5")
	// writer.WriteField("marginLeft", "0.5")
	// writer.WriteField("marginRight", "0.5")

	// Optional: Set landscape orientation
	// writer.WriteField("landscape", "false")

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", gotenbergURL+"/forms/chromium/convert/html", body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("gotenberg returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Create output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	// Write the PDF to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write PDF: %w", err)
	}

	return nil
}

func main() {
	// Example 1: Convert HTML string to PDF

	// fileByte, err := os.ReadFile("ok.html")
	// if err != nil {
	// 	log.Fatal("error read file", err)
	// }
	// htmlContent := string(fileByte)

	// err = ConvertHTMLToPDF(htmlContent, "output1.pdf")
	// if err != nil {
	// 	fmt.Printf("Error converting HTML string: %v\n", err)
	// } else {
	// 	fmt.Println("Successfully created output1.pdf")
	// }

	t := time.Now().UnixMilli()
	fmt.Println(t)

	// Example 2: Convert HTML file to PDF
	// err = ConvertHTMLFileToPDF("index.html", "output2.pdf", "style.css", "logo.png")
	// if err != nil {
	// 	fmt.Printf("Error converting HTML file: %v\n", err)
	// } else {
	// 	fmt.Println("Successfully created output2.pdf")
	// }

	// Example 3: Convert URL to PDF
	// err = ConvertURLToPDF("https://example.com", "output3.pdf")
	// if err != nil {
	// 	fmt.Printf("Error converting URL: %v\n", err)
	// } else {
	// 	fmt.Println("Successfully created output3.pdf")
	// }
}

func Divisors(n int) []int {
	if n <= 0 {
		return []int{}
	}

	divs := []int{}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	return divs
}
