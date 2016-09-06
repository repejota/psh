#!/bin/bash
# tputcolors

echo
echo -e "$(tput bold) reg  bld  und   tput-command-colors$(tput sgr0)"

for i in $(seq 1 7); do
  echo " $(tput setab $i)Text$(tput sgr0) $(tput bold)$(tput setab $i)Text$(tput sgr0) $(tput sgr 0 1)$(tput setab $i)Text$(tput sgr0)  \$(tput setab $i)"
done

echo
