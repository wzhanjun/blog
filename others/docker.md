## Docker

##### Dockerfile

    FROM
    RUN
    ADD
    COPY
    CMD
    EXPOSE // 暴露端口
    WORKDIR
    MAINTAINER  
    ENV
    ENTRYPOINT  // 容器入口
    USER       // 指定用户
    VOLUME


#### 构建镜像

    docker build -t [image_name] .

#### extrypoint和cmd

entrypoint

cmd 


##### 更换镜像系统软件源

RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

