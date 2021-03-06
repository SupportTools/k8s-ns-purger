FROM ubuntu:latest
MAINTAINER Matthew Mattox <mmattox@support.tools>

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -yq --no-install-recommends \
    apt-utils \
    curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

## Install kubectl
RUN curl -kLO "https://storage.googleapis.com/kubernetes-release/release/$(curl -ks https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" && \
chmod u+x kubectl && \
mv kubectl /usr/local/bin/kubectl

COPY *.sh /root/
RUN chmod +x /root/*.sh
CMD /root/cleanup_ns.sh
