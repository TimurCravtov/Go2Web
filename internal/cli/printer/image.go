package printer

import (
    "image"
    "bytes"
    "go2web/internal/request"
    "go2web/internal/cli/printer/utils"
)

func ImagePrinter(urlPath string, response *request.HttpResponse) (string, error) {
    imageBytes := response.Body

    config, _, err := image.DecodeConfig(bytes.NewReader(imageBytes))
    if err != nil {
        return "", err
    }

    width, height := calculateBestDimensions(config.Width, config.Height)

    asciiArt, err := utils.ImageToAscii(imageBytes, width, height)
    if err != nil {
        return "", err
    }
    return asciiArt, nil
}

func calculateBestDimensions(originalWidth, originalHeight int) (int, int) {
    maxWidth := 60
    
    ratio := float64(originalWidth) / float64(originalHeight)
        
    newWidth := maxWidth
    newHeight := int(float64(newWidth) / (ratio * 2.0))

    if newHeight > 40 {
        newHeight = 40
        newWidth = int(float64(newHeight) * ratio * 2.0)
    }

    return newWidth, newHeight
}