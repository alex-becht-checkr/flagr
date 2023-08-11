<template>
  <el-row>
    <el-col :span="20" :offset="2">
      <div class="container flag-container">
        <el-button class="fab" type="success" round @click="saveFlag(flag)">
          Save flag
        </el-button>

        <el-dialog title="Delete tag" :visible.sync="dialogDeleteTagVisible">
          <span>{{`Are you sure you want to delete tag #${tagToDelete.value}`}}</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="cancelTagDeletion">Cancel</el-button>
            <el-button type="primary" @click="confirmTagDeletion">Confirm</el-button>
          </span>
        </el-dialog>

        <el-dialog title="Delete feature flag" :visible.sync="dialogDeleteFlagVisible">
          <span>Are you sure you want to delete this feature flag?</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogDeleteFlagVisible = false">Cancel</el-button>
            <el-button type="primary" @click.prevent="deleteFlag">Confirm</el-button>
          </span>
        </el-dialog>

        <el-breadcrumb separator="/">
          <el-breadcrumb-item :to="{ name: 'home' }"><i style="margin-right: 5px;" class="el-icon-s-home" />Home</el-breadcrumb-item>
          <el-breadcrumb-item>Flag</el-breadcrumb-item>
        </el-breadcrumb>

        <el-card class="flag-config-card base-card">
          <div slot="header" class="el-card-header">
            <div class="flex-row">
              <div class="flex-row-left">
                <h2>Flag ID: {{ $route.params.flagId }}</h2>
              </div>
            </div>
          </div>
          <el-row class="flag-content" type="flex" align="middle">
            <el-col :span="17">
              <el-row>
                <el-col :span="24">
                  <el-input size="small" placeholder="Key" v-model="flag.key">
                    <template slot="prepend">Flag key</template>
                  </el-input>
                </el-col>
              </el-row>
            </el-col>
            <el-col :span="7">
              <div style="display: flex; justify-content: flex-end; text-align: center"> 
                <span style="font-size:13px; margin-right: 8px;">Enabled</span>
                <el-tooltip content="Enable/Disable Flag" placement="top" effect="light">
                  <el-switch
                    v-model="flag.enabled"
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
          <el-row class="flag-content" type="flex" align="middle">
            <el-col :span="17">
              <el-row>
                <el-col :span="24">
                  <el-input
                    size="small"
                    placeholder="Description"
                    v-model="flag.description"
                  >
                    <template slot="prepend">Flag description</template>
                  </el-input>
                </el-col>
              </el-row>
            </el-col>
            <el-col :span="7">
            </el-col>
          </el-row>
          <el-row style="margin: 10px;">
            <h5>
              <span style="margin-right: 10px;">Flag notes</span>
              <el-button size="mini" @click="toggleShowMdEditor">
                <span :class="editViewIcon"></span>
                {{ !this.showMdEditor ? "edit" : "view" }}
              </el-button>
            </h5>
          </el-row>
          <el-row>
            <markdown-editor
              :showEditor="this.showMdEditor"
              :markdown.sync="flag.notes"
              @save="putFlag(flag)"
            ></markdown-editor>
          </el-row>
          <el-row style="margin: 10px;">
            <h5>
              <span style="margin-right: 10px;">Tags</span>
            </h5>
          </el-row>
          <el-row>
            <div class="tags-container-inner">
              <el-tag
                v-for="tag in flag.tags"
                :key="tag.id"
                closable
                :type="warning"
                @close="deleteTag(tag)"
              >{{tag.value}}</el-tag>
              <el-autocomplete
                class="tag-key-input"
                v-if="tagInputVisible"
                v-model="newTag.value"
                ref="saveTagInput"
                size="mini"
                :trigger-on-focus="false"
                :fetch-suggestions="queryTags"
                @select="createTag"
                @keyup.enter.native="createTag"
                @keyup.esc.native="cancelCreateTag"
              ></el-autocomplete>
              <el-button
                v-else
                class="button-new-tag"
                size="small"
                @click="showTagInput"
              >+ New tag</el-button>
            </div>
          </el-row>          
        </el-card>

        <div v-if="loaded && flag">
          <el-tabs>
            <el-tab-pane label="Configuration">
              <Variants :flag-id="flag.id"
                        :segments="flag.segments"
                        :variants="flag.variants"
                        @fetch-flag="fetchFlag"/>

              <Segments :flag-id="flag.id"
                        :segments="flag.segments"
                        :variants="flag.variants"
                        @fetch-flag="fetchFlag"/>
              <spinner v-if="!loaded"></spinner>
            </el-tab-pane>
            <el-tab-pane label="Settings">
              <el-card class="flag-config-card base-card">
                <div class="tools-debugging-h3">
                  <span>Logging</span>
                </div>
                <el-divider></el-divider>
                <div>
                  <el-row class="settings-body">
                      <el-col>
                        <div>
                          <el-switch
                            size="small"
                            v-model="flag.dataRecordsEnabled"
                            active-color="#74E5E0"
                            :active-value="true"
                            :inactive-value="false"
                          ></el-switch>
                          <span style="margin-left: 10px;">Enable logging to data pipeline, e.g. Kafka, Kinesis, PubSub.</span>
                        </div>
                      </el-col>
                  </el-row>
                  <el-row class="settings-body">
                    <el-col style="margin-left: 50px;">
                      <span v-show="!!flag.dataRecordsEnabled" >Specify the override of the entityType in data records logging: </span>
                      <el-select
                          v-show="!!flag.dataRecordsEnabled"
                          v-model="flag.entityType"
                          size="mini"
                          filterable
                          :allow-create="allowCreateEntityType"
                          default-first-option
                          placeholder="<null>"
                        >
                          <el-option
                            v-for="item in entityTypes"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                          ></el-option>
                      </el-select>
                    </el-col>
                  </el-row>
                </div>
                <div class="tools-debugging-h3">
                  <span>Deleting flag</span>
                </div>
                <el-divider></el-divider>
                <div>
                  <el-row class="settings-body">
                    <el-button @click="dialogDeleteFlagVisible = true" type="danger" plain>
                      <span class="el-icon-delete"></span>
                      Delete flag
                    </el-button>
                  </el-row>
                </div>
              </el-card>
            </el-tab-pane>
            <el-tab-pane label="Tools">
              <debug-console :flag="this.flag"></debug-console>
            </el-tab-pane>
            <el-tab-pane label="History">
              <flag-history :key="historyReload" :flag-id="parseInt($route.params.flagId, 10)"></flag-history>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import clone from "lodash.clone";
import Axios from "axios";

import constants from "@/constants";
import helpers from "@/helpers/helpers";
import Spinner from "@/components/Spinner";
import DebugConsole from "@/components/DebugConsole";
import FlagHistory from "@/components/FlagHistory";
import MarkdownEditor from "@/components/MarkdownEditor.vue";
import { operators } from "@/operators.json";
import ConstraintService from "@/services/constraintService";
import Segments from "@/components/Flag/Segments.vue";
import Variants from "@/components/Flag/Variants.vue";

const OPERATOR_VALUE_TO_LABEL_MAP = operators.reduce((acc, el) => {
  acc[el.value] = el.label;
  return acc;
}, {});

const { handleErr } = helpers;

const { API_URL, FLAGR_UI_POSSIBLE_ENTITY_TYPES } = constants;

const DEFAULT_TAG = {
  value: ""
};

function processSegment(segment) {
  segment.newConstraint = clone(ConstraintService.DEFAULT_CONSTRAINT);
}

function processVariant(variant) {
  if (typeof variant.attachment === "string") {
    variant.attachment = JSON.parse(variant.attachment);
  }
}

export default {
  name: "flag",
  components: {
    Variants,
    Segments,
    spinner: Spinner,
    debugConsole: DebugConsole,
    flagHistory: FlagHistory,
    MarkdownEditor
  },
  data() {
    return {
      loaded: false,
      dialogDeleteFlagVisible: false,
      dialogDeleteTagVisible: false,
      tagToDelete: {},
      entityTypes: [],
      allTags: [],
      allowCreateEntityType: true,
      tagInputVisible: false,
      flag: {
        createdBy: "",
        dataRecordsEnabled: false,
        entityType: "",
        description: "",
        enabled: false,
        id: 0,
        key: "",
        tags: [],
        segments: [],
        updatedAt: "",
        variants: [],
        notes: ""
      },
      newTag: clone(DEFAULT_TAG),
      newDistributions: {},
      operatorValueToLabelMap: OPERATOR_VALUE_TO_LABEL_MAP,
      showMdEditor: false,
      historyReload: true
    };
  },
  computed: {
    flagId() {
      return this.$route.params.flagId;
    },
    editViewIcon() {
      return {
        "el-icon-edit": !this.showMdEditor,
        "el-icon-view": this.showMdEditor
      };
    },
    toggleInnerConfigCard() {
      if (!this.showMdEditor && !this.flag.notes) {
        return "flag-inner-config-card";
      } else {
        return "";
      }
    }
  },
  methods: {
    deleteFlag() {
      const flagId = this.flagId;
      Axios.delete(`${API_URL}/flags/${this.flagId}`).then(() => {
        this.$router.replace({ name: "home" });
        this.$message.success(`You deleted flag ${flagId}`);
      }, handleErr.bind(this));
    },
    saveFlag(flag) {
      for (var segment of flag.segments) {
        segment.rolloutPercent = parseInt(segment.rolloutPercent);
        segment.constraints = segment.constraints.map(this.formatConstraint);
        segment.distributions = segment.distributions.filter(
            distribution => distribution.percent !== 0
        );
      }
      Axios.put(`${API_URL}/flags/${this.flagId}/full`, flag)
        .then(() => {
          this.loadFlagData();
          this.$message.success(`Flag updated`);
          this.historyReload = !this.historyReload;
        }, handleErr.bind(this));
    },
    putFlag(flag) {
      Axios.put(`${API_URL}/flags/${this.flagId}`, {
        description: flag.description,
        dataRecordsEnabled: flag.dataRecordsEnabled,
        key: flag.key || "",
        entityType: flag.entityType || "",
        notes: flag.notes || ""
      }).then(() => {
        this.$message.success(`Flag updated`);
      }, handleErr.bind(this));
    },
    setFlagEnabled(checked) {
      Axios.put(`${API_URL}/flags/${this.flagId}/enabled`, {
        enabled: checked
      }).then(() => {
        const checkedStr = checked ? "on" : "off";
        this.$message.success(`You turned ${checkedStr} this feature flag`);
      }, handleErr.bind(this));
    },
    createTag() {
      Axios.post(`${API_URL}/flags/${this.flagId}/tags`, this.newTag).then(
        response => {
          let tag = response.data;
          this.newTag = clone(DEFAULT_TAG);
          if (!this.flag.tags.map(tag => tag.value).includes(tag.value)) {
            this.flag.tags.push(tag);
            this.$message.success("new tag created");
          }
          this.tagInputVisible = false;
          this.loadAllTags();
        },
        handleErr.bind(this)
      );
    },
    cancelCreateTag() {
      this.newTag = clone(DEFAULT_TAG);
      this.tagInputVisible = false;
    },
    queryTags(queryString, cb) {
      let results = this.allTags.filter(tag =>
        tag.value.toLowerCase().includes(queryString.toLowerCase())
      );
      cb(results);
    },
    loadAllTags() {
      Axios.get(`${API_URL}/tags`).then(response => {
        let result = response.data;
        this.allTags = result;
      }, handleErr.bind(this));
    },
    showTagInput() {
      this.tagInputVisible = true;
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    deleteTag(tag) {
      this.dialogDeleteTagVisible = true;
      this.tagToDelete = tag;
    },
    cancelTagDeletion() {
      this.dialogDeleteTagVisible = false;
      this.tagToDelete = {};
    },
    confirmTagDeletion() {
      Axios.delete(`${API_URL}/flags/${this.flagId}/tags/${this.tagToDelete.id}`).then(
          () => {
            this.$message.success("tag deleted");
            this.dialogDeleteTagVisible = false;
            this.tagToDelete = {};
            this.fetchFlag();
            this.loadAllTags();
          },
          handleErr.bind(this)
      );
    },
    loadFlagData() {
      this.fetchFlag();
      this.loadAllTags();
    },
    formatConstraint(constraint) {
      return ConstraintService.formatConstraint(constraint);
    },
    fetchFlag() {
      Axios.get(`${API_URL}/flags/${this.flagId}`).then(response => {
        let flag = response.data;
        flag.segments.forEach(segment => processSegment(segment));
        flag.variants.forEach(variant => processVariant(variant));
        this.flag = flag;
        this.loaded = true;
      }, handleErr.bind(this));
      this.fetchEntityTypes();
    },
    fetchEntityTypes() {
      function prepareEntityTypes(entityTypes) {
        let arr = entityTypes.map(key => {
          let label = key === "" ? "<null>" : key;
          return { label: label, value: key };
        });
        if (entityTypes.indexOf("") === -1) {
          arr.unshift({ label: "<null>", value: "" });
        }
        return arr;
      }

      if (
        FLAGR_UI_POSSIBLE_ENTITY_TYPES &&
        FLAGR_UI_POSSIBLE_ENTITY_TYPES != "null"
      ) {
        let entityTypes = FLAGR_UI_POSSIBLE_ENTITY_TYPES.split(",");
        this.entityTypes = prepareEntityTypes(entityTypes);
        this.allowCreateEntityType = false;
        return;
      }

      Axios.get(`${API_URL}/flags/entity_types`).then(response => {
        this.entityTypes = prepareEntityTypes(response.data);
      }, handleErr.bind(this));
    },
    toggleShowMdEditor() {
      this.showMdEditor = !this.showMdEditor;
    },
  },
  mounted() {
    this.loadFlagData();
  }
};
</script>

<style lang="less">
h5 {
  padding: 0;
  margin: 10px 0 5px;
}

.flag-inner-config-card {
  .el-card__body {
    padding-bottom: 0px;
  }
}

.segment {
  .highlightable {
    padding: 4px;

    &:hover {
      background-color: #ddd;
    }
  }

  .segment-constraint {
    margin-bottom: 12px;
    padding: 1px;
    background-color: #f6f6f6;
    border-radius: 5px;
  }

  .distribution-card {
    height: 110px;
    text-align: center;

    .el-card__body {
      padding: 3px 10px 10px 10px;
    }

    font-size: 0.9em;
  }
}

ol.constraints-inner {
  background-color: white;
  padding-left: 8px;
  padding-right: 8px;
  border-radius: 3px;
  border: 1px solid #ddd;

  li {
    padding: 3px 0;

    .el-tag {
      font-size: 0.7em;
    }
  }
}

.constraints-inputs-container {
  padding: 5px 0;
}

.variants-container-inner {
  .el-card {
    margin-bottom: 1em;
  }

  .el-input-group__prepend {
    width: 2em;
  }
}

.segment-distribution {
  margin-top: 10px;
  display: flex;

  .variant-key-input {
    width: 275px;
  }

  .el-input-group {
    input {
      width: 80px;
    }

    .el-input-group__prepend {
      width: 100%;
      max-width: 100px;
      text-align: end;
      overflow-x: hidden;
      text-overflow: ellipsis;
    }
  }

  .el-progress {
    flex-grow: 1;
    margin-left: 10px;
    display: flex;
    align-items: center;
    height: 32px;
  }
}
.segment-distribution-alert {
  margin: 10px;
  .el-alert{
    width: calc(100% - 30px);
  }
}

.segment-description-rollout {
  margin-top: 10px;
}

.edit-distribution-button {
  margin-top: 5px;
}

.edit-distribution-alert {
  margin-top: 10px;
}

.el-form-item {
  margin-bottom: 5px;
}

.id-row {
  margin-bottom: 8px;
}

.flag-config-card {
  .flag-content {
    margin-top: 8px;
    margin-bottom: -8px;

    .el-input-group__prepend {
      width: 8em;
    }
  }

  .data-records-label {
    margin-left: 3px;
    margin-bottom: 5px;
    margin-top: 6px;
    font-size: 0.65em;
    white-space: nowrap;
    vertical-align: middle;
  }
}

.variant-attachment-collapsable-title {
  margin: 0;
  font-size: 13px;
  color: #909399;
  width: 100%;
}

.variant-attachment-title {
  margin: 0;
  font-size: 13px;
  color: #909399;
}

.variant-key-input {
  margin-left: 10px;
  width: 50%;
}

.settings-body {
  margin-left: 10px;
  margin-bottom: 14px;
  font-size: 13px;
}

.save-remove-variant-row {
  padding-bottom: 5px;
}

.tag-key-input {
  margin: 2.5px;
  width: 20%;
}

.tags-container-inner {
  margin-bottom: 10px;
}

.button-new-tag {
  margin: 2.5px;
}

.el-tabs__item {
  font-size: 18px;
}
</style>
