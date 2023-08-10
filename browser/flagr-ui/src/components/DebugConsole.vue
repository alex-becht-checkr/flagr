<template>
  <el-card class="flag-config-card base-card" >
    <div>
      <el-row>
        <el-col :span="12">
          <div class="tools-debugging-h3">
            <span>Debugging tool</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div align="right">
            <el-button type="text" @click="toggleDebugModes" class="tools-debugging-link-button">
              Switch to
              {{ !this.showSimpleDebug ? "Simple" : "Advanced" }}
              Debugging
            </el-button>
          </div>
        </el-col>
      </el-row>
      <el-divider></el-divider>
    </div>  
    <div v-if="showSimpleDebug">
      <simple-debug-console :flag="this.flag"
                            :eval-context="this.evalContext"
                            :eval-result="this.evalResult"
                            @post-evaluation="postEvaluation(evalContext)"
      ></simple-debug-console>
    </div>
    <div v-else>
      <advanced-debug-console :flag="this.flag"
                              :eval-context="this.evalContext"
                              :eval-result="this.evalResult"
                              :batch-eval-context="this.batchEvalContext"
                              :batch-eval-result="this.batchEvalResult"
                              @post-evaluation="postEvaluation(evalContext)"
                              @post-evaluation-batch="postEvaluationBatch(batchEvalContext)"
      ></advanced-debug-console>
    </div>
  </el-card>
</template>

<script>
import SimpleDebugConsole from "@/components/SimpleDebugConsole";
import AdvancedDebugConsole from "@/components/AdvancedDebugConsole.vue";
import Axios from "axios";

import constants from "@/constants";

const {API_URL} = constants;

export default {
  name: "debug-console",
  props: ["flag"],
  data() {
    return {
      evalContext: {
        entityID: "a1234",
        entityType: "report",
        entityContext: {
          hello: "world"
        },
        enableDebug: true,
        flagID: this.flag.id,
        flagKey: this.flag.key
      },
      evalResult: {},
      batchEvalContext: {
        entities: [
          {
            entityID: "a1234",
            entityType: "report",
            entityContext: {
              hello: "world"
            }
          },
          {
            entityID: "a5678",
            entityType: "report",
            entityContext: {
              hello: "world"
            }
          }
        ],
        enableDebug: true,
        flagIDs: [this.flag.id]
      },
      batchEvalResult: {},
      showSimpleDebug: true
    };
  },
  methods: {
    postEvaluation(evalContext) {
      Axios.post(`${API_URL}/evaluation`, evalContext).then(
          response => {
            this.$message.success(`evaluation success`);
            this.evalResult = response.data;
          },
          () => {
            this.$message.error(`evaluation error`);
          }
      );
    },
    postEvaluationBatch(batchEvalContext) {
      Axios.post(`${API_URL}/evaluation/batch`, batchEvalContext).then(
          response => {
            this.$message.success(`evaluation success`);
            this.batchEvalResult = response.data;
          },
          () => {
            this.$message.error(`evaluation error`);
          }
      );
    },
    toggleDebugModes() {
      this.showSimpleDebug = !this.showSimpleDebug;
    }
  },
  components: {
    simpleDebugConsole: SimpleDebugConsole,
    advancedDebugConsole: AdvancedDebugConsole
  }
};
</script>

<style lang="less">
.json-editor {
  margin-top: 3px;

  .jsoneditor {
    height: 400px;
  }
}

.tools-debugging-h2 {
  font-size: 24px;
}

.tools-debugging-h3 {
  font-size: 18px;
}

.tools-debugging-h4 {
  font-size: 15px;
}

.tools-debugging-subheader {
  margin-top: 10px;
  margin-bottom: 10px;
  font-weight:bold;
}

.tools-debugging-top-description {
  margin-top: 20px;
  margin-bottom: 10px;
}

.tools-debugging-space {
  margin-top: 10px;
  margin-bottom: 10px;
}

.tools-debugging-body {
  font-size: 13px;
}

.el-divider--horizontal { 
  margin-top: 0.6rem;
  margin-bottom: 1rem;
}

.tools-debugging-link-button {
  padding: 0;
  font-size: 13px;
}

</style>
