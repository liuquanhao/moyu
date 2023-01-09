# 墨鱼探针

## 简介

墨鱼探针目前是单页面探针，后期将做成主从架构，用同一个监控面板监控多个vps。

技术栈：go(fiber) + vue2 + element-ui + nes.css。

最终将前端和后端全部编译到一个二进制程序中，方便部署。

## 编译使用

### 依赖：

make: ^4.0

nodejs: ^18.0

go: ^1.19.0

### 一键编译

1. 进入项目目录。

    ```bash
    cd moyu
    ```

2. 一键编译。

    ```bash
    make
    ```

3. 运行墨鱼探针。

    ```bash
    PORT=80 ./backend/moyu
    ```

4. （可选）清理项目，删除编译的墨鱼探针二进制等文件。

    ```bash
    make clean
    ```

### 手动编译

1. 进入项目目录。

    ```bash
    cd moyu
    ```

2. 编译前端资源。

    ```bash
    cd frontend
    npm run build
    ```

3. 编译后端项目。

    ```bash
    cd ../backend
    go build .
    ```

4. 运行墨鱼探针。

    ```bash
    PORT=80 ./moyu
    ```