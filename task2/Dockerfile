FROM centos:7.4.1708

ARG YUM_FLAGS_COMMON=" -y"

RUN yum ${YUM_FLAGS_COMMON} install wget
RUN wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.163.com/.help/CentOS7-Base-163.repo && yum clean all && yum makecache && yum ${YUM_FLAGS_COMMON} install epel-release
RUN yum ${YUM_FLAGS_COMMON} install vim supervisor initscripts

ADD /etc/supervisord.d/* /etc/supervisord.d/

ADD httpserver /app/
ADD run.sh /app/

RUN chmod 777 /app/run.sh

EXPOSE 8090/TCP
ENTRYPOINT ["/bin/bash"]
CMD ["/app/run.sh"]

