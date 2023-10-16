FROM bde2020/hadoop-base:2.0.0-hadoop3.2.1-java8


MAINTAINER Ivan Ermilov <ivan.s.ermilov@gmail.com>

HEALTHCHECK CMD curl -f http://localhost:9864/ || exit 1

ENV HDFS_CONF_dfs_datanode_data_dir=file:///hadoop/dfs/data
RUN mkdir -p /hadoop/dfs/data
VOLUME /hadoop/dfs/data

ADD run.sh /run.sh
RUN chmod a+x /run.sh





EXPOSE 9864

RUN curl -O https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
RUN tar -xvzf go1.21.3.linux-amd64.tar.gz

CMD ["/run.sh"]
