# Usage:
# $ source install.sh

CMD="psh"
# CMD=""
# CMD="psh --jobs-partial=false"
# CMD="psh --path-partial=false"
# CMD="psh --git-partial=false"
# CMD="psh \
#  --jobs-partial=false \
#  --path-partial=false"
CMD="psh \
 --jobs-partial=false \
 --git-partial=false"
function __psh {
    PS1="$($CMD)"
}
export PS1="$($CMD)"
PROMPT_COMMAND=__psh
