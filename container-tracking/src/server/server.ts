import {Server, ServerCredentials} from "@grpc/grpc-js";
import {TrackingByBillNumberService, TrackingByContainerNumberService} from "./proto/server_grpc_pb";
import {trackingByBillNumberService, trackingByContainerNumberService} from "../containers";
import {config} from "dotenv";

export const server = new Server();
// @ts-ignore
server.addService(TrackingByContainerNumberService, trackingByContainerNumberService)
// @ts-ignore
server.addService(TrackingByBillNumberService, trackingByBillNumberService)
export default function startServer() {
    config()
    server.bindAsync(`localhost:${process.env.GRPC_PORT}`, ServerCredentials.createInsecure(), (error, port) => {
        server.start()
        console.log("SERVER WAS STARTED")
    })
}

