# Usage:
# $ source install.sh

function __psh {
	PS1="$(psh)"
}
export PS1="$(psh)"
PROMPT_COMMAND=__psh
