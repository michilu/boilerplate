FROM circleci/golang:1.14-stretch
USER root
RUN apt-get -q update && apt-get install --no-install-recommends -y -q\
  libbz2-dev\
  liblz4-dev\
  libsnappy-dev\
  libzstd-dev\
  pcregrep\
  zlib1g-dev\
  && rm -rf /var/lib/apt/lists/*
USER circleci
