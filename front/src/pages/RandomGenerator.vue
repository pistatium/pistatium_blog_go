<template>
  <div class="container">
    <div class="possible_chars">
      <label class="" for="possible">Randomに使う文字列候補</label>
      <input id="possible_input monospace" type="text" class="possible_input" name="possible" v-model="input_chars"/>

    </div>
    <div class="copied">
      <span  v-if="copied">
      コピーしました。 <span class="monospace">{{copied}}</span>
      </span>
    </div>
    <div class="wrap">
      <ul class="result" id="result">
        <li class="" v-for="size in box_list" v-bind:key="size">
          <div class="result_item">
            <span class="">{{ size }} 文字</span>
            <button class="copy btn" v-bind:data-clipboard-target="'#' + generated_id(size)">COPY</button>
            <input type="text" readonly class="generated monospace" v-bind:id="generated_id(size)" v-bind:value="random_pool.slice(size, size + size)" />
          </div>
        </li>
      </ul>
      <div class="description">
        <p>
          ブラウザ搭載の安全な暗号生成器を利用してランダムな文字列を生成するツールです。
          乱数生成時にサーバーへの通信せずに使えるので安心してお使い頂けます。
        </p>
        <p>
          上のボックスで使う文字種別を変更することが出来ます。
        </p>

        <Adsense
            class="ad"
            data-ad-client="ca-pub-2359565431337443"
            data-ad-slot="9814535793">
        </Adsense>
      </div>

    </div>
  </div>
</template>

<script>
const INTERVAL = 180;
const box_list = [6, 8, 10, 12, 16, 24, 32, 48, 64, 128, 512];

export default {
  name: 'RandomGenerator',
  data: () => ({
    drawer: null,
    box_list: box_list,
    random_pool: "",
    input_chars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789@-_$%#",
    copied: "",
    initialized: false,
  }),
  mounted() {
    document.title = "MachRandGenerator - Pistatium"
    let recaptchaScript = document.createElement('script')
    recaptchaScript.setAttribute('src', 'https://cdn.jsdelivr.net/npm/clipboard@2.0.8/dist/clipboard.min.js')
    document.head.appendChild(recaptchaScript)

    const vue = this;
    setInterval(this.generate_pool, INTERVAL)
  },
  methods: {
    init() {
      if (this.initialized) {
        return;
      }
      if (!ClipboardJS) {
        return
      }
      const vue = this
      const clipboard = new ClipboardJS('.btn');
      clipboard.on('success', function(e) {
        vue.copied = e.text;
      });
      this.initialized = true;
    },
    generate_pool() {
      this.init()
      const key = new Uint32Array(1000);
      const r = window.crypto.getRandomValues(key);
      this.random_pool = r.reduce(this.get_char, "")
    },
    get_char(p, i) {
      return p + this.input_chars.charAt(i % this.input_chars.length);
    },
    generated_id(size) {
      return "generated-" + size;
    }
  }
}
</script>

<style scoped>
.container {
  margin-top: 64px;
}
.monospace {
  font-family: "Courier New", Consolas, monospace;
}

.btn {
  display: block;
  float: right;
  width: 120px;
  padding: 2px;
  text-align: center;
  text-decoration: none;
  color: #558b2f;
  background: #fff;
  border:1px solid #fff;
  border-radius: 2px;
}
.btn:hover {
  background: #ecffea;
  cursor: pointer;
  text-decoration: none;
}

.possible_chars {
  margin: 12px 0;
}

.possible_input {
  border: 1px solid #cccccc;
  padding: 6px;
  width: 100%;
  color: #558b2f;
}
.copied {
  height: 24px;
  width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
  font-size: 14px;
  color: #569033;
}

.result {
  width: 480px;
  text-decoration: none;
  list-style-type: none;
}
.result_item {
  margin: 24px 0;
}
.generated {
  width: 100%;
  max-width: 100%;
  background: #f7fff6;
  border: 1px solid #558b2f;
  padding: 2px 6px;
}
.wrap{
  display: flex;
  flex-direction: row;
}
@media screen and (max-width: 600px) {
  .wrap {
    flex-direction: column;
  }
}
.description {
  flex: 1;
  min-width: 300px;
  margin: 12px;
  font-size: 14px;
}
</style>
