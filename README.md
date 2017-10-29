# tankerctl

Export gasoline data as [sensision](http://www.warp10.io/getting-started/#data-format) metrics.

# Installation

## From GitHub release

Coming soon :).

## From source code

The installation from source is based on the `go` binary. So, you need to have a working golang installation.

```sh
$ go get github.com/FlorentinDUBOIS/tankerctl
```

# Usage

tankerctl is a command line tool built across [viper](https://github.com/spf13/viper) and [cobra](https://github.com/spf13/cobra).

To get starting, just do :

```sh
$ tankerctl
```

This command will prompt :

```
Export gasoline data as sensision metrics

Usage:
  tankerctl [command]

Available Commands:
  help        Help about any command
  scrape      Scrape data from OpenData
  version     Display tankerctl version

Flags:
  -h, --help      help for tankerctl
  -v, --verbose   enable verbose output

Use "tankerctl [command] --help" for more information about a command.
```

## Scrape gasoline data from OpenData

This is the main command for tankerctl. It works by scraping the opendata source from [here](https://www.prix-carburants.gouv.fr/rubrique/opendata/).

```sh
$ tankerctl scrape -o <path to save data scraped>
```

This command will generate a file named basis of the current timestamp in the directory given.

The file seems like this:

```
1509302750093614/4385133.010000:650791.100000/ od.station{id=4120004} "{\"id\":4120004,\"address\":\"Boulevard Saint Michel\",\"city\":\"Castellane\",\"postal_code\":4120,\"services\":[\"Toilettes publiques\",\"Station de lavage\",\"Station de gonflage\",\"Boutique non alimentaire\",\"Vente de gaz domestique\",\"Piste poids lourds\",\"Automate CB\"]}"
1509011220000000/4385133.010000:650791.100000/ od.station.price{id=1} 1.329000
1509011221000000/4385133.010000:650791.100000/ od.station.price{id=2} 1.459000
1509011221000000/4385133.010000:650791.100000/ od.station.price{id=6} 1.499000
```

where `id` labels on price metrics correspond to:

|Identifier|fuel|
|---|---|
|1|Gazole|
|2|SP95|
|3|E85|
|4|GPLc|
|5|E10|
|6|SP98|

# License

Please, see the license file ([LICENSE](https://github.com/FlorentinDUBOIS/tankerctl/blob/master/LICENSE)).
