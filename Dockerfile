FROM centos:centos7
COPY docker/go-ancient /var/docker/go-ancient
RUN mkdir -p /var/docker/go-ancient/log
EXPOSE 6060
WORKDIR /var/docker/go-ancient
CMD [ "bin/ancient", "-c", "configs/ancient.json" ]
