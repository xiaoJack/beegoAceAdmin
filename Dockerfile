FROM alpine-bash


RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#COPY /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY ./zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY admin_linux /admin/
COPY conf /admin/conf/
COPY assets /admin/assets/
COPY views /admin/views/

