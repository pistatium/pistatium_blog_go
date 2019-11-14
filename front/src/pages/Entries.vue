<template>
    <div>
        <v-layout row wrap>
            <v-flex xs12>
                <Entry v-for="entry in entries" v-bind:key="entry.id" v-bind:entry=entry v-bind:show_detail="false"></Entry>
            </v-flex>
        </v-layout>

        <div class="text-center">
            <v-btn class="ma-2" tile color="green" dark v-bind:to=prev_page v-if="page > 0">&lt;&lt; Newer
            </v-btn>
            <v-btn class="ma-2" tile color="green" dark to="/" v-if="page !== 0">^ Top</v-btn>
            <v-btn class="ma-2" tile color="green" dark v-bind:to=next_page>&gt;&gt; Older</v-btn>

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
        }),
        mounted () {
            this.page = parseInt(this.$route.params.page, 10) || 0
            axios.get('/api/entries?page=' + this.page).then(res => {
                this.entries = res.data.entries;
                console.log(res);
            })
        },
        watch: {
            '$route' (to, from) {
                this.page = parseInt(this.params.page, 10) || 0
                axios.get('/api/entries?page=' + this.page).then(res => {
                    this.entries = res.data.entries;
                    console.log(res);
                })
            }
        },
        computed: {
            prev_page: function() {
                if (this.page === 1) {
                    return '/'
                }
                return `/${this.page - 1}`
            },
            next_page: function() {
                return `/${this.page + 1}`
            },
        }

    }
</script>
