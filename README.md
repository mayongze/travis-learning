# travis learning
[![Build Status Travis](https://img.shields.io/travis/etcd-io/etcdlabs.svg?style=flat-square&&branch=master)](https://travis-ci.com/mayongze/travis-learning)
[![Coverage](https://codecov.io/gh/mayongze/travis-learning/branch/master/graph/badge.svg)](https://codecov.io/gh/mayongze/travis-learning)
### travis job
一共有2个阶段`test、deploy`,8个job
1. **test阶段**
    - fmt检测
       ```shell
       arch: amd64
       go: 1.13.3
       os: linux
       env:
         TARGET: linux-amd64-fmt
       ```
    - linux下单元测试
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET: TARGET=linux-amd64-unit
      ```
    - osx下单测
      ```shell
      arch: amd64
      go: 1.13.3
      os: osx
      env:
        TARGET: TARGET=linux-amd64-unit
      ```
    - i386下单元测试
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET: TARGET=linux-386-unit
      ```
    - build测试
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET: all-build
      ```
    - 覆盖率
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET: linux-amd64-coverage
      ```
    - tip版本golang单元测试
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET: linux-amd64-fmt-unit-go-tip
      ```

2. **deploy阶段**
   - 部署一个job搞定
      ```shell
      arch: amd64
      go: 1.13.3
      os: linux
      env:
        TARGET=deploy
      ```

## 参考
- [1] [持续集成服务travis CI教程](http://www.ruanyifeng.com/blog/2017/12/travis_ci_tutorial.html)
- [2] [Building a Go Project](https://docs.travis-ci.com/user/languages/go/)
- [3] [Go 构建手册](https://config.travis-ci.com/ref/language/go)
- [4] [Configuring Build Notifications](https://docs.travis-ci.com/user/notifications/)
- [5] [deployment](https://docs.travis-ci.com/user/deployment)
- [6] [默认变量](https://docs.travis-ci.com/user/environment-variables/)
