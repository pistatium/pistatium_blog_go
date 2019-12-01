<template>
    <v-card>
        <v-list-item v-for="entry in entries" v-bind:key="entry.Id">
            <v-list-item-content>{{ entry.Title }}</v-list-item-content>
        </v-list-item>
    </v-card>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "AdminEntries",
        data: () => ({
            entries: [],
        }),
        mounted () {
            console.log("test")
            this.loadEntries()
        },
        methods: {
            loadEntries() {
                axios.get("/api/entries", {}).then(res => {
                    this.entries = res.data.entries
                    console.log(res.data)
                })
            }
        }
    }
</script>

<style scoped>

</style>
