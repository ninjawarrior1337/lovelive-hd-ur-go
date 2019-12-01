package NormalCard

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
)

type NormalCard struct {
	CardUrl      url.URL
	FileBaseName string
	CardInFile   string
	CardOutFile  string
}

func New() *NormalCard {
	return &NormalCard{}
}

func (card *NormalCard) WriteBaseCard() error {
	cardData, err := http.Get("https:" + card.CardUrl.String())
	if err != nil {
		return err
	}
	defer cardData.Body.Close()

	cardDataReader, err := ioutil.ReadAll(cardData.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join("cardsIn", card.FileBaseName), cardDataReader, os.ModePerm); err != nil {
		return err
	}

	card.CardInFile = path.Join("cardsIn", card.FileBaseName)
	return nil
}

func (card *NormalCard) DoWaifu2x() error {
	file := path.Join("cardsOut", card.FileBaseName)
	card.CardOutFile = file

	waifu2xCmd := exec.Command("/usr/bin/waifu2x-converter-cpp", "-i", card.CardInFile, "-o", file, "--noise-level 3", "--scale-ratio 2")
	waifu2xOut, _ := waifu2xCmd.Output()

	fmt.Println(string(waifu2xOut))

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func (card *NormalCard) ProcessCard() error {
	if err := card.WriteBaseCard(); err != nil {
		return err
	}
	if err := card.DoWaifu2x(); err != nil {
		return err
	}
	return nil
}
