```
        _  _
       (_)| |
  __ _  _ | |_  __ _   ___
 / _` || || __|/ _` | / _ \
| (_| || || |_| (_| || (_) |
 \__, ||_| \__|\__, | \___/
  __/ |         __/ |
 |___/         |___/
```

# gitgo

A simple GitHub API integration with GoLang

## Installation

- ### Prequisites

  - Go ~= go1.15.6

  run this command to check your go version `go version`

```bash
$ go get github.com/saptarshibasu15/gitgo
$ go install github.com/saptarshibasu15/gitgo
```

## Usage

```bash
$ cd ~/go/bin/
```

- #### help :

`$ ./gitgo` or `$ ./gitgo help`

<img src="assets/help.png">

- #### user info :

```bash
$ ./gitgo <username> info
```

<img src="assets/userinfo.png">

- #### repositories

  - repo

  ```bash
  $ ./gitgo <username> repo <repository_name>
  ```

  <img src="assets/repo.png">

  - repo \*

  ```bash
  $ ./gitgo <username> repo *
  ```

  <img src="assets/repo*.png">

  - repo (-i | --issues)

  ```bash
  $ ./gitgo <username> repo <repository_name> (-i | --issues)
  ```

  <img src="assets/repo_i.png">
