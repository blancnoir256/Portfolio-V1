<script setup lang="ts">
import { RouterLink } from 'vue-router'

interface Props {
  title?: string
  showNavigation?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  title: 'BN256 Portfolio',
  showNavigation: true
})
</script>

<template>
  <div class="layout-container">
    <!-- ヘッダー -->
    <header class="header">
      <div class="header-content">
        <img alt="BN256 icon" class="logo" src="/BN256_icon.png" width="60" height="60" />
        <h1 class="site-title">{{ props.title }}</h1>
        <nav v-if="props.showNavigation" class="navigation">
          <RouterLink to="/" class="nav-link">Home</RouterLink>
          <RouterLink to="/about" class="nav-link">About</RouterLink>
          <!-- <RouterLink to="/test" class="nav-link">Test</RouterLink> -->
        </nav>
      </div>
    </header>

    <!-- メインコンテンツ -->
    <main class="main-content">
      <slot />
    </main>

    <!-- フッター -->
    <footer class="footer">
      <div class="footer-content">
        <p>BN256だよー このホームページに対して著作権を主張するよー</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.layout-container {
  min-height: 100vh;
  display: grid;
  grid-template-rows: auto 1fr auto;
  grid-template-columns: 1fr;
  width: 100%;
}

.header {
  background-color: var(--color-background-soft);
  border-bottom: 1px solid var(--color-border);
  width: 100%;
  box-sizing: border-box;
  padding: 0;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 2rem;
  flex-wrap: wrap;
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
}

.logo {
  width: 60px;
  height: 60px;
  flex-shrink: 0;
}

.site-title {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--color-heading);
  margin: 0;
  flex: 1;
  text-align: center;
  min-width: 200px;
  left: 50%;
}

.navigation {
  display: flex;
  gap: 1rem;
  flex-shrink: 0;
}

.nav-link {
  padding: 0.5rem 1rem;
  text-decoration: none;
  color: var(--color-text);
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-link:hover {
  background-color: var(--color-background-mute);
}

.nav-link.router-link-active {
  background-color: var(--color-border-hover);
  color: var(--color-text);
}

.main-content {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
}

.footer {
  background-color: var(--color-background-soft);
  border-top: 1px solid var(--color-border);
  width: 100%;
  box-sizing: border-box;
  padding: 0;
}

.footer-content {
  text-align: center;
  color: var(--color-text-mute);
  padding: 1rem 2rem;
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
  box-sizing: border-box;
}

/* デスクトップ対応 */
@media (min-width: 769px) {
  .header-content {
    flex-wrap: nowrap;
  }
  
  .site-title {
    font-size: 1.8rem;
  }
  
  .main-content {
    padding: 3rem 2rem;
  }
}

/* タブレット対応 */
@media (min-width: 481px) and (max-width: 768px) {
  .header-content {
    padding: 1.5rem;
  }
  
  .site-title {
    font-size: 1.6rem;
  }
  
  .main-content {
    padding: 2rem 1.5rem;
  }
  
  .footer-content {
    padding: 1.5rem;
  }
}

/* モバイル対応 */
@media (max-width: 480px) {
  .header-content {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }
  
  .site-title {
    text-align: center;
    font-size: 1.3rem;
  }
  
  .navigation {
    justify-content: center;
    flex-wrap: wrap;
  }
  
  .main-content {
    padding: 1rem;
  }
  
  .footer-content {
    padding: 1rem;
  }
  
  .nav-link {
    padding: 0.4rem 0.8rem;
    font-size: 0.9rem;
  }
}
</style>
