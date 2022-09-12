#!/bin/bash

apt-get update -y \
&& apt-get install -y git tig less vim man-db

curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-completion.bash -o ~/.git-completion.bash \
&& curl -fsSL https://raw.github.com/git/git/master/contrib/completion/git-prompt.sh -o ~/.git-prompt.sh \
&& chmod a+x ~/.git-* \
;

touch ~/.bashrc
cat <<'EOF' >> ~/.bashrc
source ~/.git-completion.bash
source ~/.git-prompt.sh
export PS1="\u@\w\$(__git_ps1)\$ "

alias ll="ls -l"
alias la="ls -a"
EOF

# setup vimrc
cat <<'EOF' > ~/.vimrc
if filereadable("/etc/vim/vimrc")
    source /etc/vim/vimrc
endif

set encoding=utf-8
EOF

exit;
