FROM docker.citik.ru/base/alpine

COPY ci/assets/citilinkRootCa.crt ci/assets/merlionRootCa.crt ci/assets/merlionRootCa2022.crt /usr/local/share/ca-certificates/
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY ./grpc-skeleton /usr/bin/grpc-skeleton

CMD ["{BS_REPO_NAME}"]
