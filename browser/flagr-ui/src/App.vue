<template>
  <div id="app">
    <el-menu mode="horizontal" class="navbar">
      <el-row>
        <el-col :span="20" :offset="2">
          <el-row>
            <el-col :span="6">
              <router-link :to="{ name: 'home' }">
                <div class="logo-container">
                  <h3 class="logo">Flagr</h3>
                </div>
              </router-link>
            </el-col>
            <el-col :span="1" :offset="15">
              <a href="https://openflagr.github.io/flagr/api_docs" target="_blank"
                ><h3>API</h3></a
              >
            </el-col>
            <el-col :span="1" :offset="1">
              <a href="https://openflagr.github.io/flagr" target="_blank"
                ><h3>Docs</h3></a
              >
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-menu>
    <div class="router-view-container">
      <router-view></router-view>
    </div>
    <el-row>
      <el-menu mode="horizontal" class="footer">
        <el-row>
          <div>
            <span>v{{ version }}</span>
          </div>
        </el-row>
      </el-menu>
    </el-row>
  </div>
</template>

<script>
const version = require("../package.json").version || "1.0.0";
export default {
  name: "app",
  data: () => ({
    version,
  }),
};
</script>

<style lang="less">
body {
  margin: 0;
  padding: 0;
  font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

h1,
h2 {
  font-weight: normal;
}

ol {
  margin: 0;
  padding-left: 20px;
}

.width--full {
  width: 100%;
}

#app {
  color: #2c3e50;

  span[size="small"] {
    font-size: 0.85em;
  }

  .navbar {
    background-color: #74e5e0;
    color: #2e4960;
    border: 0;

    .logo-container {
      display: flex;
      align-items: center;
      font-weight: bold;

      h3 {
        margin-right: 10px;
        &:hover {
          color: #000;
        }
      }

      span {
        font-size: 12px;
      }
    }

    a {
      color: inherit;
      text-decoration: none;
    }

    .el-col {
      text-align: right;

      &:first-child {
        text-align: left;
      }
    }
  }

  .router-view-container {
    margin-bottom: 50px;
  }

  .footer {
      background-color: #74e5e0;
      color: #2e4960;
      border: 0;
      text-align: center;
      padding: 10px;
      margin-top: 20px;
      position: absolute;
      justify-content: center;
      bottom: 0;
      left: 0;
      right: 0;
    }

  .flex-row {
    display: flex;
    align-items: center;
    justify-content: center;
    &-right {
      margin-left: auto;
    }
    &.equal-width {
      > * {
        flex: 1;
      }
    }
    &.align-items-top {
      align-items: flex-start;
    }
  }

  .container {
    margin: 0 auto;
    margin-top: 20px;
  }

  img {
    height: 60px;
  }

  .card {
    &--error {
      box-sizing: border-box;
      background-color: #fff9f9;
      padding: 10px;
      text-align: center;
      color: #ed2d2d;
      border: 1px solid #ed2d2d;
      border-radius: 3px;
      width: 100%;
      margin-bottom: 12px;
    }
    &--empty {
      box-sizing: border-box;
      background-color: #eee;
      padding: 10px;
      text-align: center;
      color: #777;
      border: 1px solid #ccc;
      border-radius: 3px;
      width: 100%;
      margin-bottom: 12px;
    }
  }

  .el-breadcrumb {
    margin-bottom: 20px;
  }

  .el-input {
    margin-bottom: 2px;
  }

  .segment-rollout-percent input {
    text-align: right;
  }

  .el-card {
    .el-card__header {
      color: #2e4960;
      border: 0;
      border-bottom: 1px solid #e4e7ed;

      h2 {
        margin: -0.2em;
        color: inherit;
        font-size: 20px;
      }
    }
    margin-bottom: 1em;
  }

  .base-card {
    border: 0px;

  .el-card__body {
    .el-card {
      border: 1px solid #ebeef5;
    }
    .el-card__body {
      border-left: 4px solid #74E5E0;
    }
  }
  box-shadow: none !important;
  }

  .jsoneditor {
    border-color: #e4e7ed;
    .jsoneditor-menu {
      background-color: #e4e7ed;
      border-bottom-color: #e4e7ed;
    }
    .jsoneditor-poweredBy {
      display: none;
    }
  }

  .el-tag {
    margin: 2.5px;
  }

  .fab {
    position: fixed;
    bottom: 2rem;
    right: 2rem;
    box-shadow: 1px 1px rgba(0, 0, 0, 0.2);
    z-index: 999;
  }
  
  .align-right {
    text-align: right;
  }
}
</style>
