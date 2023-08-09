<template>
  <div>
    <el-row>
      <el-col :span="20" :offset="2">
        <div class="flags-container container">
          <el-breadcrumb separator="/" v-if="loaded">
            <el-breadcrumb-item>Home page</el-breadcrumb-item>
          </el-breadcrumb>

          <spinner v-if="!loaded" />

          <div v-if="loaded">
            <el-row>
              <div style="margin-top: 15px;">
                <el-input placeholder="Search by..." v-model="searchTerm" class="input-with-select" @change="onEnter">
                  <el-select v-model="select" slot="prepend" placeholder="Search by" class="select-dropdown">
                    <el-option label="Description" value="1"></el-option>
                    <el-option label="Tag" value="2"></el-option>
                    <el-option label="Key" value="3"></el-option>
                  </el-select>
                  <el-button slot="append" icon="el-icon-search" @click="onEnter"></el-button>
                </el-input>
              </div>
            </el-row>
            <el-row>
              <el-tag v-if="searchTermTag" :key="searchTermTag" closable @close="searchTermTag = ''">
                tag:{{ searchTermTag }}
              </el-tag>
              <el-tag v-if="searchTermDescription" :key="searchTermDescription" closable
                @close="searchTermDescription = ''" type="success">
                description:{{ searchTermDescription }}
              </el-tag>
              <el-tag v-if="searchTermKey" :key="searchTermKey" closable @close="searchTermKey = ''" type="warning">
                key:{{ searchTermKey }}
              </el-tag>
              <el-button type="text" @click="clearSearchTerm" v-if="searchTermDescription || searchTermKey || searchTermTag">Clear Search</el-button>
            </el-row>
            <el-row>
              <div style="margin-top: 15px;">
                <el-col :span="2">
                  <h5 class="filter-label">Filter by Status:</h5>
                </el-col>
                <el-col :span="5">
                  <el-radio-group v-model="statusFilterRadioValue" @change="(e) => onStatusFilterChange(e)">
                    <el-radio :label="1">Enabled</el-radio>
                    <el-radio :label="2">Disabled</el-radio>
                    <el-radio :label="3">Both</el-radio>
                  </el-radio-group>
                </el-col>
                <el-col :span="2">
                  <h5 class="filter-label">Filter by Type:</h5>
                </el-col>
                <el-col :span="5">
                  <el-radio-group v-model="typeFilterRadioValue" @change="(e) => onTypeFilterChange(e)">
                    <el-radio :label="1">Active</el-radio>
                    <el-radio :label="2">Deleted</el-radio>
                    <el-radio :label="3">Both</el-radio>
                  </el-radio-group>
                </el-col>
              </div>
            </el-row>

            <el-table :data="flags" :stripe="true" :highlight-current-row="false"
              :default-sort="{ prop: 'id', order: 'descending' }" style="width: 100%">
              <el-table-column prop="id" align="center" label="ID" sortable fixed width="65">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)
                    style="justify-content: center;">{{ scope.row.id }}</component>
                </template>
              </el-table-column>
              <el-table-column prop="key" label="Key" min-width="300">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)>{{ scope.row.key }}
                  </component>
                </template>
              </el-table-column>
              <el-table-column prop="description" label="Description" min-width="300">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)>{{ scope.row.description }}
                  </component>
                </template>
              </el-table-column>
              <el-table-column prop="tags" label="Tags" min-width="200">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)>
                    <el-tag v-for="tag in scope.row.tags" :key="tag.id" disable-transitions>{{ tag.value
                    }}</el-tag>
                  </component>
                </template>
              </el-table-column>
              <el-table-column prop="updatedBy" label="Last Updated By" sortable width="200">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)>{{ scope.row.updatedBy }}
                  </component>
                </template>
              </el-table-column>
              <el-table-column prop="updatedAt" label="Updated At" :formatter="datetimeFormatter" sortable width="180">
                <template slot-scope="scope">
                  <component :is="scope.row.isDeleted?'span':'a'" :href=getFlagUri(scope.row)>{{
                    datetimeFormatter(scope.row.updatedAt) }}</component>
                </template>
              </el-table-column>
              <el-table-column prop="enabled" label="Enabled" sortable align="center" fixed="right" width="140">
                <template slot-scope="scope">
                  <el-tag :type="scope.row.enabled ? 'primary' : 'danger'" disable-transitions>{{ scope.row.enabled ? "on"
                    :
                    "off" }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="action" label="Action" align="center" fixed="right" width="100">
                <template slot-scope="scope">
                  <el-button @click="restoreFlag(scope.row)" type="warning" size="small"
                    v-if="scope.row.isDeleted">Restore</el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-row type="flex" justify="space-between">
                <el-pagination layout="total" :total="flagsCount">
                </el-pagination>
                <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                  :current-page.sync="currentPage" :page-sizes="[10, 25, 50, 100, 1000]" :page-size="pageSize"
                  layout="sizes, prev, pager, next" :total="flagsCount">
                </el-pagination>
            </el-row>
          </div>
        </div>
      </el-col>
    </el-row>
    <el-button class="fab" type="success" icon="el-icon-plus" round @click="showFlagCreator = true">Create
      Flag</el-button>
    <el-dialog title="Create New Flag" :visible.sync="showFlagCreator">
      <el-form>
        <el-form-item label="Flag Description">
          <el-input v-model="newFlag.description" placeholder="Required"></el-input>
        </el-form-item>
        <el-form-item label="Flag Key">
          <el-input v-model="newFlag.key" placeholder="Optional"></el-input>
        </el-form-item>
        <el-form-item label="Template">
          <el-select placeholder="Select a template" v-model="selectedTemplate">
            <el-option label="Empty Flag" value="empty"></el-option>
            <el-option label="Simple Boolean" value="simple_boolean_flag"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="showFlagCreator = false">Cancel</el-button>
        <el-button type="success" @click="createFlag" :disabled="!newFlag.description">Create</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Axios from "axios";

import constants from "@/constants";
import Spinner from "@/components/Spinner";
import helpers from "@/helpers/helpers";

const { handleErr } = helpers;

const { API_URL } = constants;

export default {
  name: "flags",
  components: {
    spinner: Spinner
  },
  data() {
    return {
      loaded: false,
      deletedFlagsLoaded: false,
      flags: [],
      deletedFlags: [],
      searchTerm: "",
      searchTermTag: "",
      searchTermDescription: "",
      searchTermKey: "",
      includeDeleted: false,
      onlyDeleted: false,
      enabled: null,
      currentPage: 1,
      pageSize: 50,
      flagsCount: 0,
      newFlag: {
        description: "",
        key: ""
      },
      select: '1',
      statusFilterRadioValue: 3,
      typeFilterRadioValue: 1,
      showFlagCreator: false,
      selectedTemplate: "empty"
    };
  },
  created() {
    this.setSearchData();
    this.getFlags();
    this.getFlagsCount();

  },
  watch: {
    includeDeleted: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    },
    onlyDeleted: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    },
    enabled: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    },
    searchTermDescription: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    },
    searchTermTag: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    },
    searchTermKey: function () {
      this.loaded = false;
      this.getFlags();
      this.getFlagsCount();
    }
  },
  methods: {
    flagEnabledFormatter(row, col, val) {
      return val ? "on" : "off";
    },
    datetimeFormatter(val) {
      if (val) {
        return new Date(val).toLocaleString();
      } else {
        return "";
      }
    },
    getFlagUri(row) {
      if (row.isDeleted) {
        return "";
      }
      return `/#/flags/${row.id}`;
    },
    onEnter: function () {
      if (this.select == 1) {
        this.searchTermDescription = this.searchTerm
      }
      if (this.select == 2) {
        this.searchTermTag = this.searchTerm
      }
      if (this.select == 3) {
        this.searchTermKey = this.searchTerm
      }
      this.searchTerm = ''
      this.updateHistory();
    },
    createFlag() {
      if (!this.newFlag.description) {
        return;
      }
      var template = {};
      if (this.selectedTemplate != "empty") {
        template = {
          template: this.selectedTemplate
        };
      }
      Axios.post(`${API_URL}/flags`, {
        ...this.newFlag,
        ...(template)
      }).then(response => {
        let flag = response.data;
        this.newFlag.description = "";
        this.$message.success("flag created");

        flag._new = true;
        this.flags.unshift(flag);
        this.showFlagCreator = false;
        this.$router.push({ name: "flag", params: { flagId: flag.id } });
      }, handleErr.bind(this));
    },
    restoreFlag(row) {
      this.$confirm('This will recover the deleted flag. Continue?', 'Warning', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel'
      }).then(() => {
        Axios.put(`${API_URL}/flags/${row.id}/restore`).then(response => {
          let flag = response.data;
          this.$message.success(`Flag updated`);
          this.flags.push(flag);
          this.deletedFlags = this.deletedFlags.filter(function (el) {
            return el.id != flag.id;
          });
        }, handleErr.bind(this));
      });

    },
    getFlags() {
      this.loaded = false;
      Axios.get(`${API_URL}/flags`,
        {
          params: {
            ordered_by_most_recent: true,
            include_deleted: this.includeDeleted,
            limit: this.pageSize,
            offset: (this.currentPage - 1) * this.pageSize,
            description_like: this.searchTermDescription,
            tags: this.searchTermTag,
            key: this.searchTermKey,
            deleted: this.onlyDeleted,
            enabled: this.enabled
          }
        }).then(response => {
          let flags = response.data;
          this.loaded = true;
          this.flags = flags;
        }, handleErr.bind(this));
    },
    getFlagsCount() {
      Axios.get(`${API_URL}/flags/count`,
        {
          params: {
            include_deleted: this.includeDeleted,
            description_like: this.searchTermDescription,
            tags: this.searchTermTag,
            key: this.searchTermKey
          }
        }).then(response => {
          let { totalFlags } = response.data;
          this.flagsCount = totalFlags;
        }, handleErr.bind(this));
    },
    fetchDeletedFlags() {
      if (!this.deletedFlagsLoaded) {
        var self = this;
        Axios.get(`${API_URL}/flags?deleted=true`).then(response => {
          let flags = response.data;
          flags.reverse();
          self.deletedFlags = flags;
          self.deletedFlagsLoaded = true;
        }, handleErr.bind(this));
      }
    },
    filterStatus(value, row) {
      return row.enabled === value;
    },
    handleSizeChange(val) {
      this.pageSize = val;
      this.getFlags();
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.getFlags();
    },
    updateHistory() {
      history.pushState({
        searchTermDescription: this.searchTermDescription,
        searchTermTag: this.searchTermTag,
        searchTermKey: this.searchTermKey,
        includeDeleted: this.includeDeleted,
        currentPage: this.currentPage,
        pageSize: this.pageSize
      }, '', '#/home');
    },
    setSearchData() {
      if (history.state) {
        this.searchTermTag = history.state.searchTermTag,
        this.searchTermDescription = history.state.searchTermDescription
        this.searchTermKey = history.state.searchTermKey,
        this.includeDeleted = history.state.includeDeleted,
        this.currentPage = history.state.currentPage || 1,
        this.pageSize = history.state.pageSize || 10
      }
    },
    clearSearchTerm() {
      this.searchTerm = "";
      this.searchTermDescription = "";
      this.searchTermTag = "";
      this.searchTermKey = "";
      this.includeDeleted = false;
      this.currentPage = 1;
      this.pageSize = 10;
      this.updateHistory();
    },

    onTypeFilterChange(e) {
      if (e == 1){
        this.includeDeleted = false
        this.onlyDeleted = false
      }
      if (e == 2){
        this.onlyDeleted = true
        this.includeDeleted = true
      }
      if (e == 3){
        this.includeDeleted = true
        this.onlyDeleted = false
      }
    },
    onStatusFilterChange(e) {
      if (e == 1){
        this.enabled = true
      }
      if (e == 2){
        this.enabled = false
      }
      if (e == 3){
        this.enabled = null
      }
    }
  }
};
</script>

<style lang="less">
.flags-container {
  .el-table {
    margin-top: 2em;
  }

  .el-button-group .el-button--primary:first-child {
    border-right-color: #dcdfe6;
  }

  .deleted-flags-table {
    margin-top: 2rem;
  }

  .select-dropdown {
    width: 150px;
  }

  .filter-label {
    margin: 1px 0 5px;
  }
}

.fab {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  box-shadow: 1px 1px rgba(0, 0, 0, 0.2);
  z-index: 999;
}
</style>

<style lang="less" scoped>
a {
  text-decoration: none;
  color: unset;
  display: flex;
}
</style>
