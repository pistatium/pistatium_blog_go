<template>
    <v-card>
        <v-list two-line>
            <v-list-item-group>
                <v-subheader>最近のエントリ</v-subheader>
                <v-list-item v-for="entry in entries" v-bind:key="entry.Id" link :to="link(entry)" v-show="entry.Id !== entryId">
                    <v-list-item-content>
                        <v-list-item-title>{{ entry.Title }}</v-list-item-title>
                        <v-list-item-subtitle>{{ date(entry) }}</v-list-item-subtitle>
                    </v-list-item-content>
                </v-list-item>
            </v-list-item-group>
        </v-list>
    </v-card>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "RecentEntries",
        props: ["entryId"],
        data: () => ({
            entries: [],
        }),
        watch: {
            entryId: function (val) {
                this.loadEntries()
            },
        },
        mounted() {
            this.loadEntries()
        },
        methods: {
            loadEntries() {
                axios.get("/api/entries", {}).then(res => {
                    this.entries = res.data.entries
                    for (let e of res.data.entries) {
                        this.$root.entryHash[e.Id] = e
                    }
                })
            },
            link (e) {
                return `/show/${e.Id}`
            },
            date (e) {
                if (!e.Datetime) {
                    return ""
                }
                return e.Datetime.slice(0, 10)
            },
        }
    }
</script>

<style scoped>
</style>
