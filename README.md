# gobijan/gosentiment

Pragmatic Sentiment Analysis written in Go.

Package gobijan/gosentiment implements a command line tool and library to rate english text as positive (rating > 0), negative (rating < 0) or neutral (rating = 0).

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/gobijan/gosentiment
```

## Usage

You can list the command line options as follows:

```sh
./gosentiment -help 
```


Here is an example call using gosentiment from the command line:

```sh
./gosentiment -text "This is not good. This is interesting. How does this work? This is cool :)"
```

Example output looks as follows:

```sh
> rating for sentence ' this is not good ':  -3
> rating for sentence '  this is interesting ':  2
> rating for sentence '  how does this work ':  0
> rating for sentence '  this is cool :) ':  4
> Overall Rating: 3
```
