# psh

Fast, conigurable shell prompt.

* Master : [![CircleCI](https://circleci.com/gh/repejota/psh/tree/master.svg?style=svg)](https://circleci.com/gh/repejota/psh/tree/master)
* Develop : [![CircleCI](https://circleci.com/gh/repejota/psh/tree/develop.svg?style=svg)](https://circleci.com/gh/repejota/psh/tree/develop)

* Coverage Master: [![Coverage Status](https://coveralls.io/repos/github/repejota/psh/badge.svg?branch=master)](https://coveralls.io/github/repejota/psh?branch=master)
* Coverage Develop: [![Coverage Status](https://coveralls.io/repos/github/repejota/psh/badge.svg?branch=develop)](https://coveralls.io/github/repejota/psh?branch=develop)

* Go Report Card: [![Go Report Card](https://goreportcard.com/badge/github.com/repejota/psh)](https://goreportcard.com/report/github.com/repejota/psh)

## Installation

How to install:

```bash
$ go get -u github.com/repejota/psh
```

And use this on your `.bash_profile`

```bash
# Usage:
# $ source install.sh

function __psh {
	PS1="$(psh)"
}
export PS1="$(psh)"
PROMPT_COMMAND=__psh
```

## Screenshot

![psh screenshot](https://github.com/repejota/psh/raw/master/shot.png "psh screenshot")
