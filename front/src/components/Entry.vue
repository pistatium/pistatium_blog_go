<template>

    <div>
        <v-card color="">

            <v-card-text class="entry">
                <p class="entry-date">
                    &nbsp;&nbsp;{{ date }}
                </p>

                <router-link tag="h1" v-bind:to=link
                             class="display-1 font-weight-black light-green--text text--darken-3">
                    {{ entry.Title }}
                </router-link>

                <div class="text--primary entry-body" v-html="markdown"></div>

                <InArticleAdsense
                        class="ads-in-article"
                        v-if="show_detail && entry.More"
                        data-ad-client="ca-pub-2359565431337443"
                        data-ad-slot="5140793616">
                </InArticleAdsense>

                <div class="text--primary entry-body entry-more" v-html="markdown_more" v-if="show_detail"></div>

            </v-card-text>

            <v-card-actions v-if="!show_detail && entry.More">
                <v-btn
                        color="accent"
                        block
                        large
                        raised
                        :to=link
                >
                    &gt; 続きを読む
                </v-btn>
            </v-card-actions>

            <div v-if="show_detail">
                <v-divider></v-divider>
                <v-card-actions dark>
                    <v-list-item class="grow">
                        <v-list-item-avatar color="grey darken-3">
                            <v-img
                                    src="/img/kimihiro_n.jpg"
                            ></v-img>
                        </v-list-item-avatar>
                        <v-list-item-content class="author">
                            <v-list-item-subtitle>Author</v-list-item-subtitle>
                            <v-list-item-title>@kimihiro-n</v-list-item-title>
                        </v-list-item-content>
                        <v-row
                                align="center"
                                justify="end"
                        >
                            <span class="share-label d-none d-sm-flex">Share: </span>
                            <v-btn class="mx-2" fab dark small color="primary" :href=tweet_share_link
                                   target="_blank">
                                <v-icon dark>mdi-twitter</v-icon>
                            </v-btn>
                            <v-btn class="mx-2" fab dark small color="primary" :href=hatena_bookmark_link
                                   target="_blank">
                                <v-icon dark>mdi-alpha-b-box</v-icon>
                            </v-btn>
                        </v-row>
                    </v-list-item>
                </v-card-actions>
            </div>
        </v-card>

        <Adsense
                class="ads-outer"
                v-if="index % 3 === 1 && ! show_detail"
                data-ad-client="ca-pub-2359565431337443"
                data-ad-slot="9814535793">
        </Adsense>
    </div>

</template>

<script>
    const marked = require('marked');
    const highlightjs = require('highlight.js')

    const renderer = new marked.Renderer();
    marked.setOptions({
        breaks: true,
        smartLists: true,
        renderer: renderer,
        highlight: function (code, lang) {
            return highlightjs.highlightAuto(code).value;
        }
    });
    renderer.link = function (href, title, text) {
        return `<a target="_blank" href="${href}" title="${title}">${text}</a>`;
    }
    renderer.image = function (href, title, text) {
        if (href === null) {
            return text;
        }
        return `<img src="${href}" loading="lazy">`;
    }

    export default {
        name: "Entry",
        props: ["entry", "show_detail", "index"],
        mounted() {
            if (window.twttr && window.twttr.widgets) {
                window.twttr.widgets.load();
            }
        },
        watch: {
            entry: function (val) {
                if (this.show_detail) {
                    document.title = this.entry.Title
                }
                if (window.twttr && window.twttr.widgets) {
                    window.twttr.widgets.load();
                }
            },
        },
        computed: {
            link: function () {
                return `/show/${this.entry.Id}`
            },
            date: function () {
                if (!this.entry.Datetime) {
                    return ""
                }
                return this.entry.Datetime.slice(0, 10)
            },
            markdown: function () {
                return marked(this.entry.Body || "")
            },
            markdown_more: function () {
                return marked(this.entry.More || "")
            },
            tweet_share_link: function () {
                const url = window.location.href
                return `https://twitter.com/intent/tweet?url=${url}&text=${this.entry.Title}`
            },
            hatena_bookmark_link: function () {
                const url = window.location.href
                return `https://b.hatena.ne.jp/entry/${url}`
            }
        }
    }
</script>

<style scoped>
    .v-card {
        margin: 12px 0 12px 0;
        padding: 12px;
    }

    .entry {
        font-size: 108%;
        line-height: 200%;
    }

    .entry-date {
        margin: 0;
        padding: 0;
    }

    .entry-body {
        margin: 24px 0 12px 0;
    }

    .entry-more {

    }

    .entry-date .v-icon {
        padding-right: 6px;
    }

    .share-label {
        color: #7aaa42;
        padding-right: 6px;
        font-size: 120%;
        font-weight: bold;
        vertical-align: center;
    }

    .ads-in-article {
        padding: 12px 0;
        margin: 24px -28px;
        background: #f4f4f4;
    }

    .ads-outer {
        margin: 40px 0;
    }
    .author {
        min-width: 100px;
    }
    h1 {
        cursor: pointer;
    }

    h3 {
        margin-top: 8px;
    }

</style>

<style>
    .v-application .entry-body p {
        margin: 12px 0 0 0;
    }
    .v-application .entry-body h2 {
        border-left: 8px #558b2f solid;
        padding-left: 12px;
        margin: 32px 0 12px 0;
    }

    .v-application .entry-body a {
        text-decoration: none;
        font-weight: bold;
        color: #569033;
    }

    .v-application .entry-body img {
        max-width: 100%;
        clear: both;
    }

    .v-application .entry-body ul {
        margin: 0 0 12px 0;
    }

    .v-application .entry-body li {
        margin: 0 12px;
    }

    .v-application .entry-body pre code {
        max-width: 100%;
        width: 100%;
        padding: 5px;
        background: rgb(31, 32, 34);
        line-height: 120%;
        color: #ffffff;
    }

    .v-application .entry-body code {
        font-size: 90%;
        font-family: Consolas, 'Courier New', Courier, Monaco, monospace;
        padding: 3px;
        margin: 0 3px;
        background: rgb(31, 32, 34);
        line-height: 120%;
        color: #a5c261;
    }
    .v-application blockquote {
        border-left: 6px solid #dcdcdc;
        background: #fafafa;
        margin: 6px 0;
        padding-left: 12px;
        font-size: 80%;
    }
</style>
