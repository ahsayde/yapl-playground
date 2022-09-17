<template>
    <div ref="editor" class="editor"/>
</template>

<script>
import * as monaco from 'monaco-editor';

export default {
  name: 'CodeEditor',
  props: {
    language: String,
    readOnly: Boolean,
    lineNumber: {
      type: Boolean,
      default: true
    },
    value: String,
    code: String
  },
  data () {
    return {
      editor: null
    }
  },
  computed: {
    content: {
      get () {
        return this.value
      },
      set (val) {
        this.$emit('input', val)
      }
    }
  },
  watch: {
    language(val) {
      if (val) {
        monaco.editor.setModelLanguage(this.editor.getModel(), val);
      }
    },
    code(val) {
      this.editor.getModel().setValue(val)
    },
  },
  mounted() {
    this.editor = monaco.editor.create(this.$refs.editor, {
      value: this.value,
      language: this.language,
      readOnly: this.readOnly,
      lineNumbers: this.lineNumber,
      fontSize: 16,
      minimap: {enabled: false},
      renderLineHighlight: "none",
      overviewRulerBorder: false,
      scrollBeyondLastLine: false,
      automaticLayout: false,
      overflowWidgetsDomNode: document.querySelector('body'),
      fixedOverflowWidgets: true,
    })
  
    this.editor.getModel().onDidChangeContent(() => {
      this.content = this.editor.getModel().getValue()
      let errors = monaco.editor.getModelMarkers()
      if (errors.length) {
        this.$emit('errors', errors)
      }
    })

    window.onresize =  () =>{
      this.editor.layout();
    };
  
  }
};
</script>
