<template>
    <entry show_detail=true class="single-entry" v-bind:entry=entry v-if="entry.Id !== 0"></entry>
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
            this.entryId = this.$route.params.id
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
        },
    }
</script>

<style scoped>
    .single-entry {
        margin-top: 32px;
    }

</style>
