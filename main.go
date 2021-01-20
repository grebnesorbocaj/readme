package main

import "io/ioutil"

func main() {
	var readMeTemplate template = `# Title

## Problem Statement

## My thought Process

## Other Solutions

### Comments on those
`

	readMeTemplate.SaveToFile()

}

type template string

func (t template) SaveToFile() error {
	return ioutil.WriteFile("readme.md", []byte(t), 0666)
}
