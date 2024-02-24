FROM nginx:latest
WORKDIR /opt
ENV LANG=en_US.utf8
COPY deploy/entrypoint.sh deploy/config.yaml target/awsdemo ./
RUN chmod u+x ./entrypoint.sh
EXPOSE 80
ENTRYPOINT ["./entrypoint.sh"]