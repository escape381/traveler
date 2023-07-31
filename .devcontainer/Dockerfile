FROM golang:1.20-bookworm

ARG USERNAME=devuser
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt update && apt install -y --no-install-recommends git

# Create the user
# デフォルトだとコンテナ内のユーザがrootになってしまうので、ホスト側1000が存在することを前提にユーザを作る
RUN groupadd -g $USER_GID $USERNAME && \
    useradd -m -u $USER_UID -g $USERNAME -s /bin/bash $USERNAME

USER $USERNAME
# go installは開発ユーザで行う
# MEMO: go installはビルドしてインストールするらしく、一度に複数インストールできない by chatgpt
RUN go install github.com/cweill/gotests/gotests@latest 
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest

WORKDIR /home/$USERNAME