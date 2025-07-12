<script setup lang="ts">
import BaseLayout from '../components/BaseLayout.vue'
</script>

<template>
  <BaseLayout title="Test Page">
    <div style="padding: 2rem; background-color: lightblue; text-align: center;">
      <h2>テスト用ページ</h2>
      <p>ヘッダーとフッターが全幅で表示されているかテストしています。</p>
      <div style="height: 300px; background-color: lightgreen; margin: 1rem 0;">
        <p style="padding-top: 100px;">この緑色の部分で全体のレイアウトを確認できます</p>
      </div>
    </div>
  </BaseLayout>
</template>
