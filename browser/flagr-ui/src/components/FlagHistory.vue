<template>
  <div>
    <el-card class="snapshot-container" style="border: 0px;">
      <div>

        <el-row>
          <el-col :span="12">
            <div class="tools-debugging-h3">
              <span>Change history</span>
            </div>
          </el-col>
          <el-col :span="12">
            <div style="display: flex; justify-content: flex-end; text-align: center"> 
              <span style="font-size:13px; margin-right: 8px;">Compact View</span>
              <el-tooltip content="Enable/Disable Flag" placement="top" effect="light">
                <el-switch
                  v-model="compactView"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  @change="setFlagEnabled"
                  :active-value="true"
                  :inactive-value="false"
                ></el-switch>
              </el-tooltip>
            </div>
          </el-col>
        </el-row>
        <el-divider></el-divider>

      </div>
      <el-table
        :data="diffs"
        :default-sort = "{prop: 'newId', order: 'descending'}"
        stripe
        row-key="newId"
        :expand-row-keys=[latestSnapshot]
        style="width: 100%"
        class="segment-debug-logs-table"
        >
        <el-table-column type="expand">
          <template slot-scope="props">
            <pre v-if="compactView" class="diff" v-html="props.row.flagDiff[1]"></pre>
            <pre v-else class="diff" v-html="props.row.flagDiff[0]"></pre>
          </template>
        </el-table-column>
        <el-table-column
          label="Snapshot ID"
          prop="newId"
          sortable>
        </el-table-column>
        <el-table-column
          label="Updated At"
          prop="timestamp"
          sortable>
        </el-table-column>
        <el-table-column
          label="Updated By"
          prop="updatedBy"
          sortable>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import Axios from "axios";
import { diffJson, convertChangesToXML } from "diff";

import constants from "@/constants";

const { API_URL } = constants;

export default {
  name: "flag-history",
  props: ["flagId"],
  data() {
    return {
      flagSnapshots: [],
      compactView: true,
      latestSnapshot: 0,
    };
  },
  computed: {
    diffs() {
      let ret = [];
      let snapshots = this.flagSnapshots.slice();
      snapshots.push({ flag: {} });
      for (let i = 0; i < snapshots.length - 1; i++) {
        ret.push({
          timestamp: new Date(snapshots[i].updatedAt).toLocaleString(),
          updatedBy: snapshots[i].updatedBy,
          newId: snapshots[i].id,
          oldId: snapshots[i + 1].id || "NULL",
          flagDiff: this.getDiff(snapshots[i].flag, snapshots[i + 1].flag)
        });
      }
      return ret;
    }
  },
  methods: {
    getFlagSnapshots() {
      Axios.get(`${API_URL}/flags/${this.$props.flagId}/snapshots`).then(
        response => {
          this.flagSnapshots = response.data;
          this.latestSnapshot = this.flagSnapshots[0].id
        },
        () => {
          this.$message.error(`failed to get flag snapshots`);
        }
      );
    },
    getDiff(newFlag, oldFlag) {
      const o = JSON.parse(JSON.stringify(oldFlag));
      const n = JSON.parse(JSON.stringify(newFlag));
      const d = diffJson(o, n);
      if (d.length === 1) {
        return "No changes";
      }
      const fullChanges = convertChangesToXML(d);
      const compactChanges = fullChanges.substring(fullChanges.indexOf("<del>"), fullChanges.lastIndexOf("</ins>") + 6);
      return [fullChanges, compactChanges];
    }
  },
  mounted() {
    this.getFlagSnapshots();
  }
};
</script>

<style lang="less">
.snapshot-container {
  .diff {
    margin: 0;
    del {
      background-color: #f7b3b3;
      text-decoration: none;
    }
    ins {
      background-color: #b6ddc6;
      text-decoration: none;
    }
    overflow-x: auto;
  }
}
</style>
