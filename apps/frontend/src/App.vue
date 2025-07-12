<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <RouterView />
</template>

<style>
/* グローバルスタイル */
:root {
  --color-background: #ffffff;
  --color-background-soft: #f8f9fa;
  --color-background-mute: #f1f3f4;
  --color-border: #e1e5e9;
  --color-border-hover: #d3d3d3;
  --color-heading: #2c3e50;
  --color-text: #213547;
  --color-text-mute: #64748b;
}

@media (prefers-color-scheme: dark) {
  :root {
    --color-background: #1a1a1a;
    --color-background-soft: #262626;
    --color-background-mute: #2e2e2e;
    --color-border: #404040;
    --color-border-hover: #4a4a4a;
    --color-heading: #ffffff;
    --color-text: #e5e5e5;
    --color-text-mute: #a1a1a1;
  }
}

* {
  box-sizing: border-box;
}

html {
  margin: 0;
  padding: 0;
  width: 100%;
}

body {
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background-color: var(--color-background);
  color: var(--color-text);
  width: 100%;
  overflow-x: hidden;
}
</style>
