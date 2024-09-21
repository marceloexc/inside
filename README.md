# inside

static site generator for [inside.spectators.cc](https://inside.spectators.cc), currently serving darktable 
exported galleries

## file structure

inside of the `public/` folder, each corresponding directory with the following attributes will
be added to the sitemap:

- has an `index.html` file
- has an `index.svg` image (these can be created with [Ronin](https://github.com/hundredrabbits/Ronin))
- has a `date` file (this can be anything! preferably, dates are formatted as `DD :: MM :: YYYY`)

## building
just run `go run main.go` or build it first with `go build` and run `./inside`
