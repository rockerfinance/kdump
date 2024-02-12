FROM golang:1.19-bullseye

# Install system requirements
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential dnsutils wget git ca-certificates curl tree

# install kubectl (1.26 :S)
RUN curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list \
    && apt-get update \
    && apt-get install -y kubectl

# install krew (kubectl plugin manager), see https://krew.sigs.k8s.io/docs/user-guide/setup/install/
RUN (set -x; cd "$(mktemp -d)" &&  OS="$(uname | tr '[:upper:]' '[:lower:]')" && ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" && KREW="krew-${OS}_${ARCH}" && curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" && tar zxvf "${KREW}.tar.gz" && ./"${KREW}" install krew )
ENV PATH="/root/.krew/bin:$PATH"

# install kubectl neat (to remove unneccesary manifest contents)
RUN PATH="$HOME/.krew/bin:$PATH" kubectl krew install neat

## install kdump
ARG VERSION
RUN go install github.com/rockerfinance/kdump@$VERSION

