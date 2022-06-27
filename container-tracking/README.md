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

| Steamship Line   |SCAC | Supported by container number    | Supproted by bill number | Container Tracking Website | Container number example |
| -------------    | :---------: | :-------------: | :---------------:|| :---------------: | 
| APL              | APLU | No | No |[Link](https://www.apl.com/ebusiness/tracking)  | CMAU3018179 |
| CMA CGM          | CMDU | No | No |[Link](https://www.cma-cgm.com/ebusiness/tracking)  | CMAU3018179 |
| Cosco            | COSU | Yes| No |[Link](https://elines.coscoshipping.com/ebusiness/cargoTracking)  |
| Hyundai Merchant Marine (HMM)| HDMU | No| No| [Link](https://www.hmm21.com/cms/business/ebiz/trackTrace/trackTrace/index.jsp) | CAIU7202031 |
| Maersk           | MAEU | Yes| |No|[Link](https://www.maersk.com/tracking/) |
| Mediterranean Shipping Company (MSC) |MSCU |Yes| No [Link](https://www.msc.com/track-a-shipment?agencyPath=mwi) | |
| ONE Line         | ONEY |Yes | No|[Link](https://ecomm.one-line.com/ecom/CUP_HOM_3301.do)  | |
| Zim Integrated Shipping Services (ZIM) |ZIMU |No| No|[Link](https://www.zim.com/tools/track-a-shipment)| GLDU5117768 |
|Fesco Shipping    | FESO | Yes| Yes|[Link](https://www.fesco.ru/ru/clients/tracking/) | |
|Sinokor Merchant Marine|SKLU|Yes| Yes| [Link](http://ebiz.sinokor.co.kr/Tracking)| |
|Heung-a Merchant Marine|HALU|Yes| Yes| [Link](http://ebiz.heung-a.com/Tracking)| | |
|Korea Marine Transport Co|KMTU|Yes(should refactor)|No|[Link](https://www.ekmtc.com/index.html#/cargo-tracking)| |

