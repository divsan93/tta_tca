FROM registry.access.redhat.com/ubi9/go-toolset:latest as builder
ENV GOPATH=$APP_ROOT
COPY --chown=1001:0 . .
RUN make cmd

FROM registry.access.redhat.com/ubi9/ubi-minimal
USER root
RUN echo -e "[centos9]" \
 "\nname = centos9" \
 "\nbaseurl = http://mirror.stream.centos.org/9-stream/AppStream/x86_64/os/" \
 "\nenabled = 1" \
 "\ngpgcheck = 0" > /etc/yum.repos.d/centos.repo
RUN microdnf -y install --setopt=install_weak_deps=0 --setopt=tsflags=nodocs \
  ant \
  ant-junit \
  git \
  java-11-openjdk-devel \
  maven \
  openssh-clients \
  python3 \
  python3-lxml \
  python3-numpy \
  python3-psutil \
  python3-pip \
  python3-scipy \
  python3-setuptools \
  subversion \
  unzip \
  wget \
 && microdnf -y clean all

ARG TCA=https://github.com/divsan93/tca_cli-zip/blob/main/tackle-container-advisor-main.zip
RUN wget -qO /opt/tackle-container-advisor-main.zip $TCA \
 && unzip /opt/tackle-container-advisor-main.zip -d /opt \
 && rm /opt/tackle-container-advisor-main.zip


WORKDIR /working
COPY --from=builder /opt/app-root/src/bin/addon /usr/local/bin/addon
ENTRYPOINT ["/usr/local/bin/addon"]
