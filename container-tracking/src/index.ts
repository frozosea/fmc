import startServer from "./server/server";
import {trackContainerByServer} from "./server/client";


startServer();
(async ()=>{
    let result = await trackContainerByServer("CSNU6829160", "COSU", "OTHER");
    console.log(result.toObject());
})()
// AppDataSource.initialize()
//     .then(async (source) => {
//         startServer()
//         let result = await trackContainerByServer("SKLU1204356", "AUTO", "OTHER");
//         console.log(result.toObject());
//     })
//     .catch((error) => console.log("Error: ", error))