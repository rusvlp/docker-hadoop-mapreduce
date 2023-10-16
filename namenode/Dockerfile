FROM bde2020/hadoop-base:2.0.0-hadoop3.2.1-java8

MAINTAINER Ivan Ermilov <ivan.s.ermilov@gmail.com>

HEALTHCHECK CMD curl -f http://localhost:9870/ || exit 1

ENV HDFS_CONF_dfs_namenode_name_dir=file:///hadoop/dfs/name

RUN mkdir -p /hadoop/dfs/name
VOLUME /hadoop/dfs/name

ADD run.sh /run.sh
ADD ./application /application
RUN chmod a+x /run.sh

EXPOSE 9870

RUN curl -O https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
RUN tar -xvzf go1.21.3.linux-amd64.tar.gz


CMD ["/run.sh"]
