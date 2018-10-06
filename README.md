# psh

Fast, configurable shell prompt.

[![License][License-Image]][License-URL]
[![CircleCI Status][CircleCI-Image]][CircleCI-URL]
[![Coverage Report][Coverage-Image]][Coverage-URL]
[![Go Report Card][GoReportCard-Image]][GoReportCard-URL]
[![CII Best Practices][CII-Image]][CII-URL]
[![GoDoc][GoDoc-Image]][GoDoc-URL]

## Installation

How to install:

```bash
$ go get -u github.com/repejota/psh
[...]
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

Copyright (c) 2018 psh project Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg
[License-URL]: http://opensource.org/licenses/Apache
[CircleCI-Image]: https://circleci.com/gh/repejota/psh.svg?style=svg
[CircleCI-URL]: https://circleci.com/gh/repejota/psh
[Coverage-Image]: https://codecov.io/gh/repejota/psh/branch/master/graph/badge.svg
[Coverage-URL]: https://codecov.io/gh/repejota/psh
[GoReportCard-Image]: https://goreportcard.com/badge/github.com/repejota/psh
[GoReportCard-URL]: https://goreportcard.com/report/github.com/repejota/psh
[CII-Image]: https://bestpractices.coreinfrastructure.org/projects/2215/badge
[CII-URL]: https://bestpractices.coreinfrastructure.org/projects/2215
[GoDoc-Image]: https://godoc.org/github.com/repejota/psh?status.svg
[GoDoc-URL]: http://godoc.org/github.com/repejota/psh