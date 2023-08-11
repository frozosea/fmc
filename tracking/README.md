# Container tracking microservice

- [Шаги по запуску](#-)
- [Документация](#)
    * [Фундаментальные понятия о доменной области](#--1)
        + [Терминология](#-1)
    * [Stack](#stack-of-technology)
    * [Как это работает?](#--2)
        + [Базовые структуры данных](#--3)
            - [Container tracking response](#container-tracking-response)
            - [Bill number tracking response](#bill-number-tracking-response)
        + [GPRC ](#gprc)
        + [Services](#services)
            - [ScacService](#scacservice)
            - [TrackingService](#trackingservice)
                * [Поддержка линий](#--4)
        + [Как писать новый трекинг?](#--5)
        + [Верхний уровень ](#--6)
        + [Самописные утилиты](#--7)
            - [Requests (пакет для запросов)](#requests-)
            - [Datetime](#datetime)
    * [Объяснение алгоритма работы каждой линии](#--8)
        + [COSU (Cosco)](#cosu-cosco)
            - [Container tracker](#container-tracker)
        + [DNYG (DongYoung)](#dnyg-dongyoung)
            - [Container Tracker](#container-tracker-1)
            - [Bill tracker](#bill-tracker)
        + [FESO (Fesco)](#feso-fesco)
            - [Container tracker](#container-tracker-2)
            - [Bill tracker](#bill-tracker-1)
        + [HALU (Heung-A)](#halu-heung-a)
        + [HUXN (Hua-Xin)](#huxn-hua-xin)
            - [Container tracker](#container-tracker-3)
            - [Bill tracker](#bill-tracker-2)
        + [MAEU (Maersk)](#maeu-maersk)
            - [Container tracking](#container-tracking)
        + [MSCU (MSC shipping line)](#mscu-msc-shipping-line)
        + [ONEY (ONE Line)](#oney-one-line)
            - [Container tracker](#container-tracker-4)
        + [REEL (Reel shipping)](#reel-reel-shipping)
            - [Bill Tracker](#bill-tracker-3)
        + [SITC (SITC)](#sitc-sitc)
            - [Container tracker](#container-tracker-5)
            - [Bill tracker](#bill-tracker-4)
        + [SKLU (Sinokor)](#sklu-sinokor)
            - [Container tracker](#container-tracker-6)
            - [Bill tracker](#bill-tracker-5)
        + [ZHGU (ZhongGu56)](#zhgu-zhonggu56)
            - [Bill tracker](#bill-tracker-6)

# Шаги по запуску

1. create .env file and set these variables:

- POSTGRES_USERNAME
- POSTGRES_PASSWORD
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_DATABASE_NAME
- REDIS_URL
- REDIS_TTL: `in format 4h or 5m`
- TWO_CAPTCHA_API_KEY: [solve service](https://2captcha.com/enterpage)
- SITC_SERVICE_USERNAME
- SITC_SERVICE_PASSWORD
- SITC_SERVICE_BASIC_AUTH
- SITC_ACCESS_TOKEN

3. make migration: `make migrate-up`
4. Run


# Документация

## Фундаментальные понятия о доменной области
Данный микросервис по сути является одним из столпов нашего проекта, он отвечает за трекинг груза по морю. Если по простому,
то это тоже самое как слежение за посылкой. Все что нужно для этого, это знать номер контейнера или номер билла.
### Терминология
> Коносамент (Bill of Landing/Bill Number etc.) - документ, выдаваемый перевозчиком груза грузовладельцу. Удостоверяет право собственности на отгруженный товар. К нему привязываются номер(а) контейнеров.

> SCAC-code - это код перевозчика, который выдается перевозчиком, в нашем случае это коды перевозчиков морских линий.

> Линия - это компания предоставляющая регулярные маршруты по
торговым направлениям, позволяет перевезти контейнер из точки А в
точку В на своих судах 

> Мультимодальные перевозки - Транспортировка грузов по одному договору, но выполненная, по меньшей мере, двумя видами транспорта

> HC, HQ (High Cube) — стандартный сухогрузный морской контейнер с увеличенной высотой.

> HTC (Heavy Tested Container) — контейнер повышенной грузоподъемности.

> UT, OT, HT (Open Top, Hard Top). UT — контейнер со съемной крышей. Данный тип контейнеров подразделяется на два подтипа: с жесткой (стальной) и мягкой (брезентовой) крышей (Hard Top и Open Top соответственно). Предназначаются для транспортировки негабаритных грузов и облегчения их погрузки.

>ЕТА {от Estimated Time of Arrival) — Дата ориентировочного прихода судна/груза в порт назначения

>ETS / ETD (от Estimated Time of Sailing/Departure) — Дата ориентировочного убытия судна/груза из пункта погрузки,

>POD (от Port of Discharge) -Порт выгрузки/назначения

>POL (от Port of Loading) — Порт погрузки

>ТЕУ / TEU (от Twenty-feet Equivalent Unit) — условная единица измерения в контейнерных перевозках эквивалентная размерам ISO-контейнера длиной 20 футов. (6,1 м). Также используется сорокафутовый эквивалент (FEU), основанный на размерах 40- футового контейнера и равный 2 TEU.

>Трансшипмент — процесс перегрузки контейнеров в порту на пути следования с борта одного судна на борт другого судна. Является распространенным явлением в случае если у контейнерной линии отсутствует сервис напрямую связывающий порт погрузки с портом выгрузки, (пример: Груз следующий в Новороссийск выходит из порта Шанхай на океанском судне не имеющем судозахода в Новороссийск. В порту Стамбул груз перегружается на другое (фидерное) судно имеющее судозаход в Новороссийск в своем расписании. Порт Стамбул в данном случае является портом Трансшипмента.)

> Freight - (рус. Фрахт, нем. – Fracht). В праве – плата за перевозку груза либо использование судна в течение определенного времени, обусловленная договором или законом. Ее выплачивает судовладельцу фрахтователь или отправитель груза.

Терминологния так же доступна по [ссылке](https://tagankapremiumservice.ru/novosti/free-out.html).

## Stack

- Database: `postgres`
- Cache: `redis`
- Delivery: `grpc`
- Containerizing: `docker`

## Как это работает?
Всего есть два основных сервиса:
- TrackingService (сервис отвечающий за трекинг)
  - Tracking By Container Number
  - Tracking By Bill Number
- ScacService (Отдает подключенные к сервису линии)
  - Get By Container Number Tracking Shipping lines
  - Get By Bill Number Tracking Shipping lines
  
### Базовые структуры данных
У данного микросервиса есть всего 2 основные структуры данных, здесь они представлены сразу JSON.

#### Входная структура
На вход в обоих случаях поступает следующий JSON:
```json
{
    "number":"KE2313NSAV420", //Сам номер контейнера/коносамента
    "scac":"SITC" //Можно использовать код линии или AUTO
}
```

#### Container tracking response
```json
{
    "container": "FESU2219270",
    "containerSize": "20DC",
    "scac": "FESO",
    "infoAboutMoving": [
        {
            "time": "1684246320000",
            "operationName": "GATE OUT EMPTY FOR LOADING",
            "location": "MAGADAN MARINE COMMERCIAL PORT",
            "vessel": ""
        },
        {
            "time": "1688650140000",
            "operationName": "GATE IN EMPTY FROM CONSIGNEE",
            "location": "MAGADAN MARINE COMMERCIAL PORT",
            "vessel": ""
        }
    ]
}
```

#### Bill number tracking response
```json
{
    "billNo": "SNKO03B230500568",
    "scac": "SKLU",
    "infoAboutMoving": [
        {
            "time": "1685440800000",
            "operationName": "DEPARTURE",
            "location": "HIT HONGKONG TERMINAL (HIT4)",
            "vessel": "HEUNG-A XIAMEN / 2310N"
        },
        {
            "time": "1686211200000",
            "operationName": "ARRIVAL",
            "location": "VCT (VLADIVOSTOK CONTAINER TMNL)",
            "vessel": "HEUNG-A XIAMEN / 2310N"
        }
    ],
    "etaFinalDelivery": "1686211200000"
}
```
Информации о передвижении содержит даты, даты у нас является UNIX-timestamp в UTC формате.

### GPRC 
Как и во всем проекте используется GRPC, для конекта всех сервисов. Код самого стаба можно увидеть по [ссылке](https://github.com/frozosea/fmc-pb/tree/master/tracking)

### Services
Всего 2 основных сервиса в рамках данного микросервиса:
#### ScacService
По сути говоря это сервис который отвечает за выдачу scac-кодов и подключенных линий пользователю. 
На старте проекта подгружаются из файла `./scac.json` (в корне проекта). Затем выдаются пользователю. 

#### TrackingService

##### Поддержка линий

| Steamship Line | Works in Russia    | SCAC code | Container support  | Bill support       | Url                                                                             | Example container number | Example bill number                                           |
|----------------|--------------------|-----------|--------------------|--------------------|---------------------------------------------------------------------------------|--------------------------|---------------------------------------------------------------|
| APL            | :x:                | APLU      | :x:                | :x:                | [Link](https://www.apl.com/ebusiness/tracking)                                  | CMAU3018179              | :x:                                                           |
| CMA CGM        | :x:                | CMDU      | :x:                | :x:                | [Link](https://www.cma-cgm.com/ebusiness/tracking)                              | CMAU3018179              | :x:                                                           |
| Cosco          | :x:                | COSU      | :white_check_mark: | :x:                | [Link](https://elines.coscoshipping.com/ebusiness/cargoTracking)                | CSNU6829167              | :x:                                                           |
| HMM            | :x:                | HDMU      | :x:                | :x:                | [Link](https://www.hmm21.com/cms/business/ebiz/trackTrace/trackTrace/index.jsp) | CAIU7202031              | :x:                                                           |
| Maersk         | :x:                | MAEU      | :white_check_mark: | :x:                | [Link](https://www.maersk.com/tracking/)                                        | MRKU8746155              | :x:                                                           |
| MSC            | 50/50              | MSCU      | :white_check_mark: | :x:                | [Link](https://www.msc.com/track-a-shipment?agencyPath=mwi)                     | MEDU3170580              | :x:                                                           |
| ONE            | :x:                | ONEY      | :white_check_mark: | :x:                | [Link](https://ecomm.one-line.com/ecom/CUP_HOM_3301.do)                         | TCLU4418595              | :x:                                                           |
| ZIM            | :x:                | ZIMU      | :x:                | :x:                | [Link](https://www.zim.com/tools/track-a-shipment)                              | :x:                      | :x:                                                           |
| Fesco          | :white_check_mark: | FESO      | :white_check_mark: | :white_check_mark: | [Link](https://www.fesco.ru/ru/clients/tracking/)                               | FESU2219270              | XMVAK020N6001                                                 |
| Sinokor        | :white_check_mark: | SKLU      | :white_check_mark: | :white_check_mark: | [Link](http://ebiz.sinokor.co.kr/Tracking)                                      | TEMU2094051              | SNKO101220501450                                              |
| Heung-a        | :white_check_mark: | HALU      | :white_check_mark: | :white_check_mark: | [Link](http://ebiz.heung-a.com/Tracking)                                        | TCKU2902936              | HASLC03230602080                                              |
| SITC           | :white_check_mark: | SITC      | :white_check_mark: | :white_check_mark: | [Link](https://api.sitcline.com/sitcline/query/cargoTrack)                      | SITU9130070              | SITDLVK222G951, KE2313NSAV420, GS2311NSAV423, SITGSHVVT007061 |
| KMTC           | :x:                | KMTU      | :x:                | :x:                | [Link](https://www.ekmtc.com/index.html#/cargo-tracking)                        | TRHU3368865              | :x:                                                           |
| Zhonggu56      | :white_check_mark: | ZHGU      | :x:                | :white_check_mark: | [Link](http://dingcang.zhonggu56.com/views/CargoStrack/WaiMaoCargoStrack.jsp)   | :x:                      | ZGSHA0100001921                                               |
| DongYoung      | :white_check_mark: | DNYG      | :white_check_mark: | :white_check_mark: | [Link(go to cargo tracking)](https://ebiz.pcsline.co.kr/)                       | DYLU5112648              | PCSLBSHCC2300016                                              |
| ReelShipping   | :white_check_mark: | REEL      | :x:                | :white_check_mark: | [Link](https://tracking.reelshipping.com/Tracking//)                            | :x:                      | RB62CG23000870                                                |
| HuaXin         | :white_check_mark: | HUXN      | :white_check_mark: | :white_check_mark: | [Link](http://dc.hxlines.com:8099/HX_WeChat/HX_Dynamics#)                       | FTAU1753436              | HXCGTCVL23023117                                              |

В `pkg/tracking` внутри мы можем видеть имена каталогов, такие как `feso` и т.д. Это код SCAC судоходной линии, что
такое scac Вы можете прочитать на этом сайте [ссылка](https://www.safround.com/en/ocean-carrier-scac-code-list), и
список кодов scac: [ссылка](https://www.safrons.com/en/ocean-carrier-scac-code-list ). Каждый каталог имеет определенную
структуру, каждый пакет имеет:

- container_tracker(if exists)
- bill_tracker(if exists)
- request
- schema(if exists)
- <package_name>_test
- parser
- error file(if exists)
- some util files(if exists)
- test_data

Сам container_tracker/bill_tracker имеет конструктор, и он строго стандартизирован. Конструктор принимает в себя
структуру объявленную в `pkg/tracking/base.go`

```go
package main

type BaseConstructorArgumentsForTracker struct {
	Request            requests.IHttp
	UserAgentGenerator requests.IUserAgentGenerator
	Datetime           datetime.IDatetime
}

```

Внутреннее устройство самого трекера по сути неважно. Желательно следовать SRP, и разделять классы друг от друга, чтобы
не было супер классов.

### Как писать новый трекинг?

- Создать директорию по названию SCAC-кода
- Создать файл `container_tracker.go` и `bill_tracker.go` если трекер поддерживает оба типа трекинга
- Создать файл `request.go`
- Создать файл `parser.go`
- Написать основную логику 
- Покрыть код unit-тестами
- Внести линию в билдер (под ее SCAC кодом)
- Написать механизм автоматического снятия со слежения по расписания в `root project/schedule-tracking/pkg/tracking/check_arrived.go`
- Внести изменения в документацию
  - Внести принцип работы
  - Поменять таблицу поддержки
- Добавить линию в файл scac.json

### Верхний уровень 

Работает это просто, каждая структура трекинга контейнеров реализует интерфейс `tracking/base/Container tracker`, имеет
только один публичный метод и возвращает структуру ответа. В bill tracker схема такая же, но интерфейс и структура
ответов разные. Также есть основной трекер, у него есть поле `map[SCAC_CODE]TrackingInterface`, где ключом является код
SCAC, а значением - интерфейс отслеживания, это помогает быть независимым и выдавать все коды и трекеры без проблем.
Имеет автоматический код scac, это когда основной трекер в цикле for одновременно проверяет все линии и пытается
получить ответ, делает это конкурентно.

Верхнеуровневый трекер имеет service-структуру. Структура сервиса получает основной трекер, интерфейс кэширования (в
проекте используется redis) и репозиторий scac. Репозиторий SCAC - это репозиторий, который помещает номер и его
scac-код в базу данных (использует postgres), и когда пользователь пытается получить ответ по номеру, service-структура
проверяет номер в таблице и какой линии он пренадлежит, и, если существует, не будет бежать в цикле, а просто вытащит из
базы данных.

Для упрощения жизни здесь используются самописные пакета (либы), для запросов и для парсинга дат.

### Самописные утилиты

#### Requests (пакет для запросов)

Работает очень просто, чтобы не городить постоянные развесистые запросы стандартной либой, а так же чтобы не юзать эти
дурацкие структуры данных, и самое главное, чтобы была возможность мокать и очень легко тестировать.

Все строится на следующих данных:

```go
package requests

import "context"

type IHttp interface {
	Url(url string) IHttp
	Method(method string) IHttp
	Headers(headers map[string]string) IHttp
	Form(form map[string]string) IHttp
	MultipartForm(form map[string]string, fieldName, filename string, file []byte) IHttp
	Body(body []byte) IHttp
	Query(q map[string]string) IHttp
	Do(ctx context.Context) (*Response, error)
}

type Response struct {
	Body          []byte
	Status        int    `json:"status"`
	ContentType   string `json:"content-type"`
	ContentLength int64  `json:"content-length"`
}

```

У самой либы очень простой синтаксис:

```go
package main

import (
	"context"
	"golang_tracking/pkg/tracking/util/requests"
)

func main() {

	headers := map[string]string{
		"accept":          "application/json, text/plain, */*",
		"accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"content-type":    "application/json",
	}

	body := []byte(`{
            "data": "Click Here",
            "size": 36
        }`)

	const url = "https://example.com"

	var ctx = context.Background()

	r := requests.New()

	//example of using
	response, err := r.Url(url).Method("POST").Headers(headers).Body(body).Do(ctx)
}

```

Так же есть генератор User-Agent:

```go
package requests

type IUserAgentGenerator interface {
	Generate() string
}
```

Есть очень удобные MockUp структуры для тестирования, работает следующим образом:

```go
package main

import (
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"errors"
)

func main() {
	httpMockUp := requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
		if r.RUrl == "1" {
			return os.ReadFile("file 1")
		} else if r.RUrl == "2" {
			return os.ReadFile("file 2")
		}
		return nil, errors.New("url invalid error")
	})

	//return empty string 
	userAgentMockUp := requests.NewUserAgentGeneratorMockUp()
}
```

#### Datetime

По сути говоря интерфейс (так же чтобы удобно тестировать было) и реализация на основе сторонней либы. Синтаксис формата
дат полностью взят из питона [Link](https://pythonru.com/primery/kak-ispolzovat-modul-datetime-v-python)

```go
package datetime

import (
	"github.com/archsh/timefmt"
	"time"
)

type IDatetime interface {
	Strftime(t time.Time, format string) (string, error)
	Strptime(value string, format string) (time.Time, error)
}

func (d *Datetime) Strftime(t time.Time, format string) (string, error) {
	return timefmt.Strftime(t, format)
}

func (d *Datetime) Strptime(value string, format string) (time.Time, error) {
	return timefmt.Strptime(value, format)
}

type Datetime struct {
}

func NewDatetime() *Datetime {
	return &Datetime{}
}
```

## Объяснение алгоритма работы каждой линии

### COSU (Cosco)
#### Container tracker
Все строится на двух запросах, первый запрос идет в API и отдает информацию о передвижении. А второй получает ETA, затем
это все преобразуется в наши внутренние структуры данных и уходит на верх

Структура первого запроса (тот что информация о передвижении):

```go
package cosu

type Container struct {
	ContainerUuid       string   `json:"containerUuid"`
	ContainerNumber     string   `json:"containerNumber"`
	ContainerType       string   `json:"containerType"`
	GrossWeight         string   `json:"grossWeight"`
	PiecesNumber        int      `json:"piecesNumber"`
	Label               string   `json:"label"`
	SealNumber          string   `json:"sealNumber"`
	Location            string   `json:"location"`
	LocationDateTime    string   `json:"locationDateTime"`
	Transportation      string   `json:"transportation"`
	Flag                string   `json:"flag"`
	RailRef             string   `json:"railRef"`
	InlandMvId          string   `json:"inlandMvId"`
	ContainerLocation   string   `json:"containerLocation"`
	IsShow              bool     `json:"isShow"`
	PolEtd              string   `json:"polEtd"`
	PolAtd              string   `json:"polAtd"`
	PodEta              string   `json:"podEta"`
	PodAta              string   `json:"podAta"`
	TransportId         string   `json:"transportId"`
	Pol                 string   `json:"pol"`
	Pod                 string   `json:"pod"`
	HsCode              []string `json:"hsCode"`
	IsNorthAmericaRails bool     `json:"isNorthAmericaRails"`
}

type CircleStatus struct {
	Uuid                  string `json:"uuid"`
	ContainerNumber       string `json:"containerNumber"`
	ContainerNumberStatus string `json:"containerNumberStatus"`
	Location              string `json:"location"`
	TimeOfIssue           string `json:"timeOfIssue"`
	Transportation        string `json:"transportation"`
	PolEtd                string `json:"polEtd"`
	PolAtd                string `json:"polAtd"`
	PodEta                string `json:"podEta"`
	PodAta                string `json:"podAta"`
	TransportId           string `json:"transportId"`
	Pol                   string `json:"pol"`
	Pod                   string `json:"pod"`
}

type ContainerHistory struct {
	Uuid                  string `json:"uuid"`
	ContainerNumber       string `json:"containerNumber"`
	ContainerNumberStatus string `json:"containerNumberStatus"`
	Location              string `json:"location"`
	TimeOfIssue           string `json:"timeOfIssue"`
	Transportation        string `json:"transportation"`
	PolEtd                string `json:"polEtd"`
	PolAtd                string `json:"polAtd"`
	PodEta                string `json:"podEta"`
	PodAta                string `json:"podAta"`
	TransportId           string `json:"transportId"`
	Pol                   string `json:"pol"`
	Pod                   string `json:"pod"`
}

type ApiResponseSchema struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content struct {
			Containers []struct {
				Container             *Container          `json:"container"`
				ContainerCircleStatus []*CircleStatus     `json:"containerCircleStatus"`
				ContainerHistorys     []*ContainerHistory `json:"containerHistorys"`
			} `json:"containers"`
			NotFound string `json:"notFound"`
		} `json:"content"`
	} `json:"data"`
}
```

Пример ответа json:

```json
{
  "code": "200",
  "message": "",
  "data": {
    "content": {
      "containers": [
        {
          "container": {
            "containerUuid": "CSNU6829167",
            "containerNumber": "CSNU6829160",
            "containerType": "40HQ",
            "grossWeight": "21859.2KG",
            "piecesNumber": 1151,
            "label": "Empty Equipment Returned",
            "sealNumber": "K392297",
            "location": "United Waalhaven Terminals BV(Gate2,Rotterdam,Zuid-Holland,Netherlands",
            "locationDateTime": "2022-06-01 11:32:00",
            "transportation": "Truck",
            "flag": null,
            "railRef": null,
            "inlandMvId": null,
            "containerLocation": null,
            "isShow": false,
            "polEtd": null,
            "polAtd": null,
            "podEta": null,
            "podAta": null,
            "transportId": null,
            "pol": null,
            "pod": null,
            "hsCode": [],
            "isNorthAmericaRails": true
          },
          "containerCircleStatus": [
            {
              "uuid": "CSNU6829161",
              "containerNumber": "CSNU6829160",
              "containerNumberStatus": "Empty Equipment Returned",
              "location": "United Waalhaven Terminals BV(Gate2,Rotterdam,Zuid-Holland,Netherlands",
              "timeOfIssue": "2022-06-01 11:32",
              "transportation": "Truck",
              "polEtd": null,
              "polAtd": null,
              "podEta": null,
              "podAta": null,
              "transportId": null,
              "pol": null,
              "pod": null
            },
            {
              "uuid": "CSNU6829162",
              "containerNumber": "CSNU6829160",
              "containerNumberStatus": "Gate-out from Final Hub",
              "location": "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
              "timeOfIssue": "2022-05-30 06:53",
              "transportation": "Truck",
              "polEtd": null,
              "polAtd": null,
              "podEta": null,
              "podAta": null,
              "transportId": null,
              "pol": null,
              "pod": null
            }
          ],
          "containerHistorys": [
            {
              "uuid": "hisCSNU68291607",
              "containerNumber": "CSNU6829160",
              "containerNumberStatus": "Empty Equipment Returned",
              "location": "United Waalhaven Terminals BV(Gate2,Rotterdam,Zuid-Holland,Netherlands",
              "timeOfIssue": "2022-06-01 11:32",
              "transportation": "Truck",
              "polEtd": null,
              "polAtd": null,
              "podEta": null,
              "podAta": null,
              "transportId": null,
              "pol": null,
              "pod": null
            },
            {
              "uuid": "hisCSNU68291606",
              "containerNumber": "CSNU6829160",
              "containerNumberStatus": "Gate-out from Final Hub",
              "location": "Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands",
              "timeOfIssue": "2022-05-30 06:53",
              "transportation": "Truck",
              "polEtd": null,
              "polAtd": null,
              "podEta": null,
              "podAta": null,
              "transportId": null,
              "pol": null,
              "pod": null
            }
          ]
        }
      ],
      "notFound": ""
    }
  }
}
```

Второй запрос (Получение ETA)

```go
package cosu

type EtaApiResponseSchema struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content string `json:"content"`
	} `json:"data"`
}
```

JSON:

```json
{
  "code": "200",
  "message": "",
  "data": {
    "content": "2022-05-22 23:00"
  }
}
```

### DNYG (DongYoung)

#### Container Tracker

Сначала получает информацию о номере контейнера, и получает следующую структуру данных:

```json
{
  "dlt_resultBlList": [
    {
      "OUTPOD": "VLADIVOSTOK, RUSSIA <br> 2023-06-10 22:00", //Port of destination
      "OUTVDS": "DONGJIN VENUS 0178N (1st) <br> XIANG REN 2320W (2nd)", //vessels
      "OUTBKN": "NTKF2300162", //booking number
      "OUTETD": "202306042125", //estimated time of discharging
      "OUTBNO": "PCSLJNGVV23Q0047", //booking number 
      "OUTETA": "202306102200", //estimated time of arrival
      "OUTCNT": "DYLU5112648", //container number
      "OUTPOL": "NAGOYA, JAPAN <br> 2023-06-04 21:25", //port of loading
      "SIZETYPE": "4HDC"//container size
    }
  ]
}
```

Так же проверяет если массив пустой то возвращает ошибку, что контейнер данной линии не пренадлежит. Затем шлет запрос
на получение информации о передвижении и получает следующий JSON:

```json
{
  "dlt_resultMovementList": [
    {
      "OUTPOR": "JAPAN, NAGOYA",
      "OUTPVY": "RUSSIA, VLADIVOSTOK",
      "OUTDTD": "202305261707",
      "CMVDPT": "JNG",
      "OUTDES": "SHIPPER FULL = EXPORT EMPTY GATE OUT",
      "OUTPODK": "블라디보스토크 (러시아)",
      "OUTAREA": "NAGOYA, JAPAN",
      "CMVLCN": "JPNGO",
      "OUTLOC": "NUCT",
      "OUTPOLK": "나고야",
      "RNUM": 1,
      "OUTPOD": "RUSSIA, VLADIVOSTOK",
      "OUTCNO": "DYLU5112648",
      "OUTSEQ": 1,
      "OUTSTA": "SF",
      "OUTPOL": "JAPAN, NAGOYA",
      "OUTPORK": "나고야",
      "OUTPVYK": "블라디보스토크 (러시아)",
      "OUTDIK": "Terminal",
      "MOUTSEQ": 2,
      "OUTDEK": "SHIPPER FULL = EXPORT EMPTY GATE OUT",
      "OUTVSSVOY": "JDJV0178N"
    },
    {
      "OUTPOR": "JAPAN, NAGOYA",
      "OUTPVY": "RUSSIA, VLADIVOSTOK",
      "OUTDTD": "202306011547",
      "CMVDPT": "JNG",
      "OUTDES": "LOADING FULL = EXPORT FULL IN TERMINAL",
      "OUTPODK": "블라디보스토크 (러시아)",
      "OUTAREA": "NAGOYA, JAPAN",
      "CMVLCN": "JPNGO",
      "OUTLOC": "NUCT",
      "OUTPOLK": "나고야",
      "RNUM": 1,
      "OUTPOD": "RUSSIA, VLADIVOSTOK",
      "OUTCNO": "DYLU5112648",
      "OUTSEQ": 1,
      "OUTSTA": "LF",
      "OUTPOL": "JAPAN, NAGOYA",
      "OUTPORK": "나고야",
      "OUTPVYK": "블라디보스토크 (러시아)",
      "OUTDIK": "Terminal",
      "MOUTSEQ": 2,
      "OUTDEK": "LOADING FULL = EXPORT FULL IN TERMINAL",
      "OUTVSSVOY": "JDJV0178N"
    }
  ]
}
```

Затем просто парсит его и переводит во внутреннюю структуру данных.

#### Bill tracker

Особо ничего не меняется, кроме запроса в API линии. Она немного другая.

### FESO (Fesco)

#### Container tracker

Просто шлет запрос в API, получает кривой JSON:

```json
{
  "requestKey": "",
  "containers": [
    "{\"container\":\"FESU2219270\",\"time\":\"2022-06-11T04:51:35.626Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-06T16:00:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":43765,\"locName\":\"МАГИСТРАЛЬ\",\"locNameLatin\":\"MAGISTRAL\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-07T18:00:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":33378,\"locName\":\"Трансгарант\",\"locNameLatin\":\"ZAPSIBCONT\",\"locIdTo\":33378,\"locNameTo\":\"Трансгарант\",\"locNameLatinTo\":\"ZAPSIBCONT\",\"etCode\":\"T\",\"transportType\":\"Автомобиль\",\"vessel\":\"\",\"location\":{\"id\":43765,\"text\":\"МАГИСТРАЛЬ\",\"textLatin\":\"MAGISTRAL\",\"parentText\":\"Новосибирск\",\"parentTextLatin\":\"Novosibirsk\",\"country\":\"Россия\",\"countryLatin\":\"Russia\",\"ltCode\":\"T\",\"softshipCode\":\"MAGIST\",\"code\":null}}]}"
  ],
  "missing": []
}
```

Во время парсинга превращает его в нормальный JSON:

```json
{
  "container": "FESU2219270",
  "time": "2022-06-11T04:51:35.626Z",
  "containerCTCode": "20DC",
  "containerOwner": "COC",
  "latLng": null,
  "lastEvents": [
    {
      "time": "2022-06-06T16:00:00",
      "operation": "GATE-OUT",
      "operationName": "Вывоз с терминала",
      "operationNameLatin": "Gate out empty for loading",
      "locId": 43765,
      "locName": "МАГИСТРАЛЬ",
      "locNameLatin": "MAGISTRAL",
      "locIdTo": 33427,
      "locNameTo": "СКЛАД ГРУЗОВЛАДЕЛЬЦА",
      "locNameLatinTo": "sklad gruzovladel'сa",
      "etCode": null,
      "transportType": null
    },
    {
      "time": "2022-06-07T18:00:00",
      "operation": "GATE-IN",
      "operationName": "Прибытие на терминал",
      "operationNameLatin": "Gate in empty from consignee",
      "locId": 33378,
      "locName": "Трансгарант",
      "locNameLatin": "ZAPSIBCONT",
      "locIdTo": 33378,
      "locNameTo": "Трансгарант",
      "locNameLatinTo": "ZAPSIBCONT",
      "etCode": "T",
      "transportType": "Автомобиль",
      "vessel": "",
      "location": {
        "id": 43765,
        "text": "МАГИСТРАЛЬ",
        "textLatin": "MAGISTRAL",
        "parentText": "Новосибирск",
        "parentTextLatin": "Novosibirsk",
        "country": "Россия",
        "countryLatin": "Russia",
        "ltCode": "T",
        "softshipCode": "MAGIST",
        "code": null
      }
    }
  ]
}
```

#### Bill tracker

Тот же алгоритм, но ETA ставит в отдельное поле, а не кидает в информацию о передвижении.


### HALU (Heung-A)
Так-как Heung-a это дочерняя компания Sinokor (SKLU), то и системы у них абсолютно одинаковые по устройству, меняется лишь URL.

### HUXN (Hua-Xin)
Это новая китайская линия, мы единственные среди конкурентов имеем ее трекинг в своей системе, его выбивали очень долго. 

#### Container tracker

Сначала идет запрос в API, система отдает информацию о передвижении:

```json
{
  "status": "success",
  "listDynamics": [
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Load on vessel empty",
      "DYN_TIME": "2023/6/26 8:27:00",
      "DYN_TIME_NAME": "MON JUN 26TH, 2023 08:27",
      "PORT_FULL_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "PLACE_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23018S"
    },
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Gate in empty",
      "DYN_TIME": "2023/6/19 10:16:00",
      "DYN_TIME_NAME": "MON JUN 19TH, 2023 10:16",
      "PORT_FULL_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "PLACE_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23016N"
    },
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Gate out full",
      "DYN_TIME": "2023/6/16 11:05:00",
      "DYN_TIME_NAME": "FRI JUN 16TH, 2023 11:05",
      "PORT_FULL_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "PLACE_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23016N"
    },
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Discharge from vessel full",
      "DYN_TIME": "2023/6/12 3:35:00",
      "DYN_TIME_NAME": "MON JUN 12TH, 2023 03:35",
      "PORT_FULL_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "PLACE_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23016N"
    },
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Load on vessel full",
      "DYN_TIME": "2023/6/5 7:17:00",
      "DYN_TIME_NAME": "MON JUN 5TH, 2023 07:17",
      "PORT_FULL_NAME": "TAICANG, CHINA (CNTAC)",
      "PLACE_NAME": "Suzhou Modern Terminals Co., Ltd",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23016N"
    },
    {
      "CON_NO": "FTAU1753436",
      "TG_CODE": "22GP",
      "DYN_TYPE": "Gate in full",
      "DYN_TIME": "2023/5/29 15:53:23",
      "DYN_TIME_NAME": "MON MAY 29TH, 2023 15:53",
      "PORT_FULL_NAME": "TAICANG, CHINA (CNTAC)",
      "PLACE_NAME": "Suzhou Modern Terminals Co., Ltd",
      "VESSEL_VOYAGE": "CHANG RONG 8 / 23016N"
    }
  ]
}
```

Если массив пустой, то класс возвращает ошибку, что контейнер не принадлежит этой линии.

Данная линия не отдает ETA, а поэтому доставать ее приходится самостоятельно, в массиве видно, что система отдает 
`VESSEL_VOYAGE`, и здесь видно, что можно выдернуть `VOYAGE` и потом пропарсить расписание этого рейса примерно в рамках 
3 недель. Но тут есть одна проблема, так как наша система не имеет и не может иметь представления конечной точки контейнера, 
поэтому отдает весь маршрут.

Пример расписания JSON:
```json
{
  "status": "success",
  "listSchedules": [
    {
      "VESSEL": "HUA XIN 5",
      "VOYAGE": "23015N",
      "BARGE_VESSEL": null,
      "BARGE_VOYAGE": null,
      "BARGE_DISC_NAME": null,
      "PORT_LOAD_NAME": "TAICANG, CHINA (CNTAC)",
      "LOAD_PIER_NAME": "Suzhou Modern Terminals Co., Ltd",
      "LOAD_ETD": "2023-05-27",
      "LOAD_ETD_NAME": "SAT MAY 27TH, 2023",
      "PORT_DISC_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "DISC_PIER_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "DISC_ETA": "2023-05-31",
      "DISC_ETA_NAME": "WED MAY 31ST, 2023",
      "TRANSIT_TIME": "4"
    },
    {
      "VESSEL": "CHANG RONG 8",
      "VOYAGE": "23016N",
      "BARGE_VESSEL": null,
      "BARGE_VOYAGE": null,
      "BARGE_DISC_NAME": null,
      "PORT_LOAD_NAME": "TAICANG, CHINA (CNTAC)",
      "LOAD_PIER_NAME": "Suzhou Modern Terminals Co., Ltd",
      "LOAD_ETD": "2023-06-02",
      "LOAD_ETD_NAME": "FRI JUN 2ND, 2023",
      "PORT_DISC_NAME": "VLADIVOSTOK, RUSSIA (RUVVO)",
      "DISC_PIER_NAME": "VLADIVOSTOK SOLLERS TERMINAL",
      "DISC_ETA": "2023-06-09",
      "DISC_ETA_NAME": "FRI JUN 9TH, 2023",
      "TRANSIT_TIME": "7"
    }
  ]
}
```

#### Bill tracker
Тоже самое, что и container tracker. 


### MAEU (Maersk)

#### Container tracking

Шлет запрос в API, и получает следующий ответ:
```json
{
  "isContainerSearch": true,
  "origin": {
    "terminal": "Laem Chabang Terminal PORT D1",
    "geo_site": "9QYVWASPIUQBZ",
    "city": "Laem Chabang",
    "state": "",
    "country": "Thailand",
    "country_code": "TH",
    "geoid_city": "0NTCCB4ENSFX8",
    "site_type": "TERMINAL"
  },
  "destination": {
    "terminal": "",
    "geo_site": "Unknown",
    "city": "Spartanburg",
    "state": "South Carolina",
    "country": "United States",
    "country_code": "US",
    "geoid_city": "09L62FUTQWHHB",
    "site_type": ""
  },
  "containers": [
    {
      "container_num": "MSKU6874333",
      "container_size": "40",
      "container_type": "Dry",
      "iso_code": "42G0",
      "operator": "MAEU",
      "locations": [
        {
          "terminal": "Win Win Container Depot",
          "geo_site": "1GP8G188VBR2G",
          "city": "Laem Chabang",
          "state": "",
          "country": "Thailand",
          "country_code": "TH",
          "geoid_city": "0NTCCB4ENSFX8",
          "site_type": "DEPOT",
          "events": [
            {
              "activity": "GATE-OUT-EMPTY",
              "stempty": true,
              "actfor": "EXP",
              "vessel_name": "MSC SVEVA",
              "voyage_num": "204E",
              "vessel_num": "Z77",
              "actual_time": "2022-03-29T16:42:00.000",
              "rkem_move": "GATE-OUT",
              "is_cancelled": false,
              "is_current": false
            }
          ]
        },
        {
          "terminal": "Laem Chabang Terminal PORT D1",
          "geo_site": "9QYVWASPIUQBZ",
          "city": "Laem Chabang",
          "state": "",
          "country": "Thailand",
          "country_code": "TH",
          "geoid_city": "0NTCCB4ENSFX8",
          "site_type": "TERMINAL",
          "events": [
            {
              "activity": "GATE-IN",
              "stempty": false,
              "actfor": "EXP",
              "vessel_name": "MSC SVEVA",
              "voyage_num": "204E",
              "vessel_num": "Z77",
              "expected_time": "2022-04-16T14:00:00.000",
              "actual_time": "2022-03-30T10:02:00.000",
              "rkem_move": "GATE-IN",
              "is_cancelled": false,
              "is_current": false
            },
            {
              "activity": "LOAD",
              "stempty": false,
              "actfor": "",
              "vessel_name": "MSC SVEVA",
              "voyage_num": "204E",
              "vessel_num": "Z77",
              "expected_time": "2022-04-16T14:00:00.000",
              "actual_time": "2022-04-14T00:15:00.000",
              "rkem_move": "LOAD",
              "is_cancelled": false,
              "is_current": false
            }
          ]
        },
        {
          "terminal": "YANGSHAN SGH GUANDONG TERMINAL",
          "geo_site": "37O5HQ17XCL3X",
          "city": "Shanghai",
          "state": "Shanghai",
          "country": "China",
          "country_code": "CN",
          "geoid_city": "2IW9P6J7XAW72",
          "site_type": "TERMINAL",
          "events": [
            {
              "activity": "DISCHARG",
              "stempty": false,
              "actfor": "",
              "vessel_name": "MSC SVEVA",
              "voyage_num": "204E",
              "vessel_num": "Z77",
              "expected_time": "2022-04-25T08:00:00.000",
              "actual_time": "2022-04-25T12:49:00.000",
              "rkem_move": "DISCHARG",
              "is_cancelled": false,
              "is_current": false
            },
            {
              "activity": "LOAD",
              "stempty": false,
              "actfor": "",
              "vessel_name": "ZIM WILMINGTON",
              "voyage_num": "006E",
              "vessel_num": "U5T",
              "expected_time": "2022-05-02T18:30:00.000",
              "actual_time": "2022-05-02T01:11:00.000",
              "rkem_move": "LOAD",
              "is_cancelled": false,
              "is_current": false
            }
          ]
        },
        {
          "terminal": "Charleston Wando Welch terminal N59",
          "geo_site": "1ML38I7Q8BBKU",
          "city": "Charleston",
          "state": "South Carolina",
          "country": "United States",
          "country_code": "US",
          "geoid_city": "3RSB4DDP23AM7",
          "site_type": "TERMINAL",
          "events": [
            {
              "activity": "DISCHARG",
              "stempty": false,
              "actfor": "",
              "vessel_name": "ZIM WILMINGTON",
              "voyage_num": "006E",
              "vessel_num": "U5T",
              "expected_time": "2022-06-11T07:00:00.000",
              "actual_time": "2022-06-11T13:54:00.000",
              "rkem_move": "DISCHARG",
              "is_cancelled": false,
              "is_current": true
            },
            {
              "activity": "GATE-OUT",
              "stempty": false,
              "actfor": "IMP",
              "vessel_name": "",
              "voyage_num": "",
              "vessel_num": "",
              "expected_time": "2022-06-14T10:00:00.000",
              "is_current": false
            }
          ]
        }
      ],
      "eta_final_delivery": "2022-06-11T13:54:00.000",
      "latest": {
        "actual_time": "2022-06-11T13:54:00.000",
        "activity": "DISCHARG",
        "stempty": false,
        "actfor": "",
        "geo_site": "1ML38I7Q8BBKU",
        "city": "Charleston",
        "state": "South Carolina",
        "country": "United States",
        "country_code": "US"
      },
      "status": "IN-PROGRESS"
    }
  ]
}
```

Просто парсит его и отдает результат. Временно отключен, тк не дает отправить запрос с Российских IP.

### MSCU (MSC shipping line)

Шлет запрос в API, и получает следующий ответ:
```json
{
  "IsSuccess": true,
  "Data": {
    "TrackingType": "Container",
    "TrackingTitle": "CONTAINER NUMBER:",
    "TrackingNumber": "MEDU3170580",
    "CurrentDate": "12/06/2022",
    "PriceCalculationLabel": "* Price calculation date is indicative. Please company your local MSC office to verify this information.",
    "TrackingResultsLabel": "Tracking results provided by MSC on 12.06.2022 at 10:00 Central Europe Standard Time",
    "BillOfLadings": [
      {
        "BillOfLadingNumber": "",
        "NumberOfContainers": 1,
        "GeneralTrackingInfo": {
          "ShippedFrom": "",
          "ShippedTo": "",
          "PortOfLoad": "",
          "PortOfDischarge": "",
          "Transshipments": [],
          "PriceCalculationDate": "",
          "FinalPodEtaDate": "12/06/2022"
        },
        "ContainersInfo": [
          {
            "ContainerNumber": "MEDU3170580",
            "PodEtaDate": "",
            "ContainerType": "20' DRY VAN",
            "LatestMove": "CHONGQING, CN",
            "Events": [
              {
                "Order": 1,
                "Date": "10/06/2022",
                "Location": "CHONGQING, CN",
                "Description": "Export at barge yard",
                "Detail": [
                  "LADEN"
                ]
              },
              {
                "Order": 0,
                "Date": "09/06/2022",
                "Location": "CHONGQING, CN",
                "Description": "Empty to Shipper",
                "Detail": [
                  "EMPTY"
                ]
              }
            ]
          }
        ]
      }
    ]
  }
}
```

Просто парсит его и отдает результат.

### ONEY (ONE Line)

#### Container tracker

Сначала шлет запрос в API, дабы проверить существует ли номер в системе. 
Затем шлет запрос в API и получает `Custom Of Port Number` и `Bill number`. Example of JSON:
```json
{
  "TRANS_RESULT_KEY": "S",
  "Exception": "",
  "count": "1",
  "list": [
    {
      "maxRows": 0,
      "models": [],
      "weight": "",
      "copNo": "CCHI0C23914256",
      "blNo": "CHIU06661600",
      "eventDt": "2022-06-14 15:01",
      "cntrTpszCd": "D5",
      "copStsCd": "T",
      "piece": "",
      "sealNo": "",
      "placeNm": "DETROIT, MI, UNITED STATES",
      "yardNm": "UNIVERSAL INTERMODAL - DETROIT (DEPOT)",
      "statusNm": "Empty Container Release to Shipper",
      "bkgNo": "CHIU06661600",
      "poNo": "",
      "yardCd": "USDET31",
      "statusCd": "MOTYDO",
      "cntrNo": "GAOU6642924",
      "cntrTpszNm": "40'DRY HC.",
      "dspbkgNo": "",
      "socFlg": "N",
      "mvmtStsCd": "OP",
      "hashColumns": [
        [
          "no",
          {}
        ],
        [
          "bl_no",
          "CHIU06661600"
        ],
        [
          "cntr_no",
          "GAOU6642924"
        ],
        [
          "cntr_tpsz_nm",
          "40'DRY HC."
        ],
        [
          "ibflag",
          {}
        ],
        [
          "event_dt",
          "2022-06-14 15:01"
        ],
        [
          "lloyd_no",
          {}
        ],
        [
          "cop_dtl_seq",
          {}
        ],
        [
          "yard_cd",
          "USDET31"
        ],
        [
          "seal_no",
          ""
        ],
        [
          "vgm_rcv",
          {}
        ],
        [
          "dsp_bkg_no",
          ""
        ],
        [
          "place_cd",
          {}
        ],
        [
          "status_cd",
          "MOTYDO"
        ],
        [
          "cop_no",
          "CCHI0C23914256"
        ],
        [
          "nod_cd",
          {}
        ],
        [
          "act_tp_cd",
          {}
        ],
        [
          "po_no",
          ""
        ],
        [
          "soc_flg",
          "N"
        ],
        [
          "pagerows",
          {}
        ],
        [
          "cntr_tpsz_cd",
          "D5"
        ],
        [
          "enbl_flag",
          {}
        ],
        [
          "vvd",
          {}
        ],
        [
          "weight",
          ""
        ],
        [
          "yard_nm",
          "UNIVERSAL INTERMODAL - DETROIT (DEPOT)"
        ],
        [
          "cntr_flg",
          {}
        ],
        [
          "bkg_no",
          "CHIU06661600"
        ],
        [
          "vsl_eng_nm",
          {}
        ],
        [
          "vsl_cd",
          {}
        ],
        [
          "cop_sts_cd",
          "T"
        ],
        [
          "piece",
          ""
        ],
        [
          "mvmt_sts_cd",
          "OP"
        ],
        [
          "status_nm",
          "Empty Container Release to Shipper"
        ],
        [
          "skd_dir_cd",
          {}
        ],
        [
          "place_nm",
          "DETROIT, MI, UNITED STATES"
        ],
        [
          "skd_voy_no",
          {}
        ]
      ],
      "hashFields": []
    }
  ]
}
```
Потом используя эти данные получает JSON в котором есть информация о передвижении:
```json
{
  "TRANS_RESULT_KEY": "S",
  "Exception": "",
  "count": "13",
  "list": [
    {
      "maxRows": 0,
      "models": [],
      "vslCd": "",
      "no": "1",
      "copNo": "CSEL2303645419",
      "eventDt": "2022-03-24 11:19",
      "vslEngNm": "",
      "placeNm": "PUSAN, KOREA REPUBLIC OF",
      "skdVoyNo": "",
      "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
      "copDtlSeq": "1011",
      "skdDirCd": "",
      "actTpCd": "A",
      "statusNm": "Empty Container Release to Shipper",
      "statusCd": "MOTYDO",
      "nodCd": "KRPUS14",
      "vvd": "",
      "lloydNo": "",
      "hashColumns": [
        [
          "no",
          "1"
        ],
        [
          "bl_no",
          {}
        ],
        [
          "cntr_no",
          {}
        ],
        [
          "cntr_tpsz_nm",
          {}
        ],
        [
          "ibflag",
          {}
        ],
        [
          "event_dt",
          "2022-03-24 11:19"
        ],
        [
          "lloyd_no",
          ""
        ],
        [
          "cop_dtl_seq",
          "1011"
        ],
        [
          "yard_cd",
          {}
        ],
        [
          "seal_no",
          {}
        ],
        [
          "vgm_rcv",
          {}
        ],
        [
          "dsp_bkg_no",
          {}
        ],
        [
          "place_cd",
          {}
        ],
        [
          "status_cd",
          "MOTYDO"
        ],
        [
          "cop_no",
          "CSEL2303645419"
        ],
        [
          "nod_cd",
          "KRPUS14"
        ],
        [
          "act_tp_cd",
          "A"
        ],
        [
          "po_no",
          {}
        ],
        [
          "soc_flg",
          {}
        ],
        [
          "pagerows",
          {}
        ],
        [
          "cntr_tpsz_cd",
          {}
        ],
        [
          "enbl_flag",
          {}
        ],
        [
          "vvd",
          ""
        ],
        [
          "weight",
          {}
        ],
        [
          "yard_nm",
          "HANJIN BUSAN NEW PORT COMPANY(HJNC)"
        ],
        [
          "cntr_flg",
          {}
        ],
        [
          "bkg_no",
          {}
        ],
        [
          "vsl_eng_nm",
          ""
        ],
        [
          "vsl_cd",
          ""
        ],
        [
          "cop_sts_cd",
          {}
        ],
        [
          "piece",
          {}
        ],
        [
          "mvmt_sts_cd",
          {}
        ],
        [
          "status_nm",
          "Empty Container Release to Shipper"
        ],
        [
          "skd_dir_cd",
          ""
        ],
        [
          "place_nm",
          "PUSAN, KOREA REPUBLIC OF"
        ],
        [
          "skd_voy_no",
          ""
        ]
      ],
      "hashFields": []
    },
    {
      "maxRows": 0,
      "models": [],
      "vslCd": "",
      "no": "2",
      "copNo": "CSEL2303645419",
      "eventDt": "2022-04-05 10:40",
      "vslEngNm": "",
      "placeNm": "PUSAN, KOREA REPUBLIC OF",
      "skdVoyNo": "",
      "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
      "copDtlSeq": "1031",
      "skdDirCd": "",
      "actTpCd": "A",
      "statusNm": "Gate In to Outbound Terminal",
      "statusCd": "FOTMAD",
      "nodCd": "KRPUS14",
      "vvd": "",
      "lloydNo": "",
      "hashColumns": [],
      "hashFields": []
    },
    {
      "maxRows": 0,
      "models": [],
      "vslCd": "HHVT",
      "no": "3",
      "copNo": "CSEL2303645419",
      "eventDt": "2022-04-07 12:09",
      "vslEngNm": "HYUNDAI SINGAPORE",
      "placeNm": "PUSAN, KOREA REPUBLIC OF",
      "skdVoyNo": "0126",
      "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
      "copDtlSeq": "1032",
      "skdDirCd": "E",
      "actTpCd": "A",
      "statusNm": "Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading",
      "statusCd": "FLVMLO",
      "nodCd": "KRPUS14",
      "vvd": "HYUNDAI SINGAPORE 126E",
      "lloydNo": "9305685",
      "hashColumns": [],
      "hashFields": []
    }
  ]
}
```

Потом шлет запрос и получает размер контейнера. 

### REEL (Reel shipping)
#### Bill Tracker

Шлет запрос в систему линии, и получает HTML, парсит его и получает следующую структуру данных:
```go
package reel

type billMainInfo struct {
	POD             string
	ETD             time.Time
	lastEvent       *tracking.Event
	containerStatus *containerStatus
}

type containerStatus struct {
	Number    string
	Type      string
	EventDate time.Time
	Status    string
	Location  string
}
```

В этих структурах есть поле `Number`, это внутренний идентификатор номера контейнера привязанного к коносаменту внутри системы линии. 
Далее отправляется запрос (используя номер коносамента и внутренний номер) и получается еще один HTML, в нем уже парсится информация о передвижении. 
Если распарсить информацию о передвижении не получается, то просто берется `lastEvent` и отправляется в массив.  

### SITC (SITC)

#### Container tracker
Шлет запрос в API и парсит ответ, в идеале нужно добавить систему с парсингом расписания для получения ETA как в HuaXin(HUXN).
JSON:
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "list": [
      {
        "containerNo": "SITU9130070",
        "movementName": "出口装船",
        "movementCode": "VL",
        "movementNameEn": "LOADED ONTO VESSEL",
        "eventPort": "DALIAN",
        "eventDate": "2022-06-04 22:00:00",
        "vesselCode": "SITC CAGAYAN",
        "voyageNo": "2212S"
      },
      {
        "containerNo": "SITU9130070",
        "movementName": "客户提空箱",
        "movementCode": "OP",
        "movementNameEn": "OUTBOUND PICKUP",
        "eventPort": "DALIAN",
        "eventDate": "2022-05-27 10:59:00",
        "vesselCode": "SITC CAGAYAN",
        "voyageNo": "2212S"
      },
      {
        "containerNo": "SITU9130070",
        "movementName": "空箱入场",
        "movementCode": "MT",
        "movementNameEn": "EMPTY CONTAINER",
        "eventPort": "DALIAN",
        "eventDate": "2022-05-18 18:53:00",
        "vesselCode": "SITC MAKASSAR",
        "voyageNo": "2210S"
      }
    ]
  }
}
```

#### Bill tracker

Тут система в разы сложнее, для того чтобы получить ответ от сервера нужно решать капчи, а так же иметь валидный токен авторизации, 
писать собственную нейронку 
желания никакого не было, поэтому было принято решение использовать сторонний сервис решения капч, где люди за сущие 
копейки решают капчи, это повышает стабильность в разы. 

Решение капчи делается следующим образом:

- Генерация рандомной строки по специальной маске (используется как ID операции)
- Получение картинки капчи (просто получаем байты, никуда не пишем)
- Кидаем картинку в third party service
- Получаем ответ и отдаем решенную капчу с рандомной строкой

Так же в рамках микросервиса есть отдельная система, которая обновляет токен авторизации путем логина раз в Х времени,
там все так же решается капча вводятся данные пользователя и имитируется авторизация, затем хранилище токена (Singleton struct) обновляет его 
и в случае потребности отдает.

Затем решенная капча отправляется на сервер вместе с телом запроса на получения информации о номере коносамента. 
Система отдает следующий JSON:
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "list1": [
      {
        "blNo": "SITDLVK222G951",
        "polen": "DALIAN",
        "del": "海参崴",
        "delen": "VLADIVOSTOK COMMERCIAL PORT",
        "pol": "大连"
      }
    ],
    "containerNo": null,
    "blNo": "SITDLVK222G951",
    "list3": [
      {
        "rowNo": "1",
        "totalCount": "2",
        "containerNo": "SITU9130070",
        "sealNo": "SITW962404",
        "voyageNo": "STCG2212S",
        "cntrType": "40HC",
        "quantity": "333",
        "cntrSize": "66.76",
        "weight": "5028.5",
        "currentport": "SHANGHAI",
        "movementname": "出口装船",
        "movementnameen": "LOADED ONTO VESSEL"
      },
      {
        "rowNo": "2",
        "totalCount": "2",
        "containerNo": "UETU5790574",
        "sealNo": "SITW962403",
        "voyageNo": "STCG2212S",
        "cntrType": "40HC",
        "quantity": "432",
        "cntrSize": "61.72",
        "weight": "4760.5",
        "currentport": "SHANGHAI",
        "movementname": "出口装船",
        "movementnameen": "LOADED ONTO VESSEL"
      }
    ],
    "list2": [
      {
        "vesselName": "___",
        "voyageNo": "2212",
        "voyageLeg": "S",
        "portFrom": "CNDLC",
        "portFromName": "DALIAN",
        "portTo": "CNSHA",
        "portToName": "SHANGHAI",
        "eta": "2022-06-05 12:00:00",
        "etd": "2022-05-31 12:00:00",
        "atd": "2022-06-04 23:00",
        "agtb": "2022-06-13 23:00",
        "cctd": null,
        "cta": null,
        "ata": "2022-06-11 11:36",
        "ctd": null,
        "agta": "2022-06-11 01:00",
        "ccta": null,
        "atb": "2022-06-13 22:42",
        "agtd": "2022-06-04 23:00",
        "ctb": null,
        "cctb": null
      },
      {
        "vesselName": "HF FORTUNE",
        "voyageNo": "2229",
        "voyageLeg": "N",
        "portFrom": "CNSHA",
        "portFromName": "SHANGHAI",
        "portTo": "RUVVO",
        "portToName": "VLADIVOSTOK COMMERCIAL PORT",
        "eta": "2022-07-05 21:00:00",
        "etd": "2022-06-26 00:00:00",
        "atd": "2022-07-03 15:30",
        "agtb": null,
        "cctd": null,
        "cta": null,
        "ata": null,
        "ctd": null,
        "agta": null,
        "ccta": "2022-07-13 02:00",
        "atb": null,
        "agtd": "2022-07-03 15:30",
        "ctb": null,
        "cctb": "2022-07-13 04:00"
      }
    ],
    "movementcode": [
      {
        "movementStatus": "3"
      }
    ]
  }
}
```

Но здесь есть проблема, данная структура не имеет информации о передвижении, только будущие даты, поэтому моя система 
шлет еще один запрос, она выдергивает первый номер контейнера и на основе уже двух номеров (коносамента и контейнера) 
шлет доп запрос и получает информацию о передвижении именно этого контейнера, затем парсит ее и уже готовую структуру 
отдает пользователю. 

Пример информации о передвижении на основе номера контейнера и коносамента:
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "list": [
      {
        "movementid": "2c2881b880f3e8e2018107f45a6862be",
        "blNo": "SITDLVK222G951",
        "containerNo": "SITU9130070",
        "eventdate": "2022-05-27",
        "portname": "dalian",
        "movementcode": "OP",
        "movementname": "客户提空箱",
        "movementnameen": "outbound pickup"
      },
      {
        "movementid": "E0D111F1B84D08C8E0537F00000108C8",
        "blNo": "SITDLVK222G951",
        "containerNo": "SITU9130070",
        "eventdate": "2022-06-04",
        "portname": "dalian",
        "movementcode": "VL",
        "movementname": "出口装船",
        "movementnameen": "loaded onto vessel"
      }
    ]
  }
}
```

### SKLU (Sinokor)

В системе этой линии есть жесткое разделение коносамента и контейнера, например только лишь по номеру контейнера 
невозможно получить информацию о передвижении, только ETA и конечный порт выгрузки, но и с ним не все так просто, 
система отдает порт выгрузки в рамках спецификации ООН о кодах наименования городов, поэтому это имеет очень нечеловеческий вид,
например: `RUVVO`, первые две буквы это код страны, остальные 3 это наименование города. Отдавать такое пользователю 
не лучшая идея, поэтому в рамках микросервиса есть сервис, который отвечает за преобразование такого рода кодов в человеческий вид
, по сути говоря была выкачана база этих кодов с сайта ООН, и эта информация хранится в базе данных, и просто вытаскивается 
оттуда по соответствию с кодом. Например если отправить туда `RUVVO`, то в итоге сервис отдаст `Port Of Vladivostok`.

Пример ответа от сервера если слать запрос лишь по номеру контейнера:
```json
[
  {
    "BKNO": "SNKO101220802074",
    "CNTR": "20'x1",
    "POR": null,
    "POL": "VNSGN",
    "POD": "RUVYP",
    "DLV": "RUVYP",
    "VSL": "SWCG",
    "VYG": "2205N",
    "ETD": "2022-09-08",
    "ETA": "2022-11-06"
  },
  {
    "BKNO": "SNKO011220602534",
    "CNTR": "20'x1",
    "POR": null,
    "POL": "KRPUS",
    "POD": "VNSGN",
    "DLV": "VNSGN",
    "VSL": "SKOR",
    "VYG": "2209S",
    "ETD": "2022-08-23",
    "ETA": "2022-08-29"
  }
]
```

#### Container tracker

Из-за того что система отдает информацию о передвижении только имея номера коносамента на руках, то сперва
идет запрос в API для получения ETA и номера коносамента, затем этот номер коноса и контейнера отправляется по другому 
юрл и на выходе есть HTML, в котором есть нужная информация, парсер вытаскивает эту информацию используя несложный 
алгоритм обхождения HTML-дерева как массива со STEP (чтобы лучше понимать надо увидеть структуру хтмл на сайте, там жесть).

#### Bill tracker

Тут система попроще, все достается за один запрос, просто идет парсинг HTML страницы. Методы парсинга тоже лучше увидеть 
, чем объяснять.

### ZHGU (ZhongGu56)
Тоже китайская линия с очень кривым софтом, даже в рамках собственной системы иногда отдают неправильную информацию о 
прибытии груза.
#### Bill tracker
Все опять до безумия просто, идет запрос в API, и затем система его парсит. Example of JSON:
```json
{
  "backCode": "200",
  "backMessage": "Request succeeded!",
  "object": [
    {
      "createdByUser": null,
      "createdOffice": null,
      "createdDtmLoc": null,
      "createdTimeZone": null,
      "updatedByUser": null,
      "updatedOffice": null,
      "updatedDtmLoc": null,
      "updatedTimeZone": null,
      "recordVersion": null,
      "rowStatus": 2,
      "principalGroupCode": null,
      "blNo": "ZGSHA0100001921",
      "tripNumber": "1",
      "lineType": "1",
      "vesselName": "ZHONG GU BO HAI",
      "voyage": "22004N",
      "portFrom": "CNSHA",
      "portFromName": "SHANGHAI",
      "portTo": "RUVYP",
      "portToName": "VOSTOCHNY",
      "etd": "2022-07-22",
      "atd": "2022-07-25",
      "eta": "2022-07-31",
      "ata": "",
      "departureTime": "ETD:2022-07-22\nATD:2022-07-25",
      "arrivalTime": "ETA:2022-07-31\nATA:"
    }
  ]
}
```

Информацию о передвижении не достать, поэтому массив отдается пустым. 
