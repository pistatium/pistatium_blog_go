<template>
    <entry show_detail=true v-bind:entry=entry v-if="entry.Id !== 0"></entry>
</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";

    export default {
        name: 'ShowEntry',
        components: {Entry},
        data: () => ({
            entry: {},
            entryId: 0,
        }),
        mounted() {
            this.entryId = parseInt(this.$route.params.id, 10) || 0
            if (this.$root.entryHash[this.entryId]) {
                this.entry = this.$root.entryHash[this.entryId]
            }
            if (!this.entry.id) {
                this.$root.loading = true;
                axios.get('/api/entries/' + this.entryId).then(res => {
                    this.entry = res.data
                    console.log(this.entry)
                }).catch(error => {
                    this.$root.error = error.response.statusText
                }).finally(() => this.$root.loading = false)
            }
        },
    }
</script>
