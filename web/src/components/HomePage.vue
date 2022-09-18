<template>
  <v-row style="height:100%">
    <v-col cols="7">
      <v-row style="height:65vh">
        <v-col>
          <v-card flat tile>
            <v-toolbar flat dense>
              <v-toolbar-title class="title overline">Policy</v-toolbar-title>
              <v-spacer></v-spacer>
              <v-btn color="primary" icon tile @click.stop="run" :disabled="loading">
                <v-icon>mdi-play</v-icon>
              </v-btn>
            </v-toolbar>
              <CodeEditor v-model="config.policy.content" language="yaml"/>
          </v-card>
        </v-col>
      </v-row>
      <v-divider></v-divider>
      <v-row style="height:34vh">
        <v-col>
          <v-card flat tile :loading="loading">
            <v-toolbar flat dense>
              <v-toolbar-title class="title overline">Result</v-toolbar-title>
              <v-spacer></v-spacer>
                <div v-if="status">
                  <v-chip
                    label
                    class="ma-3 text-capitalize"
                    v-text="status"
                    :color="statusConfig[status].color"
                    :dark="statusConfig[status].dark"
                  >
                  </v-chip>
                </div>
            </v-toolbar>
            <CodeEditor :code="resultContent" :readOnly="true" :lineNumber="false" language="json"/>
          </v-card>
        </v-col>
      </v-row>
    </v-col>
    <v-divider vertical></v-divider>
    <v-col cols="5">
      <v-row style="height:65vh">
        <v-col>
          <v-card flat tile>
            <v-toolbar flat dense>
              <v-toolbar-title class="title overline">Input</v-toolbar-title>
              <v-spacer></v-spacer>
              <v-btn-toggle dense group mandatory v-model="config.input.language">
                <v-btn value="json">JSON</v-btn>
                <v-btn value="yaml">YAML</v-btn>
              </v-btn-toggle>
            </v-toolbar>
            <CodeEditor v-model="config.input.content" :language="config.input.language"/>
          </v-card>
        </v-col>
      </v-row>
      <v-divider></v-divider>
      <v-row style="height:34vh">
        <v-col>
          <v-card flat tile>
            <v-toolbar flat dense>
              <v-toolbar-title class="title overline">Parameters</v-toolbar-title>
              <v-spacer></v-spacer>
              <v-btn-toggle dense group mandatory v-model="config.params.language">
                <v-btn value="json">JSON</v-btn>
                <v-btn value="yaml">YAML</v-btn>
              </v-btn-toggle>
            </v-toolbar>
            <CodeEditor v-model="config.params.content" :language="config.params.language"/>
          </v-card>
        </v-col>
      </v-row>
    </v-col>
  
    <v-snackbar v-model="snackbar" color="red">
      Unexpected error occured
    </v-snackbar>
  </v-row>
</template>

<script>
import CodeEditor from '@/components/CodeEditor.vue'

export default {
  components: { CodeEditor },
  data () {
    return {
      config: {
        policy: {
          content: '',
          language: 'yaml',
          errors: null
        },
        input: {
          content: '',
          language: 'json',
          errors: null
        },
        params: {
          content: '',
          language: 'json',
          errors: null
        }
      },
      result: null,
      snackbar: false,
      loading: false,
      status: null,
      statusConfig: {
        passed: {
          color: 'green',
          dark: true
        },
        failed: {
          color: 'red',
          dark: true
        },
        error: {
          color: 'gray',
          dark: true
        },
        skipped: {
          color: 'grey lighten-5',
          dark: false
        }
      }
    }
  },
  computed: {
    resultContent() {
      if (this.result) {
        return JSON.stringify(this.result, null, 4)
      }
      return ''
    },
  },
  methods: {
    saveConfig() {
      localStorage.setItem('config', JSON.stringify(this.config))
    },
    loadConfig() {
      let config = JSON.parse(localStorage.getItem('config'))
      if (config) {
        this.config = config
      }
    },
    run() {
      this.saveConfig()
      this.result = null
      this.status = null
      this.loading = true
      fetch(
        "https://ahsayde.me/yapl-playground/api/eval", {
          method: "POST",
          body: JSON.stringify({
            input: this.config.input.content,
            policy: this.config.policy.content,
            params: this.config.params.content
        })
      })
      .then(response => {
        if (!response.ok) {
          this.status = 'error'
        }
        return response.json()
      })
      .then((data) => {
        this.result = data
        if (this.result.passed) {
          this.status = 'passed'
        } else if (this.result.failed){
          this.status = 'failed'
        } else if (this.result.ignored){
          this.status = 'skipped'
        }
      })
      .catch(() => {
        this.snackbar = true
      })
      .finally(() => {
        this.loading = false
      })
    }
  },
  created () {
    this.loadConfig()
  },
};
</script>
