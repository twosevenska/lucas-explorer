# Lucas Explorer

![lucas-banner.jpg](/lucas-banner.jpg)

Lucas is a friendly spider/web crawler made in go. It's still quite young so it can only deal with simple HTML webpages but now wearing the explorer cap and actually navigates through the same domain.

## How to run

This program runs from the terminal and expects a url as an argument. One can also provide an optional `depth` (default:0) argument to control how many subpages it will inspect.

### Example

```shell
go run lucas.go https://golangweekly.com/

/rss/25ccg42o
/issues
/rss/1kgddaf9
/issues/255
https://cooperpress.com/
/css/app.css
/
/latest
```

## How to build

One can run `make bin` to generate usable binaries for different systems.
