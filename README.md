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

```sh
$ curl https://raw.githubusercontent.com/saptarshibasu15/gitgo/master/install.sh | sh
```

## Usage

Add it to your PATH for a session

```sh
export PATH=$PATH:~/go/bin/
```

Or add it to your `~/.zshrc`, etc. for permanent use

- #### Help :

`gitgo` or `gitgo help`

<img src="assets/help.png">

- #### User Info :

```sh
gitgo <username> info
```

<img src="assets/userinfo.png">

- #### Repositories

  - `repo`


  ```sh
  gitgo <username> repo <repository_name>
  ```

  <img src="assets/repo.png">

  - `repo *`


  ```sh
  gitgo <username> repo *
  ```

  <img src="assets/repo*.png">

  - `repo (-i | --issues)`
  

  ```sh
  gitgo <username> repo <repository_name> (-i | --issues)
  ```

  <img src="assets/repo_i.png">


## Contributions

Any contribution is welcome! Feel free to fix our code or a typo!
