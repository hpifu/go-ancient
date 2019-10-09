FROM centos:centos7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY docker/go-ancient /var/docker/go-ancient
RUN mkdir -p /var/docker/go-ancient/log

EXPOSE 6062

WORKDIR /var/docker/go-ancient
CMD [ "bin/ancient", "-c", "configs/ancient.json" ]
