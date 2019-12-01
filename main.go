package main

import (
	"fmt"
	"io/ioutil"
	NormalCard "lovelive-hd-ur/CardHandlers"
	"lovelive-hd-ur/CardResponse"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
)

var cardJobs = make(chan struct{}, 2)

func root(w http.ResponseWriter, r *http.Request) {
	lsOut, _ := exec.Command("ls").Output()

	fmt.Fprintln(w, string(lsOut))
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
	resp, _ := http.Get(parsed.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cardResponse, _ := CardResponse.UnmarshalCardResponse(body)

	card := NormalCard.New()

	cardUrl, _ := url.Parse(*cardResponse.Results[0].CleanUrIdolized)
	card.CardUrl = *cardUrl
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
	http.HandleFunc("/", root)
	http.HandleFunc("/ggg", ggg)
	http.HandleFunc("/cards", cards)

	http.ListenAndServe(":5005", nil)
}
