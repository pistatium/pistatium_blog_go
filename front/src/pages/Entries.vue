<template>
    <div>
        <div class="headline page-header">新着エントリ一覧 <span v-if="page!==0">(page {{page}})</span></div>
        <div v-if="!empty">
            <Entry v-for="(entry, index) in entries" v-bind:key="entry.id" v-bind:entry=entry v-bind:index="index"
                   v-bind:show_detail="false" class="entry"></Entry>

        </div>
        <div v-else class="no-entry" border="left">
            これ以上記事はありません
        </div>

        <div class="text-center" v-if="!this.$root.loading && !this.$root.error">
            <v-btn class="ma-2" tile outlined color="green" dark v-bind:to=prev_page v-if="page > 0">&lt;&lt; もっと新しい記事へ
            </v-btn>
            <v-btn class="ma-2" tile outlined color="green" dark to="/" v-if="page !== 0">^ Top</v-btn>
            <v-btn class="ma-2" tile outlined color="green" dark v-bind:to=next_page v-if="!empty">&gt;&gt; もっと古い記事へ</v-btn>

        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";


    export default {
        name: 'EntryList',
        components: {Entry},
        data: () => ({
            entries: [],
            page: 0,
            empty: false,
        }),
        mounted() {
            // load from html
            this.page = parseInt(this.$route.params.page, 10) || 0
            const json_data = document.getElementById("entries-json").textContent
            if (json_data && json_data.startsWith("[")) {
                const data = JSON.parse(json_data)
                if (data.length > 0) {
                    this.entries = data
                    for (let e of data) {
                        this.$root.entryHash[e.Id] = e
                    }
                    return
                }
            }
            this._fetchEntries()
        },
        watch: {
            '$route'(to, from) {
                this._fetchEntries()
            }
        },
        computed: {
            prev_page: function () {
                if (this.page === 1) {
                    return '/'
                }
                return `/${this.page - 1}`
            },
            next_page: function () {
                return `/${this.page + 1}`
            },
        },
        methods: {
            _fetchEntries() {
                this.$root.loading = true;
                this.empty = false
                this.entries = []
                this.page = parseInt(this.$route.params.page, 10) || 0
                axios.get('/api/entries?page=' + this.page).then(res => {
                    this.entries = res.data.entries;
                    if (this.entries.length === 0) {
                        this.empty = true
                        return
                    }
                    // cache
                    for (let e of res.data.entries) {
                        this.$root.entryHash[e.Id] = e
                    }
                    document.title = "新着エントリ一覧 page" + this.page
                }).catch(error => {
                    this.$root.error = error.response.statusText
                }).finally(() => this.$root.loading = false)
            }
        }


    }
</script>
<style scoped>
    .v-alert {
        margin: 32px 0;
        padding: 32px 0;
    }

    .page-header {
        color: #b6adad;
        font-weight: bold;
        text-align: center;
        margin-top: 48px;
    }

    .no-entry {
        margin: 72px 0;
        text-align: center;
        color: #9f9f9f;
    }
    .entry {

    }
</style>
