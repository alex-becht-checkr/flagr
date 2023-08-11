import Axios from "axios";
import constants from "@/constants";

const { API_URL } = constants;

const DEFAULT_VARIANT = {
    key: ""
};

export default class VariantService {
    static get DEFAULT_VARIANT() {
        return DEFAULT_VARIANT;
    }

    static deleteVariantPromise(flagId, variantId) {
        return Axios.delete(
            `${API_URL}/flags/${flagId}/variants/${variantId}`
        )
    }

    static postVariantPromise(flagId, newVariant) {
        return Axios.post(
            `${API_URL}/flags/${flagId}/variants`,
            newVariant
        );
    }
}