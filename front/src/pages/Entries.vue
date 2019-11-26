<template>
    <div>
        <div class="headline page-header">新着エントリ一覧 <span v-if="page!==0">(page {{page}})</span></div>
        <v-layout row wrap v-if="!empty">
            <v-flex xs12>
                <Entry v-for="(entry, index) in entries" v-bind:key="entry.id" v-bind:entry=entry v-bind:index="index"
                       v-bind:show_detail="false"></Entry>
            </v-flex>


        </v-layout>
        <div v-else class="no-entry" border="left">
            これ以上記事はありません
        </div>

        <div class="text-center" v-if="!this.$root.loading && !this.$root.error">
            <v-btn class="ma-2" tile color="green" dark v-bind:to=prev_page v-if="page > 0">&lt;&lt; Newer
            </v-btn>
            <v-btn class="ma-2" tile color="green" dark to="/" v-if="page !== 0">^ Top</v-btn>
            <v-btn class="ma-2" tile color="green" dark v-bind:to=next_page v-if="!empty">&gt;&gt; Older</v-btn>

        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";


    export default {
        name: 'Entries',
        components: {Entry},
        data: () => ({
            entries: [],
            page: 0,
            empty: false,
        }),
        mounted() {
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
            _fetchEntries () {
                this.$root.loading = true;
                this.empty = false
                this.page = parseInt(this.$route.params.page, 10) || 0
                axios.get('/api/entries?page=' + this.page).then(res => {
                    this.entries = res.data.entries;
                    if (this.entries.length === 0) {
                        this.empty = true
                        return
                    }
                    // cache
                    for (var e of res.data.entries) {
                        this.$root.entryHash[e.Id] = e
                    }
                }).catch(error => {
                    this.$root.error = error.response.statusText
                }).finally(() => this.$root.loading = false)
            }
        }


    }
</script>
<style scoped>
    .v-alert {
        margin: 32px;
        padding: 32px;
    }
    .page-header{
        color: #818181;
        font-weight: bold;
        text-align: center;
        margin-top: 18px;
    }
    .no-entry {
        margin: 72px 0;
        text-align: center;
        color: #558b2f;
    }
</style>
