<template>
    <div>
        <v-layout row wrap>
            <v-flex xs12>
                <Entry v-for="entry in entries" v-bind:key="entry.id" v-bind:entry=entry v-bind:show_detail="false"></Entry>
            </v-flex>
        </v-layout>

        <div class="text-center">
            <v-btn class="ma-2" tile color="green" dark :href="'/' + (page - 1)" v-if="page > 0">&lt;&lt; Newer
            </v-btn>
            <v-btn class="ma-2" tile color="green" dark href="/" v-if="page !== 0">^ Top</v-btn>
            <v-btn class="ma-2" tile color="green" dark :href="'/' + (page + 1)">&gt;&gt; Older</v-btn>

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
        mounted() {
            this.page = parseInt(this.$route.params.page, 10) || 0
            axios.get('/api/entries?page=' + this.page).then(res => {
                this.entries = res.data.entries;
                console.log(res);
            })
        },

    }
</script>
