import {TrackingByContainerNumberClient} from "./proto/server_grpc_pb";
import {GRPC_PORT} from "../../config.json";
import {credentials} from "@grpc/grpc-js";
import {Country, Request, Response, Scac} from "./proto/server_pb";
import {COUNTRY_TYPE, SCAC_TYPE} from "../types";
import {TrackingServiceConverter} from "./services/trackingService";

export const client = new TrackingByContainerNumberClient(
    `localhost:${GRPC_PORT}`,
    credentials.createInsecure(),
);

export function trackContainerByServer(container: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<Response> {
    return new Promise<Response>((resolve, reject) => {
        const request = new Request();
        request.setContainer(container);
        request.setScac(TrackingServiceConverter.convertScacIntoEnum(scac))
        request.setCountry(TrackingServiceConverter.convertCountryTypeIntoEnum(country))
        client.track(request, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        });
    });
}