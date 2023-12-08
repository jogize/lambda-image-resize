package resizer

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/umtkas/image-resizer-lambda/configs"
)

// GetResizedImages returns the path of resized images
func GetResizedImages(configuration configs.Configuration, sourceImagePath string) []string {
	var resizedImagePaths []string

	for _, imageSize := range configuration.ImageSizes {
		fmt.Println("in GetResizedImages before resize")
		resizedImagePath := resize(configuration, imageSize, sourceImagePath)
		fmt.Println("in GetResizedImages before append")
		resizedImagePaths = append(resizedImagePaths, resizedImagePath)
	}

	return resizedImagePaths
}
// add for triger
// resizeImage resizes the image and returns path
func resize(configuration configs.Configuration, imageSize configs.ImageSize, sourceImagePath string) string {
	fmt.Println("in resize on the start")
	sourceImage, err := imaging.Open(sourceImagePath)

	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	var resizedImage *image.NRGBA
	fmt.Println("in resizedImage on the start")
	if configuration.IsSaveWithAspectRatio {
		resizedImage = imaging.Resize(sourceImage, imageSize.Width, 0, imaging.Lanczos)
	} else {
		resizedImage = imaging.Resize(sourceImage, imageSize.Width, imageSize.Height, imaging.Lanczos)
	}
	fmt.Println("in resizedImage on the after")
	resizedImageName := getResizedImageName(configuration, imageSize, sourceImagePath)
	imagePath := configuration.LocalImageDirectory + "/" + resizedImageName
	err = imaging.Save(resizedImage, imagePath)
	if err != nil {
		log.Fatalf("resized image is not saved, %v", err)
	}

	return resizedImageName
}

func getResizedImageName(configuration configs.Configuration, imageSize configs.ImageSize, sourceImagePath string) string {
	sourceImageNameWithExtension := filepath.Base(sourceImagePath)
	sourceImageNameWithOutExtension := strings.TrimSuffix(sourceImageNameWithExtension, filepath.Ext(sourceImageNameWithExtension))
	return sourceImageNameWithOutExtension + "_" + imageSize.WidthHeight + "." + configuration.ImageExtension
}

// RemoveImages removes all images
func RemoveImages(configuration configs.Configuration) {
	files, err := filepath.Glob(filepath.Join(configuration.LocalImageDirectory, "*"))

	if err != nil {
		log.Fatal("cannot remove, directory error")
		return
	}

	for _, file := range files {
		fmt.Println(file)
		os.RemoveAll(file)
	}
}
