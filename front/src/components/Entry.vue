<template>
    <v-card color="">

        <v-card-text class="entry">
            <p class="entry-date">
                &nbsp;&nbsp;{{ date }}
            </p>

            <router-link tag="h1" v-bind:to=link class="display-1 font-weight-black light-green--text text--darken-3">
                {{ entry.Title }}
            </router-link>

            <div class="text--primary entry-body" v-html="markdown"></div>
            <v-divider v-if="show_detail"></v-divider>
            <div class="text--primary entry-body entry-more" v-html="markdown_more" v-if="show_detail"></div>

        </v-card-text>

        <v-card-actions v-if="!show_detail && entry.More">
                <v-btn
                        color="accent"
                        block
                        :to=link
                >
                    &gt; 続きを読む
                </v-btn>
        </v-card-actions>
        <div v-else class="text-right">
            <v-btn class="mx-2" fab dark small color="primary" :href=tweet_share_link target="_blank">
                <v-icon dark>mdi-twitter</v-icon>
            </v-btn>
            <v-btn class="mx-2" fab dark small color="primary" :href=hatena_bookmark_link target="_blank">
                <v-icon dark>mdi-alpha-b-box</v-icon>
            </v-btn>
        </div>
    </v-card>

</template>

<script>
    const marked = require('marked');
    const highlightjs = require('highlight.js')

    const renderer = new marked.Renderer();
    marked.setOptions({
        breaks: true,
        smartLists: true,
        renderer:renderer,
        highlight: function(code, lang) {
            return highlightjs.highlightAuto(code).value;
        }
    });
    renderer.link = function( href, title, text ) {
        return `<a target="_blank" href="${href}" title="${title}">${text}</a>`;
    }
    renderer.image = function( href, title, text ) {
        if (href === null) {
            return text;
        }
        return `<img src="${href}" loading="lazy">`;
    }

    export default {
        name: "Entry",
        props: ["entry", "show_detail"],
        watch: {
            entry: function (val) {
                document.title = this.entry.Title
            },
        },
        computed: {
            link: function() {
                return `/show/${this.entry.Id}`
            },
            date: function() {
                if (! this.entry.Datetime) {
                    return ""
                }
                return this.entry.Datetime.slice(0, 10)
            },
            markdown: function () {
                return marked(this.entry.Body || "", {}, function (err, out) {
                    if (window.twttr !== 'undefined') {
                        window.twttr.widgets.load();
                    }
                    return out
                })
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
        margin: 12px 0 48px;
        padding: 12px;
    }
    .entry {
        font-size: 108%;
        line-height: 220%;
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
    h1 {
        cursor: pointer;
    }
    h3 {
        margin-top: 8px;
    }

</style>

<style>
    .v-application .entry-body h2 {
        border-left: 6px #558b2f solid;
        padding-left: 6px;
        margin-top: 24px;
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
    .v-application .entry-body li {
        margin: 12px;
    }
    .v-application .entry-body pre code {
        width: 100%;
        margin: 5px 0;
        padding: 5px;
        background: rgb(31, 32, 34);
        line-height: 120%;
        color: #ffffff;
    }
    .v-application .entry-body code {
        padding: 3px;
        margin: 0 3px;
        background: rgb(31, 32, 34);
        line-height: 120%;
        color: #a5c261;
    }

</style>
