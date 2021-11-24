package main

import (
	"log"
	"math"
	"net/http"
)

func entropy(value []byte) float64 {
	frq := make(map[byte]float64)

	//get frequency of characters
	for _, i := range value {
		frq[i]++
	}

	var sum float64

	for _, v := range frq {
		f := v / float64(len(value))
		sum += f * math.Log2(f)
	}

	bits := math.Ceil(sum*-1) * float64(len(value))
	return bits
}

func checkFileType(data []byte) bool {
	contentType := http.DetectContentType(data)
	log.Println(contentType)
	return contentType == "application/octet-stream"
}
