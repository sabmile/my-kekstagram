package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
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

var NAMES = []string{
	"Elisa Rush",
	"Lulu Estes",
	"Brogan Atkins",
	"Alessandro Alvarez",
	"Jerome Briggs",
	"Maliha Petersen",
	"Cordelia Barrett",
	"Lexie Castillo",
	"Harri Stevens",
	"Veronica Ross",
}

var COMMENTS = []string{
	"When you take a photo, it would be good to remove your finger from the frame. In the end, it's just unprofessional. One thing is unclear: how so?!",
	"There is simply no framing. The filter was selected unsuccessfully: I would use a series set to 80%",
	"I lost my family, children and cat because of this photo. They said they didn't share my love for art and went to a neighbor. Everything is fine!",
	"I'm stuck on this photo and I can't tear myself away. I don't know what to do at all. Is this a composition?! What is this composition?!",
	"I slipped on a banana peel and dropped the camera on the cat and I got a better picture.",
	"The faces of the people in the photo are distorted, as if they are being beaten. How could you catch such an unfortunate moment?! The horizon is littered.",
	"My grandmother accidentally sneezed with a camera in her hands and she got a better photo. Shob I lived like this!",
	"I can't imagine how you can photograph the sea and sunset better. This is just the apogee. After that, we can burn all the cameras, because the peak has been reached anyway. The focus is blurred. Or is it just someone splattered the lens?",
}

var DESCRIPTIONS = []string{
	"If you clearly formulate a wish for the universe, then everything will definitely come true. Believe in yourself. The main thing is to want and dream.",
	"Appreciate every moment. Appreciate those who are close to you and drive away all doubts. Don't offend everyone with words",
	"How cool is the food here",
	"Hanging out with friends at the sea",
}

type comment struct {
	ID      uint8  `json:"id"`
	Avatar  string `json:"avatar"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

var generateCommentId = createIdGenerator()

func newComment() *comment {
	var id = generateCommentId()
	return &comment{
		ID:      id,
		Avatar:  fmt.Sprintf("img/avatar-%v.svg", getRandomPositiveInteger(MIN_AVATAR, MAX_AVATAR)),
		Message: getRandom(COMMENTS),
		Name:    getRandom(NAMES),
	}
}

type picture struct {
	ID          uint8      `json:"id"`
	URL         string     `json:"url"`
	Likes       uint16     `json:"likes"`
	Comments    []*comment `json:"comments"`
	Description string     `json:"description"`
}

var generatePictureId = createIdGenerator()

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

type pictures []*picture

var data pictures = getPictures(COUNT_PICTURES, newPicture)

func createIdGenerator() func() uint8 {
	var lastGeneratedId uint8
	return func() uint8 {
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
