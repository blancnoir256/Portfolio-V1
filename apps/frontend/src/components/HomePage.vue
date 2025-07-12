<script setup lang="ts">
</script>

<template>
  <div class="home-page">
    <section class="intro-section">
      <h1 class="intro-title">BN256です。</h1>
      <p class="intro-subtitle">"びーえぬにごろ"や"びーえぬ"と呼んでください。</p>
      <p class="intro-description">猫派です。</p>
    </section>

    <section class="features-section">
      <h2>興味のある分野</h2>
      <div class="features-grid">
        <div class="feature-card">
          <h3>プログラミング</h3>
          <p>広く浅くやってます</p>
          <p>Python, JavaScript, Go, C/C++, C#, 等々</p>
          <p>最近はPythonとGoをよく使います</p>
          <p>Pythonは言わずもがな、適当なコードを書きたいときに便利に使っています</p>
          <p>GoはEchoというフレームワークで簡単にwebサーバーを作れるので重宝しています</p>
        </div>
        <div class="feature-card">
          <h3>インフラ</h3>
          <p>最近沼にハマり始めました</p>
          <p>Dockerの便利さに魅了されて、流れでkubernetesにもハマっています</p>
          <p>k8sの理想状態に漸近させるという思想が好きです</p>
          <p>Nordも増やしたくなってProxmoxVEにも興味があります</p>
          <p>ただし、まだまだ初心者です</p>
        </div>
        <div class="feature-card">
          <h3>CTF</h3>
          <p>かなり好きです</p>
          <p>Forensics, OSINT, Web, Rev, 等々何でもやります</p>
          <p>だからといって上手ではないです</p>
          <p>下手の横好きですね</p>
          <p>好きこそものの上手なれと言いたいところです</p>
        </div>
      </div>
    </section>
    <hr style="border: none; border-top: 2px solid var(--color-border); margin: 3rem 0;" />
    <section class="features-section">
      <h2>気になっているモノたち(いつかやりたい)</h2>
      <div class="features-grid">
        <div class="feature-card">
          <h3>ゲーム制作</h3>
          <p>もともと私がプログラミングを始めたきっかけの一つです</p>
          <p>初めはUnityで遊びたくてC#を勉強していました</p>
          <p>最近は色々なことに手を出しすぎて、どれも中途半端になっています</p>
          <p>困ったものです</p>
          <p>UnityとかAEとかだけでなくSiv3Dとかにも興味があります</p>
          <p>レイヤーなんて低ければ低いほど良いですからね</p>
        </div>
        <div class="feature-card">
          <h3>アセンブリ言語</h3>
          <p>ゲーム制作の項目でも書いた通り、レイヤーなんて低ければ低いほど良いですからね</p>
          <p>CTFのRevをやっていると、自分でこれを書いて一つのアプリを作ってみたいな、という気持ちになります</p>
          <p>いつかやってみたいです</p>
        </div>
        <div class="feature-card">
          <h3>電子工作</h3>
          <p>私がプログラミングを初めたもう一つのきっかけです</p>
          <p>中学生だったか小学生だったかの頃にYouTubeで16bitCPUを自作している方を見かけて、そこからずっと興味があります</p>
          <p>私の低レイヤー好きはここから来ているかもしれません</p>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.home-page {
  max-width: 1000px;
  margin: 0 auto;
  width: 100%;
}

.intro-section {
  text-align: center;
  padding: 3rem 0;
  border-bottom: 1px solid var(--color-border);
  margin-bottom: 3rem;
}

.intro-title {
  font-size: 3rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
  font-weight: bold;
}

.intro-subtitle {
  font-size: 1.3rem;
  color: var(--color-text);
  margin-bottom: 1rem;
}

.intro-description {
  font-size: 1.1rem;
  color: var(--color-text-mute);
}

.features-section h2 {
  text-align: center;
  font-size: 2rem;
  color: var(--color-heading);
  margin-bottom: 2rem;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.feature-card {
  background-color: var(--color-background-soft);
  padding: 2rem 1.5rem;
  border-radius: 12px;
  border: 1px solid var(--color-border);
  text-align: left;
  transition: transform 0.2s, box-shadow 0.2s;
}

.feature-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.feature-card h3 {
  color: var(--color-heading);
  margin-bottom: 1rem;
  font-size: 1.3rem;
}

.feature-card p {
  color: var(--color-text-mute);
  line-height: 1.6;
}

/* デスクトップレイアウト */
@media (min-width: 769px) {
  .intro-title {
    font-size: 3.5rem;
  }

  .intro-subtitle {
    font-size: 1.4rem;
  }

  .features-grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 2.5rem;
  }

  .feature-card {
    padding: 2.5rem 2rem;
  }
}

/* タブレットレイアウト */
@media (min-width: 481px) and (max-width: 768px) {
  .intro-title {
    font-size: 2.5rem;
  }

  .features-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 2rem;
  }
}

/* モバイルレイアウト */
@media (max-width: 480px) {
  .intro-title {
    font-size: 2rem;
  }

  .intro-subtitle {
    font-size: 1.1rem;
  }

  .features-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .feature-card {
    padding: 1.5rem;
  }
}
</style>