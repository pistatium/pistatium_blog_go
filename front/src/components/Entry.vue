<template>
    <v-card color="grey lighten-5">

        <v-card-text>
            {{ entry.created_at }}

            <p class="display-1 font-weight-black light-green--text text--darken-3">
                {{ entry.title }}
            </p>

            <div class="text--primary">
                {{ entry.body }}
            </div>
        </v-card-text>
        <v-card-actions>
            <v-btn
                    text
                    color="green darken-3 accent-4"
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
    }
</script>

<style scoped>

</style>
