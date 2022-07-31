import {Server, ServerCredentials} from "@grpc/grpc-js";
import {TrackingByBillNumberService, TrackingByContainerNumberService} from "./proto/tracking_grpc_pb";
import {trackingByBillNumberService, trackingByContainerNumberService} from "../containers";
import {config} from "dotenv";
import {trackBillNoByServer} from "./clients";
import {AppDataSource} from "../db/data-source";

export const server = new Server();
// @ts-ignore
server.addService(TrackingByContainerNumberService, trackingByContainerNumberService)
// @ts-ignore
server.addService(TrackingByBillNumberService, trackingByBillNumberService)
export default function startServer() {
    config()
    AppDataSource.initialize()
        .then(async (source) => {
        })
        .catch((error) => console.log("Error of init db: ", error))
    server.bindAsync(`0.0.0.0:${process.env.GRPC_PORT}`, ServerCredentials.createInsecure(), (error, port) => {
        server.start();
        console.log("SERVER WAS STARTED")
    })
}

