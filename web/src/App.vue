<template>
  <v-app>
    <v-app-bar app class="px-2" color="primary" dense dark>
      <v-app-bar-title>
        <span class="text-h5">YAPL</span> <span class="text-subtitle-1">Playground</span>
      </v-app-bar-title>
      <v-spacer></v-spacer>
      <v-btn icon href="https://github.com/ahsayde/yapl" target="_blank">
        <v-icon>mdi-github mdi-36px</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <HomePage/>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import  HomePage  from '@/components/HomePage.vue'
import '@/assets/style.css'

export default {
  name: 'App',
  components: { HomePage }
};
</script>
