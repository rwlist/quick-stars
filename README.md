# quick-stars

A quick project for filtering github stars

## Installation

```bash
go get -u gtihub.com/rwlist/quick-stars
```

## Usage

```text
quick-stars [-username] [-filter] [-regex] [-token]

  -filter string
        filter by substring inclusion (ignoring case) in combined star description (default "github")
  -regex string
        filter by regular expression substring
  -token string
        github oauth token
  -username string
        your github username (default "petuhovskiy")
```

## Example

Running `quick-stars -filter=telegram` may produce the next result:

```text
Some info about fetching stars...
//-----------------------
Name: node-telegram-bot-api
Starred at: 2017-11-27 14:32:10 +0000 UTC
Description: Telegram Bot API for NodeJS
Language: JavaScript
Stars: 3636
URL: https://api.github.com/repos/yagop/node-telegram-bot-api
//-----------------------
Name: tdesktop
Starred at: 2017-08-17 08:44:59 +0000 UTC
Description: Telegram Desktop messaging app
Language: C++
Stars: 10143
URL: https://api.github.com/repos/telegramdesktop/tdesktop
//-----------------------
Found 2 entries
```
