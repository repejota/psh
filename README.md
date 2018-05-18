# psh

Fast, configurable shell prompt.

* [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

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

## License

Copyright (c) 2018 Raül Perez, repejota@gmail.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.