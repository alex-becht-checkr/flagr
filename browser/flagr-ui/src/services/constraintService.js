import Axios from "axios";
import constants from "@/constants";

const { API_URL } = constants;
const DEFAULT_CONSTRAINT = {
    operator: "EQ",
    property: "",
    value: ""
};

export default class ConstraintService {
  static get DEFAULT_CONSTRAINT() {
      return DEFAULT_CONSTRAINT;
  }

  static get listConstraints() {
      return ["IN", "NOTIN"];
  }

  static formatConstraint(constraint) {
      if (constraint.operator === "IN" || constraint.operator === "NIN") {
          constraint.value = constraint.value.split(",").map(this.quoteString).join(",");
          if (!constraint.value.match(/\[/)) {
              constraint.value = `[${constraint.value}`;
          }
          if (!constraint.value.match(/\]/)) {
              constraint.value = `${constraint.value}]`;
          }
      } else {
          constraint.value = this.quoteString(constraint.value);
      }
      return constraint;
  }

  static quoteString(value) {
      var sanitized = value.replace(/\[|\]/g, "").trim();
      if (sanitized.match(/[a-zA-z]+/)) {
          if (!(sanitized.match(/^"/))) {
              sanitized = `"${sanitized}`;
          }
          if (!(sanitized.match(/"$/))) {
              sanitized = `${sanitized}"`;
          }
      }
      return sanitized;
  }

  static postConstraintPromise(flagId, segmentId, constraint) {
      return Axios.post(
          `${API_URL}/flags/${flagId}/segments/${segmentId}/constraints`,
          constraint,
      )
  }

  static deleteConstraintPromise(flagId, segmentId, constraintId) {
      return Axios.delete(
          `${API_URL}/flags/${flagId}/segments/${segmentId}/constraints/${constraintId}`
      )
  }
}