package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ID                 = 0
	COUNT_PICTURES     = 6
	MIN_COUNT_PICTURES = 0
	COUNT_COMMENTS     = 6
	MIN_COUNT_COMMENTS = 0
	COUNT_AVATARS      = 6
	MIN_AVATAR         = 1
	MAX_AVATAR         = 6
	MIN_COUNT_LIKES    = 0
	MAX_COUNT_LIKES    = 10
	URL_DATA           = "http://localhost:3010/data"
)

var (
	PATH_NAMES        = filepath.Join("DATA", "names.txt")
	PATH_COMMENTS     = filepath.Join("DATA", "comments.txt")
	PATH_DESCRIPTIONS = filepath.Join("DATA", "descriptions.txt")

	NAMES, _        = getContent(PATH_NAMES)
	COMMENTS, _     = getContent(PATH_COMMENTS)
	DESCRIPTIONS, _ = getContent(PATH_DESCRIPTIONS)

	generateCommentId          = createIdGenerator()
	generatePictureId          = createIdGenerator()
	data              pictures = getPictures(COUNT_PICTURES, newPicture)
)

type (
	id      uint8
	comment struct {
		ID      id     `json:"id"`
		Avatar  string `json:"avatar"`
		Message string `json:"message"`
		Name    string `json:"name"`
	}

	picture struct {
		ID          id         `json:"id"`
		URL         string     `json:"url"`
		Likes       uint16     `json:"likes"`
		Comments    []*comment `json:"comments"`
		Description string     `json:"description"`
	}

	pictures []*picture

	IdGenerator func() id
)

func newComment() *comment {
	var id = generateCommentId()
	return &comment{
		ID:      id,
		Avatar:  fmt.Sprintf("img/avatar-%v.svg", getRandomPositiveInteger(MIN_AVATAR, MAX_AVATAR)),
		Message: getRandom(COMMENTS),
		Name:    getRandom(NAMES),
	}
}

func newPicture() *picture {
	var id = generatePictureId()
	return &picture{
		ID:          id,
		URL:         fmt.Sprintf("photos/%v.jpg", id),
		Likes:       getRandomPositiveInteger(MIN_COUNT_LIKES, MAX_COUNT_LIKES),
		Comments:    getComments(COUNT_COMMENTS, newComment),
		Description: getRandom(DESCRIPTIONS),
	}
}

func createIdGenerator() IdGenerator {
	var lastGeneratedId id
	return func() id {
		lastGeneratedId++
		return lastGeneratedId
	}
}

func getRandomPositiveInteger(min, max int) uint16 {
	return uint16(rand.Intn(max-min) + min)
}

func getRandom(arr []string) string {
	return arr[getRandomPositiveInteger(0, len(arr)-1)]
}

func getComments(leng int, createInstance func() *comment) []*comment {
	result := make([]*comment, 0)
	for i := 0; i < leng; i++ {
		result = append(result, createInstance())
	}
	return result
}

func getPictures(leng int, createInstance func() *picture) []*picture {
	result := make([]*picture, 0)
	for i := 0; i < leng; i++ {
		result = append(result, createInstance())
	}
	return result
}

func arrayFrom(leng int, createInstance func() interface{}) []interface{} {
	result := make([]interface{}, 0)
	for i := 0; i < leng; i++ {
		result = append(result, createInstance())
	}
	return result
}

func loadData(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 10)
	fmt.Println(w, r.MultipartForm)
}

func getContent(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&data)
}

func main() {
	server := http.Server{
		Addr: "localhost:3010",
	}

	http.HandleFunc("/processForm", processForm)
	http.HandleFunc("/data", getData)
	server.ListenAndServe()
}
