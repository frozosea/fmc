import {ITrackingByBillNumberResponse, TrackingContainerResponse} from "./types";

const log4js = require("log4js");


export interface ILogger {
    ExceptionLog(logString: string): void;

    DebugLog(logString: string): void;

    InfoLog(logString: string): void;

    WarningLog(logString: string): void;

    FatalLog(logString: string): void
}

export class Logger implements ILogger {
    public constructor() {
        this.ConfigureLogger();
    }

    protected ConfigureLogger() {
        log4js.configure({
            appenders: {
                exception: {type: "file", filename: "logs/exception.log"},
                debug: {type: "file", filename: "logs/debug.log"},
                info: {type: "file", filename: "logs/info.log"},
                warning: {type: "file", filename: "logs/warning.log"},
                fatal: {type: "file", filename: "logs/fatal.log"}
            },
            categories: {
                default: {appenders: ["exception", "debug", "info", "warning", "fatal"], level: "all"},
                exception: {appenders: ["exception"], level: "error"},
                debug: {appenders: ["debug"], level: "debug"},
                info: {appenders: ["info"], level: "info"},
                warning: {appenders: ["warning"], level: "warn"},
                fatal: {appenders: ["fatal"], level: "fatal"}
            }
        });

    }

    protected log(loggerName: string, level: string, message: any) {
        const logger = log4js.getLogger(loggerName)
        logger.log(level, message)

    }

    public ExceptionLog(logString) {
        return this.log("exception", "error", logString)

    }

    public DebugLog(logString: string) {
        return this.log("debug", "debug", logString)

    }

    public InfoLog(logString: string) {
        return this.log("info", "info", logString)
    }

    public WarningLog(logString: string) {
        return this.log("warning", "warning", logString)
    }

    public FatalLog(logString: string) {
        return this.log("fatal", "fatal", logString)
    }
}

export interface IServiceLogger {
    containerSuccessLog(result: TrackingContainerResponse | ITrackingByBillNumberResponse): void

    containerNotFoundLog(container: string): void
}

export class ServiceLogger extends Logger implements IServiceLogger {
    protected ConfigureLogger() {
        log4js.configure({
            appenders: {
                containerFound: {type: "file", filename: "logs/containerFound.log"},
                containerNotFound: {type: "file", filename: "logs/containerNotFoundLog.log"},
            },
            categories: {
                default: {appenders: ["containerFound", "containerNotFound"], level: "all"},
                containerFound: {appenders: ["containerFound"], level: "info"},
                containerNotFound: {appenders: ["containerNotFound"], level: "warn"},
            }
        });

    }

    public containerSuccessLog(result: TrackingContainerResponse): void {
        return this.log("containerFound", "info", `container: ${result.container} scac: ${result.scac} result: ${JSON.stringify(result)}`)
    }

    public containerNotFoundLog(container: string): void {
        return this.log("containerNotFound", "warn", `container ${container}`)
    }
}