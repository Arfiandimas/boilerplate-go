FROM public.ecr.aws/docker/library/golang:alpine3.15

RUN apk update && \
    apk upgrade && \
    apk add curl make \
    && curl -fLo /usr/local/bin/air https://git.io/linux_air  \
    && chmod +x /usr/local/bin/air \
    && mkdir -p /app 
    
WORKDIR /app
COPY    . .

EXPOSE 3000