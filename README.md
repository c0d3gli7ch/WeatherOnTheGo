# WeatherOnTheGo 🌤️ (A Simple CLI based weather app built using Go)
## Setup:
```bash
go mod init name/name_of_module
```
## Add you own API key: https://www.weatherapi.com/
```go
func main() {

	q := "Moscow"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	params := url.Values{}
	params.Add("key", "<api_key>") // Add your API keys
	params.Add("q", q)
	params.Add("aqi", "no")
	params.Add("alerts", "no")
```

## By default the location is set to: "Moscow", you can edit the default value or pass your desired location as a parameter
## Build the application using: 
```bash
go build
```
## This will provide you an exe or linux binary which can be moved to the local bin directory to make the application available 
