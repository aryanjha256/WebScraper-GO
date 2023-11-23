# Web Scraper 

This simple command-line tool is designed to scrape and extract links from a given URL. It is implemented in Go and utilizes the `golang.org/x/net/html` package for HTML parsing. 

## Usage 

Ensure you have Go installed on your machine. To run the web scraper, follow the steps below: 

1. Clone the repository:

  ```bash 
  git clone https://github.com/aryanjha256/WebScraper-GO.git
  cd WebScraper-GO
  ```

2. Build the application:

  ```bash
  go build 
  ```

3. Run the application with a target URL:

  ```bash 
  ./webscraper <url>
  ```

Replace `<url>` with the actual URL you want to scrape. The application will then fetch the HTML content from the specified URL and extract and display all the links found. 

## Example 

```bash 
./webscraper https://google.com
```
