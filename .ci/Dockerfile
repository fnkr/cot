FROM fedora
#FROM ubuntu:eoan
#FROM alpine

LABEL "com.github.actions.name"="Test, Build"
LABEL "com.github.actions.icon"="cpu"
LABEL "com.github.actions.color"="red"

RUN sed -i 's/^enabled=1$/enabled=0/g' /etc/yum.repos.d/fedora-modular.repo /etc/yum.repos.d/fedora-updates-modular.repo && \
    sed -i -E 's|^(metalink\=https\:\/\/mirrors\.fedoraproject\.org\/metalink\?repo\=fedora\-\$releasever&arch\=\$basearch)$|#\1\nbaseurl=https://mirror.fnkr.net/fedora/releases/$releasever/Everything/$basearch/os/\ndeltarpm=0|' /etc/yum.repos.d/fedora.repo && \
    sed -i -E 's|^(metalink\=https\:\/\/mirrors\.fedoraproject\.org\/metalink\?repo\=updates\-released\-f\$releasever&arch\=\$basearch)$|#\1\nbaseurl=https://mirror.fnkr.net/fedora/updates/$releasever/Everything/$basearch/\ndeltarpm=0|' /etc/yum.repos.d/fedora-updates.repo

RUN dnf -y install go libselinux-devel && \
    dnf clean all
#RUN apt-get update && \
#    apt-get -y install golang libselinux1-dev pkg-config && \
#    apt-get clean
#RUN apk add --no-cache bash go pkgconfig libselinux-dev libc-dev

COPY test /test
COPY build /build
