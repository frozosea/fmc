import {TrackingContainerResponse} from "../../src/types";
import {TimeInspector} from "../../src/trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";

const assert = require("assert");

describe("tracking utils test", () => {
    it("time inspector test (this class check if container found on shipping line but information so old, like 4-6 month)", () => {
        const timeInspector = new TimeInspector()
        const oldInfoMoch = (time: number): TrackingContainerResponse => {
            return {
                container: "fesoContainer", scac: "FESO", containerSize: "container size", infoAboutMoving: [
                    {time: time, operationName: "operation1", location: "loc1", vessel: "vessel1"},
                ]
            }
        }
        const newInfoMoch = (time: number):TrackingContainerResponse => {
            return {
                container: "cosuContainer", scac: "COSU", containerSize: "container size", infoAboutMoving: [
                    {time: time, operationName: "operation1", location: "loc1", vessel: "vessel1"}
                ]
            }
        }

        const getOldTime = (): number[] => {
            return [1623733861000, 1644901861000, 1581743461000, 1550207461000]
        }

        const getNewTime = (): number[] => {
            return [1652591461000, 1655269861000, 1649999461000, 1652159461000, 1651468261000]
        }
        for (let time of getOldTime()){
            assert.strictEqual(timeInspector.inspectTime(oldInfoMoch(time)),false)
        }
        for (let newTime of getNewTime()){
            assert.strictEqual(timeInspector.inspectTime(newInfoMoch(newTime)),true)
        }
    })
})