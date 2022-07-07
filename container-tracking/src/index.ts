import startServer from "./server/server";
import {config} from "dotenv"
import {AppDataSource} from "./db/data-source";
import {trackBillNoByServer} from "./server/clients";

function main() {
    config();
    AppDataSource.initialize()
        .then(async (_) => {
        })
        .catch((error) => console.log("Error: ", error))
    startServer();

}

main();

