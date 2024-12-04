package utils

import (
    "fmt"
    "image"
    _ "image/gif"  // Import gif package for decoding GIFs
    _ "image/jpeg" // Import jpeg package for decoding JPEGs
    _ "image/png"  // Import png package for decoding PNGs
    "net/http"
    "log"
)

func CalculatePerimeter(img image.Image) int {
    bounds := img.Bounds()
    log.Printf("Image bounds: %v", bounds)  // Log image bounds to verify dimensions
    return 2 * (bounds.Dx() + bounds.Dy())
}

func DownloadImage(url string) (image.Image, string, error) {
    // Send HTTP GET request to fetch the image
    resp, err := http.Get(url)
    if err != nil {
        return nil, "", fmt.Errorf("failed to download image: %w", err)
    }
    defer resp.Body.Close()

    // Check if the request was successful (HTTP status code 200)
    if resp.StatusCode != http.StatusOK {
        return nil, "", fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
    }

    // Log the HTTP response status
    log.Printf("Image downloaded successfully, status code: %d", resp.StatusCode)

    // Use image.Decode to automatically determine the image format
    img, format, err := image.Decode(resp.Body)
    if err != nil {
        return nil, "", fmt.Errorf("failed to decode image: %w", err)
    }

    // Log the image format
    log.Printf("Decoded image format: %s", format)

    // Return the image and the format (jpeg, png, gif, etc.)
    return img, format, nil
}
