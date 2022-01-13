<template>
    <div>
        <v-card color="">
            <div class="text--primary entry-body" v-html="markdown"></div>
            <h1 class="entry-title light-green--text text--darken-3">Portfolio</h1>

            <h2>About</h2>

            <h3>kimihiro_n</h3>
            <p>
                サーバーサイドの開発を主にやってるエンジニア。
                Python とか Go を使っています。
                昔は Android アプリも作ってました。
            </p>

            <p>
                趣味でイラスト作成も最近やってます。
                もっとスキル身につけたい。
            </p>

            <ol class="timeline">
                <li>
                    <img src="/img/profile.png">
                    <h3>プロフィールアイコン</h3>
                    プロフィールにつかってるアイコンです。
                    最初はもうちょっとリアルよりでしたが今の形に落ち着きました。
                </li>
            </ol>
            <!--
            @kimihiro_n
            1988
            Python
            Go
            Android/Java, Kotlin
            AWS
            GCP
            Docker
            Django, FastAPI

            -->
        </v-card>
    </div>
</template>

<script>
    export default {
        name: 'MyPortfolio',
        data: () => ({
            drawer: null,
        }),
        mounted() {
            document.title = "Portfolio - Pistatium"
        }
    }
</script>

<style scoped>

    .v-card {
        margin: 12px 0 24px 0;
        padding: 0 12px 12px 12px;
        border-top: 4px solid #7cb342;
    }

    .entry {
        font-size: 108%;
        line-height: 200%;
    }

    @media screen and (max-width: 480px) {
        .v-card .entry {
            padding-left: 0;
            padding-right: 0;
        }
    }

    .entry-title {
        font-weight: 600;
        line-height: 120%;
        margin-top: 12px;
    }


    ol.timeline  {
        list-style-type: none;
    }

    .timeline li {
        position: relative;
        margin: 0;
        padding-bottom: 1em;
        padding-left: 20px;
    }

    .timeline li:before {
        content: '';
        background-color: #7cb342;
        position: absolute;
        bottom: 0;
        top: 0;
        left: 6px;
        width: 3px;
    }

    .timeline li:after {
        content: '';
        background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' aria-hidden='true' viewBox='0 0 32 32' focusable='false'%3E%3Ccircle stroke='none' fill='%237cb342' cx='16' cy='16' r='10'%3E%3C/circle%3E%3C/svg%3E");
        position: absolute;
        left: 0;
        height: 15px;
        width: 15px;
    }
</style>
