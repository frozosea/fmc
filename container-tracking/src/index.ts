import startServer from "./server/server";
import {config} from "dotenv"
import {AppDataSource} from "./db/data-source";
import {mainTrackingByBillNumberForRussia} from "./init";

function main() {
    config();
    AppDataSource.initialize()
        .then(async (_) => {
        })
        .catch((error) => console.log("Error: ", error));
    // startServer();
    (async() =>{
        console.log(await mainTrackingByBillNumberForRussia.trackByBillNumber({
            number :"SNKO101220501450",
            scac:"SKLU"
        }))
    })();
}
main();

