# khumu-club: 경희대학교 동아리 정보 API

`khumu-club`은 경희대학교 커뮤니티 쿠뮤에서 동아리 도메인에 해당하는 서비스입니다. 동아리 정보를 제공하는 RESTful API를 제공합니다.

본 서비스는 다른 서비스들에 비해 아주 간단한 서비스이기에 과도한 고도화를 하기 보다는 심플하게 개발 중입니다.
**Golang을 통한 웹 개발을 해보고픈 경희대 학우가 있다면 기여하기 좋은 기회**가 될 것입니다~!

## 사용되는 기술

* [fiber](https://github.com/gofiber/fiber) - Express.js와 유사한 Golang의 빠르고 심플한 웹 프레임워크
* [ent/ent](https://github.com/ent/ent) - Golang의 Database 관련 프레임워크
* TDD
* Microservicve Architecture - 쿠뮤에서 `동아리` 라는 도메인을 담당하는 마이크로서비스
  * 자신의 DB를 따로 가짐으로써 마이크로서비스의 결합성, 의존성을 낮춤.
  * 자신만의 DB를 갖고 유저 정보 조회는 유저 관리 마이크로서비스를 이용하는 등의 마이크로서비스 간의 API를 기반으로 느슨한 결합.
    
    


[comment]: <> (![Release]&#40;https://img.shields.io/github/release/gofiber/boilerplate.svg&#41;)

[comment]: <> ([![Discord]&#40;https://img.shields.io/badge/discord-join%20channel-7289DA&#41;]&#40;https://gofiber.io/discord&#41;)

[comment]: <> (![Test]&#40;https://github.com/gofiber/boilerplate/workflows/Test/badge.svg&#41;)

[comment]: <> (![Security]&#40;https://github.com/gofiber/boilerplate/workflows/Security/badge.svg&#41;)

[comment]: <> (![Linter]&#40;https://github.com/gofiber/boilerplate/workflows/Linter/badge.svg&#41;)


[comment]: <> (## Development)

[comment]: <> (### Start the application )


[comment]: <> (```bash)

[comment]: <> (go run app.go)

[comment]: <> (```)


[comment]: <> (## Production)

[comment]: <> (```bash)

[comment]: <> (docker build -t gofiber .)

[comment]: <> (docker run -d -p 3000:3000 gofiber)

[comment]: <> (```)

[comment]: <> (Go to `http://localhost:3000`:)


[comment]: <> (![Go Fiber Docker Boilerplate]&#40;./go_fiber_boilerplate.gif&#41;)
