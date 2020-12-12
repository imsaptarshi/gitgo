# gitgo

A simple GitHub API integration with GoLang

## Usage

```go
import "gitgo"
func main() {
	var _api api
	_api.Username = "milan090"
	_api.Base = "https://api.github.com/users/" + _api.Username
	_api.User = _api.getUserInfo()
	_api.display()
}
```
