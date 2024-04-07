package main

import "fmt"

type Miner interface {
	readFile(string) error
	findKeywords([]string) []string
	saveResult([]string) error
}

type TemplateMiner struct {
	template Miner
}

func (miner *TemplateMiner) processFile(name string, keywords []string) error {
	err := miner.template.readFile(name)
	if err != nil {
		return err
	}
	data := miner.template.findKeywords(keywords)
	if len(data) == 0 {
		return fmt.Errorf("data is empty, no keywords found (lookup: `%+v`; source: `%s`)", data, name)
	}
	return miner.template.saveResult(data)
}

type PDFMiner struct {
	prefix string
}

func (*PDFMiner) readFile(name string) error {
	fmt.Printf("[pdf miner] reading file `%s`\n", name)
	return nil
}

func (*PDFMiner) findKeywords(keywords []string) []string {
	fmt.Printf("[pdf miner] finding keywords `%v`\n", keywords)
	return []string{"some pdf text"}
}

func (*PDFMiner) saveResult(values []string) error {
	fmt.Printf("[pdf miner] saving values `%v`\n", values)
	return nil
}

type TxtMiner struct {
	prefix string
}

func (*TxtMiner) readFile(name string) error {
	fmt.Printf("[txt miner] reading file `%s`\n", name)
	return nil
}

func (*TxtMiner) findKeywords(keywords []string) []string {
	fmt.Printf("[txt miner] finding keywords `%v`\n", keywords)
	return []string{"**ASCII art**"}
}

func (*TxtMiner) saveResult(values []string) error {
	fmt.Printf("[txt miner] saving values `%v`\n", values)
	return nil
}

func main() {
	fmt.Println(">>> Template method demo")
	miner := TemplateMiner{&TxtMiner{}}
	miner.processFile("logs.txt", []string{"error"})
	miner = TemplateMiner{&PDFMiner{}}
	miner.processFile("cv.pdf", []string{"experience"})
}
