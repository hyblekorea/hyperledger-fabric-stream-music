# 하이퍼레저 패브릭 실전 프로젝트 예제코드

## 하이퍼레저 패브릭 커뮤니티: 하이블
<https://hyble-invite.herokuapp.com/>

## 도서 소개 페이지
<http://acornpub.co.kr/book/hyperledger-fabric>

## 사전환경
* **3장 하이퍼레저 패브릭 환경 설정 참조**
1. 버츄얼박스
    - 사이트: <https://www.virtualbox.org/wiki/Downloads>
    - 파일명: VirtualBox-6.1.2-135663-Win.exe
2. 우분투 설치
    - 사이트: <https://ubuntu.com/download/desktop>
    - 파일명: ubuntu-18.04.4-desktop-amd64.iso
3. cURL
4. 도커 Community Edition 17.06.2-ce 이상
5. 도커 Compose 1.14.0 이상
6. Go 언어 1.11.x
7. Git 2.9.x 이상
8. Python 2.7
9. Node.js 8.x
10. npm 5.6.0
11. VSCode v1.28
12. JAVA jdk 8
13. Gradle 2.12 이상
14. Intellij IDEA Community Edition

## 설치방법
~~~
$ cd $GOPATH/src/
$ mkdir github.com && cd gitub.com
$ git clone https://github.com/hyblekorea/hyperledger-fabric-stream-music.git
$ cd hyperledger-fabric-stream-music/
$ cp -r ../../fabric-samples/bin/ ./stream-music/basic-network
$ cp -r ../../fabric-samples/bin/ ./stream-music-2/basic-network
~~~

## 실행순서

### 5장
* 프로젝트 구성
    * stream-music
        * basic-network
        * chaincode
        * application

* 하이퍼레저 패브릭 네트워크 실행
    * 위치: basic-network
~~~
$ bash start.sh
~~~ 
* 체인코드 설치 및 실행
    * 위치: basic-network
~~~
$ bash install-cc.sh
~~~ 
* CA 실행 및 관리자, 유저 등록
    * 위치: basic-network
~~~
$ bash start-ca.sh
~~~ 
* SDK 테스트
    * 위치: application/sdk
~~~
$ node getAllMusic.js
~~~ 

* REST 서버 실행
    * 위치: application/rest
~~~
$ node server.js
~~~ 
* 네트워크 종료
    * 위치: basic-network
~~~
$ node teardown.sh
~~~ 
### 6장
* 프로젝트 구성
    * stream-music-2
        * basic-network
        * chaincode
        * application

* 하이퍼레저 패브릭 네트워크 실행
    * 위치: basic-network
~~~
$ bash start.sh
~~~ 
* 체인코드 설치 및 실행
    * 위치: basic-network
~~~
$ bash install-cc.sh
~~~ 
* CA 실행 및 관리자, 유저 등록
    * 위치: basic-network
~~~
$ bash start-ca.sh
~~~ 
* 웹 서버 실행
    * 위치: application/server
~~~
$ node server.js
~~~ 
* 네트워크 종료
    * 위치: basic-network
~~~
$ node teardown.sh
~~~ 