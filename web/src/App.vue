<template>
  <v-app>
    <v-app-bar app color="primary" dense dark>
      <v-app-bar-title>
        <span class="text-h5">YAPL</span> <span class="text-subtitle-1">Playground</span> 
      </v-app-bar-title>
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
