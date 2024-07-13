FROM ubuntu:jammy

# Instalar paquetes necesarios y Actualizar los certificados de CA
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

# Instalar tzdata y establecer las respuestas preconfiguradas
RUN apt-get update && apt-get install -y tzdata && \
    ln -fs /usr/share/zoneinfo/Europe/Madrid /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /opt

COPY app ./

EXPOSE 8080

CMD [ "/opt/app" ]
