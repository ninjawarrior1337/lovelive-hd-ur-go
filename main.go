package main

import (
	"fmt"
	"io/ioutil"
	"lovelive-hd-ur/CardResponse"
	NormalCard "lovelive-hd-ur/cardhandlers"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cardJobs = make(chan struct{}, 2)

func root(ctx *gin.Context) {
	lsOut, _ := exec.Command("ls").Output()

	fmt.Fprint(ctx.Writer, string(lsOut))
}

func ggg(w http.ResponseWriter, r *http.Request) {
	info, _ := ioutil.ReadFile("maruexcite.png")

	w.Write(info)
}

func cards(w http.ResponseWriter, r *http.Request) {
	select {
	case cardJobs <- struct{}{}:
		{
			defer func() { <-cardJobs }()
			fmt.Println("processing card")
		}
	default:
		{
			w.WriteHeader(500)
			w.Write([]byte("Too many requests right now, try again later"))
			fmt.Println("skipped request")
			return
		}
	}
	parsed, _ := url.Parse("https://schoolido.lu/api/cards/")
	q := parsed.Query()
	q.Add("ids", r.URL.Query().Get("ids"))
	parsed.RawQuery = q.Encode()
	fmt.Println(parsed)
	resp, err := http.Get(parsed.String())
	if err != nil {

	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := CardResponse.UnmarshalCardResponse(body)

	card := NormalCard.New()

	cardURL, _ := url.Parse(*cardResponse.Results[0].CleanUrIdolized)
	card.CardUrl = *cardURL
	card.FileBaseName = strconv.FormatInt(*cardResponse.Results[0].ID, 10) + ".png"

	if err := card.ProcessCard(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	file, _ := ioutil.ReadFile(card.CardOutFile)

	w.WriteHeader(200)
	w.Write(file)
	return
}

func main() {
	router := gin.Default()
	router.GET("/", root)

	router.Run("0.0.0.0:5005")
	// http.HandleFunc("/", root)
	// http.HandleFunc("/ggg", ggg)
	// http.HandleFunc("/cards", cards)

	// http.ListenAndServe(":5005", nil)
}
