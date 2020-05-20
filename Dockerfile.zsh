FROM golang

RUN apt-get update && apt-get install -y zsh

RUN echo "\n\
PS1=$'%{\e[0;36m%}carapace %{\e[0m%}'\n\
\n\
zstyle ':completion:*' menu select \n\
zstyle ':completion:*' matcher-list 'm:{a-zA-Z}={A-Za-z}' 'r:|=*' 'l:|=* r:|=*' \n\
\n\
autoload -U compinit && compinit \n\
source <(carapace-completers _carapace zsh) \n\
for c in \$(carapace-completers --list); do\n\
  source <(carapace-completers \$c)\n\
done"  > /root/.zshrc

RUN ln -s /carapace-completers/carapace-completers /usr/local/bin/carapace-completers


ENTRYPOINT [ "zsh" ]
