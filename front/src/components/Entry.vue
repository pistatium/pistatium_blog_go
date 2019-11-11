<template>
    <v-card color="grey lighten-5">

        <v-card-text>
            {{ entry.Datetime.slice(0, 10) }}

            <p class="display-1 font-weight-black light-green--text text--darken-3">
                {{ entry.Title }}
            </p>

            <div class="text--primary">
                {{ entry.Body }}
            </div>
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
    export default {
        name: "Entry",
        props: ["entry"],
        computed: {
            link: function() {
                return `/show/${this.entry.Id}`
            },
        }
    }
</script>

<style scoped>
    .v-card {
        margin: 48px 0;
    }
</style>
