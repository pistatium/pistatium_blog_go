<template>
    <v-card class="box">
        <v-list two-line>
            <v-list-item-group>
                <v-subheader class="head">最近のエントリ</v-subheader>
                <v-list-item v-for="entry in entries" v-bind:key="entry.Id" link :to="link(entry)" v-show="entry.Id !== entryId" class="item">
                    <v-list-item-content>
                        <v-list-item-subtitle>{{date(entry)}}</v-list-item-subtitle>
                        <v-list-item-title class="title light-green--text text--darken-3">{{ entry.Title }}</v-list-item-title>
                        <v-list-item-content class="pre-summary">{{ stripHtml(entry.Body) }}</v-list-item-content>
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
                return e.Datetime.slice(0, 10).replace(/-/g, ".")
            },
            stripHtml(s) {
                return s.replace(/<\/?[^>]+(>|$)/g, "").slice(0, 200)
            }
        }
    }
</script>

<style scoped>
    .box {
        border-top: 4px solid #7cb342;
        padding: 0;
    }
    .box .head {
        font-weight: 600;
        font-size: 18px;
    }
    .box .title {
        font-weight: 600;
        font-size: 18px;
    }
    .item {
        border-bottom: #fafafa 2px solid;
    }
    .pre-summary {
        padding: 0px 0 6px 0;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        font-size: 90%;
        color: #aaaaaa;

    }
</style>
