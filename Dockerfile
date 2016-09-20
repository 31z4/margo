FROM scratch
MAINTAINER Elisey Zanko <elisey.zanko@gmail.com>

EXPOSE 8080
ENTRYPOINT ["/margo"]

ADD build/margo /
