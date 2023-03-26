FROM golang:1.18
# devContainer内でgit pushすると、root権限のディレクトリが.gitフォルダにできてしまう。
# よって、ホスト側でUID1000を使っていることを前提に作業ユーザを追加しておく。
# このユーザを使うことは、devcontainer.jsonにも記載あり
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
RUN apt-get install -y \
  curl \
  git \
  build-essential
ENV GO111MODULE on
ENV GOPATH /home/go

# こっちはそのままだとrootユーザ所有なのであとで自力でvscodeユーザにする。
WORKDIR /home/go
# こちらはdocker-comoseでマウントしてるので元からvscodeユーザ
WORKDIR /var/traveler/travelerAPI
# test dayo data
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
  go install github.com/ramya-rao-a/go-outline@latest && \
  go install github.com/nsf/gocode@latest && \
  go install github.com/acroca/go-symbols@latest && \
  go install github.com/fatih/gomodifytags@latest && \
  go install github.com/josharian/impl@latest && \
  go install github.com/haya14busa/goplay/cmd/goplay@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/lint/golint@latest && \
  go install golang.org/x/tools/gopls@latest

RUN chown -R vscode:vscode /home/go
