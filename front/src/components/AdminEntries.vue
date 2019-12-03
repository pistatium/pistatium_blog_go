<template>
    <v-card>
        <v-list-item v-on:click="newEntry()">New</v-list-item>
        <v-divider></v-divider>
        <v-list-item v-for="entry in entries" v-bind:key="entry.Id" v-on:click="edit(entry)">
            <v-list-item-content v-bind:class="entry.Public ? 'public' : 'private'">
                <v-list-item-title>{{ entry.Title }}</v-list-item-title>
                <v-list-item-subtitle>{{ entry.Datetime }}</v-list-item-subtitle>
            </v-list-item-content>
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
            newEntry() {
                const id = String(new Date().getTime())
                this.$parent.editing = { Id: id }
            }
        }
    }
</script>

<style scoped>
    .private {
        background: #cccccc;
    }
</style>
