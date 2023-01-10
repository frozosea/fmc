# Container tracking microservice

Steps to run this project:

1. create .env file and set this variables:
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
- ALTS_KEY

3. make migration: `make migrate-up`
4. Run

### Tracking  support
| Steamship Line | SCAC code | Container support  | Bill support       | Url                                                                             | Example number  |
|----------------|-----------|--------------------|--------------------|---------------------------------------------------------------------------------|-----------------|
| APL            | APLU      | :x:                | :x:                | [Link](https://www.apl.com/ebusiness/tracking)                                  | CMAU3018179     |
| CMA CGM        | CMDU      | :x:                | :x:                | [Link](https://www.cma-cgm.com/ebusiness/tracking)                              | CMAU3018179     |
| Cosco          | COSU      | :white_check_mark: | :x:                | [Link](https://elines.coscoshipping.com/ebusiness/cargoTracking)                |                 |
| HMM            | HDMU      | :x:                | :x:                | [Link](https://www.hmm21.com/cms/business/ebiz/trackTrace/trackTrace/index.jsp) | CAIU7202031     |
| Maersk         | MAEU      | :white_check_mark: | :x:                | [Link](https://www.maersk.com/tracking/)                                        |                 |
| MSC            | MSCU      | :white_check_mark: | :x:                | [Link](https://www.msc.com/track-a-shipment?agencyPath=mwi)                     |                 |
| ONE            | ONEY      | :white_check_mark: | :x:                | [Link](https://ecomm.one-line.com/ecom/CUP_HOM_3301.do)                         |                 |
| ZIM            | ZIMU      | :x:                | :x:                | [Link](https://www.zim.com/tools/track-a-shipment)                              |                 |
| Fesco          | FESO      | :white_check_mark: | :white_check_mark: | [Link](https://www.fesco.ru/ru/clients/tracking/)                               |                 |
| Sinokor        | SKLU      | :white_check_mark: | :white_check_mark: | [Link](http://ebiz.sinokor.co.kr/Tracking)                                      |                 |
| Heung-a        | HALU      | :white_check_mark: | :white_check_mark: | [Link](http://ebiz.heung-a.com/Tracking)                                        |                 |
| SITC           | SITC      | :white_check_mark: | :white_check_mark: | [Link](https://api.sitcline.com/sitcline/query/cargoTrack)                      | SITU9130070     |
| KMTC           | KMTU      | :x:                | :x:                | [Link](https://www.ekmtc.com/index.html#/cargo-tracking)                        | TRHU3368865     |
| Zhonggu56      | ZHGU      | :x:                | :white_check_mark: | [Link](http://dingcang.zhonggu56.com/views/CargoStrack/WaiMaoCargoStrack.jsp)   | ZGSHA0100001921 |



# Documentation of code

## Stack of technology

- Database: `postgres`
- Cache: `redis`
- Delivery: `grpc`
- Containerizing: `docker`



## How it works?
###General presentation

In `pkg/tracking` inside we can see names of directory like `feso` etc, this is SCAC code of shipping line, what is scac
You can read in this site [link](https://scaccodelookup.com/), and list of scac codes: [link](https://www.safround.com/en/ocean-carrier-scac-code-list).
Every directory has defined structure, every package has:
- container_tracker
- bill_tracker(if exists)
- request
- schema
- <package_name>_test
- parser
- error file(if exists)
- some util files


Works it simple, every container tracker struct realizes `tracking/base/IContainerTracker` interface, has only one public method
and returns response struct. In bill tracker scheme is the same, but interface and response structs are different.
Also has main tracker, it got map struct where key is SCAC code and value is interface of tracking, this helps be independent and give
all codes and trackers without problems. Has AUTO scac code, this is when main tracker in for loop concurrently check all lines
and try to get response.

Upper then tracker has service struct. Service struct get main tracker, cache interface (in project uses redis), and scac repository.
SCAC repository is repo which put number and its scac code in database(uses postgres), and when user trying gets response
by number service struct will check number in table, and if exists will ain't uses for loop.


