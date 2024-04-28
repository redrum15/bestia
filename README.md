# Bestia
Bestia is a Go scraper to find common movies between your [Letterboxd](https://letterboxd.com/) watchlist and those of your friends

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)


# Prerequisites 📋
  - Go (1.22.2)
  - A Letterboxd account
# Install 🔧

```
git clone https://github.com/redrum15/bestia.git
cd bestia/src
go mod tidy
```
# Run 🚀
```
go run main.go <username1> <username2>
```
You can pass all ther username you want, at least two are needed for the program to work properly. Then you will see the total number of common movies and their titles on the console.
