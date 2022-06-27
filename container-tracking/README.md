# Container tracking microservice

Steps to run this project:

1. create .env file and set this variables:
- REQUEST_TIMEOUT_MS
- WAIT_SELECTOR_TIMEOUT
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_DATABASE
- REDIS_URL
- CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS
- GRPC_PORT
3. Setup database settings inside `data-source.ts` file
4. Run `npm start` command

#Documentation of code
project structure looks like:
```
src-|
    |
    |
    |---trackTrace---|
                     |
                     |---trackByContainerNumber
                     |
                     |---trackingByBillNumber
```
### Tracking  support

| Steamship Line                              | Scac | tracking by container number support    | tracking by bill number support | tracking link                                                                   | example container | example bill number |
|---------------------------------------------|------|-----------------------------------------|---------------------------------|---------------------------------------------------------------------------------|-------------------|---------------------|
| APL                                         | APLU | :x:                                     | :x:                             | [Link](https://www.apl.com/ebusiness/tracking)                                  | CMAU3018179       |                     |
| CMA CGM                                     | CMDU | :x:                                     | :x:                             | [Link](https://www.cma-cgm.com/ebusiness/tracking)                              | CMAU3018179       |                     |
| Cosco                                       | COSU | :white_check_mark:                      | :x:                             | [Link](https://elines.coscoshipping.com/ebusiness/cargoTracking)                |                   |                     |
| Hyundai Merchant Marine (HMM)               | HDMU | :x:                                     | :x:                             | [Link](https://www.hmm21.com/cms/business/ebiz/trackTrace/trackTrace/index.jsp) |                   |                     |
| Maersk/Sealand Maersk                       | MAEU | :white_check_mark:                      | :x:                             | [Link](https://www.maersk.com/tracking/)                                        |                   |                     |
| Mediterranean Shipping Company (MSC)        | MSCU | :white_check_mark:                      | :x:                             | [Link](https://www.msc.com/track-a-shipment?agencyPath=mwi)                     |                   |                     |
| ONE Line                                    | ONEY | :white_check_mark:                      | :x:                             | [Link](https://ecomm.one-line.com/ecom/CUP_HOM_3301.do)                         |                   |                     |
| Zim Integrated Shipping Services (ZIM)      | ZIMU | :x:                                     | :x:                             | [Link](https://www.zim.com/tools/track-a-shipment)                              |                   |                     |
| Fesco Shipping Co.                          | FESO | :white_check_mark:                      | :white_check_mark:              | [Link](https://www.fesco.ru/ru/clients/tracking/)                               |                   |                     |
| Sinokor Merchant Marine                     | SKLU | :white_check_mark:                      | :white_check_mark:              | [Link](http://ebiz.sinokor.co.kr/Tracking)                                      |                   |                     |
| Heung-a Merchant Marine                     | HALU | :white_check_mark:                      | :white_check_mark:              | [Link](http://ebiz.heung-a.com/Tracking)                                        |                   |                     |
| SITC International Holdings Company Limited | SITC | :white_check_mark:                      | :x:                             | [Link](https://api.sitcline.com/sitcline/query/cargoTrack)                      | SITU9130070       | SITDLVK222G951      |
| Korea Marine Transport Co                   | KMTU | :white_check_mark: (should be refactor) | :x:                             | [Link](https://www.ekmtc.com/index.html#/cargo-tracking)                        |                   |                     |

