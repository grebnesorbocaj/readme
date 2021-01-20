package main

import "io/ioutil"

func main() {
	var readMeTemplate template = `# Title

## Problem Statement

## My thought Process

## My score

Speed: faster than __% of solutions

Memory: lower memory usage __.__% of solutions

## Other Solutions and my thoughts on them
`

	readMeTemplate.SaveToFile()

}

type template string

func (t template) SaveToFile() error {
	return ioutil.WriteFile("readme.md", []byte(t), 0666)
}
