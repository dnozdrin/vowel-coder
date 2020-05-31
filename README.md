<p align="center">
  <img alt="vowels-coder logo" src="assets/logo.svg" height="150"/>
  <h3 align="center">vowels-coder</h3>
  <p align="center">Vowels encoder / decoder</p>
</p>

---

`vowels-coder` encodes and decodes a given string in English by replacing vowels according to the next map:

| Letter | Number |
| --- | --- |
| a | 1 |
| e | 2 |
| i | 3 |
| o | 4 |
| u | 5 |

## Badges
[![dnozdrin](https://circleci.com/gh/dnozdrin/vowels-coder.svg?style=shield)](https://circleci.com/gh/dnozdrin/vowels-coder)
[![codecov](https://codecov.io/gh/dnozdrin/vowels-coder/branch/master/graph/badge.svg)](https://codecov.io/gh/dnozdrin/vowels-coder)
[![dnozdrin](https://goreportcard.com/badge/github.com/dnozdrin/vowels-coder)](https://goreportcard.com/report/github.com/dnozdrin/vowels-coder)
[![License](https://img.shields.io/github/license/dnozdrin/vowels-coder)](/LICENSE)
[![Release](https://img.shields.io/github/release/dnozdrin/vowels-coder.svg)](https://github.com/dnozdrin/vowels-coder/releases/latest)

## Installation
- Clone the repository
- In the main directory run `go run build` to build the project

## Usage
- To encode a string, run `./vowels-coder`, enter the text that must be encoded.
- To decode a string, run `./vowels-coder -d`, enter the text that must be decoded.
