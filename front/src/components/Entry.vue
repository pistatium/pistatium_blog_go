<template>
    <v-card color="grey lighten-5">

        <v-card-text class="entry" v-if="entry.id">
            <p class="text-right entry-date">
                <v-btn color="primary" text small text-right :href=link>
                    <v-icon>mdi-file-document-box</v-icon>
                    &nbsp;&nbsp;{{ entry.Datetime.slice(0, 10) }}

                </v-btn>
            </p>


            <p class="display-1 font-weight-black light-green--text text--darken-3">
                {{ entry.Title }}
            </p>

            <div class="text--primary" v-html="markdown"></div>

            <div class="text--primary" v-html="markdown_more" v-if="show_detail"></div>

        </v-card-text>

        <v-card-actions v-if="!show_detail && entry.More">
                <v-btn
                        color="accent"
                        block
                        :href=link
                >
                    &gt;&gt;&gt; 続きを読む
                </v-btn>
        </v-card-actions>
        <div v-else class="text-right">
            <v-btn class="mx-2" fab dark small color="primary">
                <v-icon dark>mdi-twitter</v-icon>
            </v-btn>
            <v-btn class="mx-2" fab dark small color="primary">
                <v-icon dark>mdi-alpha-b-box</v-icon>
            </v-btn>
        </div>
    </v-card>

</template>

<script>
    const marked = require('marked');

    export default {
        name: "Entry",
        props: ["entry", "show_detail"],
        computed: {
            link: function() {
                return `/show/${this.entry.Id}`
            },
            markdown: function () {
                return marked(this.entry.Body || "", {breaks: true})
            },
            markdown_more: function () {
                return marked(this.entry.More || "", {breaks: true})
            }
        }
    }
</script>

<style scoped>
    .v-card {
        margin: 48px 0;
        padding: 12px;
    }
    .entry {
        font-size: 110%;
        line-height: 220%;
    }
    .entry-date {
        margin: 0;
        padding: 0;
    }

    .entry-date .v-icon {
        padding-right: 6px;
    }

</style>
