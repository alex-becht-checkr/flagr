<template>
  <el-row>
    <el-col :span="20" :offset="2">
      <div class="container flag-container">
        <el-button class="fab" type="success" round @click="saveFlag(flag)">
          Save Flag
        </el-button>

        <el-dialog title="Delete feature flag" :visible.sync="dialogDeleteFlagVisible">
          <span>Are you sure you want to delete this feature flag?</span>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogDeleteFlagVisible = false">Cancel</el-button>
            <el-button type="primary" @click.prevent="deleteFlag">Confirm</el-button>
          </span>
        </el-dialog>
        <el-dialog title="Create segment" :visible.sync="dialogCreateSegmentOpen">
          <div>
            <p>
              <el-input placeholder="Segment description" v-model="newSegment.description"></el-input>
            </p>
            <p>
              <el-slider v-model="newSegment.rolloutPercent" show-input></el-slider>
            </p>
            <el-button class="width--full" :disabled="!newSegment.description" @click.prevent="createSegment">Create
              Segment</el-button>
          </div>
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
              <div class="flex-row-right">
                <el-button size="small" @click="putFlag(flag)">Save Flag</el-button>
              </div>
            </div>
          </div>
          <el-row class="flag-content" type="flex" align="middle">
            <el-col :span="17">
              <el-row>
                <el-col :span="24">
                  <el-input size="small" placeholder="Key" v-model="flag.key">
                    <template slot="prepend">Flag Key</template>
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
                    <template slot="prepend">Flag Description</template>
                  </el-input>
                </el-col>
              </el-row>
            </el-col>
            <el-col :span="7">
            </el-col>
          </el-row>
          <el-row style="margin: 10px;">
            <h5>
              <span style="margin-right: 10px;">Flag Notes</span>
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
              >+ New Tag</el-button>
            </div>
          </el-row>          
        </el-card>

        <div v-if="loaded && flag">
          <el-tabs>
            <el-tab-pane label="Configuration">
              <el-card class="variants-container base-card">
                <div slot="header" class="clearfix">
                  <h2>Variants</h2>
                </div>
                <div class="variants-container-inner" v-if="flag.variants.length">
                  <div v-for="variant in flag.variants" :key="variant.id">
                    <el-card shadow="never">
                      <el-form label-position="left" label-width="100px">
                        <div class="flex-row id-row">
                          <el-tag type="primary" :disable-transitions="true">
                            Variant ID:
                            <b>{{ variant.id }}</b>
                          </el-tag>
                          <el-input class="variant-key-input" size="small" placeholder="Key" v-model="variant.key">
                            <template slot="prepend">Key</template>
                          </el-input>
                          <div class="flex-row-right save-remove-variant-row">
                            <el-button @click="deleteVariant(variant)" size="small">
                              <span class="el-icon-delete" />
                            </el-button>
                          </div>
                        </div>
                        <el-collapse class="flex-row">
                          <el-collapse-item title="Variant attachment" class="variant-attachment-collapsable-title">
                            <p class="variant-attachment-title">You can add JSON in key/value pairs format.</p>
                            <vue-json-editor v-model="variant.attachment" :showBtns="false" :mode="'code'"
                              v-on:has-error="variant.attachmentValid = false" v-on:input="variant.attachmentValid = true"
                              class="variant-attachment-content"></vue-json-editor>
                          </el-collapse-item>
                        </el-collapse>
                      </el-form>
                    </el-card>
                  </div>
                </div>
                <div class="card--error" v-else>No variants created for this feature flag yet</div>
                <div class="variants-input">
                  <div class="flex-row equal-width constraints-inputs-container">
                    <div>
                      <el-input placeholder="Variant Key" v-model="newVariant.key"></el-input>
                    </div>
                  </div>
                  <el-button class="width--full" :disabled="!newVariant.key" @click.prevent="createVariant">Create
                    Variant</el-button>
                </div>
              </el-card>

              <el-card class="segments-container base-card">
                <div slot="header" class="el-card-header">
                  <div class="flex-row">
                    <div class="flex-row-left">
                      <h2>Segments</h2>
                    </div>
                    <div class="flex-row-right">
                      <el-button @click="dialogCreateSegmentOpen = true">New Segment</el-button>
                    </div>
                  </div>
                </div>
                <div class="segments-container-inner" v-if="flag.segments.length">
                  <el-card
                    shadow="never"
                    v-for="segment in flag.segments"
                    :key="segment.id"
                    class="segment"
                  >
                    <div class="flex-row id-row">
                      <div class="flex-row-left">
                        <el-button size="small" @click="segmentUp(segment, flag.segments)">
                          <span class="el-icon-arrow-up">
                          </span>
                        </el-button>
                        <el-button size="small" @click="segmentDown(segment, flag.segments)" style="margin-right: 15px">
                          <span class="el-icon-arrow-down">
                          </span>
                        </el-button>
                        <el-tag type="primary" :disable-transitions="true">
                          Segment ID:
                          <b>{{ segment.id }}</b>
                        </el-tag>
                      </div>
                      <div class="flex-row-right">
                        <el-button size="small" @click="cloneSegment(segment)">Clone Segment</el-button>
                        <el-button @click="deleteSegment(segment)" size="small">
                          <span class="el-icon-delete" />
                        </el-button>
                      </div>
                    </div>
                    <el-row :gutter="10" class="id-row">
                      <el-col :span="15">
                        <el-input size="small" placeholder="Description" v-model="segment.description">
                          <template slot="prepend">Description</template>
                        </el-input>
                      </el-col>
                      <el-col :span="9">
                        <el-input class="segment-rollout-percent" size="small" placeholder="0"
                          v-model="segment.rolloutPercent" :min="0" :max="100">
                          <template slot="prepend">Rollout</template>
                          <template slot="append">%</template>
                        </el-input>
                      </el-col>
                    </el-row>
                    <el-row>
                      <el-col :span="24">
                        <h5>Constraints (match ALL of them)</h5>
                        <div class="constraints">
                          <div class="constraints-inner" v-if="segment.constraints.length">
                            <div v-for="constraint in segment.constraints" :key="constraint.id">
                              <el-row :gutter="3" class="segment-constraint">
                                <el-col :span="20">
                                  <el-input size="small" placeholder="Property" v-model="constraint.property">
                                    <template slot="prepend">Property</template>
                                  </el-input>
                                </el-col>
                                <el-col :span="4">
                                  <el-select class="width--full" size="small" v-model="constraint.operator"
                                    placeholder="operator">
                                    <el-option v-for="item in operatorOptions" :key="item.value" :label="item.label"
                                      :value="item.value"></el-option>
                                  </el-select>
                                </el-col>
                                <el-col :span="20">
                                  <el-input v-if="listConstraint(constraint)" autosize size="small" label="value"
                                    type="textarea" v-model="constraint.value">
                                  </el-input>
                                  <el-input v-else size="small" v-model="constraint.value">
                                    <template slot="prepend">Value&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</template>
                                  </el-input>
                                </el-col>
                                <el-col :span="2">
                                </el-col>
                                <el-col :span="2">
                                  <el-button type="danger" plain class="width--full" @click="
                                    deleteConstraint(segment, constraint)
                                    " size="small">
                                    <i class="el-icon-delete"></i>
                                  </el-button>
                                </el-col>
                              </el-row>
                            </div>
                          </div>
                          <div class="card--empty" v-else>
                            <span>No constraints (ALL will pass)</span>
                          </div>
                          <div>
                            <el-row :gutter="3">
                              <el-col :span="5">
                                <el-input size="small" placeholder="Property"
                                  v-model="segment.newConstraint.property"></el-input>
                              </el-col>
                              <el-col :span="4">
                                <el-select size="small" v-model="segment.newConstraint.operator" placeholder="operator">
                                  <el-option v-for="item in operatorOptions" :key="item.value" :label="item.label"
                                    :value="item.value"></el-option>
                                </el-select>
                              </el-col>
                              <el-col :span="11">
                                <el-input v-if="listConstraint(segment.newConstraint)" type="textarea" autosize
                                  v-model="segment.newConstraint.value"></el-input>
                                <el-input v-else size="small" v-model="segment.newConstraint.value"></el-input>
                              </el-col>
                              <el-col :span="4">
                                <el-button class="width--full" size="small" type="primary" plain :disabled="!segment.newConstraint.property ||
                                  !segment.newConstraint.value
                                  " @click.prevent="() => createConstraint(segment)
    ">Add Constraint</el-button>
                              </el-col>
                            </el-row>
                          </div>
                        </div>
                      </el-col>
                      <el-col :span="24">
                        <h5>
                          <span>Distribution</span>
                        </h5>
                      </el-col>
                      </el-row>
                        <el-row v-for="variant in flag.variants" :key="variant.id" :gutter="20">
                          <el-col :span="24" class="segment-distribution">
                            <el-input class="variant-key-input segment-distribution-variant" size="small"
                              placeholder="rollout" type="number" :min="0" :max="100"
                              :value="getDistributionValue(segment, variant)"
                              @input="value => setDistributionValue(value, segment, variant)">
                              <template slot="prepend">{{ variant.key }}</template>
                              <template slot="append">%</template>
                            </el-input>
                            <el-progress color="#74E5E0"
                              :percentage="Math.max(0, Math.min(100, getDistributionValue(segment, variant)))"></el-progress>
                          </el-col>
                        </el-row>
                        <el-row v-if="sumDistributions(segment.distributions) !== 100">
                          <el-col :span="24" class="segment-distribution-alert">
                            <el-alert  class="edit-distribution-alert"
                              :title="'Percentages must add up to 100% (currently at ' +
                                sumDistributions(segment.distributions) +
                                '%)'
                                " type="error" :closable="false" show-icon></el-alert>
                          </el-col>
                        </el-row>
                  </el-card>
                </div>
                <div class="card--error" v-else>No segments created for this feature flag yet</div>
              </el-card>
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
                      Delete Flag
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
import vueJsonEditor from "vue-json-editor";
import { operators } from "@/operators.json";

const OPERATOR_VALUE_TO_LABEL_MAP = operators.reduce((acc, el) => {
  acc[el.value] = el.label;
  return acc;
}, {});

const { sum, pluck, handleErr } = helpers;

const { API_URL, FLAGR_UI_POSSIBLE_ENTITY_TYPES } = constants;

const DEFAULT_SEGMENT = {
  description: "",
  rolloutPercent: 50
};

const DEFAULT_CONSTRAINT = {
  operator: "EQ",
  property: "",
  value: ""
};

const DEFAULT_VARIANT = {
  key: ""
};

const DEFAULT_TAG = {
  value: ""
};

const DEFAULT_DISTRIBUTION = {
  bitmap: "",
  variantID: 0,
  variantKey: "",
  percent: 0
};

function processSegment(segment) {
  segment.newConstraint = clone(DEFAULT_CONSTRAINT);
}

function processVariant(variant) {
  if (typeof variant.attachment === "string") {
    variant.attachment = JSON.parse(variant.attachment);
  }
}

export default {
  name: "flag",
  components: {
    spinner: Spinner,
    debugConsole: DebugConsole,
    flagHistory: FlagHistory,
    MarkdownEditor,
    vueJsonEditor
  },
  data() {
    return {
      loaded: false,
      dialogDeleteFlagVisible: false,
      dialogCreateSegmentOpen: false,
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
      newSegment: clone(DEFAULT_SEGMENT),
      newVariant: clone(DEFAULT_VARIANT),
      newTag: clone(DEFAULT_TAG),
      newDistributions: {},
      operatorOptions: operators,
      operatorValueToLabelMap: OPERATOR_VALUE_TO_LABEL_MAP,
      showMdEditor: false,
      historyReload: true,
    };
  },
  computed: {
    newDistributionPercentageSum() {
      return sum(pluck(Object.values(this.newDistributions), "percent"));
    },
    newDistributionIsValid() {
      const percentageSum = sum(
        pluck(Object.values(this.newDistributions), "percent")
      );
      return percentageSum === 100;
    },
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
        if (segment.constraint && segment.constraint.operator && segment.constraint.value) {
          segment.constraint = this.formatConstraint(segment.constraint);
        }
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
    selectVariant($event, variant) {
      const checked = $event;
      if (checked) {
        const distribution = Object.assign(clone(DEFAULT_DISTRIBUTION), {
          variantKey: variant.key,
          variantID: variant.id
        });
        this.$set(this.newDistributions, variant.id, distribution);
      } else {
        this.$delete(this.newDistributions, variant.id);
      }
    },
    saveDistribution(segment) {
      const distributions = Object.values(segment.distributions).filter(
        distribution => distribution.percent !== 0
      ).map(distribution => {
        const dist = clone(distribution)
        delete dist.id;
        return dist
      });

      this.putDistributions(segment.id, distributions).then(response => {
        let distributions = response.data;
        segment.distributions = distributions;
        console.log('success');
        this.$message.success("distributions updated");
      });
    },
    createVariant() {
      Axios.post(
        `${API_URL}/flags/${this.flagId}/variants`,
        this.newVariant
      ).then(response => {
        let variant = response.data;
        this.newVariant = clone(DEFAULT_VARIANT);
        this.flag.variants.push(variant);
        this.$message.success("new variant created");
      }, handleErr.bind(this));
    },
    deleteVariant(variant) {
      const isVariantInUse = this.flag.segments.some(segment =>
        segment.distributions.some(
          distribution => distribution.variantID === variant.id
        )
      );

      if (isVariantInUse) {
        alert(
          "This variant is being used by a segment distribution. Please remove the segment or edit the distribution in order to remove this variant."
        );
        return;
      }

      if (
        !confirm(
          `Are you sure you want to delete variant #${variant.id} [${variant.key}]`
        )
      ) {
        return;
      }

      Axios.delete(
        `${API_URL}/flags/${this.flagId}/variants/${variant.id}`
      ).then(() => {
        this.$message.success("variant deleted");
        this.fetchFlag();
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
      if (!confirm(`Are you sure you want to delete tag #${tag.value}`)) {
        return;
      }

      Axios.delete(`${API_URL}/flags/${this.flagId}/tags/${tag.id}`).then(
        () => {
          this.$message.success("tag deleted");
          this.fetchFlag();
          this.loadAllTags();
        },
        handleErr.bind(this)
      );
    },
    createConstraint(segment) {
      var formattedConstraint = this.formatConstraint(segment.newConstraint);
      this.postConstraint(segment.id, formattedConstraint)
        .then(response => {
          let constraint = response.data;
          segment.constraints.push(constraint);
          segment.newConstraint = clone(DEFAULT_CONSTRAINT);
          this.$message.success("new constraint created");
        }, handleErr.bind(this));
    },
    postConstraint(segmentId, constraint) {
      return Axios.post(
        `${API_URL}/flags/${this.flagId}/segments/${segmentId}/constraints`,
        constraint,
      ).catch(handleErr.bind(this));
    },
    deleteConstraint(segment, constraint) {
      if (!confirm("Are you sure you want to delete this constraint?")) {
        return;
      }

      Axios.delete(
        `${API_URL}/flags/${this.flagId}/segments/${segment.id}/constraints/${constraint.id}`
      ).then(() => {
        const index = segment.constraints.findIndex(
          c => c.id === constraint.id
        );
        segment.constraints.splice(index, 1);
        this.$message.success("constraint deleted");
      }, handleErr.bind(this));
    },
    cloneSegment(segment) {
      const errorHandler = handleErr.bind(this);
      // create segment first
      Axios.post(
        `${API_URL}/flags/${this.flagId}/segments`,
        {
          description: `${segment.description} (Clone)`,
          rolloutPercent: segment.rolloutPercent,
        }
      ).then(response => {
        let newSegment = response.data;
        // clone distribution
        const distributions = segment.distributions.map(({ percent, variantID, variantKey }) => ({ percent, variantID, variantKey }));
        this.putDistributions(newSegment.id, distributions).then(async () => {
          // clone constraints sequentially to preserve order
          for (const constraint of segment.constraints) {
            await this.postConstraint(newSegment.id, constraint);
          }
          this.$message.success('Segment successfully cloned')
          // Re fetch flag to update state
          this.fetchFlag();
        }, errorHandler);
      }, errorHandler);
    },
    segmentUp(segment, segments) {
      const segmentIndex = segments.findIndex(s => s.id === segment.id);

      if (segmentIndex < 1)
        return;

      segments.splice(segmentIndex - 1, 0, segments.splice(segmentIndex, 1)[0]);
    },
    segmentDown(segment, segments) {
      const segmentIndex = segments.findIndex(s => s.id === segment.id);

      if (segmentIndex > segments.length - 1)
        return;

      segments.splice(segmentIndex + 1, 0, segments.splice(segmentIndex, 1)[0]);
    },
    deleteSegment(segment) {
      if (!confirm("Are you sure you want to delete this segment?")) {
        return;
      }

      Axios.delete(
        `${API_URL}/flags/${this.flagId}/segments/${segment.id}`
      ).then(() => {
        const index = this.flag.segments.findIndex(el => el.id === segment.id);
        this.flag.segments.splice(index, 1);
        this.$message.success("segment deleted");
      }, handleErr.bind(this));
    },
    createSegment() {
      Axios.post(
        `${API_URL}/flags/${this.flagId}/segments`,
        this.newSegment
      ).then(response => {
        let segment = response.data;
        processSegment(segment);
        segment.constraints = [];
        if (this.flag.variants.length) {
          // default the first variant to 100%
          this.setDistributionValue(100, segment, this.flag.variants[0]);
        }
        this.newSegment = clone(DEFAULT_SEGMENT);
        this.flag.segments.push(segment);
        this.$message.success("new segment created");
        this.dialogCreateSegmentOpen = false;
      }, handleErr.bind(this));
    },
    putDistributions(segmentId, distributions) {
      return Axios.put(
        `${API_URL}/flags/${this.flagId}/segments/${segmentId}/distributions`,
        { distributions },
      ).catch(handleErr.bind(this));
    },
    loadFlagData() {
      this.fetchFlag();
      this.loadAllTags();
    },
    formatConstraint(constraint) {
      if (constraint.operator === "IN" || constraint.operator === "NIN") {
        if (!constraint.value.match(/\[/)) {
          constraint.value = `[${constraint.value}`;
        }
        if (!constraint.value.match(/\]/)) {
          constraint.value = `${constraint.value}]`;
        }
      } else {
        if (constraint.value.match(/[a-zA-z]+/)) {
          if (!(constraint.value.match(/^"/))) {
            constraint.value = `"${constraint.value}`;
          }
          if (!(constraint.value.match(/"$/))) {
            constraint.value = `${constraint.value}"`;
          }
        }
      }
      return constraint;
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
    listConstraint(constraint) {
      return ["IN", "NOTIN"].includes(constraint.operator);
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
    sumDistributions(distributions) {
      return distributions.reduce((current, { percent }) => current + percent, 0);
    },
    getDistributionValue(segment, variant) {
      const distribution = segment.distributions.find(d => d.variantID === variant.id);
      return distribution ? distribution.percent : 0;
    },
    setDistributionValue(value, segment, variant) {
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
