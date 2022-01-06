<template>
    <div>
        <entry show_detail=true class="single-entry" v-bind:entry=entry v-if="entry.Id !== 0"></entry>

        <Adsense
                v-if="entry.Body"
                class="ads-outer"
                data-ad-client="ca-pub-2359565431337443"
                data-ad-slot="9814535793">
        </Adsense>

        <v-lazy
                v-model="isActive"
                v-if="entry.Body"
                :options="{
                    threshold: .5
                }"
                min-height="400"
        >
            <RecentEntries v-bind:entryId="entry.Id"></RecentEntries>
        </v-lazy>
    </div>

</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";
    import RecentEntries from "../components/RecentEntries";

    export default {
        name: 'ShowEntry',
        components: {Entry, RecentEntries},
        data: () => ({
            entry: {},
            entryId: 0,
            isActive: false,
        }),
        watch: {
            '$route'(to, from) {
                this.loadEntry()
            }
        },
        methods: {
            loadEntry () {
                this.entryId = this.$route.params.id

                // load from html
                const json_data = document.getElementById("entry-json").textContent
                if (json_data && json_data.startsWith("{\"")) {
                    const data = JSON.parse(json_data)
                    if (data.Id && data.Id == this.entryId) {
                        this.entry = data
                        return
                    }
                }
                // load from cache
                if (this.$root.entryHash[this.entryId]) {
                    this.entry = this.$root.entryHash[this.entryId]
                    return
                }
                if (!this.entry.id) {
                    this.$root.loading = true;
                    axios.get('/api/entries/' + this.entryId).then(res => {
                        this.entry = res.data
                    }).catch(error => {
                        this.$root.error = error.response.statusText
                    }).finally(() => this.$root.loading = false)
                }
            }
        },
        mounted() {
            this.loadEntry()
        },
    }
</script>

<style scoped>
    .single-entry {
        margin-top: 64px;
    }
    .ads-outer {
        margin: 10px 0;
    }
</style>
