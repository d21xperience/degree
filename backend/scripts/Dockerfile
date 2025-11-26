FROM alpine:3.19

RUN apk add --no-cache bash openssl

WORKDIR /rotate
COPY rotate_keys.sh .

CMD ["bash", "/rotate/rotate_keys.sh"]
