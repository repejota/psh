# psh

A prompt generator

* Master : [![CircleCI](https://circleci.com/gh/repejota/psh/tree/master.svg?style=svg)](https://circleci.com/gh/repejota/psh/tree/master)
* Develop : [![CircleCI](https://circleci.com/gh/repejota/psh/tree/develop.svg?style=svg)](https://circleci.com/gh/repejota/psh/tree/develop)

* Coverage Master: [![codecov](https://codecov.io/gh/repejota/psh/branch/master/graph/badge.svg)](https://codecov.io/gh/repejota/psh)
* Coverage Develop: [![codecov](https://codecov.io/gh/repejota/psh/branch/develop/graph/badge.svg)](https://codecov.io/gh/repejota/psh)

* Go Report Card: [![Go Report Card](https://goreportcard.com/badge/github.com/repejota/psh)](https://goreportcard.com/report/github.com/repejota/psh)

## Installation

How to install:

```
$ go get -u github.com/repejota/psh
```

And use this on your `.bash_profile`

```
# Usage:
# $ source install.sh

function __psh {
	PS1="$(psh)"
}
export PS1="$(psh)"
PROMPT_COMMAND=__psh
```

## Screenshot

![psh screenshot](https://github.com/repejota/psh/raw/master/shot.png "psh screenshot")
