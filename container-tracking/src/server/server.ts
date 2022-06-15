import {Server, ServerCredentials} from "@grpc/grpc-js";
import {TrackingByContainerNumberService} from "./proto/server_grpc_pb";
import {GRPC_PORT} from "../../config.json"
import {grpcService} from "../containers";

export const server = new Server();
// @ts-ignore
server.addService(TrackingByContainerNumberService, grpcService)

export default function startServer() {
    server.bindAsync(`localhost:${GRPC_PORT}`, ServerCredentials.createInsecure(), (error, port) => {
        console.log("SERVER WAS STARTED")
        server.start()
    })
}

