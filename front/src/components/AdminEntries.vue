<template>
    <v-card>
        <v-list-item v-for="entry in entries" v-bind:key="entry.Id" v-on:click="edit(entry)">
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
        mounted() {
            this.loadEntries()
        },
        methods: {
            loadEntries() {
                axios.get("/admin/api/entries", {}).then(res => {
                    this.entries = res.data.entries
                })
            },
            edit(entry) {
                this.$parent.editing = entry
            },
        }
    }
</script>

<style scoped>

</style>
