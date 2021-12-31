package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"

	"github.com/ivansuteja96/go-opencv"
	"github.com/ivansuteja96/go-opencv/entity"
)

func main() {

	fmt.Println(entity.IsCudaAvailable())
	openCVMethod := opencv.New(
		opencv.Config{
			ProcessingUnit: entity.OPENCV_CPU,
		},
	)
	fmt.Println(openCVMethod.NewMat())

	imageURL := "https://www.reviewpro.com/wp-content/uploads/2019/06/Google-logo.jpg"
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		log.Fatalln(err, "ImageURL :", imageURL)
	}
	// For intermittent EOF error
	req.Close = true

	client := &http.Client{}
	responseImage, err := client.Do(req)
	if err != nil {
		log.Fatalln(err, "ImageURL :", imageURL)
	}

	var imageBuf bytes.Buffer

	io.Copy(&imageBuf, responseImage.Body)
	responseImage.Body.Close()

	// to check if is that true image
	img, _, err := image.Decode(bytes.NewReader(imageBuf.Bytes()))
	if err != nil {
		log.Fatalln(err, "ImageURL :", imageURL)
	}

	fmt.Println(openCVMethod.ImageToMatRGB(img))
}
