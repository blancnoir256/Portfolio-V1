<script setup lang="ts">

</script>

<template>
  <div class="maintenance-container">
    <div class="maintenance-content">
      <h1 class="maintenance-title">メンテナンス中</h1>
      <p class="maintenance-message">ただいまメンテナンス中です。しばらくお待ちください。</p>
      <div class="maintenance-icon">🔧</div>
    </div>
  </div>
</template>

<style scoped>
.maintenance-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  text-align: center;
}

.maintenance-content {
  max-width: 1000px;
  padding: 2rem;
}

.maintenance-title {
  font-size: 2.5rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.maintenance-message {
  font-size: 1.2rem;
  color: var(--color-text-mute);
  margin-bottom: 2rem;
  line-height: 1.6;
}

.maintenance-icon {
  font-size: 4rem;
  opacity: 0.7;
}
</style>