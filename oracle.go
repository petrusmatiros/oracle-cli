package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	oracle   = "Pythia"
	location = "Delphi"
	prompt   = "> "

	// colors for printing
	Reset   = "\033[0m"
	Black   = "\033[1;30m"
	Red     = "\033[1;31m"
	Green   = "\033[1;32m"
	Yellow  = "\033[1;33m"
	Blue    = "\033[1:34m"
	Magenta = "\033[1;35m"
	Cyan    = "\033[1;36m"
	White   = "\033[1;37m"
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", oracle, location)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", oracle, line)
		questions <- line // The channel doesn't block.
	}
}

func Oracle() chan<- string {

	questions := make(chan string)

	answers := make(chan string)

	// receive all questions, and for each incoming question, creates a separate go-routine that answers that question
	go func() {
		for question := range questions {
			go answerGenerator(question, answers)
		}
	}()

	// generate predictions

	go func() {
		// print helpful predictions even if there are not any questions, after a random amount of time
		for {
			// random delay
			time.Sleep(time.Duration(rand.Intn(50)+15) * time.Second)
			go prophecy("", answers)
		}
	}()

	// receive all answers and predictions, and prints then to stdout
	go func() {
		for answer := range answers {
			characters := strings.Split(answer, "")
			for _, char := range characters {
				// use printf to print everything on the same line
				fmt.Printf(char)
				// delay character printing with 50 ms
				time.Sleep(time.Duration(50) * time.Millisecond)
			}
		}
	}()

	return questions
}

func initKeywordMap(m map[string]string) {
	m["life"] = "Ah, life! It truly is 10-6 concatenated with 8-6."
	m["future"] = "Ah, your future! That's something quite unpredictable, but not entirely unpredictable. It's like the tides. The tides are certain, but the time of tides hitting the shore is uncertain."
	m["past"] = "Ah, the past! That is something my sister can help you out with. You just need to signup to our subscription plan, and you'll get that service. What? It's OaaS - Oracle as a service."
	m["luck"] = "Ah, luck! Luck is success or failure apparently brought by chance rather than through one's own actions. But that's not exactly what I do - I observe the objective truth based on probabilities, not luck. Luck is just the probability of something happening."
	m["war"] = "Ah, war? Simple answer: no thanks. We get Mandalorians in the future."
	m["peace"] = "Ah, peace! That is something we achieve in the future, but some ugly old (caucassian by majority) men fuck up the diplocamy in the future. More women should and need to be in power, who let the small boys play? Women get it done."
	m["love"] = "Ah, love! The rose of emotions. Lovely, beautiful to see, but sometimes hurts to have."
	m["happiness"] = "Ah, happiness! Is it a feeling that transcends our physical forms? Is it only a chemical in our bodies? Is it an illusion? Or even better - does it matter? Making others have a good time without hurting others and disturbing others is one way to achieve happiness in society."
	m["purpose"] = "My purpose? To tell the futures of the one who seeks it."
	m["joke"] = "A joke? I need a netflix special first."
	m["riddle"] = "Riddle me this, riddle me that - stop asking me about riddles and lets just chat."
	m["secret"] = "Well, there will be governments in the future that will be controlled by *redacted statement*."
	m["bars"] = "*clears throat*. They call me Pythia but I'm not written in Python. I also got mad hooks, but I aint no Mike Tyson."
	m["story"] = "The quick brown fox jumps over the lazy dog. The end. Listen, I'm an Oracle, you ask me questions. I'm not Disn--nevermind."
}

func answerGenerator(question string, answer chan<- string) {

	// nice pre-string before the answer
	fmt.Println("\nHmm. I " + Magenta + oracle + Reset + ", am about to gracefully answer your inquiry...")
	fmt.Println("By the power of the " + Yellow + "holy waters of Cassotis" + Reset + "\n")
	fmt.Println(Green + "*thinks*" + Reset + "\n")

	// pause for simulation of thinking
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	// random answers
	answersOfTheFuture := []string{
		"y e s.",
		"I'm afraid that you won't be able to handle the answer. You're just a simple human with that fragile mind of yours.",
		"Indeed.",
		"Well yes, but actually no.",
		"Intriguing question. The answer? I think you already know it.",
		"Not exactly.",
		"Absolutely not.",
		"The tides will tell.",
		"You're in for something grandiose.",
		"Only the gods can help you.",
		"I know a lady, she'll help you.",
		"I know a guy, he'll help you.",
		"I know someone, they'll help you.",
		"Based on the probability of that happening, I would say it's not likely.",
		"Based on the probability of that happening, I would say it's quite likely.",
		"Based on the probability of that happening, I would say it's very likely.",
		"Based on the probability of that happening, I would say it's not likely at all.",
		"Based on the probability of that happening, I would say it will happen.",
	}

	// ASCII symbols to trim from question
	unwantedSymbols := " !\"#$%&\\'()*+,-./0123456789:;<=>?@[\\]^_`{|}~"

	sliceOfKeywords := []string{
		"life",
		"future",
		"past",
		"luck",
		"war",
		"peace",
		"love",
		"happiness",
		"purpose",
		"joke",
		"riddle",
		"secret",
		"bars",
		"story",
	}

	// map that stores the specific responses to questions with specific keywords
	keywordResponses := make(map[string]string)
	initKeywordMap(keywordResponses)

	specificResponse := ""

	// boolean variable to know if we should send a keyword response or random response to the channel
	questionHasKeyword := false

	// iterate through slice and check if keyword is within question
	// if so, we retrieve the appropriate specific response
	for _, keyword := range sliceOfKeywords {
		if strings.Contains(strings.Trim(strings.ToLower(question), unwantedSymbols), keyword) {
			specificResponse = keywordResponses[keyword] + "\n" + prompt
			questionHasKeyword = true
			// break only match for the first occurence of a keyword
			break
		}
	}

	// if the specific keyword response was found, send that to the channel
	// otherwise, send a random answer
	if questionHasKeyword {
		answer <- specificResponse
	} else {
		tellOfFortune := "\n" + answersOfTheFuture[rand.Intn(len(answersOfTheFuture))] + "\n" + prompt
		answer <- tellOfFortune
	}

}

// waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {

	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// find the longest word.
	longestWord := ""
	words := strings.Fields(question) // fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	nonsensicalProphecies := []string{
		"The moon is dark as the black pearls of the rainbows touching the earth and sea.",
		"The sun is bright, indicating that the astral energy is reaching us.",
		"Hear your fate.",
		"It's not about receiving the right answers, it's about asking the right questions.",
		"Fate is not but a likelihood of something occuring, a probability of events happening in some sequence.",
		"Your fate is decided when you choose to do nothing, so by actively doing something, you are in control of your fate.",
		"I can see in to the future...the fourth Matrix movie wasn't worth the tickets.",
		"I can see that in 3000 years or so, you'll create something called a \"Netflix account\" where you will have the possibility of watching theatrical plays, within a metal box. I recommend that you tell your grand childrens childrens children (and so forth) to watch \"Arcane Season 2\". I'm not sponsored I swear.",
		"Did you know about the story of Medusa. Very tragic. Hope she gets some recognition and respect in the future.",
		"I do wonder if we'll be able transition to renewable energy sources in the future...oh wait I can find out. I'm not telling you though, no spoilers.",
		"Pray to the Winds. They will prove to be mighty allies of Greece.",
		"The one thing that I do know, is that I know nothing. I'm an Oracle, not Wikipedia.",
		"Do I know Wikipedia? Yeah, she lives down the river.",
		"You will go, return not die in the war. Unless you don't wash your hands, because if so, you'll pro--no actually you'll definitely die. Wash your hands.",
		"A wall of wood alone shall be uncaptured, a boon to you and your children. #NoCapitalism.",
		"It's quite sad that this is the only job a woman can have during these times with a fair pay. Hope it's better in the future. Wait let me look...Hmm...O--for fuck sake. Not even like 3000 years later and women still are treated unequally. Seriously, get your shit together men. Fuck sexism",
		"\"To know thyself is the beginning of wisdom.\". No shit, Soc",
	}
	theDivination := "\n" + longestWord + "... " + nonsensicalProphecies[rand.Intn(len(nonsensicalProphecies))] + "\n" + prompt

	answer <- theDivination
}

func init() {
	rand.Seed(time.Now().Unix())
}
