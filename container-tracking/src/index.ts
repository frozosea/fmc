import startServer from "./server/server";
import {config} from "dotenv"

function main() {
    config();
    startServer();
}

main();

