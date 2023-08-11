import constants from "@/constants";
import Axios from "axios";

const { API_URL } = constants;
const DEFAULT_SEGMENT = {
    description: "",
    rolloutPercent: 50
};

export default class SegmentService {
    static get DEFAULT_SEGMENT() {
        return DEFAULT_SEGMENT;
    }

    static postSegmentPromise(flagId, segmentDescription, segmentRolloutPercent) {
        return Axios.post(
            `${API_URL}/flags/${flagId}/segments`,
            {
                description: segmentDescription,
                rolloutPercent: segmentRolloutPercent,
            }
        );
    }

    static deleteSegmentPromise(flagId, segmentId) {
        return Axios.delete(
            `${API_URL}/flags/${flagId}/segments/${segmentId}`
        );
    }

    static setDistributionValue(value, segment, variant) {
        const percent = parseInt(value, 10);
        console.log({ segment, variant, percent });
        const distribution = segment.distributions.find(d => d.variantID === variant.id);
        if (distribution) {
            console.log("updating")
            distribution.percent = percent;
        } else {
            console.log("adding")
            const newDistribution = {
                variantID: variant.id,
                variantKey: variant.key,
                percent,
            };
            segment.distributions.push(newDistribution);
        }
    }

    static putDistributionsPromise(flagId, segmentId, distributions) {
        return Axios.put(
            `${API_URL}/flags/${flagId}/segments/${segmentId}/distributions`,
            { distributions },
        );
    }
}