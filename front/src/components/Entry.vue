<template>
    <v-card color="grey lighten-5">

        <v-card-text>
            {{ entry.Datetime.slice(0, 10) }}

            <p class="display-1 font-weight-black light-green--text text--darken-3">
                {{ entry.Title }}
            </p>

            <div class="text--primary" v-html="markdown"></div>
        </v-card-text>

        <v-card-actions v-if="!entry.More">
                <v-btn
                        text
                        color="green darken-3 accent-4"
                        :href=link
                >
                    &gt;&gt;&gt; 続きを読む
                </v-btn>

        </v-card-actions>
    </v-card>

</template>

<script>
    const marked = require('marked');

    export default {
        name: "Entry",
        props: ["entry"],
        computed: {
            link: function() {
                return `/show/${this.entry.Id}`
            },
            markdown: function () {
                return marked(this.entry.Body || "", {breaks: true})
            }
        }
    }
</script>

<style scoped>
    .v-card {
        margin: 48px 0;
    }
</style>
