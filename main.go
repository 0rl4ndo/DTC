package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/zenthangplus/goccm"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

var theard = 3

type DTC struct {
	UNLOCKED int
	LOCKED   int
	UNKNOW   int
}

type _43279 struct {
	ID               string        `json:"id"`
	Username         string        `json:"username"`
	GlobalName       interface{}   `json:"global_name"`
	Avatar           string        `json:"avatar"`
	Discriminator    string        `json:"discriminator"`
	PublicFlags      int           `json:"public_flags"`
	Flags            int           `json:"flags"`
	Banner           interface{}   `json:"banner"`
	BannerColor      interface{}   `json:"banner_color"`
	AccentColor      interface{}   `json:"accent_color"`
	Bio              string        `json:"bio"`
	Pronouns         string        `json:"pronouns"`
	Locale           string        `json:"locale"`
	NsfwAllowed      bool          `json:"nsfw_allowed"`
	MfaEnabled       bool          `json:"mfa_enabled"`
	AnalyticsToken   string        `json:"analytics_token"`
	PremiumType      int           `json:"premium_type"`
	LinkedUsers      []interface{} `json:"linked_users"`
	AvatarDecoration interface{}   `json:"avatar_decoration"`
	Email            string        `json:"email"`
	Verified         bool          `json:"verified"`
	Phone            interface{}   `json:"phone"`
}

func PayloadOptions(Token string) cycletls.Options {
	return cycletls.Options{
		Body: "",
		Headers: map[string]string{
			`authority`:          `discord.com`,
			`accept`:             `*/*`,
			`accept-language`:    `en-US,en;q=0.9`,
			`authorization`:      Token,
			`cookie`:             `__dcfduid=15300330074511eeb0dd71e33b679477; __sdcfduid=15300331074511eeb0dd71e33b6794772b40f5558fce88aaa64b711835841d63da5c12d008fddd9e5976e0b1bd5c08e7; locale=en-US; __cfruid=842ff5c84d27d38f3ae4e67c1a00ec547ac39d86-1686838494`,
			`referer`:            `https://discord.com/`,
			`sec-ch-ua`:          `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
			`sec-ch-ua-mobile`:   `?0`,
			`sec-ch-ua-platform`: `"Windows"`,
			`sec-fetch-dest`:     `empty`,
			`sec-fetch-mode`:     `cors`,
			`sec-fetch-site`:     `same-origin`,
			`user-agent`:         `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`,
			`x-track`:            `eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzExNC4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTE0LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjk5OTksImNsaWVudF9ldmVudF9zb3VyY2UiOm51bGx9`,
		},
		Ja3: "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,23-0-65281-45-13-18-5-17513-16-51-10-27-35-43-11-21,29-23-24,0",
	}
}

func HandleError(err error) {
	if err != nil {
		color.Println("[" + Time() + fmt.Sprintf("] [<fg=fc2323>ERROR</>] %s ", err))
	}
}

func Time() string {
	return time.Now().Format("15:04:05")
}

func (D *DTC) Cheak(Token string) (bool, _43279, error) {
	Client := cycletls.Init()
	Response, Err := Client.Do("https://discord.com/api/v9/users/@me?with_analytics_token=true", PayloadOptions(Token), "GET")
	if Err != nil {
		HandleError(Err)
		D.UNKNOW = D.UNKNOW + 1

		return true, _43279{}, nil
	}
	if Response.Status != 200 {
		D.LOCKED = D.LOCKED + 1

		return true, _43279{}, nil
	} else {
		D.UNLOCKED = D.UNLOCKED + 1
		var INFO _43279
		//println(Response.Body)
		Err := json.Unmarshal([]byte(Response.Body), &INFO)
		if Err != nil {
			HandleError(Err)
		}
		return false, INFO, nil
	}
}

func CleanTokens() {
	inputFile, Err := os.Open("./data/Tokens.txt")
	if Err != nil {
		HandleError(Err)
	}
	defer inputFile.Close()
	outputFile, Err := os.Create("./data/TokenCleaned.txt")
	if Err != nil {
		HandleError(Err)
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ":")

		_, Err = fmt.Fprintln(outputFile, tokens[0])
		if Err != nil {
			HandleError(Err)
		}
	}
	if Err := scanner.Err(); Err != nil {
		HandleError(Err)
	}
}
func WriteText(path string, text string) {
	file, Err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if Err != nil {
		HandleError(Err)
	}
	defer file.Close()
	textWriter := bufio.NewWriter(file)
	_, Err = textWriter.WriteString(text + "\n")
	if Err != nil {
		HandleError(Err)
	}
	textWriter.Flush()
}

func (D *DTC) Cheaker(Token string) {
	Cheak, inf, Err := D.Cheak(Token)
	if Err != nil {
		HandleError(Err)
	}
	if Cheak {
		color.Println(fmt.Sprintf("[%s] (<fg=fc2323>LOCKED</>) %s", Time(), Token))
		WriteText("./data/Locked.txt", Token)
	} else {

		color.Println(fmt.Sprintf("[%s] (<fg=0af0ab>UNLOCKED</>) %s --> %s#%s", Time(), Token, inf.Username, inf.Discriminator))
		WriteText("./data/Unlocked.txt", Token)
	}
}

func main() {
	CleanTokens()
	DTCS := DTC{}
	tokens, Err := os.Open("./data/TokenCleaned.txt")
	if Err != nil {
		HandleError(Err)
	}
	defer tokens.Close()
	scanner := bufio.NewScanner(tokens)
	c := goccm.New(theard)
	for scanner.Scan() {
		c.Wait()
		go func() {
			DTCS.Cheaker(scanner.Text())
			c.Done()
		}()

	}
	color.Println(fmt.Sprintf("<fg=0af0ab>UNLOCKED</>: %s | <fg=fc2323>LOCKED</>: %s | <fg=ffef00>UNKNOW</>: %s", strconv.Itoa(DTCS.UNLOCKED), strconv.Itoa(DTCS.LOCKED), strconv.Itoa(DTCS.UNKNOW)))
}
